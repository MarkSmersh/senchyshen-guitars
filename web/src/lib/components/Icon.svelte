<script lang="ts">
	interface IconProps {
		name: string;
		height: number;
		alt?: boolean;
	}

	const { name, height, alt }: IconProps = $props();

	const icons: Record<string, string> = import.meta.glob(`$lib/assets/*.svg`, {
		eager: true,
		import: 'default'
	});

	let prefix = $derived(name + (alt ? '-alt' : '') + '.svg');

	let src = $derived(
		icons[Object.keys(icons).find((i) => i.split('/').at(-1) == prefix) as string]
	);
</script>

<img {height} alt={name} {src} />

<style>
	img {
		aspect-ratio: 1 / 1;
	}
</style>
