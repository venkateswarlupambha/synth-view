import { useEffect, useState } from "react";
import { Pie, Bar } from "react-chartjs-2";
import {
  Chart as ChartJS,
  ArcElement,
  Tooltip,
  Legend,
  CategoryScale,
  LinearScale,
  BarElement
} from "chart.js";

ChartJS.register(
  ArcElement,
  Tooltip,
  Legend,
  CategoryScale,
  LinearScale,
  BarElement
);

function Card({ title, value }) {
  return (
    <div style={{
      background: "#f4f6f8",
      padding: "20px",
      borderRadius: "10px",
      boxShadow: "0 2px 5px rgba(0,0,0,0.1)",
      textAlign: "center",
      width: "200px"
    }}>
      <h3>{title}</h3>
      <h2>{value}</h2>
    </div>
  );
}

function App() {
  const [metrics, setMetrics] = useState(null);
  const [activeTab, setActiveTab] = useState("dashboard");

  const fetchMetrics = () => {
    fetch("http://localhost:8080/metrics")
      .then(res => res.json())
      .then(data => setMetrics(data))
      .catch(err => console.error(err));
  };

  useEffect(() => {
    fetchMetrics();

    // Auto refresh every 5 seconds
    const interval = setInterval(fetchMetrics, 5000);

    return () => clearInterval(interval);

  }, []);

  if (!metrics) return <h2 style={{padding:"20px"}}>Loading metrics...</h2>;

  /* -------------------- PIE CHART -------------------- */

  const protocolData = {
    labels: ["DNS-like", "HTTP-like"],
    datasets: [
      {
        data: [
          metrics.DNSFlows,
          metrics.HTTPFlows
        ],
        backgroundColor: [
          "#4CAF50",
          "#2196F3"
        ]
      }
    ]
  };

  /* -------------------- BAR CHART -------------------- */

  const fidelityData = {
    labels: ["Average Fidelity"],
    datasets: [
      {
        label: "Fidelity Score",
        data: [metrics.AvgFidelity],
        backgroundColor: "#3b82f6"
      }
    ]
  };

  return (
    <div style={{ fontFamily: "Arial" }}>
      {/* Main Content */}
      <div style={{ padding: "30px" }}>
        <h1 style={{ marginBottom: "30px" }}>
          Network Traffic Analysis Dashboard
        </h1>

        {/* Metric Cards */}
        <div style={{
          display: "flex",
          gap: "20px",
          flexWrap: "wrap",
          marginBottom: "40px"
        }}>
          <Card title="Total Flows" value={metrics.TotalFlows} />
          <Card title="Average Fidelity" value={metrics.AvgFidelity.toFixed(2)} />
          <Card title="Low Fidelity Flows" value={metrics.LowFidelity} />
        </div>

        {/* Charts Section */}
        <div style={{
          display: "flex",
          gap: "40px",
          flexWrap: "wrap",
          marginBottom: "40px"
        }}>
          <div style={{ width: "400px" }}>
            <h2>Protocol Distribution</h2>
            <Pie data={protocolData} />
          </div>

          <div style={{ width: "400px" }}>
            <h2>Fidelity Score</h2>
            <Bar data={fidelityData} />
          </div>
        </div>

        {/* Traffic Type Cards */}
        <h2>Traffic Classification</h2>

        <div style={{
          display: "flex",
          gap: "20px",
          marginTop: "20px",
          marginBottom: "40px"
        }}>
          <Card title="DNS-like Flows" value={metrics.DNSFlows} />
          <Card title="HTTP-like Flows" value={metrics.HTTPFlows} />
        </div>

        {/* Anomaly Section */}
        <h2>Anomaly Indicators</h2>

        <div style={{
          background: "#fff3f3",
          border: "1px solid #ffcccc",
          padding: "20px",
          borderRadius: "10px",
          width: "400px"
        }}>
          <p><strong>Low Fidelity Flows:</strong> {metrics.LowFidelity}</p>
          <p>
            Flows with low fidelity may indicate synthetic or abnormal
            traffic patterns.
          </p>
        </div>
      </div>
    </div>
  );
}

export default App;
