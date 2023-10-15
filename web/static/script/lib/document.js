export const onClick = (element, clickAction) => {
    checkHTMLElement(element);
    checkAction(clickAction);


    element.addEventListener("click", clickAction);
}

const checkHTMLElement = (element) => {
    if (element instanceof HTMLElement) {
        return;
    }
    throw new Error("Not a HTML element.")
}

const checkAction = (action) => {
    if (typeof action === "function") {
        return;
    }
    throw new Error(`Invalid action type ${typeof action}.`)
}