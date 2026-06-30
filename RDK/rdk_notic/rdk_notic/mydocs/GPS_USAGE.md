# GPS 使用说明

本文说明如何在 RDK X5 上启动 `rdk-webrtc` 服务、访问 `/gps` 接口，以及如何判断 GPS 是否已经定位成功。

## 1. 前提条件

使用前请先确认：

- ATGM336H 已正确连接到 RDK X5。
- 串口设备节点存在，默认设备为：`/dev/serial/by-id/usb-1a86_USB_Serial-if00-port0`
- 天线已接好，建议在室外或靠近窗边等空旷环境测试。
- 首次上电后可能需要等待一段时间才能完成定位。

可先检查设备节点是否存在：

```bash
ls -l /dev/serial/by-id/
```

如果能看到 `usb-1a86_USB_Serial-if00-port0`，说明系统已经识别到 GPS 串口设备。

## 2. 启动服务

远端项目目录：`/root/rdk_notic`

进入目录后启动：

```bash
cd /root/rdk_notic
./.local-go/go/bin/go run ./cmd/rdk-webrtc
```

如果希望改用 `18080` 端口启动：

```bash
cd /root/rdk_notic
LISTEN_ADDR=:18080 ./.local-go/go/bin/go run ./cmd/rdk-webrtc
```

当前 GPS 默认就是启用状态，对应默认参数：

- `GPS_ENABLED=true`
- `GPS_DEVICE=/dev/serial/by-id/usb-1a86_USB_Serial-if00-port0`
- `GPS_BAUD=9600`

如果设备路径不是默认值，可以自行覆盖：

```bash
cd /root/rdk_notic
GPS_DEVICE=/dev/serial/by-id/你的设备名 ./.local-go/go/bin/go run ./cmd/rdk-webrtc
```

## 3. 访问 `/gps`

服务启动后，可以直接访问 GPS 接口。

如果服务监听默认 `:8080`：

```bash
curl http://127.0.0.1:8080/gps
```

如果服务监听 `:18080`：

```bash
curl http://127.0.0.1:18080/gps
```

浏览器访问示例：

- `http://板子IP:8080/gps`
- `http://板子IP:18080/gps`

例如板子 IP 是 `192.168.3.142` 时：

- `http://192.168.3.142:8080/gps`
- `http://192.168.3.142:18080/gps`

## 4. 查看 `/health`

`/health` 会返回服务整体状态，其中也包含简要的 GPS 状态，适合快速确认服务是否正常启动。

```bash
curl http://127.0.0.1:8080/health
```

浏览器访问示例：

- `http://板子IP:8080/health`
- `http://板子IP:18080/health`

## 5. `/gps` 返回字段说明

`/gps` 返回的是当前缓存的最新 GPS 快照。常用字段如下。

### 顶层状态字段

- `connected`
  - 是否已经成功打开 GPS 串口。
  - `true` 表示串口连接成功。
  - `false` 表示串口未连上，通常是设备节点不存在、设备被占用或模块未接好。

- `valid`
  - 当前定位是否有效。
  - `true` 表示已经得到可用定位。
  - `false` 表示还没有拿到有效定位，即使串口已经连通也可能出现这种情况。

- `antennaStatus`
  - 来自 NMEA `TXT` 信息的天线状态文本。
  - 当前实测常见值为：`ANTENNA OK`
  - 它表示天线状态正常，但不等于已经定位成功。

- `stale`
  - 数据是否过期。
  - `true` 表示最近一段时间没有收到新的 GPS 语句，当前缓存可能已经不新鲜。
  - `false` 表示最近仍在持续收到数据。

### `fix` 定位字段

当 `valid=true` 时，`fix` 中的内容才值得上层程序使用。

- `fix.latitude`
  - 纬度，十进制度。

- `fix.longitude`
  - 经度，十进制度。

- `fix.satellites`
  - 当前参与解算或被报告的卫星数量。

- `fix.fixQuality`
  - 定位质量，来自 GGA。
  - 一般来说，`0` 可理解为未定位，其它值表示已有不同等级的定位结果。

- `fix.hdop`
  - 水平精度因子。
  - 通常越小越好。

- `fix.speedKnots`
  - 速度，单位节。

- `fix.speedKph`
  - 速度，单位公里每小时。

- `fix.course`
  - 航向角，单位度。

- `fix.altitudeMeters`
  - 海拔高度，单位米。

补充说明：`fix` 里还可能包含 `timestampUtc`、`status`、`mode`、`navStatus`、`source` 等辅助字段，用于诊断数据来源和语句类型。

## 6. 三种常见状态判断

### 情况一：串口未连接

典型特征：

- `connected=false`
- `valid=false`
- `fix` 通常为空
- 可能同时带有 `lastError`

说明：

- 服务没有成功打开 GPS 串口。
- 先检查设备节点、接线和串口是否被其他程序占用。

