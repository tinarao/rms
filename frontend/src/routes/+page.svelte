<script>
	import AuthDialog from '$lib/components/auth-dialog.svelte';
	import ModeToggle from '$lib/components/mode-toggle.svelte';
</script>

<header class="container flex justify-end gap-x-2 border-b py-2">
	<ModeToggle>Тема</ModeToggle>
	<AuthDialog />
</header>
<main class="container">
	<div>featured</div>
</main>
