let navbarNav: string = "navbar-nav"
let navbarItem: string = "navbar-item"
let navbar: HTMLElement | null = document.getElementById(navbarNav) as HTMLElement | null;
let navbarItems: HTMLCollectionOf<HTMLElement> = document.getElementsByClassName(navbarItem) as HTMLCollectionOf<HTMLElement>
if (navbar) {
    navbar.style.cssText = `--items: ${navbarItems.length}`
}
