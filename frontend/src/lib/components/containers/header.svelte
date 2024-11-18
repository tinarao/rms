<script>
	import { userStore } from '$lib/auth/store';
	import AuthDialog from '../auth-dialog.svelte';
	import ModeToggle from '../mode-toggle.svelte';
	import ProfileDropdown from '../profile-dropdown.svelte';
</script>

<header class="flex justify-end gap-x-2 border-b px-8 py-2">
	<ModeToggle>Тема</ModeToggle>
	{#if $userStore}
		<ProfileDropdown />
	{:else}
		<AuthDialog />
	{/if}
</header>
