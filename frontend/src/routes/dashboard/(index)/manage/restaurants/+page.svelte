<script lang="ts">
	import type { PageData } from './$types';
	import Plus from 'lucide-svelte/icons/plus';

	let { data }: { data: PageData } = $props();
	import Button from '$lib/components/ui/button/button.svelte';
</script>

<div class="flex items-center gap-x-2">
	<Button size="sm" href="/dashboard/create/restaurant"
		><Plus class="mr-2 size-4" /> Добавить</Button
	>
	<span class="text-sm font-medium">
		Всего: {data.restaurants.length}
	</span>
</div>
<div class="grid grid-cols-4 gap-2 py-4">
	{#each data.restaurants as restaurant}
		<a href={`/dashboard/restaurants/${restaurant.slug}`} class="group rounded-md">
			<img
				class="h-20 w-full rounded-t-lg object-cover"
				src="https://img.freepik.com/free-vector/abstract-organic-pattern-design-background_1048-19286.jpg"
				alt={`Ресторан "${restaurant.name}"`}
			/>
			<div
				class="rounded-b-lg bg-neutral-50 p-3 transition group-hover:bg-neutral-100 dark:bg-neutral-900 group-hover:dark:bg-neutral-800"
			>
				<h3 class="text-primary text-2xl font-bold">
					{restaurant.name}
				</h3>
				<p class="line-clamp-1 overflow-hidden text-ellipsis">
					{restaurant.description}
				</p>
			</div>
		</a>
	{/each}
</div>
