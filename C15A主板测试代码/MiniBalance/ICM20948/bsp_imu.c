#include "bsp_imu.h"

#include "icm20948_reg.h"
#include "bsp_siic.h"

IMU_DATA_t axis_9Val = { 0 };
ATTITUDE_DATA_t AttitudeVal = { 0 };

static IMU_DATA_t IMU_HW_ZeroPoint = { 0 };

uint8_t IMU_HW_Init(void)
{
	pIICInterface_t iicdev = &UserII2Dev;
	IIC_Status_t check_state = IIC_OK;
	uint8_t writebuf = 0;

	writebuf = REG_VAL_SELECT_BANK_0;
	check_state += iicdev->write_reg(ICM20948_DEV << 1, REG_BANK_SEL, &writebuf, 1, 500);

	check_state += iicdev->read_reg(ICM20948_DEV << 1, WHO_AM_I, &writebuf, 1, 500);
	if(writebuf != 0xEA)
	{
		return 1;
	}

	writebuf = (1 << 7);
	check_state += iicdev->write_reg(ICM20948_DEV << 1, PWR_MGMT_1, &writebuf, 1, 500);
	iicdev->delay_ms(100);

	writebuf = 0x00;
	check_state += iicdev->write_reg(ICM20948_DEV << 1, USER_CTRL, &writebuf, 1, 500);

	writebuf = 0x01;
	check_state += iicdev->write_reg(ICM20948_DEV << 1, PWR_MGMT_1, &writebuf, 1, 100);

	writebuf = REG_VAL_SELECT_BANK_2;
	check_state += iicdev->write_reg(ICM20948_DEV << 1, REG_BANK_SEL, &writebuf, 1, 100);

	writebuf = 0x01;
	check_state += iicdev->write_reg(ICM20948_DEV << 1, GYRO_SMPLRT_DIV, &writebuf, 1, 100);

	writebuf = (3 << 1) | (1 << 0) | (5 << 3);
	check_state += iicdev->write_reg(ICM20948_DEV << 1, GYRO_CONFIG_1, &writebuf, 1, 100);

	writebuf = 0x01;
	check_state += iicdev->write_reg(ICM20948_DEV << 1, ACCEL_SMPLRT_DIV_2, &writebuf, 1, 100);

	writebuf = (0 << 1) | (1 << 0) | (5 << 3);
	check_state += iicdev->write_reg(ICM20948_DEV << 1, ACCEL_CONFIG, &writebuf, 1, 100);

	writebuf = REG_VAL_SELECT_BANK_0;
	check_state += iicdev->write_reg(ICM20948_DEV << 1, REG_BANK_SEL, &writebuf, 1, 100);

	writebuf = (1 << 1);
	check_state += iicdev->write_reg(ICM20948_DEV << 1, INT_PIN_CFG, &writebuf, 1, 100);

	check_state += iicdev->read_reg(AK09916_DEV << 1, WIA, &writebuf, 1, 100);
	if(writebuf != 0x09)
	{
		return 1;
	}

	writebuf = (1 << 3);
	check_state += iicdev->write_reg(AK09916_DEV << 1, CNTL2, &writebuf, 1, 100);

	if(check_state != 0)
	{
		return 1;
	}

	IMU_HW_ClearZeroPoint();
	return 0;
}

uint8_t IMU_HW_DeInit(void)
{
	pIICInterface_t iicdev = &UserII2Dev;
	IIC_Status_t check_state = IIC_OK;
	uint8_t writebuf = 0;

	writebuf = REG_VAL_SELECT_BANK_0;
	check_state += iicdev->write_reg(ICM20948_DEV << 1, REG_BANK_SEL, &writebuf, 1, 500);

	check_state += iicdev->read_reg(ICM20948_DEV << 1, WHO_AM_I, &writebuf, 1, 500);
	if(writebuf != 0xEA)
	{
		return 1;
	}

	writebuf = (1 << 6);
	check_state += iicdev->write_reg(ICM20948_DEV << 1, PWR_MGMT_1, &writebuf, 1, 500);

	if(check_state != 0)
	{
		return 1;
	}

	return 0;
}

void IMU_HW_SetZeroPoint(const IMU_DATA_t *point)
{
	if(point == 0)
	{
		return;
	}

	IMU_HW_ZeroPoint = *point;
}

void IMU_HW_ClearZeroPoint(void)
{
	IMU_DATA_t zero_point = {0};
	IMU_HW_SetZeroPoint(&zero_point);
}

void IMU_HW_Read9Axis(IMU_DATA_t *data)
{
	uint8_t magnbuf[8];
	uint8_t tmpbuf[12];
	pIICInterface_t iicdev = &UserII2Dev;

	if(data == 0)
	{
		return;
	}

	iicdev->read_reg(ICM20948_DEV << 1, ACCEL_XOUT_H, tmpbuf, 12, 100);

	data->accel.x = (short)(tmpbuf[0] << 8 | tmpbuf[1]);
	data->accel.y = (short)(tmpbuf[2] << 8 | tmpbuf[3]);
	data->accel.z = (short)(tmpbuf[4] << 8 | tmpbuf[5]);

	data->accel.x *= 0.00059814453125f;
	data->accel.y *= 0.00059814453125f;
	data->accel.z *= 0.00059814453125f;

	data->accel.x -= IMU_HW_ZeroPoint.accel.x;
	data->accel.y -= IMU_HW_ZeroPoint.accel.y;
	data->accel.z -= IMU_HW_ZeroPoint.accel.z;

	data->gyro.x = (short)(tmpbuf[6] << 8 | tmpbuf[7]);
	data->gyro.y = (short)(tmpbuf[8] << 8 | tmpbuf[9]);
	data->gyro.z = (short)(tmpbuf[10] << 8 | tmpbuf[11]);

	data->gyro.x *= 0.06103515625f;
	data->gyro.y *= 0.06103515625f;
	data->gyro.z *= 0.06103515625f;

	data->gyro.x *= 0.01745329252f;
	data->gyro.y *= 0.01745329252f;
	data->gyro.z *= 0.01745329252f;

	data->gyro.x -= IMU_HW_ZeroPoint.gyro.x;
	data->gyro.y -= IMU_HW_ZeroPoint.gyro.y;
	data->gyro.z -= IMU_HW_ZeroPoint.gyro.z;

	iicdev->read_reg(AK09916_DEV << 1, HXL, magnbuf, 8, 100);

	if(((magnbuf[7] >> 3) & 0x01) == 0)
	{
		data->magn.x = (short)(magnbuf[1] << 8 | magnbuf[0]);
		data->magn.y = (short)(magnbuf[3] << 8 | magnbuf[2]);
		data->magn.z = (short)(magnbuf[5] << 8 | magnbuf[4]);

		data->magn.x *= 0.1495361328125f;
		data->magn.y *= 0.1495361328125f;
		data->magn.z *= 0.1495361328125f;
	}
}
