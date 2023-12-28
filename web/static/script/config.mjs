export const endpoints = {
    newProject: "/project",
    onboardWith: (access, provider) => {
        if (access !== "login" && access !== "signup"){
            console.error(`Unknown access modifier ${access}`)
            return
        }
        return `/${access}-with?${provider}`
    },
}