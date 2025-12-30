<script lang="ts">
	import type { ProductType } from '$lib/api/products';
	import Icon from './Icon.svelte';

	const { type, height, alt }: { type: ProductType; height: number; alt?: boolean } = $props();

	let name = $derived.by(() => {
		switch (type) {
			case 'guitar':
				return 'guitars';
			case 'pickup':
				return 'pickups';
			case 'bodyshape':
				return 'bodyshapes';
			case 'amplifier':
				return 'amplifiers';
			case 'crafted':
				return 'guitars';
		}
	});
</script>

<Icon {name} {height} {alt} />
