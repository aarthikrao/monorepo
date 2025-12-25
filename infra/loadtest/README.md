# WebSocket Loadtest

This folder contains all the scripts and documentation related to load testing the monopoly services.

| Service | Feature | Script | Expected Latency (p99) | Expected Throughput | Expected Concurrency |
|---|---|---:|---:|---:|---:|
| Notifier | Real-time WebSocket notifications | notifier_load_test.js | < 40 ms | 5K RPS | 1M connections |


## Prerequisites
- Install k6 (macOS):

```bash
brew install k6
```

Quick run (defaults: 50 VUs, 30s):

```bash
k6 run <script.js>
```

Refer the comments in the script for customizing VUs, duration, and other options.

Optional
- To add thresholds or more advanced scenarios, edit the `options` in
  `notifier_load_test.js` or add additional k6 metrics and checks.