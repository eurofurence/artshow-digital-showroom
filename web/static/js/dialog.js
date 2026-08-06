const dialog = document.getElementById("media-dialog");

const title = document.getElementById("dialog-title");
const creator = document.getElementById("dialog-creator");
const description = document.getElementById("dialog-description");
const contact = document.getElementById("dialog-contact");
const contactQr = document.getElementById("dialog-contact-qr");

const video = document.getElementById("dialog-video");
const source = document.getElementById("dialog-video-source");
const duration = document.getElementById("dialog-duration");

let currentVideo = "";

function openMediaDialog(card) {
    currentVideo = card.dataset.video;

    title.textContent = card.dataset.title;
    creator.textContent = card.dataset.creator;
    contact.textContent = card.dataset.contact;
    contactQr.src = card.dataset.contactQr;
    description.textContent = card.dataset.description;
    duration.textContent = "Full runtime: " + card.dataset.duration;

    // Stop any currently playing video
    video.pause();

    // Load the selected video
    source.src = card.dataset.preview;
    video.load();

    dialog.showModal();

    // Autoplay after loading
    video.play().catch(() => {
        // Autoplay may be blocked by the browser.
    });
}

async function playCurrentVideo() {
    await fetch("/play", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            video: currentVideo
        })
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
