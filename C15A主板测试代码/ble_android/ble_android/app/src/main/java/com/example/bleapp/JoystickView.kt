package com.example.bleapp

import android.content.Context
import android.graphics.Canvas
import android.graphics.Color
import android.graphics.Paint
import android.util.AttributeSet
import android.view.MotionEvent
import android.view.View
import kotlin.math.atan2
import kotlin.math.cos
import kotlin.math.min
import kotlin.math.sin
import kotlin.math.sqrt

/**
 * 虚拟摇杆自定义 View。
 *
 * 对外暴露归一化坐标：
 *   x ∈ [-1.0, 1.0]  →  左负右正
 *   y ∈ [-1.0, 1.0]  →  上正(前进)下负(后退)
 *
 * 注意 y 轴正方向是"屏幕上方 = 前进 = +1"，与 Android 屏幕坐标 y 轴相反。
 */
class JoystickView @JvmOverloads constructor(
    context: Context,
    attrs: AttributeSet? = null,
    defStyleAttr: Int = 0
) : View(context, attrs, defStyleAttr) {

    // ── 回调 ──────────────────────────────────────────────────────────────────

    interface OnJoystickMoveListener {
        /** @param x [-1,1] 左负右正   @param y [-1,1] 上正下负 */
        fun onJoystickMove(x: Float, y: Float)
    }

    var listener: OnJoystickMoveListener? = null

    // ── 画笔 ──────────────────────────────────────────────────────────────────

    private val paintBase = Paint(Paint.ANTI_ALIAS_FLAG).apply {
        color = Color.parseColor("#1a2a3a")
        style = Paint.Style.FILL
    }

    private val paintBaseRing = Paint(Paint.ANTI_ALIAS_FLAG).apply {
        color = Color.parseColor("#2a4a5a")
        style = Paint.Style.STROKE
        strokeWidth = 3f
    }

    private val paintCross = Paint(Paint.ANTI_ALIAS_FLAG).apply {
        color = Color.parseColor("#1e3a4a")
        style = Paint.Style.STROKE
        strokeWidth = 2f
    }

    private val paintThumb = Paint(Paint.ANTI_ALIAS_FLAG).apply {
        color = Color.parseColor("#44aaff")
        style = Paint.Style.FILL
    }

    private val paintThumbRing = Paint(Paint.ANTI_ALIAS_FLAG).apply {
        color = Color.parseColor("#88ccff")
        style = Paint.Style.STROKE
        strokeWidth = 3f
    }

    // ── 几何参数 ──────────────────────────────────────────────────────────────

    private var centerX = 0f
    private var centerY = 0f
    private var baseRadius = 0f
    private var thumbRadius = 0f

    // 当前摇杆头位置（像素）
    private var thumbX = 0f
    private var thumbY = 0f

    // ── 尺寸计算 ──────────────────────────────────────────────────────────────

    override fun onSizeChanged(w: Int, h: Int, oldW: Int, oldH: Int) {
        super.onSizeChanged(w, h, oldW, oldH)
        centerX = w / 2f
        centerY = h / 2f
        baseRadius = min(w, h) / 2f * 0.85f
        thumbRadius = baseRadius * 0.28f
        thumbX = centerX
        thumbY = centerY
    }

    // ── 绘制 ──────────────────────────────────────────────────────────────────

    override fun onDraw(canvas: Canvas) {
        super.onDraw(canvas)

        // 底盘
        canvas.drawCircle(centerX, centerY, baseRadius, paintBase)
        canvas.drawCircle(centerX, centerY, baseRadius, paintBaseRing)

        // 十字线
        canvas.drawLine(centerX - baseRadius, centerY, centerX + baseRadius, centerY, paintCross)
        canvas.drawLine(centerX, centerY - baseRadius, centerX, centerY + baseRadius, paintCross)

        // 摇杆头
        canvas.drawCircle(thumbX, thumbY, thumbRadius, paintThumb)
        canvas.drawCircle(thumbX, thumbY, thumbRadius, paintThumbRing)
    }

    // ── 触摸事件 ──────────────────────────────────────────────────────────────

    override fun onTouchEvent(event: MotionEvent): Boolean {
        when (event.action) {
            MotionEvent.ACTION_DOWN,
            MotionEvent.ACTION_MOVE -> {
                val dx = event.x - centerX
                val dy = event.y - centerY
                val dist = sqrt(dx * dx + dy * dy)

                if (dist <= baseRadius) {
                    thumbX = event.x
                    thumbY = event.y
                } else {
                    // 限制在底盘边缘
                    val angle = atan2(dy, dx)
                    thumbX = centerX + baseRadius * cos(angle)
                    thumbY = centerY + baseRadius * sin(angle)
                }

                invalidate()
                reportPosition()
                return true
            }

            MotionEvent.ACTION_UP,
            MotionEvent.ACTION_CANCEL -> {
                thumbX = centerX
                thumbY = centerY
                invalidate()
                listener?.onJoystickMove(0f, 0f)
                return true
            }
        }
        return super.onTouchEvent(event)
    }

    private fun reportPosition() {
        if (baseRadius <= 0f) return
        val nx = (thumbX - centerX) / baseRadius          // [-1, 1] 左负右正
        val ny = -(thumbY - centerY) / baseRadius          // [-1, 1] 上正下负（翻转 y）
        listener?.onJoystickMove(nx, ny)
    }
}
