<script>
	import { GetFilesInDir } from '#backend'

	export let on_select = (file) => {}
	export let open_at = '.'

	let fileList = GetFilesInDir(open_at).catch((e) => {
		console.error(e)
	})
</script>

<div class="file-selector"></div>

<style>
	.file-selector {
		overflow-x: hidden;
		overflow-y: scroll;
	}
</style>
