#include "imu_app.h"

#include <math.h>
#include <string.h>

static IMU_DATA_t IMU_App_Data = {0};
static ATTITUDE_DATA_t IMU_App_Attitude = {0};
static IMU_DATA_t IMU_App_ZeroPoint = {0};
static ATTITUDE_DATA_t IMU_App_ZeroAttitude = {0};

static uint8_t IMU_App_AttitudeStateInitialized;
static float IMU_App_AttitudeIntegral[3];
static float IMU_App_Quaternion[4] = {1.0f, 0.0f, 0.0f, 0.0f};
static float IMU_App_RelativeYaw;
static float IMU_App_OnlineGyroBiasZ;
static uint16_t IMU_App_StaticSampleCount;
static uint8_t IMU_App_MagYawReferenceReady;
static float IMU_App_MagYawReference;

static const float IMU_App_SamplePeriod = 0.005f;
static const float IMU_App_Kp = 8.0f;
static const float IMU_App_Ki = 0.008f;
static const float IMU_App_IntegralLimit = 2.0f;
static const float IMU_App_RelativeYawLimit = 3.1415926f;
static const float IMU_App_StaticAccelNorm = 9.8f;
static const float IMU_App_StaticAccelThreshold = 0.25f;
static const float IMU_App_StaticGyroThreshold = 0.8f * 0.01745329252f;
static const uint16_t IMU_App_StaticDetectCountThreshold = 40;
static const float IMU_App_OnlineBiasAdaptRate = 0.02f;
static const float IMU_App_MagNormMin = 20.0f;
static const float IMU_App_MagNormMax = 120.0f;
static const float IMU_App_YawMagErrorLimit = 8.0f * 0.01745329252f;
static const float IMU_App_YawMagPullGain = 0.04f;

static float IMU_App_ClampFloat(float value, float min_value, float max_value)
{
	if(value < min_value)
	{
		return min_value;
	}

	if(value > max_value)
	{
		return max_value;
	}

	return value;
}

static float IMU_App_AbsFloat(float value)
{
	return (value < 0.0f) ? -value : value;
}

static float IMU_App_WrapAnglePi(float angle)
{
	while(angle > IMU_App_RelativeYawLimit)
	{
		angle -= 2.0f * IMU_App_RelativeYawLimit;
	}

	while(angle < -IMU_App_RelativeYawLimit)
	{
		angle += 2.0f * IMU_App_RelativeYawLimit;
	}

	return angle;
}

static void IMU_App_ResetAttitudeState(void)
{
	IMU_App_AttitudeStateInitialized = 0;
	IMU_App_AttitudeIntegral[0] = 0.0f;
	IMU_App_AttitudeIntegral[1] = 0.0f;
	IMU_App_AttitudeIntegral[2] = 0.0f;
	IMU_App_Quaternion[0] = 1.0f;
	IMU_App_Quaternion[1] = 0.0f;
	IMU_App_Quaternion[2] = 0.0f;
	IMU_App_Quaternion[3] = 0.0f;
	IMU_App_RelativeYaw = 0.0f;
	IMU_App_OnlineGyroBiasZ = 0.0f;
	IMU_App_StaticSampleCount = 0;
	IMU_App_MagYawReferenceReady = 0;
	IMU_App_MagYawReference = 0.0f;
}

static void IMU_App_QuaternionToEuler(ATTITUDE_DATA_t *attitude)
{
	float q0 = IMU_App_Quaternion[0];
	float q1 = IMU_App_Quaternion[1];
	float q2 = IMU_App_Quaternion[2];
	float q3 = IMU_App_Quaternion[3];

	attitude->roll = atan2(2.0f * q2 * q3 + 2.0f * q0 * q1,
	                      -2.0f * q1 * q1 - 2.0f * q2 * q2 + 1.0f);
	attitude->pitch = asin(IMU_App_ClampFloat(-2.0f * q1 * q3 + 2.0f * q0 * q2, -1.0f, 1.0f));
	attitude->yaw = atan2(2.0f * (q1 * q2 + q0 * q3),
	                     q0 * q0 + q1 * q1 - q2 * q2 - q3 * q3);
}

