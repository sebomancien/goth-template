function toggleTheme() {
    const attribute = "data-bs-theme"
    const html = document.documentElement;
    const theme = html.getAttribute(attribute)
    html.setAttribute(attribute, theme === "dark" ? "light" : "dark");
}
