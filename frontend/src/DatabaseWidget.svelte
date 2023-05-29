<script>
	import { OpenDatabase } from '#backend'
	import DatabaseBrowserModal from './DatabaseBrowserModal.svelte'

	let browsingForDB = false
	let dbFile = null

	const openFileBrowser = (e) => {
		browsingForDB = !browsingForDB
	}

	const openDatabaseFile = (file) => {
		browsingForDB = false
		OpenDatabase(file)
			.catch((e) => {
				console.error(e)
			})
			.finally(() => {
				browsingForDB = false
			})
	}

	const closeDatabaseBrowser = () => {
		browsingForDB = false
	}
</script>

{#if browsingForDB}
	<DatabaseBrowserModal
		on_select="{openDatabaseFile}"
		on_close="{closeDatabaseBrowser}" />
{/if}

<div
	on:click|stopPropagation="{openFileBrowser}"
	class="database-widget"
	class:db-open="{!!dbFile}">
	DB
</div>

<style>
	.database-widget {
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

	.db-open {
		background: darkgreen;
	}
</style>
