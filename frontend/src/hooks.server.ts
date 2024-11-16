import type { Handle } from "@sveltejs/kit";

export const handle: Handle = async ({ event, resolve }) => {
    if (event.url.pathname === "/dashboard") {
        console.log("protect~!")
    }

    return await resolve(event);
}