<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';
	import Button from '$lib/components/ui/button/button.svelte';
	import Input from '$lib/components/ui/input/input.svelte';
	import Label from '$lib/components/ui/label/label.svelte';
	import Textarea from '$lib/components/ui/textarea/textarea.svelte';
	import Salad from 'lucide-svelte/icons/salad';
	import LoaderCircle from 'lucide-svelte/icons/loader-circle';
	import axios from 'axios';
	import { toast } from 'svelte-sonner';
	import { buttonVariants } from '../ui/button';

	let isLoading = $state(false);
	let { restaurantId, token, children }: { restaurantId: number; token: string; children: any } =
		$props();

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

			if (!restaurantId) {
				toast.error('Ресторан не выбран!');
				return;
			}

			const res = await axios.post(
				'http://localhost:3000/api/p/dishes',
				{
					name,
					description,
					restaurantId: restaurantId
				},
				{
					validateStatus: () => true,
					headers: {
						Authorization: `Bearer ${token}`
					}
				}
			);

			if (res.status !== 201) {
				toast.error(res.data.message);
				return;
			}

			toast.success('Блюдо успешно создано!');
		} finally {
			isLoading = false;
		}
	};
</script>

<Dialog.Root>
	<Dialog.Trigger class={buttonVariants({ variant: 'default' })}>
		{@render children()}
	</Dialog.Trigger>
	<Dialog.Content>
		<Dialog.Header>
			<Dialog.Title>Форма создания блюда</Dialog.Title>
			<Dialog.Description>
				Создайте карточку товара и внесите в неё основную информацию. Более детальные параметры
				можно будет изменить в панели управления.
			</Dialog.Description>
		</Dialog.Header>
		<form onsubmit={handleSubmit}>
			<div>
				<Label>Название</Label>
				<Input disabled={isLoading} name="name" required placeholder="Название блюда" />
			</div>
			<div>
				<Label>Описание</Label>
				<Textarea disabled={isLoading} name="description" maxlength={200} class="max-h-32" />
			</div>
			<Dialog.Footer class="justify-between pt-4">
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
			</Dialog.Footer>
		</form>
	</Dialog.Content>
</Dialog.Root>
