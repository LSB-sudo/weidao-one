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

class JoystickView @JvmOverloads constructor(
    context: Context,
    attrs: AttributeSet? = null,
    defStyleAttr: Int = 0
) : View(context, attrs, defStyleAttr) {

    interface OnJoystickMoveListener {
        fun onJoystickMove(x: Float, y: Float)
    }

    var listener: OnJoystickMoveListener? = null

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

    private var centerX = 0f
    private var centerY = 0f
    private var baseRadius = 0f
    private var thumbRadius = 0f
    private var thumbX = 0f
    private var thumbY = 0f

    override fun onSizeChanged(w: Int, h: Int, oldW: Int, oldH: Int) {
        super.onSizeChanged(w, h, oldW, oldH)
        centerX = w / 2f
        centerY = h / 2f
        baseRadius = min(w, h) / 2f * 0.85f
        thumbRadius = baseRadius * 0.28f
        thumbX = centerX
        thumbY = centerY
    }

    override fun onDraw(canvas: Canvas) {
        super.onDraw(canvas)
        canvas.drawCircle(centerX, centerY, baseRadius, paintBase)
        canvas.drawCircle(centerX, centerY, baseRadius, paintBaseRing)
        canvas.drawLine(centerX - baseRadius, centerY, centerX + baseRadius, centerY, paintCross)
        canvas.drawLine(centerX, centerY - baseRadius, centerX, centerY + baseRadius, paintCross)
        canvas.drawCircle(thumbX, thumbY, thumbRadius, paintThumb)
        canvas.drawCircle(thumbX, thumbY, thumbRadius, paintThumbRing)
    }

    override fun onTouchEvent(event: MotionEvent): Boolean {
        when (event.action) {
            MotionEvent.ACTION_DOWN, MotionEvent.ACTION_MOVE -> {
                val dx = event.x - centerX
                val dy = event.y - centerY
                val dist = sqrt(dx * dx + dy * dy)

                if (dist <= baseRadius) {
                    thumbX = event.x
                    thumbY = event.y
                } else {
                    val angle = atan2(dy, dx)
                    thumbX = centerX + baseRadius * cos(angle)
                    thumbY = centerY + baseRadius * sin(angle)
                }

                invalidate()
                reportPosition()
                return true
            }

            MotionEvent.ACTION_UP, MotionEvent.ACTION_CANCEL -> {
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
        val nx = (thumbX - centerX) / baseRadius
        val ny = -(thumbY - centerY) / baseRadius
        listener?.onJoystickMove(nx, ny)
    }
}
