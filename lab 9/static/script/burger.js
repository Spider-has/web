const burgerButtonId = "burger-menu-open-button";
const burgerMenuId = "burger-menu";
const burgerCloseButtonId = "burger-menu-close-button";

window.onload = () => {
    let burgerMenu = document.getElementById(burgerMenuId);
    let burgerOpenButton = document.getElementById(burgerButtonId);
    let closeMenuButton = document.getElementById(burgerCloseButtonId);

    if(burgerOpenButton){
        burgerOpenButton.addEventListener('click', () => {
            openBurger(burgerMenuId, burgerButtonId);
        })
    }

    if(closeMenuButton){
        closeMenuButton.addEventListener('click', () => {
            closeBurger(burgerMenuId, burgerButtonId);
        })
    }
}

const openBurger = (burgerMenuId, burgerButtonId) => {
    let burgerMenu = document.getElementById(burgerMenuId);
    let burgerOpenButton = document.getElementById(burgerButtonId);

    if(burgerMenu){
        burgerMenu.classList.add("burger-menu-navigation-list_open");
    }

    if(burgerOpenButton){
        burgerOpenButton.classList.add("burger-menu-open-icon_hidden")
    }
}

const closeBurger = (burgerMenuId, burgerButtonId) => {
    let burgerMenu = document.getElementById(burgerMenuId);
    let burgerOpenButton = document.getElementById(burgerButtonId);

    if(burgerMenu){
        burgerMenu.classList.remove("burger-menu-navigation-list_open");
    }

    if(burgerOpenButton){
        burgerOpenButton.classList.remove("burger-menu-open-icon_hidden")
    }
}