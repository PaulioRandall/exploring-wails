<script>
	import { ToAbsPath, ListFilesInDir } from '#backend'

	export function reset() {
		selectedIndex = null
		selected = null
	}

	export function refresh() {
		open_at = open_at
	}

	export let open_at = '.'
	export let selected = null
	export let onDoubleClick = () => {}

	let files = []
	let selectedIndex = null

	const select = (i) => {
		selectedIndex = i

		if (selectedIndex == null) {
			selected = null
		} else {
			selected = files[selectedIndex]
		}
	}

	const selectAndOpen = (i) => {
		select(i)
		onDoubleClick()
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

	$: ToAbsPath(open_at)
		.then((path) => (open_at = path))
		.then(ListFilesInDir)
		.then(sortByDirThenAlphaNumerically)
		.then((res) => (files = res))
		.then(reset)
		.catch((e) => {
			console.error(e)
			files = []
		})
</script>

<div class="file-selector">
	{#each files as f, i (f.name)}
		<div
			on:click|stopPropagation={select(i)}
			on:dblclick|stopPropagation={selectAndOpen(i)}
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
	.file-selector {
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
