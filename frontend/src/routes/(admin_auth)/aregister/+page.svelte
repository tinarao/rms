<script lang="ts">
	import RegisterForm from './_forms/register-form.svelte';
	import * as Card from '$lib/components/ui/card';
</script>

<!-- 
    Temporary form
    Admin' registration functionality will be moved to the Dashboard page
 -->

<Card.Root>
	<Card.Header>
		<Card.Title>Форма регистрации администратора</Card.Title>
		<Card.Description>Создайте учётную запись для нового администратора</Card.Description>
	</Card.Header>
	<Card.Content>
		<RegisterForm />
	</Card.Content>
</Card.Root>
