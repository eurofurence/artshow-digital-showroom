function openMediaDialog(card) {
    const dialog = document.getElementById("media-dialog");

    document.getElementById("dialog-image").src =
        card.dataset.image;

    document.getElementById("dialog-title").textContent =
        card.dataset.title;

    document.getElementById("dialog-creator").textContent =
        card.dataset.creator;

    document.getElementById("dialog-description").textContent =
        card.dataset.description;

    dialog.showModal();
}

// Close dialog when clicking on backdrop
document
    .getElementById("media-dialog")
    .addEventListener("click", function(event) {
        if (event.target === this) {
            this.close();
        }
    });
