document.addEventListener("DOMContentLoaded", function () {
    var isDarkMode = document.body.classList.contains("dark-mode");

    var sunIcon = document.querySelector("#darkModeToggle .fa-sun");
    var moonIcon = document.querySelector("#darkModeToggle .fa-moon");

    if (isDarkMode) {
        sunIcon.style.display = "none"; 
        moonIcon.style.display = "inline";
    } else {
        sunIcon.style.display = "inline";
        moonIcon.style.display = "none"; 
    }
});

function myFunction() {
    var element = document.body;
    element.classList.toggle("dark-mode");

    var sunIcon = document.querySelector("#darkModeToggle .fa-sun");
    var moonIcon = document.querySelector("#darkModeToggle .fa-moon");

    var isDarkMode = element.classList.contains("dark-mode");

    if (isDarkMode) {
        sunIcon.style.display = "none";
        moonIcon.style.display = "inline"; 
        sunIcon.style.color = "#fff"; 
        moonIcon.style.color = "#fff";
    } else {
        sunIcon.style.display = "inline"; 
        moonIcon.style.display = "none";
        sunIcon.style.color = "#000"; 
        moonIcon.style.color = "#000";
    }
}
