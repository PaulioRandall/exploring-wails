<script>
	import { CreateDatabase, OpenDatabase } from '#backend'
	import { message, newErrorMsg, newSuccessMsg } from '#lib/messenger/store.js'
	import Modal from '#lib/Modal.svelte'
	import FileSelector from '#lib/file-selector/FileSelector.svelte'
	import DatabaseControls from './DatabaseControls.svelte'

	let browsingForDB = false
	let managingDB = false // TODO
	let dbFile = null

	const toggleManager = () => {
		browsingForDB = !browsingForDB // false
		//managingDB = !managingDB // TODO
	}

	const closeManager = () => {
		browsingForDB = false
		managingDB = false
	}

	const createDatabase = (file) => {
		CreateDatabase(file)
			.then(() => {
				dbFile = file
				message.set(newSuccessMsg(`Database created: ${dbFile}`))
			})
			.catch((e) => {
				console.error(e)
				message.set(newErrorMsg(e))
			})
			.finally(closeManager)
	}

	const openDatabase = (file) => {
		browsingForDB = false
		OpenDatabase(file)
			.then(() => {
				dbFile = file
				message.set(newSuccessMsg(`Database opened: ${dbFile}`))
			})
			.catch((e) => {
				console.error(e)
				message.set(newErrorMsg(e))
			})
			.finally(closeManager)
	}
</script>

{#if managingDB}
	<Modal>
		<!--
	
		-->
		<DatabaseControls />
	</Modal>
{/if}

{#if browsingForDB}
	<Modal>
		<FileSelector
			do_create={createDatabase}
			on_open={openDatabase}
			on_close={closeManager} />
	</Modal>
{/if}

<div
	on:click|stopPropagation={toggleManager}
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
