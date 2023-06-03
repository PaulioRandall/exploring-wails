<script>
	import Button from '#lib/button/Button.svelte'
	import ButtonSpacer from '#lib/button/ButtonSpacer.svelte'
	import ButtonBar from '#lib/button/ButtonBar.svelte'
	import FileSelector from '#lib/FileSelector.svelte'

	import CreateDatabaseBar from './CreateDatabaseBar.svelte'

	export let on_select = (dbFile) => {}
	export let on_close = () => {}

	let dir = '.'
	let reset, refresh

	let selected
	let buttonText

	const useSelectedFile = () => {
		if (selected.isDir) {
			dir = selected.absPath
		} else {
			on_select(selected.absPath)
		}
	}

	const onDatabaseCreated = (file) => {
		on_select(selected.absPath)
	}

	$: if (!selected) {
		buttonText = '...'
	} else if (selected.isDir) {
		buttonText = 'Open'
	} else {
		buttonText = 'Connect'
	}
</script>

<div class="database-browser" on:click|stopPropagation={reset}>
	<div class="vertical-container">
		<div class="header">{dir}</div>
		<FileSelector
			onDoubleClick={useSelectedFile}
			bind:open_at={dir}
			bind:selected
			bind:reset
			bind:refresh />
		<div class="spacer" />
		<ButtonBar>
			<Button disabled={!selected} on_click={useSelectedFile}>
				{buttonText}
			</Button>
			<ButtonSpacer />
			<Button type="cancel" on_click={on_close}>Close</Button>
		</ButtonBar>
		<CreateDatabaseBar {dir} on_create={refresh} />
	</div>
</div>

<style>
	.database-browser {
		position: absolute;
		top: 0;
		left: 0;

		width: 100%;
		height: 100%;

		background: rgba(0, 0, 0, 0.3);
	}

	.vertical-container {
		display: flex;
		flex-direction: column;

		--vertical-container-margin: 4rem;
		height: calc(100% - var(--vertical-container-margin) * 2);
		margin: var(--vertical-container-margin);

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

		height: 4rem;
		background: darkgrey;

		border-bottom: 2px solid goldenrod;
	}

	.spacer {
		flex-grow: 1;

		border-bottom: 2px solid goldenrod;
	}
</style>
