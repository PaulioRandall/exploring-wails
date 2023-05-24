<script>
	import Node from './Node.svelte'

	export let task

	export let whenDeletePressed = (task) => {}
	export let whenEditComplete = (task) => {}

	let editMode = false

	const deleteMe = () => whenDeletePressed(task)
	const editMe = () => (editMode = true)

	const enterKey = 13
	const keyPressed = (e) => {
		if (editMode && e.keyCode === enterKey) {
			e.preventDefault()
			editComplete()
		}
	}

	const editComplete = () => {
		editMode = false
		whenEditComplete(task)
	}

	const getTaskText = () => (task.text ? task.text : '')
</script>

<Node>
	<div
		id="{task.id}"
		class="task"
		on:click|preventDefault="{editMe}"
		on:keydown="{keyPressed}"
		on:focusout|preventDefault="{editComplete}">
		{#if editMode}
			<input
				autofocus
				type="text"
				class="edit-text-input"
				bind:value="{task.text}" />
		{:else}
			<div class="text">{getTaskText()}</div>
		{/if}
		<div class="delete-btn" on:click|preventDefault="{deleteMe}">🗑</div>
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