static uint8_t IMU_App_TryGetMagYaw(IMU_DATA_t *data, float *mag_yaw)
{
	float roll;
	float pitch;
	float mx;
	float my;
	float magn_norm;

	magn_norm = sqrt(data->magn.x * data->magn.x +
	                 data->magn.y * data->magn.y +
	                 data->magn.z * data->magn.z);
	if(magn_norm < IMU_App_MagNormMin || magn_norm > IMU_App_MagNormMax)
	{
		return 0;
	}

	roll = atan2(data->accel.y, data->accel.z);
	pitch = atan2(-data->accel.x,
	              sqrt(data->accel.y * data->accel.y + data->accel.z * data->accel.z));

	mx = data->magn.x * cosf(pitch)
	   + data->magn.y * sinf(roll) * sinf(pitch)
	   + data->magn.z * cosf(roll) * sinf(pitch);
	my = data->magn.y * cosf(roll) - data->magn.z * sinf(roll);

	*mag_yaw = atan2(-my, mx);
	return 1;
}

static void IMU_App_UpdateOnlineGyroBias(IMU_DATA_t *data)
{
	float accel_norm;

	accel_norm = sqrt(data->accel.x * data->accel.x +
	                  data->accel.y * data->accel.y +
	                  data->accel.z * data->accel.z);

	if(IMU_App_AbsFloat(accel_norm - IMU_App_StaticAccelNorm) < IMU_App_StaticAccelThreshold &&
	   IMU_App_AbsFloat(data->gyro.x) < IMU_App_StaticGyroThreshold &&
	   IMU_App_AbsFloat(data->gyro.y) < IMU_App_StaticGyroThreshold &&
	   IMU_App_AbsFloat(data->gyro.z) < IMU_App_StaticGyroThreshold)
	{
		if(IMU_App_StaticSampleCount < 65535)
		{
			IMU_App_StaticSampleCount++;
		}
	}
	else
	{
		IMU_App_StaticSampleCount = 0;
	}

	if(IMU_App_StaticSampleCount >= IMU_App_StaticDetectCountThreshold)
	{
		IMU_App_OnlineGyroBiasZ += IMU_App_OnlineBiasAdaptRate * data->gyro.z;
	}

	data->gyro.z -= IMU_App_OnlineGyroBiasZ;
}

static void IMU_App_ApplyZeroPoint(const IMU_DATA_t *zero_point, const ATTITUDE_DATA_t *zero_attitude)
{
	memcpy(&IMU_App_ZeroPoint, zero_point, sizeof(IMU_DATA_t));
	memcpy(&IMU_App_ZeroAttitude, zero_attitude, sizeof(ATTITUDE_DATA_t));
	IMU_HW_SetZeroPoint(&IMU_App_ZeroPoint);
	IMU_App_ResetAttitudeState();
}

