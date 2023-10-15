const requestOptions = (method, body) => {
    if (typeof method !== "string") {
        throw new Error(`Invalid method ${method}.`)
    }
    return {
        method: method,
        mode: "cors",
        cache: "no-cache",
        credentials: "same-origin",
        headers: {
            "Content-Type": "application/json",
        },
        body: body ? JSON.stringify(body) : ""
    }
}

export const Post = (url, body) => {
    checkUrl(url);
    checkBody(body);

    return fetch(url, requestOptions("POST", body));
}

export const Get = (url) => {
    checkUrl(url);

    return fetch(url, requestOptions("GET"))
}

const checkUrl = (url) => {
    if (typeof url === "string") {
        return;
    }
    throw new Error("Invalid url.")
}

const checkBody = (body) => {
    if (typeof body === "object") {
        return;
    }
    throw new Error(`Invalid body type ${typeof body}.`)
}