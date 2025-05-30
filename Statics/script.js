document.addEventListener("DOMContentLoaded", () => {
    fetch("/data")
      .then(res => res.json())
      .then(data => {
        const list = document.getElementById("data-cell-id");
        data.forEach(sat => {
          const div = document.createElement("div");
          div.className = "satellite-card";
          div.innerHTML = `
            <h3>${sat.id} - ${sat.country}</h3>
            <p><strong>Launch Date:</strong> ${sat.launch_date}</p>
            <p><strong>Launcher:</strong> ${sat.launcher}</p>
          `;
          list.appendChild(div);
        });
      })
      .catch(err => {
        console.error("Failed to fetch satellite data:", err);
      });
  });
  