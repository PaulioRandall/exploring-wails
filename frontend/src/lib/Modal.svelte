<div class="modal">
	<slot />
</div>

<style>
	.modal {
		position: absolute;
		top: 0;
		left: 0;

		width: 100vw;
		height: 100vh;

		background: rgba(0, 0, 0, 0.3);
	}

	.modal > :global(*) {
		--modal-pad: 3rem;
		height: calc(100% - var(--modal-pad) * 2);
		margin: var(--modal-pad);
	}
</style>
