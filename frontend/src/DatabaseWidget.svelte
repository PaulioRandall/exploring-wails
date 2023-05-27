<script>
	import { OpenDatabase } from '#backend'
	import DatabaseBrowserModal from './DatabaseBrowserModal.svelte'

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

	let browsingForDB = false
</script>

{#if browsingForDB}
	<DatabaseBrowserModal
		on_select="{openDatabaseFile}"
		on_close="{closeDatabaseBrowser}" />
{/if}

<div class="database-widget" on:click|stopPropagation="{openFileBrowser}">
	DB
</div>

<style>
	.database-widget {
		position: absolute;
		bottom: 0;
		right: 0;

		display: flex;
		justify-content: center;
		align-items: center;

		width: 100px;
		height: 100px;

		font-weight: bold;
		font-size: 32px;
		cursor: pointer;
		user-select: none;

		background: darkgreen;
		border-top-left-radius: 2rem;
	}
</style>
