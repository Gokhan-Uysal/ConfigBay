import {endpoints} from "./config.mjs";
import {Post} from "./lib/fetch.mjs";
import {onClick} from "./lib/document.js";

const newProjectBtnName = "new-project-btn";

const createProject = async () => {
    try {
        let response = await Post(endpoints.newProject, {});
        let data = await response.json();
        console.log(data);
    }
    catch (err) {
        console.error(err);
    }
};

const newProjectBtnElement = document.getElementById(newProjectBtnName);
onClick(newProjectBtnElement, createProject);