<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
	import * as Select from '$lib/components/ui/select';
	import RenameRestaurant from '$lib/components/admin/rename-restaurant.svelte';

	import List from 'lucide-svelte/icons/list';
	import Soup from 'lucide-svelte/icons/soup';

	let { data } = $props();

	const navMenuItems = [
		{
			id: 0,
			name: 'Заказы',
			icon: List,
			slug: 'orders'
		},
		{
			id: 1,
			name: 'Продукты',
			icon: Soup,
			slug: 'products'
		}
	];
</script>

<svelte:head>
	<title>Ресторан "{data.restaurant.name}" - RMS Dashboard</title>
</svelte:head>

<div class="flex gap-x-2 py-2">
	<h1 class="text-4xl" title="Название ресторана">{data.restaurant.name}</h1>
	<RenameRestaurant restaurant={data.restaurant} token={data.token} />
</div>
<div class="border-y py-2">
	<Select.Root>
		<Select.Trigger class="w-[180px]">
			<Select.Value placeholder="Блоки" />
		</Select.Trigger>
		<Select.Content>
			{#each navMenuItems as NavItem}
				<Select.Item value={NavItem.name}
					><NavItem.icon class="mr-2 size-4" /> {NavItem.name}</Select.Item
				>
			{/each}
		</Select.Content>
	</Select.Root>
</div>
