import { redirect, type Handle } from '@sveltejs/kit';
import axios from 'axios';

export const handle: Handle = async ({ event, resolve }) => {
    if (event.url.pathname.startsWith('/dashboard')) {
        const token = event.cookies.get("access_token")
        if (!token) {
            return redirect(302, "/alogin")
        }

        const response = await axios.get('http://localhost:3000/api/averify', {
            headers: {
                "Authorization": `Bearer ${token}`
            },
            validateStatus: () => true
        })
        if (response.status !== 200) {
            event.cookies.delete("access_token", { path: '/' })
            return redirect(302, "/alogin")
        }
    }

    const response = await resolve(event);
    return response;
};