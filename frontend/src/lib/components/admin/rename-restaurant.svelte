<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';
	import type { Restaurant } from '$lib/validators/restaurants';
	import { buttonVariants } from '../ui/button';
	import PenLine from 'lucide-svelte/icons/pen-line';
	import Label from '../ui/label/label.svelte';
	import Input from '../ui/input/input.svelte';
	import Button from '../ui/button/button.svelte';
	import { cn } from '$lib/utils';
	import axios from 'axios';
	import { toast } from 'svelte-sonner';
	import { goto, invalidate } from '$app/navigation';

	let { restaurant, token }: { restaurant: Restaurant; token: string } = $props();
	let newName = $state(restaurant.name);
	let isLoading = $state(false);

	const handleSaveChanges = async () => {
		isLoading = true;
		try {
			const res = await axios.patch(
				'http://localhost:3000/api/p/restaurants/edit',
				{
					name: newName,
					description: restaurant.description,
					id: restaurant.id
				},
				{
					validateStatus: () => true,
					headers: {
						Authorization: `Bearer ${token}`
					}
				}
			);

			if (res.status !== 200) {
				toast.error(res.data.message);
				return;
			}

			toast.success(res.data.message);
			invalidate('data:restaurants');
		} finally {
			isLoading = false;
		}

		goto(`/dashboard`);
	};
</script>

<Dialog.Root>
	<Dialog.Trigger class={buttonVariants({ variant: 'ghost', size: 'icon' })}>
		<PenLine class="size-4" /></Dialog.Trigger
	>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Переименовать ресторан</Dialog.Title>
		</Dialog.Header>
		<div>
			<Label>Новое название</Label>
			<Input disabled={isLoading} bind:value={newName} />
		</div>
		<Dialog.Footer>
			<Button
				disabled={isLoading}
				onclick={handleSaveChanges}
				class="hover:bg-primary hover:border-transparent hover:text-white"
				variant="outline">Изменить</Button
			>
			<Dialog.Close
				class={cn(
					buttonVariants({ variant: 'outline' }),
					'hover:border-transparent hover:bg-red-500 hover:text-white'
				)}>Отмена</Dialog.Close
			>
		</Dialog.Footer>
	</Dialog.Content>
</Dialog.Root>
