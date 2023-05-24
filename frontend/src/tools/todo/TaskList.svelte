<script>
	import { tasks } from './store.js'

	import Task from './Task.svelte'
	import NewTaskButton from './NewTaskButton.svelte'

	let idPool = 0
	const genId = () => {
		idPool++
		return idPool
	}

	tasks.update((list) => {
		list.push({
			id: genId(),
			text: 'Setup SQLite database to read/write tasks',
		})

		list.push({
			id: genId(),
			text: 'Tasks should update themselves',
		})

		return list
	})

	const taskIdx = (list, task) => {
		for (let i = 0; i < list.length; i++) {
			if (list[i].id === task.id) {
				return i
			}
		}
		return -1
	}

	const newTask = () =>
		tasks.update((list) => {
			list.push({
				id: genId(),
			})
			return list
		})

	const deleteTask = (task) =>
		tasks.update((list) => {
			const i = taskIdx(list, task)
			list.splice(i, 1)
			return list
		})

	const updateTask = (task) =>
		tasks.update((list) => {
			const i = taskIdx(list, task)
			list[i] = task
			return list
		})
</script>

<div class="task-list">
	{#each $tasks as task (task.id)}
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
