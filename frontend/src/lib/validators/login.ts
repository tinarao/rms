import { z } from "zod";

export const loginFormSchema = z.object({
    email: z
        .string({ message: 'Адрес электронной почты не указан!' })
        .email('Некорректный адрес электронной почты!'),
    password: z
        .string({ message: 'Пароль не указан!' })
        .min(8, 'Слишком короткий пароль!')
        .max(32, 'Слишком длинный пароль!')
});

export type RegisterFormSchema = z.infer<typeof loginFormSchema>