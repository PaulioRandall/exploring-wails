<script>
	import { tasks } from './store.js'
	import Node from './Node.svelte'

	export let task

	let editMode = false
	const startEdit = () => (editMode = true)

	const enterKey = 13
	const keyPressed = (e) => {
		if (editMode && e.keyCode === enterKey) {
			e.preventDefault()
			endEdit()
		}
	}

	const endEdit = () => {
		editMode = false
		updateTask()
	}

	const getTaskText = () => (task.text ? task.text : '')

	const taskIdx = () => {
		for (let i = 0; i < $tasks.length; i++) {
			if ($tasks[i].id === task.id) {
				return i
			}
		}
		return -1
	}

	const deleteTask = () =>
		tasks.update((list) => {
			const i = taskIdx(list, task)
			list.splice(i, 1)
			return list
		})

	const updateTask = () =>
		tasks.update((list) => {
			const i = taskIdx(list, task)
			list[i] = task
			return list
		})
</script>

<Node>
	<div
		id={task.id}
		class="task"
		on:click|preventDefault={startEdit}
		on:keydown={keyPressed}
		on:focusout|preventDefault={endEdit}>
		{#if editMode}
			<textarea
				autofocus
				type="text"
				class="edit-text-input"
				bind:value={task.text} />
		{:else}
			<div class="text">{getTaskText()}</div>
		{/if}
		<div class="delete-btn" on:click|preventDefault={deleteTask}>🗑</div>
	</div>
</Node>

<style>
	.task {
		width: 100%;
		height: 100%;

		display: flex;
	}

	.text {
		flex-grow: 1;

		display: flex;
		justify-content: center;
		align-items: center;

		padding: 4px;
		cursor: text;
	}

	.text:hover {
		box-shadow: inset 0 0 4px 2px goldenrod;
	}

	.edit-text-input {
		flex-grow: 1;

		resize: none;
		margin: 3px;
		text-align: center;
	}

	.delete-btn {
		margin: 3px;
		padding: 25px;

		display: flex;
		justify-content: center;
		align-items: center;

		cursor: pointer;

		box-shadow: 0 0 1px 1px gray;
	}

	.delete-btn:hover {
		box-shadow: 0 0 4px 2px red;
	}
</style>
