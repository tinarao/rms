<script>
	let { children } = $props();
</script>

<div class="grid h-screen grid-cols-2">
	<div
		class="col-span-1 hidden h-full flex-col items-center justify-center bg-neutral-100 md:flex dark:bg-neutral-900"
	>
		<p>Форма создания ресторана</p>
	</div>
	<div class="col-span-2 md:col-span-1">
		{@render children()}
	</div>
</div>