### 情况二：串口已连接，但还未定位

典型特征：

- `connected=true`
- `valid=false`
- `antennaStatus` 可能为 `ANTENNA OK`
- 可能已经持续收到语句，但 `fix` 里的经纬度还不能作为有效定位使用
- 即使 `fix` 中已经出现部分缓存或辅助字段，例如速度、航向、HDOP，也不能据此认定已经定位成功；上层仍应以 `connected=true && valid=true` 作为使用经纬度的条件

这正是当前实测过的状态：

- `connected: true`
- `valid: false`
- `antennaStatus: "ANTENNA OK"`

说明：

- 串口通信已经正常。
- 模块还没有完成定位。
- 这种情况通常需要把天线放到更空旷的位置，并等待一段时间。
### 情况三：定位成功

典型特征：

- `connected=true`
- `valid=true`
- `fix.latitude` 和 `fix.longitude` 有有效值
- `fix.satellites`、`fix.fixQuality`、`fix.hdop` 等字段也会更有参考意义

说明：

- 此时可以把 `/gps` 的经纬度等信息交给上层程序使用。

## 7. 排查建议

如果 `/gps` 结果不符合预期，可以按下面顺序排查：

### 1）检查设备节点

```bash
ls -l /dev/serial/by-id/
```

确认默认设备：

```bash
ls -l /dev/serial/by-id/usb-1a86_USB_Serial-if00-port0
```

如果节点不存在，先检查 USB 转串口是否被系统识别。

### 2）检查天线连接

- 确认天线已接好。
- `antennaStatus=ANTENNA OK` 只能说明天线状态正常，不代表已经定位成功。

### 3）检查环境

- 优先在室外、空旷处测试。
- 室内、金属遮挡、贴近墙体或没有天空视野时，容易长时间无法定位。

### 4）给模块足够等待时间

- 冷启动时，GPS 可能需要几十秒到几分钟才能完成首次定位。
- 看到 `connected=true` 但 `valid=false` 时，不要立刻判定模块异常。

### 5）检查端口占用或服务端口冲突

如果服务本身没有启动成功，先检查监听端口是否冲突。比如想使用 `18080` 端口时：

```bash
LISTEN_ADDR=:18080 ./.local-go/go/bin/go run ./cmd/rdk-webrtc
```

如果已有其他进程占用该端口，需要换一个端口再启动。

## 8. 给上层程序的调用建议

上层程序建议轮询 `/gps`，而不是假定服务启动后立刻就有有效定位。

建议逻辑：

1. 周期性请求 `/gps`
2. 先判断 `connected`
3. 再判断 `valid`
4. 只有在 `connected=true && valid=true` 时，才使用 `fix.latitude` 和 `fix.longitude`
5. 如果 `stale=true`，即使之前定位成功，也建议视为数据可能已过期

一个简单示意：

```python
resp = GET /gps
if resp.status_code != 200:
    记录错误，稍后重试
    return

data = resp.json()
if not data.get("connected", False):
    提示 GPS 串口未连接，稍后重试
    return

if not data.get("valid", False):
    # 这里的 fix 可能带有速度、航向、HDOP 等缓存/辅助字段，
    # 但不能据此认为已经拿到有效经纬度
    提示 GPS 尚未定位成功，继续等待
    return

fix = data.get("fix") or {}
lat = fix.get("latitude")
lng = fix.get("longitude")
if lat is None or lng is None:
    记录异常：valid=true 但缺少经纬度
    return

使用 lat / lng 进入你的上层业务逻辑
```

如需排查问题，还可以关注这些诊断字段：

- `enabled`：GPS 功能是否启用
- `device`：当前串口设备路径
- `baud`：当前串口波特率
- `staleAfterSec`：超过多久未收到新数据会判定为 stale
- `lastReadAt`：最近一次读取串口数据的时间
- `lastSentenceAt`：最近一次收到 NMEA 语句的时间
- `lastValidAt`：最近一次拿到有效定位的时间
- `lastSentenceType`：最近一次语句类型，例如 GGA / RMC / TXT
- `lastSentence`：最近一次收到的原始语句
- `lastText`：最近一次 TXT 文本内容
- `lastError`：最近一次错误信息
## 9. 最小使用步骤

如果你只想快速验证 GPS 是否工作，可以直接按这个顺序操作：

```bash
cd /root/rdk_notic
LISTEN_ADDR=:18080 ./.local-go/go/bin/go run ./cmd/rdk-webrtc
```

新开一个终端查询：

```bash
curl http://127.0.0.1:18080/health
curl http://127.0.0.1:18080/gps
```

判断原则：

- `connected=false`：先查串口和接线
- `connected=true && valid=false`：先把天线移到空旷环境并继续等待
- `connected=true && valid=true`：说明 GPS 可以正常给上层使用
