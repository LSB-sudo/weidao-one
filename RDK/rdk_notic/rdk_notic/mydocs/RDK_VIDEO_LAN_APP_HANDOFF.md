# RDK 视频回传同局域网 APP 交接说明

## 1. 文档目的

本文档用于给 Android APP 侧做交接，说明当前 RDK 视频回传方案的网络前提、APP 侧需要补充的能力，以及现有 `android/RdkCameraViewer` 工程中默认 IP 地址对应的位置。

本文档只做说明，不修改现有源码。

## 2. 当前方案结论

当前 RDK 视频预览方案基于以下前提：

- RDK 板端服务运行在 `/root/rdk_notic`
- Android APP 当前不是原生 WebRTC 客户端，而是一个最小化 WebView 壳
- APP 打开的是板端现有网页：`/viewer`
- 当前设计默认用于同一个局域网内访问
- 当前默认不依赖公网 STUN/TURN，首要场景就是 same-LAN preview

这与远端 `AGENTS.md` 中的长期工程决策一致：

- Android 第一版客户端位于 `/root/rdk_notic/android/RdkCameraViewer`
- 这是一个单 Activity 的 Android WebView 壳
- 通过明文 HTTP 在同局域网内加载 `/viewer`

## 3. 交接时需要明确告诉 APP 侧的事项

### 3.1 必须在同一个局域网

APP 手机和 RDK 板子必须处于同一个局域网，否则当前默认方案无法直接访问板端页面。

典型要求：

- 手机和 RDK 连接同一个 Wi-Fi
- 或者二者位于同一网段且可以互相访问
- 手机浏览器能直接打开 `http://<RDK_IP>:8080/viewer`

如果手机浏览器本身打不开这个地址，APP 内 WebView 也通常打不开。

### 3.2 APP 侧应该加入“手动输入 IP 地址”能力

当前 Android 工程里，板子的访问地址是写死的默认值，不适合作为正式交接方案。

因此建议 APP 侧增加：

- 手动输入 RDK IP 地址
- 可选输入端口，默认 `8080`
- 拼接得到最终地址：`http://<IP>:8080/viewer`
- 保存上次输入的 IP，避免每次重新输入

建议原因：

- 同局域网下，RDK IP 可能因为 DHCP 变化而变化
- 不同现场网络环境里，板子 IP 不一定固定为当前默认值
- 仅靠写死 IP，不利于现场部署和售后交接

## 4. 当前 Android 工程里 IP 地址对应哪里

远端工程路径：

- `android/RdkCameraViewer`

当前默认地址定义在：

- `android/RdkCameraViewer/app/src/main/res/values/strings.xml`

当前内容是：

```xml
<string name="viewer_url">http://192.168.3.142:8080/viewer</string>
```

这说明：

- 当前写死的 IP 是 `192.168.3.142`
- 端口是 `8080`
- APP 实际加载的是板端的 `/viewer` 页面

也就是说，APP 当前显示的视频页面地址，本质上就是：

```text
http://192.168.3.142:8080/viewer
```

## 5. 这个 IP 在代码里是怎么被使用的

### 5.1 启动时加载入口

文件：

- `app/src/main/java/com/rdk/cameraviewer/MainActivity.java`

关键调用：

```java
webView.loadUrl(getString(R.string.viewer_url));
```

这表示 APP 启动后，WebView 会直接读取 `strings.xml` 里的 `viewer_url`，然后加载该地址。

所以：

- `viewer_url` 是当前 APP 默认连接地址的唯一核心入口
- 如果只改默认 IP，优先就是改这里

### 5.2 报错提示里显示的地址也来自这里

同一个文件 `MainActivity.java` 中，加载失败时会执行：

```java
showError(getString(R.string.error_loading, getString(R.string.viewer_url)));
```

而 `strings.xml` 中对应错误文案是：

```xml
<string name="error_loading">Unable to open %1$s. Confirm the phone and RDK board are on the same LAN and the board service is running.</string>
```

这表示：

- 报错时界面上显示出来的 URL，也是 `viewer_url`
- 所以当前 APP 中“显示出来的那个 IP 地址”，本质上也是从 `viewer_url` 这个资源项取出来的

换句话说：

- 默认访问地址来自 `viewer_url`
- 报错时显示给用户看的地址，也来自 `viewer_url`

## 6. 当前界面里哪些文件和“手动输入 IP”改造有关

虽然这次不改源码，但如果 APP 侧后续要做“手动输入 IP”，主要会涉及以下文件：

### 6.1 布局文件

- `app/src/main/res/layout/activity_main.xml`

当前这里只有：

- 一个 `WebView`
- 一个顶部错误提示 `TextView`

如果以后要支持手动输入 IP，一般会在这里增加：

- `EditText`，用于输入 IP
- `Button`，用于确认或连接
- 如有需要，可再增加一个端口输入框，默认 `8080`

### 6.2 页面逻辑文件

- `app/src/main/java/com/rdk/cameraviewer/MainActivity.java`

后续如果要支持手动输入 IP，通常需要在这里增加：

- 读取输入框内容
- 组装 `http://<IP>:8080/viewer`
- 调用 `webView.loadUrl(...)`
- 可选保存到 `SharedPreferences`
- 启动时优先读取用户上次保存的 IP

### 6.3 默认资源文件

- `app/src/main/res/values/strings.xml`

即使后续支持手动输入 IP，这里仍建议保留一个默认地址，作为：

- 首次安装的默认值
- 输入为空时的回退值
- 调试用默认地址

## 7. 推荐给 APP 侧的实现口径

建议 APP 侧按下面的口径理解当前项目：

- 当前 APP 不是自己直接拉视频流，而是 WebView 打开板端 `/viewer`
- 板端地址本质上是 `http://<RDK_IP>:8080/viewer`
- 当前工程里这个地址被集中定义在 `viewer_url`
- 现场部署时，最需要补的能力是“手动输入 IP 地址”
- 同一局域网是当前方案成立的前提条件

建议最小改造目标：

1. 增加一个 IP 输入框
2. 默认端口先固定 `8080`
3. 点击连接后拼成 `http://<输入IP>:8080/viewer`
4. WebView 加载该地址
5. 报错提示中同步显示当前实际访问地址
6. 保存最近一次成功或手动输入的 IP

## 8. 与当前远端记录对齐的状态说明

本次编写文档前，已核对远端 `/root/rdk_notic` 内记录，结论如下：

- `AGENTS.md` 已明确 Android 第一版客户端是最小化 WebView 壳，加载 `/viewer`，同局域网明文 HTTP 访问
- 最新 journal 中与本交接最相关的记录是 2026-05-27 的 Android Viewer WebView 方案说明
- 当前 `.trellis/tasks/05-29-mqtt-phase1-gps-publisher/task.json` 状态仍为 `in_progress`，但该任务是 MQTT/GPS 方向，和本次视频回传 APP 交接文档不是同一主题

因此，本文档的交接依据主要来自：

- 远端 `AGENTS.md`
- `android/RdkCameraViewer` 当前源码
- 远端 journal 中关于 Android Viewer 和 same-LAN WebRTC viewer 的既有记录

## 9. 本次未做的事情

本次没有执行以下操作：

- 没有修改 Android 源码
- 没有新增手动输入 IP 的 UI
- 没有调整当前默认 `viewer_url`
- 没有修改板端 Go 服务

本次仅补充交接文档，供 APP 侧按此说明继续实现。
