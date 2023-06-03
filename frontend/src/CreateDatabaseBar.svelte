<script>
	import { CreateDatabase } from '#backend'

	import Button from '#lib/button/Button.svelte'
	import TextInput from '#lib/TextInput.svelte'

	export let on_create = () => {}
	export let dir = '.'
	let name = ''

	// TODO: validate name is not already taken by calling backend
	// TODO: Disable when typing, add debounce, when stop typing check name
	//       is not empty and also not already taken
	$: nameIsValid = !!name

	const createDatabase = () => {
		const file = `${dir}/${name}`
		console.log('Creating new database:', file)

		CreateDatabase(file).then(on_create).catch(console.error)
	}
</script>

<div class="create-database-bar">
	<Button disabled={!nameIsValid} on_click={createDatabase}
		>Create database</Button>
	<TextInput bind:value={name} />
</div>

<style>
	.create-database-bar {
		align-self: stretch;

		display: flex;
		align-items: center;
		column-gap: 1rem;

		padding: 0.8rem 1rem;

		background: darkgrey;
	}
</style>