static void IMU_App_InitializeAttitudeState(IMU_DATA_t *imudata, ATTITUDE_DATA_t *attitude)
{
	float accel_norm;
	float magn_norm;
	float roll;
	float pitch;
	float yaw;
	float mag_yaw;

	accel_norm = sqrt(imudata->accel.x * imudata->accel.x +
	                  imudata->accel.y * imudata->accel.y +
	                  imudata->accel.z * imudata->accel.z);
	magn_norm = sqrt(imudata->magn.x * imudata->magn.x +
	                 imudata->magn.y * imudata->magn.y +
	                 imudata->magn.z * imudata->magn.z);

	if(accel_norm < 1e-6f || magn_norm < 1e-6f)
	{
		IMU_App_QuaternionToEuler(attitude);
		return;
	}

	imudata->accel.x /= accel_norm;
	imudata->accel.y /= accel_norm;
	imudata->accel.z /= accel_norm;
	imudata->magn.x /= magn_norm;
	imudata->magn.y /= magn_norm;
	imudata->magn.z /= magn_norm;

	roll = atan2(imudata->accel.y, imudata->accel.z);
	pitch = atan2(-imudata->accel.x,
	              sqrt(imudata->accel.y * imudata->accel.y + imudata->accel.z * imudata->accel.z));

	if(IMU_App_TryGetMagYaw(imudata, &mag_yaw))
	{
		yaw = mag_yaw;
		IMU_App_MagYawReference = mag_yaw;
		IMU_App_MagYawReferenceReady = 1;
	}
	else
	{
		yaw = 0.0f;
		IMU_App_MagYawReference = 0.0f;
		IMU_App_MagYawReferenceReady = 0;
	}

	IMU_App_Quaternion[0] = cosf(roll * 0.5f) * cosf(pitch * 0.5f) * cosf(yaw * 0.5f)
	                      + sinf(roll * 0.5f) * sinf(pitch * 0.5f) * sinf(yaw * 0.5f);
	IMU_App_Quaternion[1] = sinf(roll * 0.5f) * cosf(pitch * 0.5f) * cosf(yaw * 0.5f)
	                      - cosf(roll * 0.5f) * sinf(pitch * 0.5f) * sinf(yaw * 0.5f);
	IMU_App_Quaternion[2] = cosf(roll * 0.5f) * sinf(pitch * 0.5f) * cosf(yaw * 0.5f)
	                      + sinf(roll * 0.5f) * cosf(pitch * 0.5f) * sinf(yaw * 0.5f);
	IMU_App_Quaternion[3] = cosf(roll * 0.5f) * cosf(pitch * 0.5f) * sinf(yaw * 0.5f)
	                      - sinf(roll * 0.5f) * sinf(pitch * 0.5f) * cosf(yaw * 0.5f);

	IMU_App_AttitudeIntegral[0] = 0.0f;
	IMU_App_AttitudeIntegral[1] = 0.0f;
	IMU_App_AttitudeIntegral[2] = 0.0f;
	IMU_App_RelativeYaw = 0.0f;
	IMU_App_AttitudeStateInitialized = 1;

	IMU_App_QuaternionToEuler(attitude);
	attitude->yaw = 0.0f;
}

