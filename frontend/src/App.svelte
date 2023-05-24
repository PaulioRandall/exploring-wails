<script>
	import menuTool from '#menu/+meta.js'
	import { currTool } from '#tools/store.js'

	if (!$currTool) {
		currTool.set(menuTool)
	}
</script>

<svelte:component this="{$currTool.component}" />
