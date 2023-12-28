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
        body: body ? JSON.stringify(body) : undefined
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

export const checkResponse = (response) => {
    if (!(response instanceof Response)) {
        throw new Error(`Unknown response type ${typeof response}`);
    }

    const statusCode = response.status;

    if (statusCode >= 100 && statusCode < 200) {
        console.debug(`Information response with [${statusCode}] from [${response.url}]`);
        return;
    }

    if (statusCode >= 200 && statusCode < 300) {
        console.debug(`Successful response with [${statusCode}] from [${response.url}]`)
        return;
    }

    if (statusCode >= 300 && statusCode < 400) {
        console.debug(`Redirection message response with [${statusCode}] from [${response.url}]`)
        return;
    }

    if (statusCode >= 400 && statusCode < 500) {
        throw new Error(`Client error response with [${statusCode}] from [${response.url}]`);
    }

    if (statusCode >= 500) {
        throw new Error(`Server error response with [${statusCode}] from [${response.url}]`);
    }
}

export const urlSearchParams = (params) => {
    if (typeof params !== "object"){
        throw new Error(`Unknown parameter type ${typeof params}`)
    }
    return new URLSearchParams(params);
}

const checkUrl = (url) => {
    if (typeof url === "string") {
        return;
    }
    if (url.startsWith("/")) {
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