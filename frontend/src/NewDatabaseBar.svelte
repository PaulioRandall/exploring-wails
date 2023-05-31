<script>
	import Button from '#lib/button/Button.svelte'
	import TextInput from '#lib/TextInput.svelte'

	export let on_create = (name) => {}

	let name = ''
	const acceptName = () => on_create(name)

	// TODO: validate name is not already taken by calling backend
	// TODO: Disable when typing, add debounce, when stop typing check name
	//       is not empty and also not already taken
	$: nameIsValid = !!name
</script>

<div class="new-database-bar">
	<Button disabled={!nameIsValid} on_click={acceptName}>New database</Button>
	<TextInput bind:value={name} />
</div>

<style>
	.new-database-bar {
		align-self: stretch;

		display: flex;
		align-items: center;
		column-gap: 1rem;

		padding: 0.8rem 1rem;

		background: darkgrey;
	}
</style>