static void IMU_App_UpdateAttitude(IMU_DATA_t *imudata, ATTITUDE_DATA_t *attitude)
{
	float q0 = IMU_App_Quaternion[0];
	float q1 = IMU_App_Quaternion[1];
	float q2 = IMU_App_Quaternion[2];
	float q3 = IMU_App_Quaternion[3];
	float norm;
	float vx, vy, vz;
	float ex, ey;
	float pa, pb, pc, pd;
	float mag_yaw;
	float yaw_error;
	float q0q0;
	float q0q1;
	float q0q2;
	float q1q1;
	float q1q3;
	float q2q2;
	float q2q3;
	float q3q3;

	if(!IMU_App_AttitudeStateInitialized)
	{
		IMU_App_InitializeAttitudeState(imudata, attitude);
		attitude->roll -= IMU_App_ZeroAttitude.roll;
		attitude->pitch -= IMU_App_ZeroAttitude.pitch;
		attitude->yaw -= IMU_App_ZeroAttitude.yaw;
		return;
	}

	norm = sqrt(imudata->accel.x * imudata->accel.x +
	            imudata->accel.y * imudata->accel.y +
	            imudata->accel.z * imudata->accel.z);
	if(norm < 1e-10f)
	{
		return;
	}

	imudata->accel.x /= norm;
	imudata->accel.y /= norm;
	imudata->accel.z /= norm;

	q0q0 = q0 * q0;
	q0q1 = q0 * q1;
	q0q2 = q0 * q2;
	q1q1 = q1 * q1;
	q1q3 = q1 * q3;
	q2q2 = q2 * q2;
	q2q3 = q2 * q3;
	q3q3 = q3 * q3;

	vx = 2.0f * (q1q3 - q0q2);
	vy = 2.0f * (q0q1 + q2q3);
	vz = q0q0 - q1q1 - q2q2 + q3q3;

	ex = imudata->accel.y * vz - imudata->accel.z * vy;
	ey = imudata->accel.z * vx - imudata->accel.x * vz;

	IMU_App_AttitudeIntegral[0] = IMU_App_ClampFloat(IMU_App_AttitudeIntegral[0] + ex * IMU_App_SamplePeriod,
	                                                  -IMU_App_IntegralLimit, IMU_App_IntegralLimit);
	IMU_App_AttitudeIntegral[1] = IMU_App_ClampFloat(IMU_App_AttitudeIntegral[1] + ey * IMU_App_SamplePeriod,
	                                                  -IMU_App_IntegralLimit, IMU_App_IntegralLimit);
	IMU_App_AttitudeIntegral[2] = 0.0f;

	imudata->gyro.x = imudata->gyro.x + IMU_App_Kp * ex + IMU_App_Ki * IMU_App_AttitudeIntegral[0];
	imudata->gyro.y = imudata->gyro.y + IMU_App_Kp * ey + IMU_App_Ki * IMU_App_AttitudeIntegral[1];

	pa = q0;
	pb = q1;
	pc = q2;
	pd = q3;
	q0 = q0 + (-q1 * imudata->gyro.x - q2 * imudata->gyro.y - q3 * imudata->gyro.z) * (0.5f * IMU_App_SamplePeriod);
	q1 = pb + (pa * imudata->gyro.x + pc * imudata->gyro.z - pd * imudata->gyro.y) * (0.5f * IMU_App_SamplePeriod);
	q2 = pc + (pa * imudata->gyro.y - pb * imudata->gyro.z + pd * imudata->gyro.x) * (0.5f * IMU_App_SamplePeriod);
	q3 = pd + (pa * imudata->gyro.z + pb * imudata->gyro.y - pc * imudata->gyro.x) * (0.5f * IMU_App_SamplePeriod);

	norm = sqrt(q0 * q0 + q1 * q1 + q2 * q2 + q3 * q3);
	q0 /= norm;
	q1 /= norm;
	q2 /= norm;
	q3 /= norm;

	IMU_App_Quaternion[0] = q0;
	IMU_App_Quaternion[1] = q1;
	IMU_App_Quaternion[2] = q2;
	IMU_App_Quaternion[3] = q3;

	IMU_App_QuaternionToEuler(attitude);
	IMU_App_RelativeYaw += imudata->gyro.z * IMU_App_SamplePeriod;
	IMU_App_RelativeYaw = IMU_App_WrapAnglePi(IMU_App_RelativeYaw);

	if(IMU_App_TryGetMagYaw(imudata, &mag_yaw))
	{
		if(!IMU_App_MagYawReferenceReady)
		{
			IMU_App_MagYawReference = mag_yaw;
			IMU_App_MagYawReferenceReady = 1;
		}

		yaw_error = IMU_App_WrapAnglePi((mag_yaw - IMU_App_MagYawReference) - IMU_App_RelativeYaw);
		yaw_error = IMU_App_ClampFloat(yaw_error, -IMU_App_YawMagErrorLimit, IMU_App_YawMagErrorLimit);
		IMU_App_RelativeYaw = IMU_App_WrapAnglePi(IMU_App_RelativeYaw + yaw_error * IMU_App_YawMagPullGain);
	}

	attitude->roll -= IMU_App_ZeroAttitude.roll;
	attitude->pitch -= IMU_App_ZeroAttitude.pitch;
	attitude->yaw = IMU_App_RelativeYaw;
}

