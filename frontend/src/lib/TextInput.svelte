<script>
	export let value = ''
	export let label = ''
</script>

<label>
	{label}
	<input class="text-input" bind:value />
</label>

<style>
	.text-input {
		width: 100%;
	}
</style>
