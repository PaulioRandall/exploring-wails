<script>
	import { tasks, genId } from './store.js'

	import Task from './Task.svelte'
	import NewTaskButton from './NewTaskButton.svelte'

	tasks.update((list) => {
		list.push({
			id: genId(),
			text: 'Setup SQLite database to read/write tasks',
		})

		list.push({
			id: genId(),
			text: 'Add top menu bar componnet with back to main menu button',
		})

		return list
	})
</script>

<div class="task-list">
	{#each $tasks as task (task.id)}
		<Task task="{task}" />
	{/each}
	<NewTaskButton />
</div>

<style>
	.task-list {
		display: flex;
		flex-direction: column;
		row-gap: 1rem;

		overflow: scroll;
		max-height: 100%;
	}
</style>
