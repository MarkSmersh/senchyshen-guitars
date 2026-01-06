<script lang="ts">
	import { goto } from '$app/navigation';
	import type { ProductType } from '$lib/api/products';
	import { ProductTypeToName } from '$lib/dictionary';

	const { type }: { type: ProductType } = $props();

	let name = $derived(ProductTypeToName(type));

	function gotoReload() {
		goto(`/catalog?type=${type}`).then(() => location.reload());
	}
</script>

<a onclick={() => gotoReload()} href={`/catalog?type=${type}`}>
	{name}
</a>

<style>
	a {
		color: inherit;
		cursor: pointer;
	}
</style>
