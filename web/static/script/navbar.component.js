"use strict";
var navbarNav = "navbar-nav";
var navbarItem = "navbar-item";
var navbar = document.getElementById(navbarNav);
var navbarItems = document.getElementsByClassName(navbarItem);
if (navbar) {
    navbar.style.cssText = "--items: ".concat(navbarItems.length);
}
//# sourceMappingURL=navbar.component.js.map