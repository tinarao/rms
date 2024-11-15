<script lang="ts">
	import { setMode } from 'mode-watcher';
</script>

<button on:click={() => setMode('light')}>светлая</button>
<button on:click={() => setMode('dark')}>тёмная</button>
