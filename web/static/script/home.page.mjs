import {endpoints} from "./config.mjs";
import {checkResponse, Post} from "./lib/fetch.mjs";
import {onClick} from "./lib/document.js";

const newProjectBtnId = "new-project-btn";

const createProject = async () => {
    try {
        let response = await Post(endpoints.newProject, {});
        checkResponse(response);
    }
    catch (err) {
        console.error(err);
    }
};

const newProjectBtnElement = document.getElementById(newProjectBtnId);
onClick(newProjectBtnElement, createProject);