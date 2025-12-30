<script lang="ts">
	interface IconProps {
		name: string;
		height: number;
		alt?: boolean;
	}

	const { name, height, alt }: IconProps = $props();
</script>

<img {height} alt={name} src={'/src/lib/assets/' + name + (alt ? '-alt' : '') + '.svg'} />

<style>
	img {
		aspect-ratio: 1 / 1;
	}
</style>
