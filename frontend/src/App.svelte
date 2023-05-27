<script>
	import DatabaseWidget from './DatabaseWidget.svelte'
	import menuTool from '#menu/+meta.js'
	import { currTool } from '#tools/store.js'

	if (!$currTool) {
		currTool.set(menuTool)
	}
</script>

<div class="app">
	<svelte:component this="{$currTool.component}" />
	<DatabaseWidget />
</div>

<style>
	.app {
		width: 100vw;
		height: 100vh;

		position: relative;
	}
</style>
