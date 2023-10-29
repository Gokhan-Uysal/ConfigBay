import {onClick} from "./lib/document.js";
import {Get, checkResponse, urlSearchParams} from "./lib/fetch.mjs";
import {endpoints, providers} from "./config.mjs";

const googleId = "Google";
const githubId = "Github";

const signupWith = async (provider) => {
    try {
        let params = urlSearchParams({provider: provider})
        const url = endpoints.signupWith(params.toString());
        let response = await Get(url);
        checkResponse(response);
    }
    catch (err){
        console.error(err);
    }
}

const loginWith = async (provider) => {
    try {
        let params = urlSearchParams({provider: provider})
        const url = endpoints.loginWith(params.toString());
        let response = await Get(url);
        checkResponse(response);
    }
    catch (err){
        console.error(err);
    }
}

const googleBtnElement = document.getElementById(googleId);
const githubBtnElement = document.getElementById(githubId);

if (access === "login"){
    onClick(googleBtnElement, () => loginWith(providers.google))
    onClick(githubBtnElement    , () => loginWith(providers.github))
}
else if (access === "signup"){
    onClick(googleBtnElement, () => signupWith(providers.google))
    onClick(githubBtnElement    , () => signupWith(providers.github))
}
else {
    console.error("Access modifier not found in template.")
}
