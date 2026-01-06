<script lang="ts">
	import Image from './Image.svelte';

	let { images }: { images: string[] } = $props();

	let selected = $state(0);

	let image = $derived(images[selected]);
</script>

<div class="selector">
	<div class="container">
		<div class="image">
			<Image src={image} alt={image} />
		</div>
	</div>
	<div class="images">
		{#each images as i, index}
			<button class="container" onclick={() => (selected = index)}>
				<div class="image">
					<Image src={i} alt={i} />
				</div>
			</button>
		{/each}
	</div>
</div>

<style>
	.selector {
		display: flex;
		flex-direction: column;
		justify-content: flex-start;
		gap: 20px;
		align-items: center;
		height: 100%;

		.container {
			aspect-ratio: 1 / 1;
			background: var(--accent);
			display: flex;
			padding: 20px;
			/* I am so bad terrible at css... */
			/* height: 100%; */
			width: min-content;
			/* flex: 1; */
			height: 70%;

			.image {
				aspect-ratio: 1 / 1;
				background: var(--secondary);
				display: flex;
				height: 100%;
				width: min-content;
			}
		}

		/* @media only screen and (max-width: 1200px) { */
		/* 	.container { */
		/* 		width: calc(70% - 20px * 2); */
		/* 		height: unset; */
		/* 		flex: 0; */
		/* 	} */
		/* } */

		.images {
			display: flex;
			flex-wrap: wrap;
			justify-content: center;
			gap: 10px;
			/* height: 100%; */
			width: 100%;
			flex: 0 1 0;

			.container {
				height: 60px;
				padding: 4px;
				cursor: pointer;
				width: fit-content;
				flex-grow: 0;
			}
		}
	}
</style>
