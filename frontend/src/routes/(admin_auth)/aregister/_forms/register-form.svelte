<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
	import * as Form from '$lib/components/ui/form';
	import * as Card from '$lib/components/ui/card';
	import Input from '$lib/components/ui/input/input.svelte';

	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { registerFormSchema, type RegisterFormSchema } from '$lib/validators/register';

	export let data: SuperValidated<Infer<RegisterFormSchema>>;
	const form = superForm(data, {
		validators: zodClient(registerFormSchema)
	});

	const { form: formData, enhance } = form;
</script>

<form method="POST" use:enhance>
	<Form.Field {form} name="firstName">
		<Form.Control let:attrs>
			<div>
				<Form.Label>Имя</Form.Label>
				<Input {...attrs} bind:value={$formData.firstName} />
				<Form.FieldErrors class="text-xs" />
			</div>
		</Form.Control>
	</Form.Field>
	<Form.Field {form} name="lastName">
		<Form.Control let:attrs>
			<div>
				<Form.Label>Фамилия</Form.Label>
				<Input {...attrs} bind:value={$formData.lastName} />
				<Form.FieldErrors class="text-xs" />
			</div>
		</Form.Control>
	</Form.Field>
	<Form.Field {form} name="email">
		<Form.Control let:attrs>
			<div>
				<Form.Label>Адрес электронной почты</Form.Label>
				<Input type="email" {...attrs} bind:value={$formData.email} />
				<Form.FieldErrors class="text-xs" />
			</div>
		</Form.Control>
	</Form.Field>
	<Form.Field {form} name="password">
		<Form.Control let:attrs>
			<div>
				<Form.Label>Пароль</Form.Label>
				<Input type="password" {...attrs} bind:value={$formData.password} />
				<Form.FieldErrors class="text-xs" />
			</div>
		</Form.Control>
	</Form.Field>
	<hr class="my-3" />
	<Form.Button>Создать</Form.Button>
</form>
