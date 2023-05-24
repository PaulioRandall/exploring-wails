<script>
	import Task from './Task.svelte'
	import NewTaskButton from './NewTaskButton.svelte'

	let idPool = 0
	const genId = () => {
		idPool++
		return idPool
	}

	let tasks = [
		{
			id: genId(),
			text: 'Move tasks to svelte store',
		},
		{
			id: genId(),
			text: 'Setup SQLite database to read/write tasks',
		},
	]

	const taskIdx = (task) => {
		for (let i = 0; i < tasks.length; i++) {
			if (tasks[i].id === task.id) {
				return i
			}
		}
		return -1
	}

	const newTask = () => {
		tasks.push({
			id: genId(),
		})
		tasks = tasks
	}

	const deleteTask = (task) => {
		const i = taskIdx(task)
		console.log(i)
		tasks.splice(i, 1)
		tasks = tasks
	}

	const updateTask = (task) => {
		const i = taskIdx(task)
		tasks[i] = task
		tasks = tasks
	}
</script>

<div class="task-list">
	{#each tasks as task (task.id)}
		<Task
			task="{task}"
			whenDeletePressed="{deleteTask}"
			whenEditComplete="{updateTask}" />
	{/each}
	<NewTaskButton onClick="{newTask}" />
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
