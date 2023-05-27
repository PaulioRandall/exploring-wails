<script>
	import Button from '#lib/button/Button.svelte'
	import ButtonSpacer from '#lib/button/ButtonSpacer.svelte'
	import ButtonBar from '#lib/button/ButtonBar.svelte'
	import FileSelector from '#lib/FileSelector.svelte'

	export let on_select = (dbFile) => {}
	export let on_close = () => {}

	let dbFile = ''

	const fileSelected = (file) => {
		dbFile = file
	}

	const connectToDB = () => {
		on_select(dbFile)
	}
</script>

<div class="database-browser-modal">
	<div class="window">
		<FileSelector on_select="{fileSelected}" />
		<ButtonBar>
			<Button disabled="{!dbFile}" on_click="{connectToDB}">Connect</Button>
			<ButtonSpacer />
			<Button type="cancel" on_click="{on_close}">Close</Button>
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

	.window {
		flex-grow: 1;

		display: flex;
		flex-direction: column;

		margin: 20%;

		background: grey;
		border: 2px solid goldenrod;
		border-radius: 0.4rem;
	}

	.window > :global(:first-child) {
		flex-grow: 1;
	}
</style>
