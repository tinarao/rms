
import { registerFormSchema } from "$lib/validators/register.js";
import { fail, type Actions } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types.js";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";

export const load: PageServerLoad = async () => {
    return {
        form: await superValidate(zod(registerFormSchema)),
    };
};

export const actions: Actions = {
    default: async (event) => {
        const form = await superValidate(event, zod(registerFormSchema));
        if (!form.valid) {
            return fail(400, {
                form,
            });
        }

        // logic here

        return {
            form,
        };
    },
};