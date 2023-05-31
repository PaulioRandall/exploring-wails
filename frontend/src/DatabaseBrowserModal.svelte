<script>
	import { onMount } from 'svelte'
	import { ToAbsPath } from '#backend'

	import Button from '#lib/button/Button.svelte'
	import ButtonSpacer from '#lib/button/ButtonSpacer.svelte'
	import ButtonBar from '#lib/button/ButtonBar.svelte'
	import FileSelector from '#lib/FileSelector.svelte'

	export let on_select = (dbFile) => {}
	export let on_close = () => {}

	let dir = '.'
	let selected
	let reset
	let buttonText

	$: if (!selected) {
		buttonText = '...'
	} else if (selected.IsDir) {
		buttonText = 'Open'
	} else {
		buttonText = 'Connect'
	}

	const useSelectedFile = () => {
		if (selected.IsDir) {
			dir = selected.AbsPath
		} else {
			on_select(selected.AbsPath)
		}
	}

	onMount(async () => {
		dir = await ToAbsPath(dir)
	})
</script>

<div class="database-browser-modal">
	<div class="modal" on:click|stopPropagation={reset}>
		<div class="header">{dir}</div>
		<FileSelector open_at={dir} bind:selected bind:reset />
		<div class="spacer" />
		<ButtonBar>
			<Button disabled={!selected} on_click={useSelectedFile}
				>{buttonText}</Button>
			<ButtonSpacer />
			<Button type="cancel" on_click={on_close}>Close</Button>
		</ButtonBar>
	</div>
</div>

<style>
	.database-browser-modal {
		position: absolute;
		top: 0;
		left: 0;

		display: flex;

		width: 100vw;
		height: 100vh;

		background: rgba(0, 0, 0, 0.3);
	}

	.modal {
		flex-grow: 1;

		display: flex;
		flex-direction: column;

		margin: 4rem;

		background: lightgrey;
		border: 2px solid goldenrod;
		border-radius: 0.4rem;
	}

	.spacer {
		flex-grow: 1;
		border-bottom: 2px solid goldenrod;
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

	.divider {
		margin: 0 0 2px 0;
		padding: 0;

		border-bottom: 2px solid goldenrod;
	}
</style>
