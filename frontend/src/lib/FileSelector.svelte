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

	$: ToAbsPath(open_at)
		.then((path) => {
			open_at = path
		})
		.then(() => {
			return ListFilesInDir(open_at)
		})
		.then((res) => {
			files = res
		})
		.catch((e) => {
			console.error(e)
			files = []
		})
</script>

<div class="file-selector">
	{#each files as f, i (f.Name)}
		<div
			on:click|stopPropagation={select(i)}
			class="file"
			class:selected-file={selectedIndex === i}>
			{f.Name}
			{#if f.IsDir}
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
