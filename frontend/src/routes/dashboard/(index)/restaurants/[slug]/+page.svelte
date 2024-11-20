<script>
	import RenameRestaurant from '$lib/components/admin/rename-restaurant.svelte';

	let { data } = $props();
</script>

<svelte:head>
	<title>{data.restaurant.name} - RMS Dashboard</title>
</svelte:head>

<div class="flex items-center gap-x-3">
	<h1 class="text-4xl" title="Название ресторана">{data.restaurant.name}</h1>
	<RenameRestaurant restaurant={data.restaurant} token={data.token} />
</div>
