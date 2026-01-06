<script lang="ts">
	import type { ProductDetails } from '$lib/api/products';
	import { ProductDetailToName } from '$lib/dictionary';

	const { name }: { name: ProductDetails } = $props();
</script>

{ProductDetailToName[name]}
