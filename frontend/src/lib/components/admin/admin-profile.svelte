<script lang="ts">
	import { adminStore } from '$lib/auth/admin-store';
	import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { buttonVariants } from '../ui/button';
</script>

<DropdownMenu.Root>
	<DropdownMenu.Trigger class={buttonVariants({ variant: 'outline', size: 'sm' })}
		>{$adminStore?.firstName} {$adminStore?.lastName}</DropdownMenu.Trigger
	>
	<DropdownMenu.Content>
		<DropdownMenu.Group>
			<DropdownMenu.Label>
				<p class="text-lg">
					{$adminStore?.firstName}
					{$adminStore?.lastName}
				</p>
				<p class="text-muted-foreground text-xs font-medium">
					{$adminStore?.email}
				</p>
			</DropdownMenu.Label>
			<DropdownMenu.Separator />
			<DropdownMenu.Item>Profile</DropdownMenu.Item>
			<DropdownMenu.Item>Billing</DropdownMenu.Item>
			<DropdownMenu.Item>Team</DropdownMenu.Item>
			<DropdownMenu.Item>Subscription</DropdownMenu.Item>
		</DropdownMenu.Group>
	</DropdownMenu.Content>
</DropdownMenu.Root>
