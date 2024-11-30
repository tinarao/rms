<script lang="ts">
	import * as Select from '$lib/components/ui/select';
	import RenameRestaurant from '$lib/components/admin/rename-restaurant.svelte';

	import List from 'lucide-svelte/icons/list';
	import Soup from 'lucide-svelte/icons/soup';
	import Cog from 'lucide-svelte/icons/cog';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import DishesBlock from './_components/dishes-block.svelte';

	let { data } = $props();
	type Blocks = 'orders' | 'dishes' | 'settings' | undefined;
	const blocksArr = ['orders', 'dishes', 'settings'];
	let selectedBlock = $state<Blocks>(undefined);

	$effect(() => {
		goto($page.url.pathname + `?block=${selectedBlock}`);
	});

	onMount(() => {
		let blockQuery = $page.url.searchParams.get('block');
		if (!selectedBlock || !blockQuery) {
			selectedBlock = 'settings';
		}

		if (blockQuery) {
			if (blocksArr.includes(blockQuery)) {
				selectedBlock = blockQuery as Blocks;
			}
		}
	});

	const navMenuItems = [
		{
			id: 0,
			name: 'Заказы',
			icon: List,
			slug: 'orders'
		},
		{
			id: 1,
			name: 'Блюда',
			icon: Soup,
			slug: 'dishes'
		},
		{
			id: 2,
			name: 'Настройки',
			icon: Cog,
			slug: 'settings'
		}
	];
</script>

<svelte:head>
	<title>Ресторан "{data.restaurant.name}" - RMS Dashboard</title>
</svelte:head>

<h1 class="my-2 text-4xl" title="Название ресторана">
	Ресторан <span class="font-medium">"{data.restaurant.name}"</span>
</h1>
<div class="border-y py-2">
	<Select.Root
		onSelectedChange={(v) => {
			if (v) {
				selectedBlock = v.value as Blocks;
			}
		}}
	>
		<Select.Trigger class="w-[180px]">
			<Select.Value placeholder="Блоки" />
		</Select.Trigger>
		<Select.Content>
			{#each navMenuItems as NavItem}
				<Select.Item value={NavItem.slug}
					><NavItem.icon class="mr-2 size-4" /> {NavItem.name}</Select.Item
				>
			{/each}
		</Select.Content>
	</Select.Root>
</div>

<div>
	{#if selectedBlock === 'dishes'}
		<DishesBlock
			token={data.token}
			restaurantId={data.restaurant.id}
			dishes={data.restaurant.dishes}
		/>
	{:else if selectedBlock === 'orders'}
		<p>Заказы</p>
	{:else}
		<p>Общие настройки</p>
	{/if}
</div>
