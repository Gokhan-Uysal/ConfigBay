export const onClick = (element, clickAction) => {
    checkHTMLElement(element);
    checkAction(clickAction);
    
    element.addEventListener("click", clickAction, false);
}

export const onMouseOver = (element, overAction) => {
    checkHTMLElement(element);
    checkAction(overAction);

    element.addEventListener("mouseover", overAction, false);
}

export const onMouseOut = (element, outAction) => {
    checkHTMLElement(element);
    checkAction(outAction);

    element.addEventListener("mouseout", outAction, false)
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