<script>
	import Button from '#lib/button/Button.svelte'
	import TextInput from '#lib/TextInput.svelte'

	export let do_create = (file) => {}
	export let dir = '.'
	let name = ''

	// TODO: validate name is not already taken by calling backend
	// TODO: Disable when typing, add debounce, when stop typing check name
	//       is not empty and also not already taken
	$: nameIsValid = !!name

	const createFile = () => {
		const file = `${dir}/${name}`
		console.log('Creating new file:', file)
		do_create(file)
	}
</script>

<div class="file-creator">
	<Button disabled={!nameIsValid} on_click={createFile}>New</Button>
	<TextInput bind:value={name} />
</div>

<style>
	.file-creator {
		align-self: stretch;

		display: flex;
		align-items: center;
		column-gap: 1rem;

		padding: 0.8rem 1rem;
		background: darkgrey;
	}
</style>
