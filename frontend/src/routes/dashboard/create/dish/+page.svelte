<script lang="ts">
	import * as Card from '$lib/components/ui/card';
	import * as Select from '$lib/components/ui/select';
	import Button from '$lib/components/ui/button/button.svelte';
	import Input from '$lib/components/ui/input/input.svelte';
	import Label from '$lib/components/ui/label/label.svelte';
	import Textarea from '$lib/components/ui/textarea/textarea.svelte';
	import Salad from 'lucide-svelte/icons/salad';
	import LoaderCircle from 'lucide-svelte/icons/loader-circle';
	import axios from 'axios';
	import { toast } from 'svelte-sonner';
	import { goto } from '$app/navigation';
	import { type Restaurant } from '$lib/validators/restaurants.js';

	let { data } = $props();
	let selectedRestaurant = $state<Restaurant | undefined>(undefined);
	let isLoading = $state(false);

	const handleSubmit = async (e: SubmitEvent) => {
		isLoading = true;
		e.preventDefault();

		try {
			// @ts-ignore
			const formData = new FormData(e.currentTarget);
			const name = formData.get('name')!.toString();
			const description = formData.get('description')?.toString() ?? '';

			if (name.length < 4) {
				toast.error('Слишком короткое название!');
				return;
			}

			if (name.length > 32) {
				toast.error('Слишком длинное название!');
				return;
			}

			if (!selectedRestaurant) {
				toast.error('Ресторан не выбран!');
				return;
			}

			const res = await axios.post(
				'http://localhost:3000/api/p/dishes',
				{
					name,
					description,
					restaurantId: selectedRestaurant.id
				},
				{
					validateStatus: () => true,
					headers: {
						Authorization: `Bearer ${data.token}`
					}
				}
			);

			if (res.status !== 201) {
				toast.error(res.data.message);
				return;
			}

			toast.success('Блюдо успешно создано!');
			goto(`/dashboard/restaurants/${selectedRestaurant.slug}`);
		} finally {
			isLoading = false;
		}
	};
</script>

<svelte:head>
	<title>Создание блюда - RMS Dashboard</title>
</svelte:head>
<Card.Root class="w-[34rem]">
	<Card.Header>
		<Card.Title>Форма создания блюда</Card.Title>
		<Card.Description
			>Создайте карточку товара и внесите в неё основную информацию. Более детальные параметры можно
			будет изменить в панели управления.</Card.Description
		>
	</Card.Header>
	<form onsubmit={handleSubmit}>
		<Card.Content class="space-y-1.5">
			<div>
				<Label>Название</Label>
				<Input disabled={isLoading} name="name" required placeholder="Название блюда" />
			</div>
			<div>
				<Label>Описание</Label>
				<Textarea disabled={isLoading} name="description" maxlength={200} class="max-h-32" />
			</div>
			<div>
				<Label>Выберите ресторан, в который нужно добавить блюдо</Label>
				<Select.Root onSelectedChange={(e) => (selectedRestaurant = e!.value as Restaurant)}>
					<Select.Trigger disabled={isLoading} class="w-48">
						<Select.Value placeholder="Ресторан" />
					</Select.Trigger>
					<Select.Content>
						{#each data.restaurants as restaurant}
							<Select.Item value={restaurant}>{restaurant.name}</Select.Item>
						{/each}
					</Select.Content>
				</Select.Root>
			</div>
		</Card.Content>
		<Card.Footer class="justify-between">
			<Button type="submit" disabled={isLoading}>
				{#if isLoading}
					<LoaderCircle class="mr-2 size-4 animate-spin" />
				{:else}
					<Salad class="mr-2 size-4" />
				{/if}
				Создать</Button
			>
			<Button
				type="button"
				disabled={isLoading}
				variant="outline"
				class="hover:border-transparent hover:bg-red-500 hover:text-white">Отменить</Button
			>
		</Card.Footer>
	</form>
</Card.Root>
