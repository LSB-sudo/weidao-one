# RDK Camera Viewer

Minimal Android Studio project that wraps the existing board-side WebRTC viewer in a WebView.

## What it does

- Single Activity Android app
- Loads `http://192.168.3.142:8080/viewer` inside a WebView
- Enables JavaScript, DOM storage, WebChromeClient, WebViewClient, and media playback without a user gesture
- Uses cleartext HTTP so the current board-side service can be opened without changing the server

## Prerequisites

- Phone and RDK board must be on the same LAN
- The board-side service in `/root/rdk_notic` must already be running
- Android Studio on the development machine should open this directory and sync dependencies automatically

## Open in Android Studio

1. Copy or open `/root/rdk_notic/android/RdkCameraViewer` on your development machine.
2. In Android Studio, choose Open and select the `RdkCameraViewer` directory.
3. Let Android Studio sync the Gradle project.
4. Build or run the `app` configuration to install the APK on a phone.

## Change the board IP

Edit `app/src/main/res/values/strings.xml` and update `viewer_url`.

## APK build note

This repository does not include a Gradle wrapper and does not build the APK on the board. The intended workflow is to open the project in Android Studio and let it provision the matching Gradle/Android Gradle Plugin versions during sync.

## Gradle sync troubleshooting

If Android Studio reports that plugin `com.android.application` version `8.5.2` was not found and only searched `Gradle Central Plugin Repository`, check the root [settings.gradle](C:\Users\Lenovo\Desktop\卫稻一号远程\android\RdkCameraViewer\settings.gradle). It must define:

- `pluginManagement.repositories` with `google()`, `mavenCentral()`, and `gradlePluginPortal()`
- `dependencyResolutionManagement.repositories` with `google()` and `mavenCentral()`

The Android Gradle Plugin is resolved from Google's Maven repository, not only from the Gradle Plugin Portal. After updating `settings.gradle`, re-sync the project in Android Studio.

## First version scope

This first Android client is intentionally a WebView shell over the existing `/viewer` page. If later versions need tighter device integration, background behavior, or native controls, that can be handled in a future native WebRTC client.
