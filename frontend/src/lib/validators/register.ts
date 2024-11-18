import { z } from "zod";

export const registerFormSchema = z.object({
    firstName: z
        .string({ message: 'Имя не указано!' })
        .min(2, 'Слишком короткое имя!')
        .max(50, 'Слишком длинное имя!'),
    lastName: z
        .string({ message: 'Фамилия не указана!' })
        .min(2, 'Слишком короткая фамилия!')
        .max(50, 'Слишком длинная фамилия!'),
    email: z
        .string({ message: 'Адрес электронной почты не указан!' })
        .email('Некорректный адрес электронной почты!'),
    password: z
        .string({ message: 'Пароль не указан!' })
        .min(8, 'Слишком короткий пароль!')
        .max(32, 'Слишком длинный пароль!')
});

export type RegisterFormSchema = typeof registerFormSchema