void IMU_App_Init(void)
{
	IMU_HW_ClearZeroPoint();
	memset(&IMU_App_Data, 0, sizeof(IMU_App_Data));
	memset(&IMU_App_Attitude, 0, sizeof(IMU_App_Attitude));
	memset(&IMU_App_ZeroPoint, 0, sizeof(IMU_App_ZeroPoint));
	memset(&IMU_App_ZeroAttitude, 0, sizeof(IMU_App_ZeroAttitude));
	IMU_App_ResetAttitudeState();
}

void IMU_App_CalibrateZero(u16 sample_count, u16 delay_ms_per_sample)
{
	IMU_DATA_t zero_point_sum = {0};
	IMU_DATA_t zero_point_avg = {0};
	ATTITUDE_DATA_t attitude_sum = {0};
	ATTITUDE_DATA_t attitude_avg = {0};
	u16 sample_index;

	if(sample_count == 0)
	{
		return;
	}

	for(sample_index = 0; sample_index < sample_count; sample_index++)
	{
		IMU_HW_Read9Axis(&IMU_App_Data);
		IMU_App_UpdateAttitude(&IMU_App_Data, &IMU_App_Attitude);

		zero_point_sum.accel.x += IMU_App_Data.accel.x;
		zero_point_sum.accel.y += IMU_App_Data.accel.y;
		zero_point_sum.accel.z += IMU_App_Data.accel.z;
		zero_point_sum.gyro.x += IMU_App_Data.gyro.x;
		zero_point_sum.gyro.y += IMU_App_Data.gyro.y;
		zero_point_sum.gyro.z += IMU_App_Data.gyro.z;
		zero_point_sum.magn.x += IMU_App_Data.magn.x;
		zero_point_sum.magn.y += IMU_App_Data.magn.y;
		zero_point_sum.magn.z += IMU_App_Data.magn.z;

		attitude_sum.roll += IMU_App_Attitude.roll;
		attitude_sum.pitch += IMU_App_Attitude.pitch;
		attitude_sum.yaw += IMU_App_Attitude.yaw;

		delay_ms(delay_ms_per_sample);
	}

	zero_point_avg.accel.x = zero_point_sum.accel.x / sample_count;
	zero_point_avg.accel.y = zero_point_sum.accel.y / sample_count;
	zero_point_avg.accel.z = zero_point_sum.accel.z / sample_count;
	zero_point_avg.gyro.x = zero_point_sum.gyro.x / sample_count;
	zero_point_avg.gyro.y = zero_point_sum.gyro.y / sample_count;
	zero_point_avg.gyro.z = zero_point_sum.gyro.z / sample_count;
	zero_point_avg.magn.x = zero_point_sum.magn.x / sample_count;
	zero_point_avg.magn.y = zero_point_sum.magn.y / sample_count;
	zero_point_avg.magn.z = zero_point_sum.magn.z / sample_count;

	attitude_avg.roll = attitude_sum.roll / sample_count;
	attitude_avg.pitch = attitude_sum.pitch / sample_count;
	attitude_avg.yaw = attitude_sum.yaw / sample_count;

	zero_point_avg.accel.z -= 9.8f;
	IMU_App_ApplyZeroPoint(&zero_point_avg, &attitude_avg);
	IMU_App_Update();
}

void IMU_App_Update(void)
{
	IMU_HW_Read9Axis(&IMU_App_Data);
	IMU_App_UpdateOnlineGyroBias(&IMU_App_Data);
	IMU_App_UpdateAttitude(&IMU_App_Data, &IMU_App_Attitude);

	axis_9Val = IMU_App_Data;
	AttitudeVal = IMU_App_Attitude;
}

void IMU_App_GetAttitude(ATTITUDE_DATA_t *attitude)
{
	if(attitude == 0)
	{
		return;
	}

	*attitude = IMU_App_Attitude;
}

void IMU_App_GetGyro(IMU_DATA_t *imu_data)
{
	if(imu_data == 0)
	{
		return;
	}

	*imu_data = IMU_App_Data;
}
