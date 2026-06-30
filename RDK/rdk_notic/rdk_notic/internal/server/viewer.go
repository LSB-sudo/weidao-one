package server

const viewerHTML = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>RDK USB Camera Viewer</title>
  <style>
    :root {
      color-scheme: dark;
      font-family: Arial, sans-serif;
      background: #111827;
      color: #e5e7eb;
    }
    body {
      margin: 0;
      padding: 24px;
      background: #111827;
    }
    .wrap {
      max-width: 960px;
      margin: 0 auto;
    }
    h1 {
      margin-top: 0;
      font-size: 28px;
    }
    h2 {
      margin-bottom: 8px;
      font-size: 18px;
    }
    p, pre {
      line-height: 1.5;
    }
    .toolbar {
      display: flex;
      gap: 12px;
      margin: 16px 0;
      flex-wrap: wrap;
    }
    button {
      border: 0;
      border-radius: 6px;
      padding: 10px 16px;
      font-size: 15px;
      cursor: pointer;
      background: #2563eb;
      color: white;
    }
    button.secondary {
      background: #374151;
    }
    video {
      width: 100%;
      aspect-ratio: 16 / 9;
      background: black;
      border-radius: 8px;
    }
    .panel {
      background: #1f2937;
      padding: 12px;
      border-radius: 8px;
      overflow: auto;
      white-space: pre-wrap;
    }
    .grid {
      display: grid;
      gap: 16px;
      grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
      margin-top: 16px;
    }
    a { color: #93c5fd; }
  </style>
</head>
<body>
  <div class="wrap">
    <h1>USB Camera WebRTC Preview</h1>
    <p>Use this page from the same LAN as the board. The browser pulls a single H.264 video track over WebRTC.</p>
    <div class="toolbar">
      <button id="connectBtn">Connect</button>
      <button id="disconnectBtn" class="secondary">Disconnect</button>
      <a href="/health" target="_blank" rel="noreferrer">/health</a>
      <a href="/devices" target="_blank" rel="noreferrer">/devices</a>
    </div>
    <div id="status" class="panel">Idle</div>
    <video id="video" autoplay playsinline controls muted></video>
    <div class="grid">
      <div>
        <h2>WebRTC Stats</h2>
        <pre id="stats" class="panel">No active peer connection.</pre>
      </div>
      <div>
        <h2>Debug</h2>
        <pre id="debug" class="panel"></pre>
      </div>
    </div>
  </div>
  <script>
    const video = document.getElementById('video');
    const statusEl = document.getElementById('status');
    const statsEl = document.getElementById('stats');
    const debugEl = document.getElementById('debug');
    let pc;
    let statsTimer;

    function setStatus(message) {
      statusEl.textContent = message;
    }

    function setDebug(value) {
      debugEl.textContent = typeof value === 'string' ? value : JSON.stringify(value, null, 2);
    }

    function setStats(lines) {
      statsEl.textContent = Array.isArray(lines) ? lines.join('\n') : String(lines);
    }

    function stopStatsLoop() {
      if (statsTimer) {
        clearInterval(statsTimer);
        statsTimer = null;
      }
    }

    function startStatsLoop() {
      stopStatsLoop();
      updateStats();
      statsTimer = setInterval(updateStats, 1000);
    }

    function formatBytes(bytes) {
      if (!Number.isFinite(bytes) || bytes < 0) {
        return 'n/a';
      }
      if (bytes < 1024) {
        return bytes + ' B';
      }
      if (bytes < 1024 * 1024) {
        return (bytes / 1024).toFixed(1) + ' KiB';
      }
      return (bytes / (1024 * 1024)).toFixed(2) + ' MiB';
    }

    async function waitForICEGathering(peer) {
      if (peer.iceGatheringState === 'complete') {
        return;
      }
      await new Promise((resolve) => {
        const check = () => {
          if (peer.iceGatheringState === 'complete') {
            peer.removeEventListener('icegatheringstatechange', check);
            resolve();
          }
        };
        peer.addEventListener('icegatheringstatechange', check);
      });
    }

    async function updateStats() {
      if (!pc) {
        setStats('No active peer connection.');
        return;
      }

      try {
        const report = await pc.getStats();
        let inbound;
        let trackStats;
        let selectedPair;

        report.forEach((entry) => {
          if (entry.type === 'inbound-rtp' && entry.kind === 'video') {
            inbound = entry;
          }
          if (entry.type === 'track' && entry.kind === 'video') {
            trackStats = entry;
          }
          if (entry.type === 'candidate-pair' && entry.nominated && entry.state === 'succeeded') {
            selectedPair = entry;
          }
        });

        const playback = typeof video.getVideoPlaybackQuality === 'function'
          ? video.getVideoPlaybackQuality()
          : null;

        const width = trackStats?.frameWidth || video.videoWidth || 0;
        const height = trackStats?.frameHeight || video.videoHeight || 0;
        const framesDecoded = trackStats?.framesDecoded ?? inbound?.framesDecoded;
        const framesDropped = trackStats?.framesDropped ?? playback?.droppedVideoFrames;
        const bytesReceived = inbound?.bytesReceived;
        const jitterSeconds = inbound?.jitter;
        const packetsLost = inbound?.packetsLost;
        const hints = [];
        if ((framesDropped ?? 0) > 0 && (packetsLost ?? 0) === 0) {
          hints.push('high dropped frames with zero packet loss usually indicates decoder pressure or H.264 sample or packetization issues.');
        }

        setStats([
          'connectionState: ' + pc.connectionState,
          'iceConnectionState: ' + pc.iceConnectionState,
          'iceGatheringState: ' + pc.iceGatheringState,
          'videoResolution: ' + (width && height ? width + 'x' + height : 'n/a'),
          'framesDecoded: ' + (framesDecoded ?? 'n/a'),
          'framesDropped: ' + (framesDropped ?? 'n/a'),
          'bytesReceived: ' + formatBytes(bytesReceived ?? NaN),
          'jitter: ' + (jitterSeconds != null ? (jitterSeconds * 1000).toFixed(1) + ' ms' : 'n/a'),
          'packetsLost: ' + (packetsLost ?? 'n/a'),
          'selectedCandidatePair: ' + (selectedPair ? selectedPair.localCandidateId + ' -> ' + selectedPair.remoteCandidateId : 'n/a'),
          ...(hints.length ? ['hint: ' + hints.join(' ')] : []),
        ]);
      } catch (error) {
        setStats('Stats unavailable: ' + (error.message || String(error)));
      }
    }

    async function connect() {
      disconnect();
      setStatus('Creating offer...');
      setDebug('');
      setStats('Creating peer connection...');

      pc = new RTCPeerConnection({ iceServers: [] });
      pc.addTransceiver('video', { direction: 'recvonly' });
      pc.ontrack = (event) => {
        if (event.streams && event.streams[0]) {
          video.srcObject = event.streams[0];
        } else {
          video.srcObject = new MediaStream([event.track]);
        }
        setStatus('Streaming');
        startStatsLoop();
      };
      pc.onconnectionstatechange = () => {
        setStatus('Peer state: ' + pc.connectionState);
        updateStats();
      };
      pc.oniceconnectionstatechange = () => updateStats();

      const offer = await pc.createOffer();
      await pc.setLocalDescription(offer);
      await waitForICEGathering(pc);

      setStatus('Sending offer...');
      const response = await fetch('/offer', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(pc.localDescription),
      });

      const text = await response.text();
      let payload;
      try {
        payload = JSON.parse(text);
      } catch (err) {
        throw new Error(text || err.message);
      }
      setDebug(payload);

      if (!response.ok) {
        throw new Error(payload.error || 'offer failed');
      }

      await pc.setRemoteDescription(payload);
      setStatus('Waiting for video track...');
      startStatsLoop();
    }

    function disconnect() {
      stopStatsLoop();
      if (pc) {
        pc.close();
        pc = null;
      }
      if (video.srcObject) {
        for (const track of video.srcObject.getTracks()) {
          track.stop();
        }
        video.srcObject = null;
      }
      setStatus('Idle');
      setStats('No active peer connection.');
    }

    document.getElementById('connectBtn').addEventListener('click', async () => {
      try {
        await connect();
      } catch (error) {
        setStatus('Error');
        setDebug(error.message || String(error));
        disconnect();
      }
    });
    document.getElementById('disconnectBtn').addEventListener('click', () => disconnect());
  </script>
</body>
</html>`
