<script>
	import { ToAbsPath, ListFilesInDir } from '#backend'
	import { message, newErrorMsg, newSuccessMsg } from '#lib/messenger/store.js'

	export const reset = () => (selectedIndex = null)
	export const refresh = () => (dir = dir)

	export let dir = '.'
	export let on_click = (selected) => {}
	export let on_double_click = (selected) => {}

	let files = []
	let selectedIndex = null

	const onClick = (i) => () => {
		selectedIndex = i
		on_click(files[i])
	}

	const onDoubleClick = (i) => () => {
		selectedIndex = i
		on_double_click(files[i])
	}

	const sortByDirThenAlphaNumerically = (files) => {
		files.sort((a, b) => {
			const A_FIRST = -1
			const B_FIRST = 1

			if (a.name === '..') return A_FIRST
			if (b.name === '..') return B_FIRST
			if (a.isDir && !b.isDir) return A_FIRST
			if (!a.isDir && b.isDir) return B_FIRST
			return a.name < b.name ? A_FIRST : B_FIRST
		})

		return files
	}

	$: ToAbsPath(dir)
		.then((path) => (dir = path))
		.then(ListFilesInDir)
		.then(sortByDirThenAlphaNumerically)
		.then((res) => (files = res))
		.then(reset)
		.catch((e) => {
			console.error(e)
			message.set(newErrorMsg(e))
			files = []
		})
</script>

<div class="file-list">
	{#each files as f, i (f.name)}
		<div
			on:click|stopPropagation={onClick(i)}
			on:dblclick|stopPropagation={onDoubleClick(i)}
			class="file"
			class:selected-file={selectedIndex === i}>
			{f.name}
			{#if f.isDir}
				<span>📁</span>
			{/if}
		</div>
	{/each}
</div>

<style>
	.file-list {
		overflow-x: hidden;
		overflow-y: scroll;
	}

	.file {
		display: flex;
		justify-content: space-between;

		color: black;
		cursor: pointer;
		user-select: none;
		text-align: left;

		padding: 1rem;
		border-bottom: 1px dotted black;
	}

	.file:last-child {
		border-bottom: none;
	}

	.selected-file {
		color: white;
		font-weight: bold;
		background: forestgreen;
	}
</style>
