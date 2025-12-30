<script lang="ts">
	import { env } from '$env/dynamic/public';

	const { src, alt }: { src: string; alt: string } = $props();

	let serverSrc = $derived(env.PUBLIC_SERVER + '/assets/' + src);
</script>

<img src={serverSrc} {alt} />

<style>
	img {
		height: inherit;
	}
</style>
