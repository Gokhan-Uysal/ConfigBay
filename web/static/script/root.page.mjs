import {onClick} from "./lib/document.js";
import {Get, checkResponse, urlSearchParams} from "./lib/fetch.mjs";
import {endpoints, providers} from "./config.mjs";

const googleId = "google";
const githubId = "github";

const signupWith = async (provider) => {
    try {
        let params = urlSearchParams({provider: provider})
        const url = endpoints.signup(params.toString());
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
        const url = endpoints.signup(params.toString());
        console.log(url);
        let response = await Get(url);
        checkResponse(response);
    }
    catch (err){
        console.error(err);
    }
}

const googleBtnElement = document.getElementById(googleId);
onClick(googleBtnElement, () => signupWith(providers.google))