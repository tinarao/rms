<script lang="ts">
	import { setMode, mode } from 'mode-watcher';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { buttonVariants } from './ui/button';

	let { children, size = 'sm' }: { children: any; size?: 'sm' | 'lg' | 'default' } = $props();
</script>

<DropdownMenu.Root>
	<DropdownMenu.Trigger class={buttonVariants({ variant: 'outline', size })}>
		{@render children()}
	</DropdownMenu.Trigger>
	<DropdownMenu.Content>
		<DropdownMenu.Group>
			<DropdownMenu.Item on:click={() => setMode('light')}>Светлая</DropdownMenu.Item>
			<DropdownMenu.Item on:click={() => setMode('dark')}>Тёмная</DropdownMenu.Item>
		</DropdownMenu.Group>
	</DropdownMenu.Content>
</DropdownMenu.Root>
