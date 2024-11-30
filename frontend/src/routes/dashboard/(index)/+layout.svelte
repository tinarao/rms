<script lang="ts">
	import ModeToggle from '$lib/components/mode-toggle.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import PlusIcon from 'lucide-svelte/icons/plus';
	import { buttonVariants } from '$lib/components/ui/button/index.js';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import AdminProfile from '$lib/components/admin/admin-profile.svelte';

	let { children, data } = $props();
</script>

<div class="flex h-screen flex-col">
	<header class="flex items-center justify-between border-b px-8 py-2">
		{#if data.restaurants.length}
			<DropdownMenu.Root>
				<DropdownMenu.Trigger
					title="Выберите ресторан для редактирования"
					class={buttonVariants({ variant: 'outline' })}>Рестораны</DropdownMenu.Trigger
				>
				<DropdownMenu.Content>
					{#each data.restaurants as restaurant}
						<DropdownMenu.Sub>
							<DropdownMenu.SubTrigger>
								<span>{restaurant.name}</span>
							</DropdownMenu.SubTrigger>
							<DropdownMenu.SubContent>
								<DropdownMenu.Item href="/dashboard/restaurants/{restaurant.slug}?block=dishes"
									>Блюда</DropdownMenu.Item
								>
								<DropdownMenu.Item href="/dashboard/restaurants/{restaurant.slug}?block=orders"
									>Заказы</DropdownMenu.Item
								>
							</DropdownMenu.SubContent>
						</DropdownMenu.Sub>
					{/each}

					<DropdownMenu.Group>
						<DropdownMenu.Separator />
						<DropdownMenu.Item href="/dashboard">
							<PlusIcon class="mr-2 size-4" /> Управление ресторанами
						</DropdownMenu.Item>
					</DropdownMenu.Group>
				</DropdownMenu.Content>
			</DropdownMenu.Root>
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
		<div>
			<ModeToggle size="sm">Тема</ModeToggle>
			<AdminProfile />
		</div>
	</header>
	<main class="flex-1 px-8 py-4">
		{@render children()}
	</main>
</div>
