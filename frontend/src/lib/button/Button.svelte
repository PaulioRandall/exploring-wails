<script>
	export let on_click = (e) => {}

	export let type = 'submit' // "submit", "cancel"
	export let disabled = false
</script>

<button
	{disabled}
	class="button"
	class:submit-type={type === 'submit'}
	class:cancel-type={type === 'cancel'}
	on:click={on_click}>
	<slot />
</button>

<style>
	.button {
		display: inline-block;

		appearance: none;
		border-style: none;

		padding: 16px 20px;

		color: #ffffff;
		font-size: 18px;
		font-weight: bold;
		text-align: center;

		cursor: pointer;
		user-select: none;
		touch-action: manipulation;

		box-shadow: 0 0 1px 1px #505050;
		border-radius: 8px;

		transition: all 0.2s;
	}

	.button:hover {
		transform: scale(1.05);
		box-shadow: 2px 2px 4px 1px #303030;
	}

	.submit-type {
		background-color: forestgreen;
	}

	.cancel-type {
		background-color: slategrey;
	}

	.button:active {
		transition: all 0.05s;
		transform: scale(0.95);
		box-shadow: inset 2px 2px 4px 1px #303030;
	}

	.button:disabled {
		transition: none !important;
		transform: none !important;
		box-shadow: 2px 2px 4px 1px #303030 !important;
		filter: opacity(40%) brightness(50%);
	}
</style>
