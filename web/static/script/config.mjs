export const endpoints = {
    newProject: "/project",
    signupWith: (provider) => `/signup-with?${provider}`,
    loginWith: (provider) => `/login-with?${provider}`
}

export const providers = {
    google: "google",
    github: "github"
}