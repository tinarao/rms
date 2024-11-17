<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';
	import { buttonVariants } from './ui/button';
	import Button from './ui/button/button.svelte';
	import Input from './ui/input/input.svelte';
	import Label from './ui/label/label.svelte';

	import LoaderIcon from 'lucide-svelte/icons/loader-circle';
	import LoginIcon from 'lucide-svelte/icons/log-in';

	import { toast } from 'svelte-sonner';

	let phoneInput = $state('');
	let isLoading = $state(false);
	const phoneNumberRegexp =
		/^\s*(?:\+?(\d{1,3}))?[-. (]*(\d{3})[-. )]*(\d{3})[-. ]*(\d{4})(?: *x(\d+))?\s*$/;

	const handleSubmit = async () => {
		isLoading = true;
		try {
			console.log('try block');
			const isValidPhoneNumber = phoneNumberRegexp.test(phoneInput);
			if (!isValidPhoneNumber) {
				toast.error('Некорректный номер телефона!');
				return;
			}

			console.log('validated phone number');

			const response = await fetch('http://localhost:3000/api/auth', {
				method: 'POST',
				body: JSON.stringify({
					phone: phoneInput
				})
			});

			console.log(response);
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
