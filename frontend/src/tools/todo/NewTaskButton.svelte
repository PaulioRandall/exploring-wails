<script>
	import { tasks, genId } from './store.js'
	import Node from './Node.svelte'

	const addEmptyTask = () =>
		tasks.update((list) => {
			list.push({
				id: genId(),
			})
			return list
		})
</script>

<Node>
	<div class="new-task-button" on:click|preventDefault="{addEmptyTask}">＋</div>
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
