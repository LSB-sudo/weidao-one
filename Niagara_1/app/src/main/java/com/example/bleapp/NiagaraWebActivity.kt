package com.example.bleapp

import android.annotation.SuppressLint
import android.graphics.Bitmap
import android.os.Bundle
import android.view.View
import android.webkit.WebChromeClient
import android.webkit.WebResourceRequest
import android.webkit.WebView
import android.webkit.WebViewClient
import android.widget.ImageButton
import android.widget.ProgressBar
import android.widget.TextView
import androidx.appcompat.app.AppCompatActivity

class NiagaraWebActivity : AppCompatActivity() {

    private lateinit var webView: WebView
    private lateinit var progressBar: ProgressBar
    private lateinit var tvUrl: TextView
    private lateinit var tvError: TextView
    private lateinit var btnBack: ImageButton
    private lateinit var btnRefresh: ImageButton

    @SuppressLint("SetJavaScriptEnabled")
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_niagara_web)

        webView = findViewById(R.id.webNiagara)
        progressBar = findViewById(R.id.progressNiagara)
        tvUrl = findViewById(R.id.tvNiagaraUrl)
        tvError = findViewById(R.id.tvNiagaraError)
        btnBack = findViewById(R.id.btnNiagaraBack)
        btnRefresh = findViewById(R.id.btnNiagaraRefresh)

        val targetUrl = getString(R.string.niagara_url)
        tvUrl.text = targetUrl

        btnBack.setOnClickListener {
            if (webView.canGoBack()) {
                webView.goBack()
            } else {
                finish()
            }
        }

        btnRefresh.setOnClickListener {
            tvError.visibility = View.GONE
            webView.reload()
        }

        webView.settings.apply {
            javaScriptEnabled = true
            domStorageEnabled = true
            builtInZoomControls = false
            displayZoomControls = false
            loadWithOverviewMode = true
            useWideViewPort = true
            mediaPlaybackRequiresUserGesture = false
        }

        webView.webChromeClient = object : WebChromeClient() {
            override fun onProgressChanged(view: WebView?, newProgress: Int) {
                progressBar.progress = newProgress
                progressBar.visibility = if (newProgress >= 100) View.GONE else View.VISIBLE
            }
        }

        webView.webViewClient = object : WebViewClient() {
            override fun shouldOverrideUrlLoading(view: WebView?, request: WebResourceRequest?): Boolean {
                return false
            }

            override fun onPageStarted(view: WebView?, url: String?, favicon: Bitmap?) {
                progressBar.visibility = View.VISIBLE
                tvError.visibility = View.GONE
                if (!url.isNullOrBlank()) {
                    tvUrl.text = url
                }
            }

            override fun onPageFinished(view: WebView?, url: String?) {
                progressBar.visibility = View.GONE
                if (!url.isNullOrBlank()) {
                    tvUrl.text = url
                }
            }

            @Deprecated("Deprecated in Java")
            override fun onReceivedError(
                view: WebView?,
                errorCode: Int,
                description: String?,
                failingUrl: String?
            ) {
                showError(failingUrl ?: targetUrl)
            }
        }

        webView.loadUrl(targetUrl)
    }

    override fun onBackPressed() {
        if (webView.canGoBack()) {
            webView.goBack()
        } else {
            super.onBackPressed()
        }
    }

    private fun showError(url: String) {
        progressBar.visibility = View.GONE
        tvUrl.text = url
        tvError.text = getString(R.string.niagara_error_loading, url)
        tvError.visibility = View.VISIBLE
    }
}
