import { z } from "zod";

export const dishSchema = z.object({
    id: z.number().positive(),
    name: z.string(),
    description: z.string(),
    tags: z.array(z.string()).nullable().optional(),
    pictures: z.array(z.string()).nullable().optional(),
    purchases: z.number(),
    isPublished: z.boolean(),
    orders: z.array(z.any()).optional().nullable(),
    createdAt: z.string(),
    updatedAt: z.string(),
    deletedAt: z.string().nullable(),
})

export type Dish = z.infer<typeof dishSchema>