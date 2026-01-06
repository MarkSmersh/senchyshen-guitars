<script lang="ts">
	let { color, size }: { color: string; size: number } = $props();
</script>

<div class="color" style:height={size + 'px'} style:background={'#' + color}></div>

<style>
	div {
		aspect-ratio: 1 / 1;
		border: 2px solid var(--accent);
	}
</style>
