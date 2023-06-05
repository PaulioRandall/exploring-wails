<script>
	import { fade } from 'svelte/transition'
	import { message, msgTypes } from './store.js'

	let msg = null
	let timeout = null

	$: if ($message) {
		msg = $message
		message.set(null)

		clearTimeout(timeout)
		timeout = setTimeout(() => (msg = null), 1500)
	}
</script>

<section class="messenger">
	{#if msg}
		<div
			class="message"
			style="background: {msgTypes[msg.type].bgColor}"
			out:fade={{ duration: 1500 }}>
			{@html msg.text.replace(/\n/g, '<br/>')}
		</div>
	{/if}
</section>

<style>
	.messenger {
		position: absolute;
		top: 1rem;
		left: 1rem;

		pointer-events: none;
	}

	.message {
		text-align: left;
		font-size: 24px;
		padding: 0.2rem 0.5rem;

		border-radius: 1rem;
	}
</style>
