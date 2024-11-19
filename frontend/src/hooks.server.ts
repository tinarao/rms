import { redirect, type Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
    if (event.url.pathname.startsWith('/dashboard')) {
        console.log("Protect!~")
        const token = event.cookies.get("access_token")
        if (!token) {
            throw redirect(302, "/")
        }

        // request to /me that returns first and last name
        // compare with adminStore.firstName and adminStore.lastName
        // if false => redirect
    }

    const response = await resolve(event);
    return response;
};