import { persisted } from 'svelte-persisted-store';

interface Session {
    phone: string,
}

export const userStore = persisted<Session | undefined>("session", undefined)