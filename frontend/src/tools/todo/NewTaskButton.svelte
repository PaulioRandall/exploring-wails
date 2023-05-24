<script>
	import Node from './Node.svelte'

	export let onClick = () => {}
</script>

<Node>
	<div class="new-task-button" on:click|preventDefault="{onClick}">＋</div>
</Node>

<style>
	.new-task-button {
		width: 100%;
		height: 100%;

		display: flex;
		justify-content: center;
		align-items: center;

		font-size: 42px;

		cursor: pointer;
	}

	.new-task-button:hover {
		box-shadow: inset 0 0 6px 2px goldenrod;
	}
</style>
