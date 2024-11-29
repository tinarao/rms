import { z } from "zod";
import { dishSchema } from "./dish";

export const restaurantSchema = z.object({
    id: z.number().positive(),
    name: z.string(),
    slug: z.string(),
    description: z.string(),
    creatorId: z.number().positive(),

    creator: z.any(),
    orders: z.array(z.any()).nullable(),
    dishes: z.array(z.any()).nullable(),
    pictures: z.array(z.any()).nullable(),

    createdAt: z.string(),
    updatedAt: z.string(),
    deletedAt: z.string().nullable(),
})

export const createRestaurantFormSchema = z.object({
    name: z
        .string({ message: 'Вы не указали название ресторана!' })
        .min(2, "Слишком короткое название ресторана")
        .max(32, "Слишком длинное название канала"),
    description: z
        .string({ message: 'Описание не указано!' })
        .min(16, 'Слишком короткое описание!')
        .max(500, 'Слишком длинное описание!')
});

export type Restaurant = z.infer<typeof restaurantSchema>