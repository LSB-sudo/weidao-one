# OneNET APP HTTP API Access

## 1. Problem Statement

The current OneNET device MQTT identity is bound to the device-side session and requires `ClientID=FieldGuard_I`.
If the APP or a desktop MQTT client reuses that same device identity, the active RDK device connection is forced offline.
That makes direct device MQTT access unsuitable as the APP data access path.

## 2. Selected Approach

The selected approach is:

- RDK remains the only device-side MQTT client connected to OneNET.
- RDK publishes device properties and receives platform replies on the device MQTT path.
- APP reads latest device properties and, when required, historical data through OneNET HTTP API or another application-side API.
- APP does not log in as the device MQTT client.

## 3. Data Flow Diagram

```text
+-------------+        Device MQTT         +-----------+        HTTP API / App-side API        +-------------+
|     RDK     | -------------------------> |  OneNET   | <------------------------------------ |     APP     |
|             | property publish           |           | latest property / history query       |             |
| - owns MQTT | <------------------------- |           |                                       | - reads data |
| - publishes | platform reply / downlink  |           |                                       | - shows UI   |
+-------------+                            +-----------+                                       +-------------+
```

## 4. What RDK Does

- Own the device MQTT session.
- Use the required device identity, including `ClientID=FieldGuard_I`.
- Publish device properties such as `battery_voltage`.
- Subscribe to its own `$sys/.../#` topics for OneNET replies or downstream instructions.
- Keep the device online and avoid concurrent login conflicts.

## 5. What APP Does

- Query the latest property values through OneNET HTTP API or a backend API.
- Query historical property data if the product needs trends or charts.
- Poll at a defined interval if no push channel exists yet.
- Present stale or missing data clearly in the UI.
- Avoid any direct device MQTT login with the device identity.

## 6. API/Data Requirements to Confirm from OneNET Docs

- Latest property query endpoint for the target product and device.
- Historical property query endpoint, time range parameters, and granularity limits.
- Required authentication mode for application-side access.
- Token scope, expiration, refresh flow, and quota limits.
- Response field names, timestamps, units, and error codes.
- Whether OneNET provides server-side push, rule engine forwarding, or WebSocket options for later realtime use.

## 7. Security Notes

- Use an application-side token or another app-scoped credential for APP data access.
- Do not embed the device token or device-secret-equivalent credential in the APP.
- Keep device MQTT credentials exclusive to the RDK runtime.
- Prefer a backend proxy if direct APP credential distribution creates operational or security risk.
- Redact sensitive tokens from logs, screenshots, and shared documents.

## 8. Validation Checklist

- RDK property publish succeeds and logs show `code=200`.
- OneNET console property page reflects the new property value.
- APP-side HTTP query returns the same latest property value.
- APP does not attempt device MQTT login with `FieldGuard_I`.
- Sensitive application or device credentials are not exposed in APP artifacts or logs.

## 9. Next Implementation Tasks

1. Confirm the exact OneNET HTTP API endpoints and authentication headers.
2. Confirm which properties the first APP screen must display.
3. Define APP polling frequency and stale-data behavior.
4. Decide whether APP calls OneNET directly or through a backend proxy.
5. Implement one end-to-end property read path and capture request/response examples.
6. Evaluate whether a later realtime path should use backend push, rule engine forwarding, or WebSocket.
