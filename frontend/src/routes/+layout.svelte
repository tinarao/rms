<script lang="ts">
	import '../app.css';
	import { ModeWatcher } from 'mode-watcher';
	let { children } = $props();
	import { Toaster } from '$lib/components/ui/sonner';
</script>

<Toaster />
<ModeWatcher defaultMode="dark" />
{@render children()}
