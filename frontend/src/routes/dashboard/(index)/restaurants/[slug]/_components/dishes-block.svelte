<script lang="ts">
	import Placeholder from '$lib/assets/pic-placeholder.jpeg';
	import type { Dish } from '$lib/validators/dish';
	import CreateDish from '$lib/components/forms/create-dish.svelte';

	let {
		dishes,
		restaurantId,
		token
	}: { dishes: Dish[] | null; restaurantId: number; token: string } = $props();
</script>

{#if !dishes || !dishes.length}
	<div class="flex h-72 flex-col items-center justify-center">
		<p>Ни одного блюда не добавлено для этого ресторана.</p>
		<CreateDish {restaurantId} {token}>Добавить</CreateDish>
	</div>
{:else}
	<div class="py-2">
		<p>
			Всего позиций: {dishes.length}
		</p>
	</div>
	<div class="grid grid-cols-4 gap-2">
		{#each dishes as dish}
			<a href="/dashboard/dish/{dish.slug}" class="hover:bg-accent rounded-md border transition">
				<img
					src={dish.pictures && dish.pictures.length ? dish.pictures[0] : Placeholder}
					alt={dish.name}
					class="aspect-square h-36 w-full overflow-hidden rounded-t-sm object-cover"
				/>
				<div class="rounded-b-md p-3 pb-6 pt-4">
					<h5 class="text-primary text-xl font-medium">{dish.name}</h5>
					<p class="text-sm text-neutral-700 dark:text-neutral-300">
						{dish.description}
					</p>
				</div>
			</a>
		{/each}
	</div>
{/if}
