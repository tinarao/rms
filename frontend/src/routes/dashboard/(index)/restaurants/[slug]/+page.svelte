<script>
	import RenameRestaurant from '$lib/components/admin/rename-restaurant.svelte';
	import Button from '$lib/components/ui/button/button.svelte';

	let { data } = $props();
	import * as Card from '$lib/components/ui/card';
	import ChevronRight from 'lucide-svelte/icons/chevron-right';
</script>

<svelte:head>
	<title>{data.restaurant.name} - RMS Dashboard</title>
</svelte:head>

<div class="flex items-center gap-x-3">
	<h1 class="text-4xl" title="Название ресторана">{data.restaurant.name}</h1>
	<RenameRestaurant restaurant={data.restaurant} token={data.token} />
</div>
<div class="grid grid-cols-3 gap-4 py-4">
	<Card.Root>
		<Card.Header>
			<Card.Title>Продукты</Card.Title>
		</Card.Header>
		<Card.Content>
			<p>Card Content</p>
		</Card.Content>
		<Card.Footer>
			<Button><ChevronRight class="mr-2 size-4" /> Посмотреть все</Button>
		</Card.Footer>
	</Card.Root>
</div>
