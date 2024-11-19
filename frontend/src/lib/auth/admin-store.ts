import { persisted } from 'svelte-persisted-store';

interface Admin {
    firstName: string,
    lastName: string,
    email: string,
}

export const adminStore = persisted<Admin | undefined>("admin", undefined)