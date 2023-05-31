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
			text: 'Convert DatabaseWidget to AppMenuBar with tabs that allow user to navigate between different apps.',
		})

		list.push({
			id: genId(),
			text: 'Replace intro menu page with splash screen + add splash screen to AppMenuBar',
		})

		list.push({
			id: genId(),
			text: 'Allow tasks to have sub-tasks.',
		})

		list.push({
			id: genId(),
			text: 'Represent tasks & their sub-tasks as horizontal tree.',
		})

		return list
	})
</script>

<div class="task-list">
	{#each $tasks as task (task.id)}
		<Task {task} />
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
