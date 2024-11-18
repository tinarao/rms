import { writable } from 'svelte/store';

interface Session {
    phone: string,
}

export const userStore = writable<Session>()