<script>
	import { currTool } from '#tools/store.js'
	import toolList from '#tools/tool-list.js'

	const navToTool = (tool) => {
		currTool.set(tool)
	}
</script>

<main class="main">
	<div class="menu">
		{#each toolList as tool}
			<div class="tool" on:click={() => navToTool(tool)}>
				<p class="tool-name">{tool.name}</p>
			</div>
		{/each}
	</div>
</main>

<style>
	.main {
		max-width: 100%;
		padding: 1rem;
	}

	.menu {
		display: flex;
		flex-wrap: wrap;
		justify-content: center;
		gap: 2rem;
	}

	.tool {
		width: 16rem;
		height: 16rem;

		border: 1px solid goldenrod;
		border-radius: 0.4rem;

		cursor: pointer;
	}

	.tool-name {
		text-align: center;
		font-size: 2rem;
	}
</style>
