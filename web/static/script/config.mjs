export const endpoints = {
    newProject: "/project",
    signup: (provider) => `/signup?${provider}`,
    login: (provider) => `/login?${provider}`
}

export const providers = {
    google: "google",
    github: "github"
}