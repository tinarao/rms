<script lang="ts">
	import * as Card from '$lib/components/ui/card';
	import Button from '$lib/components/ui/button/button.svelte';
	import Input from '$lib/components/ui/input/input.svelte';
	import Label from '$lib/components/ui/label/label.svelte';
	import { toast } from 'svelte-sonner';
	import { loginFormSchema } from '$lib/validators/login';
	import axios from 'axios';
	import { adminStore } from '$lib/auth/admin-store';
	import { goto } from '$app/navigation';

	const handleSubmit = async (event: SubmitEvent) => {
		if (!event.currentTarget) {
			return;
		}

		// @ts-ignore
		const formData = new FormData(event.currentTarget);

		const { success, error, data } = loginFormSchema.safeParse({
			email: formData.get('email'),
			password: formData.get('password')
		});
		if (!success) {
			toast.error(error.errors[0].message);
			return;
		}

		const response = await axios.post('api/alogin', data, {
			validateStatus: () => true,
			withCredentials: true
		});
		if (response.status !== 200) {
			toast.error(response.data.message);
			return;
		}

		adminStore.set(response.data.admin);

		toast.success(response.data.message);
		goto('/dashboard');
		return;
	};
</script>

<Card.Root>
	<Card.Header>
		<Card.Title>Авторизация</Card.Title>
	</Card.Header>
	<Card.Content class="w-96">
		<form method="POST" on:submit|preventDefault={(e) => handleSubmit(e)}>
			<div>
				<Label>Адрес электронной почты</Label>
				<Input required name="email" type="email" />
			</div>
			<div>
				<Label>Пароль</Label>
				<Input required name="password" type="password" />
			</div>
			<hr class="my-3" />
			<Button type="submit">Войти</Button>
		</form>
	</Card.Content>
</Card.Root>
