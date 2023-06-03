<script>
	import { OpenDatabase } from '#backend'
	import DatabaseBrowser from './DatabaseBrowser.svelte'

	let browsingForDB = false
	let dbFile = null

	const toggleBrowser = () => (browsingForDB = !browsingForDB)
	const closeBrowser = () => (browsingForDB = false)

	const openDatabase = (file) => {
		browsingForDB = false
		OpenDatabase(file).catch(console.error).finally(closeBrowser)
	}
</script>

{#if browsingForDB}
	<DatabaseBrowser on_select={openDatabase} on_close={closeBrowser} />
{/if}

<div
	on:click|stopPropagation={toggleBrowser}
	class="database-manager"
	class:database-open={!!dbFile}>
	DB
</div>

<style>
	.database-manager {
		position: fixed;
		bottom: 0;
		right: 0;

		display: flex;
		justify-content: center;
		align-items: center;

		width: 3rem;
		height: 3rem;
		padding-left: 0.4rem;

		font-weight: bold;
		font-size: 24px;
		cursor: pointer;
		user-select: none;

		background: indianred;
		border-top-left-radius: 1.2rem;
	}

	.database-open {
		background: darkgreen;
	}
</style>
