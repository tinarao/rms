<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
	import * as Card from '$lib/components/ui/card';
	import Input from '$lib/components/ui/input/input.svelte';
	import Label from '$lib/components/ui/label/label.svelte';
	import Textarea from '$lib/components/ui/textarea/textarea.svelte';
	import { createRestaurantFormSchema } from '$lib/validators/restaurants';
	import axios from 'axios';
	import { toast } from 'svelte-sonner';
	import type { LayoutData } from './$types';
	import { goto } from '$app/navigation';
	import LoaderCircle from 'lucide-svelte/icons/loader-circle';

	let { data: serverLoadData }: { data: LayoutData } = $props();
	let isLoading = $state(false);

	const handleSubmit = async (event: SubmitEvent) => {
		console.log('Token', serverLoadData.token);
		isLoading = true;
		event.preventDefault();

		try {
			// @ts-ignore
			const formData = new FormData(event.currentTarget);

			const { success, error, data } = createRestaurantFormSchema.safeParse({
				name: formData.get('name'),
				description: formData.get('description')
			});
			if (!success) {
				toast.error(error.errors[0].message);
				return;
			}

			const response = await axios.post('http://localhost:3000/api/p/restaurants/create', data, {
				validateStatus: () => true,
				withCredentials: true,
				headers: {
					Authorization: `Bearer ${serverLoadData.token}`
				}
			});
			console.log(response);
			if (response.status !== 201) {
				toast.error(response.data.message);
				return;
			}

			toast.success(response.data.message);
			goto(`/dashboard/restaurants/${response.data.slug}`);
			return;
		} finally {
			isLoading = false;
		}
	};
</script>

<div class="flex h-full flex-col items-center justify-center">
	<Card.Root class="md:w-96">
		<Card.Header>
			<Card.Title>Создание ресторана</Card.Title>
			<Card.Description>Форма создания ресторана</Card.Description>
		</Card.Header>
		<form action="POST" onsubmit={(e) => handleSubmit(e)}>
			<Card.Content class="space-y-2">
				<div>
					<Label>Название</Label>
					<Input
						required
						name="name"
						disabled={isLoading}
						placeholder="Название нового ресторана"
					/>
				</div>
				<div>
					<Label>Описание</Label>
					<Textarea
						required
						name="description"
						disabled={isLoading}
						class="max-h-72"
						maxlength={500}
					/>
				</div>
			</Card.Content>
			<Card.Footer class="justify-between">
				<Button type="submit" disabled={isLoading}
					>{#if isLoading}
						<LoaderCircle class="size-4 animate-spin" />
					{:else}
						Создать
					{/if}</Button
				>
				<Button
					href="/dashboard"
					disabled={isLoading}
					variant="ghost"
					class="hover:bg-red-500 hover:text-white"
				>
					Отменить</Button
				>
			</Card.Footer>
		</form>
	</Card.Root>
</div>
