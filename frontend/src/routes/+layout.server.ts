import { redirect } from "@sveltejs/kit";
import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = async ({ cookies }) => {
    const token = cookies.get('access_token')
    if (!token) {
        cookies.delete('access_token', { path: '/' })
        throw redirect(302, "/")
    }

    return { token }
}