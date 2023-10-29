import {Get, checkResponse, urlSearchParams} from "./lib/fetch.mjs";
import {endpoints} from "./config.mjs";
import {onClick} from "./lib/document.js";

export const onboardWith = async (access, provider) => {
    try {
        let params = urlSearchParams({provider: provider})
        const url = endpoints.onboardWith(access, params.toString());
        let response = await Get(url);
        checkResponse(response);
    }
    catch (err){
        console.error(err);
    }
}

const initializeOnboardItems = (access) => {
    document.querySelectorAll('.onboard-item').forEach(item => {
        const provider = item.getAttribute('id');
        onClick(item, () => onboardWith(access, provider))
    });
}

initializeOnboardItems(access)

