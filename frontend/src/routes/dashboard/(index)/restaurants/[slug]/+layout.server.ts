import axios from "axios";
import type { LayoutServerLoad } from "./$types";
import { redirect } from "@sveltejs/kit";
import { restaurantSchema } from "$lib/validators/restaurants";
import { dishSchema, type Dish } from "$lib/validators/dish";
import { z } from "zod";

export const load: LayoutServerLoad = async ({ params }) => {
    if (!params.slug) {
        console.error("No slug")
        return redirect(302, "/dashboard")
    }

    const res = await axios.get(`http://localhost:3000/api/restaurants/slug/${params.slug}`)

    if (res.status > 299) {
        console.error("Bad status")
        return redirect(302, "/dashboard")
    }

    const { success, data: restaurant, error } = restaurantSchema.safeParse(res.data.restaurant)
    if (!success) {
        console.error("Validation error", error.errors[0])
        return redirect(302, "/dashboard")
    }

    let dishesResult = z.array(dishSchema).safeParse(restaurant.dishes)
    if (!dishesResult.success) {
        console.error(dishesResult.error.errors[0])
        // blyat what to do here
        return redirect(302, "/dashboard")
    }

    return {
        restaurantSlug: params.slug,
        restaurant: restaurant,
        dishes: dishesResult.data
    }
}