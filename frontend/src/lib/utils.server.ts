import { NODE_ENV } from "$env/static/private"

export const getBackendUrl = () => {
    return NODE_ENV === "development" ? "http://localhost:3000" : ""
}