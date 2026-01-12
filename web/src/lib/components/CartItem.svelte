<script lang="ts">
	import type { ProductType as IProductType } from '$lib/api/products';
	import ProductItem from './ProductItem.svelte';

	interface CartItemProps {
		id: number;
		image: string;
		title: string;
		type?: IProductType;
		count: number;
		price: number;
		onremove?: (id: number) => void;
		onclick?: () => void;
	}

	const {
		id,
		title,
		price,
		type,
		image,
		count,
		onremove,
		onclick = () => {}
	}: CartItemProps = $props();
</script>

<ProductItem {type} {onclick} {count} {image} {price} {title} {id} {onremove} />
