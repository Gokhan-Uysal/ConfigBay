let navbarNav = "navbar-nav";
let navbarItem = "navbar-item";
let navbar = document.getElementById(navbarNav);
let navbarItems = document.getElementsByClassName(navbarItem);
if (navbar) {
    navbar.style.cssText = `--items: ${navbarItems.length}`;
}