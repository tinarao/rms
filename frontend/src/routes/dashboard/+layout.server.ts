import { fail, redirect } from "@sveltejs/kit";
import type { LayoutServerLoad } from "./$types";
import { z } from "zod";
import axios from "axios";
import { restaurantSchema, type Restaurant } from "$lib/validators/restaurants";

export const load: LayoutServerLoad = async ({ cookies, depends }) => {
    depends('data:restaurants')
    const token = cookies.get('access_token')
    if (!token) {
        cookies.delete('access_token', { path: '/' })   // На всякий случай
        return redirect(302, "/")
    }

    const response = await axios.get("http://localhost:3000/api/restaurants/all");
    const { data: restaurants, success, error } = z.array(restaurantSchema).safeParse(response.data.restaurants);
    if (!success) {
        console.error(error.errors[0].message);
        console.log(response.data.restaurants)
        fail(500)
    }

    return { token, restaurants }
}