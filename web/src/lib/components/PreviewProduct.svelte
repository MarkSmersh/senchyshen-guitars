<script lang="ts">
	import type { ProductType as IProductType } from '$lib/api/products';
	import IconType from './IconType.svelte';
	import Image from './Image.svelte';
	import Price from './Price.svelte';
	import ProductType from './ProductType.svelte';

	interface PreviewProductsProps {
		id: number;
		image: string;
		title: string;
		type: IProductType;
		price: number;
	}

	const { id, image, type, price, title }: PreviewProductsProps = $props();
</script>

<div class="container">
	<div class="image-container">
		<Image src={image} alt={title} />
	</div>
	<div class="data">
		<a href={'/products/' + id}><h4>{title}</h4></a>
		<div class="type">
			<IconType {type} height={24} />
			<ProductType {type} />
		</div>
	</div>
	<p>
		<Price {price} />
	</p>
</div>

<style>
	.container {
		background: var(--accent);
		width: 300px;
		box-sizing: border-box;
		aspect-ratio: 9 / 13;
		padding: 16px;
		position: relative;
		display: flex;
		flex-direction: column;
		transition: 0.2s;

		.image-container {
			aspect-ratio: 1 / 1;
			background: var(--secondary);
			width: 100%;
			display: flex;
		}

		.data {
			padding: 16px 8px;
			display: flex;
			flex-direction: column;
			justify-content: start;
			gap: 8px;

			h4 {
				overflow-x: hidden;
				text-overflow: ellipsis;
				white-space: nowrap;
				font-size: 20px;
			}

			a {
				color: var(--primary);
			}

			.type {
				display: flex;
				align-items: center;
				gap: 6px;
				font-size: 18px;
				color: var(--secondary);
			}
		}

		p {
			color: var(--primary);
			display: flex;
			font-size: 24px;
			font-weight: 600;
			padding: 8px;
		}
	}

	.container:hover {
		box-shadow: -4px 0px 12px rgba(0, 0, 0, 0.25);
	}
</style>
