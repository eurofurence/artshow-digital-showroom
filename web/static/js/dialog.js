const dialog = document.getElementById("media-dialog");

const title = document.getElementById("dialog-title");
const creator = document.getElementById("dialog-creator");
const description = document.getElementById("dialog-description");

const video = document.getElementById("dialog-video");
const source = document.getElementById("dialog-video-source");

function openMediaDialog(card) {
    title.textContent = card.dataset.title;
    creator.textContent = card.dataset.creator;
    description.textContent = card.dataset.description;

    // Stop any currently playing video
    video.pause();

    // Load the selected video
    source.src = card.dataset.video;
    video.load();

    dialog.showModal();

    // Autoplay after loading
    video.play().catch(() => {
        // Autoplay may be blocked by the browser.
    });
}

dialog.addEventListener("close", () => {
    video.pause();
    video.currentTime = 0;
});

// Close dialog when clicking on backdrop
document
    .getElementById("media-dialog")
    .addEventListener("click", function(event) {
        if (event.target === this) {
            this.close();
        }
    });
