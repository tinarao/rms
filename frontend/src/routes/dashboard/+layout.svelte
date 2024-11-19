<script lang="ts">
	let { children } = $props();
</script>

<div class="flex h-screen flex-col">
	<header class="py-4"></header>
	<main class="flex-1 bg-indigo-400">
		{@render children()}
	</main>
</div>
