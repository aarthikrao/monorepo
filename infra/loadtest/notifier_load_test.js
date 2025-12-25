import ws from 'k6/ws';
import { check, sleep } from 'k6';

// Configuration from CLI args (with sensible defaults)
// Usage:
//   k6 run notifier_load_test.js
//   k6 run --env WS_URL=wss://example.com/ws --env VUS=200 --env DURATION=1m notifier_load_test.js
//   k6 run --vus 100 --duration 2m notifier_load_test.js  (also works)

const DEFAULT_VUS = 50;
const DEFAULT_DURATION = '30s';
const DEFAULT_URL = 'ws://localhost:8080/ws';

const vus = __ENV.VUS ? parseInt(__ENV.VUS, 10) : DEFAULT_VUS;
const duration = __ENV.DURATION || DEFAULT_DURATION;
const wsUrl = __ENV.WS_URL || DEFAULT_URL;

export let options = {
  vus: vus,
  duration: duration,
};

// This script opens a WebSocket to the server and sends a small
// "method payload" formatted message that matches the server's expected
// framing: "echo {\"msg\":\"hello\"}". Each VU creates a single
// connection and sends periodic messages, then closes the connection.

export default function () {
  const url = wsUrl;

  const res = ws.connect(url, null, function (socket) {
    socket.on('open', function () {
      // send an initial message including VU/ITER to help correlate logs
      socket.send(`echo {"msg":"hello","vu":${__VU},"iter":${__ITER}}`);
    });

    socket.on('message', function (data) {
      // optional lightweight check - ensure we receive something back
      // (we don't parse JSON here to keep allocations low)
    });

    socket.on('close', function () {
      // closed by server or client
    });

    // send a few periodic messages (1 per second) to simulate activity
    for (let i = 0; i < 10; i++) {
      socket.send(`echo {"msg":"ping","count":${i},"vu":${__VU}}`);
      sleep(1);
    }

    socket.close();
  });

  // check the upgrade result (res may be null on failed connect)
  check(res, {
    'connected (101)': (r) => r && r.status === 101,
  });

  // small sleep to pace iterations (not strictly necessary when using duration)
  sleep(1);
}
