# 🛡️ SynthView - Network Traffic Validation Tool

![University](https://img.shields.io/badge/University-Hertfordshire-002f6c?style=for-the-badge)
![Go](https://img.shields.io/badge/Go-1.23-00ADD8?style=for-the-badge&logo=go)
![React](https://img.shields.io/badge/React-18-61DAFB?style=for-the-badge&logo=react)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

---

## 📌 Project Overview

**SynthView** is a lightweight, high-throughput network monitoring tool that validates **synthetic traffic realism**. It gives a **fidelity score from 0 to 100** - higher score means more realistic traffic.

> **Problem:** Traditional tools like Wireshark monitor organic traffic well but cannot validate synthetic traffic used in stress testing and cybersecurity drills.

> **Solution:** SynthView fills this gap with a quantifiable fidelity scoring mechanism.

---

## 🎯 Key Features

| Feature | Description |
|---------|-------------|
| 📡 **PCAP Analysis** | Reads and processes PCAP files |
| 🔀 **Flow Reconstruction** | Groups packets using 5-tuple (SrcIP, DstIP, SrcPort, DstPort, Protocol) |
| 📊 **Feature Extraction** | Extracts 5 statistical features |
| 📈 **Baseline Modeling** | Creates separate TCP and UDP baselines |
| 🎯 **Fidelity Scoring** | Calculates realism score 0-100 |
| 🏷️ **Fingerprinting** | Labels traffic as DNS-like, HTTP-like, or Anomalous |
| 🖥️ **React Dashboard** | Real-time visualization with charts |

---

## 🧠 Algorithms Used

| # | Algorithm | Formula |
|---|-----------|---------|
| 1 | **Flow Reconstruction** | `FlowID = Hash(SrcIP + DstIP + SrcPort + DstPort + Protocol)` |
| 2 | **Feature Extraction** | `Duration = T_last - T_first` , `PacketRate = Packets / Duration` , `ByteRate = Bytes / Duration` , `IAT Variance = Σ(IAT-μ)² / n` , `Burstiness = σ / μ` |
| 3 | **Baseline Modeling** | `Baseline = (Feature₁ + ... + Featureₙ) / n` (per protocol) |
| 4 | **Fidelity Scoring** | `Distance = √[(d₁)² + (d₂)² + (d₃)²]` , `Score = 100 × e^(-Distance)` |
| 5 | **Fingerprinting** | `UDP + ≤2 packets → DNS-like` , `TCP + ≥4 packets + Burstiness≥0.5 → HTTP-like` , `Fidelity < 60 → Anomalous` |

---

## 🔄 How It Works

```
PCAP File → Packet Reader → Flow Reconstruction → Feature Extraction → Baseline Model → Fidelity Score → Fingerprinting → Metrics API → Dashboard
```

| Step | Component | Output |
|------|-----------|--------|
| 1 | PCAP Reader | Raw packets |
| 2 | Flow Table | Network flows (5-tuple) |
| 3 | Feature Extraction | Duration, Packet Rate, Byte Rate, IAT Variance, Burstiness |
| 4 | Baseline Model | TCP/UDP averages |
| 5 | Fidelity Engine | Score (0-100) |
| 6 | Fingerprinting | DNS-like, HTTP-like, Anomalous |
| 7 | Metrics API | JSON data |
| 8 | Dashboard | Visual charts |

---

## 🛠️ Technologies Used

| Technology | Purpose |
|------------|---------|
| **Go (Golang)** | Backend - high performance packet processing |
| **gopacket** | PCAP file reading and packet decoding |
| **React** | Frontend dashboard |
| **Chart.js** | Interactive charts (Pie, Bar) |
| **REST API** | Backend-Frontend communication |
| **Ubuntu Linux** | Development environment |

---

## 📁 Project Structure

```
synth-view/
├── cmd/synthview/
│   └── main.go
├── internal/
│   ├── api/server.go
│   ├── baseline/
│   ├── capture/pcap_reader.go
│   ├── features/
│   ├── fidelity/fidelity.go
│   ├── fingerprint/
│   ├── flow/
│   └── metrics/
└── ui/synthview-dashboard/
```

---

## 📦 Dependencies

### Backend (Go)

- gopacket v1.1.19

### Frontend (React)

- react-chartjs-2
- chart.js

### Install All Dependencies

```bash
# Backend dependencies (auto-downloaded)
go mod tidy

# Frontend dependencies
cd ui/synthview-dashboard
npm install
```

---

## 🚀 Installation & Setup

### Prerequisites

| Requirement | Version |
|-------------|---------|
| Ubuntu Linux | 22.04 or 24.04 |
| Go | 1.23 or higher |
| Node.js | 18.x or higher |
| npm | 9.x or higher |

---

### Step 1: Install Go

```bash
wget https://go.dev/dl/go1.23.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
go version
```

---

### Step 2: Install Node.js and npm

```bash
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
sudo apt install -y nodejs
node --version
npm --version
```

---

### Step 3: Clone the Repository

```bash
git clone https://github.com/venkateswarlupambha/synth-view.git
cd synth-view
```

---

### Step 4: Run Backend Server

```bash
go run cmd/synthview/main.go
```

**Expected Output:**

```
=== FLOW FEATURES ===
10.0.2.15 → 10.0.2.3 | UDP | dur=0.001s pkts=2 bytes=124 ...

=== BASELINE MODEL ===
TCP | mean_dur=2.7315 mean_pkt_rate=10447.71 mean_burst=2.010
UDP | mean_dur=0.0123 mean_pkt_rate=234.56 mean_burst=0.345

=== FIDELITY SCORES ===
10.0.2.15 → 10.0.2.3 | UDP | Fidelity Score = 100.00

=== FLOW FINGERPRINTING ===
10.0.2.15 → 10.0.2.3 | UDP | Fidelity=100.00 | Fingerprint=DNS-like

=== SYSTEM METRICS ===
Total Flows: 622
Average Fidelity: 22.03
Low-Fidelity Flows (<60): 457
DNS-like Flows: 120
HTTP-like Flows: 45

[✔] API server running at http://localhost:8080/metrics
```

---

### Step 5: Run Frontend Dashboard

Open a **new terminal** and run:

```bash
cd ~/synth-view/ui/synthview-dashboard
npm install
npm start
```

> ✅ Dashboard opens at: `http://localhost:3000`

---

## 📊 Dashboard Features

| Component | What it shows |
|-----------|---------------|
| **Total Flows** | Number of network conversations |
| **Average Fidelity** | Overall realism score (0-100) |
| **Low Fidelity Flows** | Flows with score below 60 |
| **Anomalous %** | Percentage of unrealistic traffic |
| **Pie Chart** | DNS-like vs HTTP-like distribution |
| **Bar Chart** | Average fidelity score visualization |
| **Alert Box** | Anomaly detection warnings |

---

## 🖼️ Dashboard Screenshots

### Overview - Dashboard
<img width="600" alt="SynthView Dashboard Overview" src="https://github.com/user-attachments/assets/18591f38-a59b-46a3-b1d3-8a9466f8c0ff" />

---

### Charts - Pie & Bar Visualization
<img width="600" alt="SynthView Charts View" src="https://github.com/user-attachments/assets/19e60e80-2d7b-48e4-b1fd-e76d9d035403" />

---

### About
<img width="600" alt="SynthView Anomaly Alerts" src="https://github.com/user-attachments/assets/ec635457-d495-42b0-ad5a-128a088dd1b7" />

---

## ⚠️ Troubleshooting

| Problem | Solution |
|---------|----------|
| `go: command not found` | Go not installed. Run Step 1 again |
| `npm: command not found` | Node.js not installed. Run Step 2 again |
| `Connection refused` | Backend not running. Run `go run cmd/synthview/main.go` |
| `Dashboard shows no data` | Check if backend is running on port 8080 |
| `PCAP file not found` | Place sample.pcap in datasets/ folder |

---

## 📈 Sample Results

| Metric | Value | Interpretation |
|--------|-------|----------------|
| Total Flows | 622 | 622 conversations analyzed |
| Average Fidelity | 22.03 | Mostly synthetic traffic |
| Low Fidelity (<60) | 457 | 73% unrealistic |
| DNS-like Flows | 120 | UDP with ≤2 packets |
| HTTP-like Flows | 45 | TCP with ≥4 packets, bursty |

---

## 🔮 Future Work

- [ ] Machine learning based baselines instead of simple averages
- [ ] Encrypted traffic analysis (TLS/SSL) using header-only methods
- [ ] Time-series database (Prometheus/InfluxDB) for historical storage
- [ ] High-throughput performance testing at enterprise scale
- [ ] Live traffic capture (instead of offline PCAP)

---

## 👨‍🎓 Student Information

| Field | Details |
|-------|---------|
| **Name** | Venkateswarlu Pambha |
| **Student ID** | 24094190 |
| **Program** | MSc Computer Science |
| **University** | University of Hertfordshire |
| **Supervisor** | Dr. Kufreh Sampson |
| **Module** | 7COM1040 - Masters Project |

---

## 📝 Project Objectives

- ✅ Develop synthetic traffic validation system using statistical comparison
- ✅ Create fidelity scoring mechanism (0-100)
- ✅ Real-time traffic analysis with immediate feedback
- ✅ Visualize results on user-friendly dashboard

---

## 🔬 Research Gap

| Problem | Solution |
|---------|----------|
| Traditional tools monitor organic traffic only | SynthView validates synthetic traffic |
| No tool measures synthetic traffic realism | Fidelity score from 0 to 100 |
| Network simulations lack validation | Protocol-specific baselines (TCP/UDP) |

---

## 📄 License

This project is submitted in partial fulfillment of the requirements for the degree of **Master of Science in Computer Science** at the **University of Hertfordshire**.

© 2026 Venkateswarlu Pambha

---

## 🙏 Acknowledgments

- Dr. Kufreh Sampson for supervision and guidance
- University of Hertfordshire, School of Physics, Engineering, and Computer Science
- Open-source community (Go, React, gopacket, Chart.js)

---

## 📞 Contact

| Platform | Link |
|----------|------|
| **GitHub** | [venkateswarlupambha](https://github.com/venkateswarlupambha) |
| **University** | University of Hertfordshire |

---

## ⭐ Show Your Support

If you find this project useful for your research or learning, please consider giving it a star on GitHub!
