<script lang="ts">
	let { data } = $props();
</script>

<svelte:head>
	<title>{data.dish.name} - RMS Dashboard</title>
</svelte:head>

<h1 title="Название блюда" class="my-2 text-4xl font-medium">{data.dish.name}</h1>
