import axios from "axios";
import type { LayoutServerLoad } from "./$types";
import { z } from "zod";
import { restaurantSchema } from "$lib/validators/restaurants";

export const load: LayoutServerLoad = async () => {
    const response = await axios.get("http://localhost:3000/api/restaurants/all")
    const { data: restaurants, success } = z.array(restaurantSchema).safeParse(response.data.restaurants)
    if (!success) {
        return {
            restaurants: []
        }
    }

    return {
        restaurants
    }
}