async function playCurrentVideo(id) {
  await fetch("/play", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ id: id }),
  });
}

async function stopVideo() {
  await fetch("/play", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({}),
  });
}

// Close dialog when clicking on backdrop
document.addEventListener("click", function (event) {
  const dialog = document.getElementById("media-dialog");

  if (dialog && event.target === dialog) {
    dialog.close();
  }
});

// horizontal scroll on main
document.querySelector("main").addEventListener(
  "wheel",
  (e) => {
    e.preventDefault();
    e.currentTarget.scrollLeft += e.deltaY;
  },
  { passive: false },
);
