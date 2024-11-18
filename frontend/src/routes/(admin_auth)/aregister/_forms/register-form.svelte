<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
	import Input from '$lib/components/ui/input/input.svelte';
	import Label from '$lib/components/ui/label/label.svelte';
	import { registerFormSchema } from '$lib/validators/register';
	import axios from 'axios';
	import { toast } from 'svelte-sonner';

	const handleSubmit = async (event: SubmitEvent) => {
		if (!event.currentTarget) {
			return;
		}

		// @ts-ignore
		const formData = new FormData(event.currentTarget);

		const { success, error, data } = registerFormSchema.safeParse({
			firstName: formData.get('firstName'),
			lastName: formData.get('lastName'),
			email: formData.get('email'),
			password: formData.get('password')
		});
		if (!success) {
			toast.error(error.errors[0].message);
			return;
		}

		const response = await axios.post('api/aregister', data, {
			validateStatus: () => true
		});
		if (response.status !== 201) {
			toast.error(response.data.message);
			return;
		}

		toast.success(response.data.message);
		return;
	};
</script>

<form method="POST" on:submit|preventDefault={(e) => handleSubmit(e)}>
	<div>
		<Label>Имя</Label>
		<Input required name="firstName" />
	</div>
	<div>
		<Label>Фамилия</Label>
		<Input required name="lastName" />
	</div>
	<div>
		<Label>Адрес электронной почты</Label>
		<Input required name="email" type="email" />
	</div>
	<div>
		<Label>Пароль</Label>
		<Input required name="password" type="password" />
	</div>
	<hr class="my-3" />
	<Button type="submit">Создать</Button>
</form>
