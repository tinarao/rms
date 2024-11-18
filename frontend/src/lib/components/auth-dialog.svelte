<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';
	import { buttonVariants } from './ui/button';
	import Button from './ui/button/button.svelte';
	import Input from './ui/input/input.svelte';
	import Label from './ui/label/label.svelte';

	import LoaderIcon from 'lucide-svelte/icons/loader-circle';
	import LoginIcon from 'lucide-svelte/icons/log-in';

	import { toast } from 'svelte-sonner';
	import axios from 'axios';
	import { PUBLIC_TELEGRAM_BOT_URL } from '$env/static/public';
	import { userStore } from '$lib/auth/store';

	let phoneInput = $state('');
	let isLoading = $state(false);
	const phoneNumberRegexp =
		/^\s*(?:\+?(\d{1,3}))?[-. (]*(\d{3})[-. )]*(\d{3})[-. ]*(\d{4})(?: *x(\d+))?\s*$/;

	interface ResponseType {
		message: string;
		auth_redirect: boolean;
	}

	const handleSubmit = async () => {
		isLoading = true;
		try {
			if (!phoneInput) {
				toast.error('Вы не ввели номер телефона!');
				return;
			}
			const isValidPhoneNumber = phoneNumberRegexp.test(phoneInput);
			if (!isValidPhoneNumber) {
				toast.error('Некорректный номер телефона!');
				return;
			}

			const response = await axios.post(
				'http://localhost:3000/api/auth',
				{
					phone: phoneInput
				},
				{ withCredentials: true }
			);

			if (response.status > 210) {
				toast.error('Возникла ошибка при авторизации. Попробуйте позже');
				return;
			}

			if (response.status === 201) {
				window.open(PUBLIC_TELEGRAM_BOT_URL, '_blank');
				return;
			}

			userStore.set({
				phone: response.data.phone
			});

			toast.success('Вход выполнен успешно!');
		} finally {
			isLoading = false;
		}
	};
</script>

<Dialog.Root>
	<Dialog.Trigger class={buttonVariants({ variant: 'default', size: 'sm' })}>Войти</Dialog.Trigger>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Авторизация</Dialog.Title>
			<Dialog.Description>
				Войдите, чтобы оставлять заказы и отслеживать их статус. Для авторизации необходим Telegram
				- это быстро и безопасно.
			</Dialog.Description>
		</Dialog.Header>
		<div>
			<Label>Номер телефона</Label>
			<Input bind:value={phoneInput} />
		</div>
		<Button disabled={isLoading} on:click={handleSubmit}>
			{#if isLoading}
				<LoaderIcon class="mr-2 size-4 animate-spin" /> Загрузка...
			{:else}
				<LoginIcon class="mr-2 size-4" /> Войти
			{/if}
		</Button>
	</Dialog.Content>
</Dialog.Root>
