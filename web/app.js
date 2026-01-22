function bytesToGB(b) {
  return (b / 1024 / 1024 / 1024).toFixed(1) + " GB";
}

function secToTime(s) {
  const h = Math.floor(s / 3600);
  const m = Math.floor((s % 3600) / 60);
  return `${h}h ${m}m`;
}

async function load() {
  const r = await fetch("/api/status");
  const d = await r.json();

  document.getElementById("cpu").textContent = d.cpu.toFixed(1);
  document.getElementById("temp").textContent = d.temp.toFixed(1);
  document.getElementById("mem").textContent =
    `${bytesToGB(d.mem_used)} / ${bytesToGB(d.mem_total)}`;
  document.getElementById("disk").textContent =
    `${bytesToGB(d.disk_used)} / ${bytesToGB(d.disk_total)}`;
  document.getElementById("uptime").textContent = secToTime(d.uptime);
  document.getElementById("host").textContent = d.hostname;
}

load();
setInterval(load, 2000);
