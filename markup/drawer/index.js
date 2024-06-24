const drawerToggleButton = document.getElementById("drawer-toggle-button");

const content = document.getElementById("content");

drawerToggleButton.addEventListener("click", () => {
    toggleDrawer();
});

function toggleDrawer() {
    content.classList.toggle("drawer-open");

    if (content.classList.contains("drawer-open")) {
        drawerToggleButton.innerHTML = "詳細を閉じる";
    } else {
        drawerToggleButton.innerHTML = "詳細を開く";
    }
}
