<script lang="ts">
	import Image from './Image.svelte';

	const { images }: { images: string[] } = $props();

	let container: HTMLDivElement | undefined = $state();

	let image = $state(0);

	function swapImage(e: MouseEvent) {
		if (container) {
			const parent = container.getBoundingClientRect();

			const perBlock = parent.width / images.length;
			const x = Math.floor((e.clientX - parent.left) / perBlock);

			image = x;
		}
	}
</script>

<!-- svelte-ignore a11y_no_static_element_interactions, a11y_mouse_events_have_key_events -->
<div
	bind:this={container}
	class="image-container"
	onmousemove={(e) => swapImage(e)}
	onmouseleave={() => (image = 0)}
>
	<Image src={images[image]} alt={images[image]} />
</div>

<style>
	.image-container {
		aspect-ratio: 1 / 1;
		background: var(--secondary);
		width: 100%;
		display: flex;
	}
</style>
