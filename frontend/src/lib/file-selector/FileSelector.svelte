<script>
	import Button from '#lib/button/Button.svelte'
	import ButtonSpacer from '#lib/button/ButtonSpacer.svelte'
	import ButtonBar from '#lib/button/ButtonBar.svelte'

	import FileList from './FileList.svelte'
	import FileCreator from './FileCreator.svelte'

	export let do_create = (file) => {}
	export let on_open = (file) => {}
	export let on_close = () => {}

	let dir = '.'
	let reset, refresh

	let selected = null

	const unselect = () => {
		reset()
		selectFile(null)
	}

	const selectFile = (file) => {
		selected = file
	}

	const selectFileAndOpen = (file) => {
		selected = file
		openFileOrDir()
	}

	const openFileOrDir = () => {
		if (selected.isDir) {
			dir = selected.absPath
		} else {
			on_open(selected.absPath)
		}
	}
</script>

<div class="file-selector" on:click|stopPropagation={unselect}>
	<div class="header">{dir}</div>
	<FileList
		bind:dir
		bind:reset
		bind:refresh
		on_click={selectFile}
		on_double_click={selectFileAndOpen} />
	<div class="spacer" />
	<ButtonBar>
		<Button disabled={!selected} on_click={openFileOrDir}>Open</Button>
		<ButtonSpacer />
		<Button type="cancel" on_click={on_close}>Close</Button>
	</ButtonBar>
	<FileCreator {dir} {do_create} />
</div>

<style>
	.file-selector {
		display: flex;
		flex-direction: column;

		background: lightgrey;
		border: 2px solid goldenrod;
		border-radius: 0.3rem;
	}

	.header {
		display: flex;
		justify-content: center;
		align-items: center;

		font-size: 18px;
		color: black;

		padding: 1rem;

		background: darkgrey;
		border-bottom: 2px solid goldenrod;
	}

	.spacer {
		flex-grow: 1;
		border-bottom: 2px solid goldenrod;
	}
</style>
