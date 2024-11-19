import { z } from "zod";

export const restaurantSchema = z.object({
    id: z.number().positive(),
    name: z.string(),
    description: z.string(),
    orders: z.any(),
    pictures: z.any(),
    creator: z.any(),
    creatorId: z.number().positive()
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

export type RegisterFormSchema = z.infer<typeof createRestaurantFormSchema>