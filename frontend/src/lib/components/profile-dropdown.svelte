<script lang="ts">
	import { userStore } from '$lib/auth/store';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { buttonVariants } from './ui/button';
</script>

<DropdownMenu.Root>
	<DropdownMenu.Trigger class={buttonVariants({ variant: 'default', size: 'sm' })}
		>{$userStore!.phone}</DropdownMenu.Trigger
	>
	<DropdownMenu.Content>
		<DropdownMenu.Group>
			<DropdownMenu.Item>Мои заказы</DropdownMenu.Item>
		</DropdownMenu.Group>
	</DropdownMenu.Content>
</DropdownMenu.Root>
