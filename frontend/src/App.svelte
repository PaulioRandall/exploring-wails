<script>
	import DatabaseManager from './DatabaseManager.svelte'
	import Messenger from '#lib/messenger/Messenger.svelte'
	import menuTool from '#menu/+meta.js'
	import { currTool } from '#tools/store.js'

	if (!$currTool) {
		currTool.set(menuTool)
	}
</script>

<div class="app">
	<svelte:component this={$currTool.component} />
	<DatabaseManager />
	<Messenger />
</div>

<style>
	.app {
		position: relative;

		width: 100vw;
		height: 100vh;

		overflow: hidden;
	}
</style>
