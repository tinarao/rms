import { fail, redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import axios from "axios";
import { dishSchema } from "$lib/validators/dish";

export const load: PageServerLoad = async ({ params, cookies }) => {
    if (!params.dishSlug) {
        console.error("No slug")
        return redirect(302, "/dashboard")
    }

    const token = cookies.get('access_token')
    if (!token) {
        return redirect(302, "/")
    }

    const res = await axios.get(`http://localhost:3000/api/p/dishes/${params.dishSlug}`, {
        headers: {
            Authorization: `Bearer ${token}`
        }
    })
    if (res.status !== 200) {
        return redirect(302, '/dashboard')
    }

    const { data: dish, error, success } = dishSchema.safeParse(res.data.dish)
    if (!success) {
        console.error(error)
        return redirect(302, '/dashboard')
    }

    return {
        dishSlug: params.dishSlug,
        dish
    }
}