<script lang="ts">
	import ModeToggle from '$lib/components/mode-toggle.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import * as Select from '$lib/components/ui/select';
	import PlusIcon from 'lucide-svelte/icons/plus';
	let { children, data } = $props();
</script>

<div class="flex h-screen flex-col">
	<header class="flex items-center gap-x-2 border-b px-8 py-2">
		{#if data.restaurants.length}
			<Select.Root>
				<Select.Trigger class="w-64">
					<Select.Value placeholder="Выбор ресторана" />
				</Select.Trigger>
				<Select.Content>
					{#each data.restaurants as restaurant}
						<Select.Item value="light">{restaurant.name}</Select.Item>
					{/each}
				</Select.Content>
			</Select.Root>
		{:else}
			<div class="flex items-center gap-x-1">
				<Button class="w-64" variant="outline" disabled>Список ресторанов пуст</Button>
				<Button
					href="/dashboard/create/restaurant"
					title="Нажмите, чтобы добавить ресторан"
					variant="ghost"
					size="icon"
				>
					<PlusIcon class="size-4" />
				</Button>
			</div>
		{/if}
		<div class="h-full w-0.5 bg-neutral-100 dark:bg-neutral-900"></div>
		<ModeToggle size="default">Тема</ModeToggle>
	</header>
	<main class="flex-1 px-8 py-4">
		{@render children()}
	</main>
</div>
