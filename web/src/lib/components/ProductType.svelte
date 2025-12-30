<script lang="ts">
	import type { ProductType } from '$lib/api/products';

	const { type }: { type: ProductType } = $props();

	let name = $derived.by(() => {
		switch (type) {
			case 'guitar':
				return 'Gitara';
			case 'pickup':
				return 'Przetwornik';
			case 'bodyshape':
				return 'Kstałt';
			case 'amplifier':
				return 'Wzmacniacz';
			case 'crafted':
				return 'Na zamówienie';
		}
	});
</script>

<a href={`/catalog?type=${type}`}>
	{name}
</a>

<style>
	a {
		color: inherit;
	}
</style>
