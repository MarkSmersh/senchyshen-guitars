<script lang="ts">
	import type { ProductType as IProductType } from '$lib/api/products';
	import Icon from './Icon.svelte';
	import IconType from './IconType.svelte';
	import Image from './Image.svelte';
	import Price from './Price.svelte';
	import ProductType from './ProductType.svelte';

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

<div class="cart-item">
	<button onclick={() => onclick()} class="image">
		<Image src={image} alt={title} />
	</button>
	<div class="text">
		<h3>{title}</h3>
		{#if type}
			<div class="type">
				<IconType {type} height={20} />
				<ProductType {type} />
			</div>
		{/if}
	</div>
	<div class="numbers">
		<p>{count} szt.</p>
		<h4><Price {price} /></h4>
	</div>
	{#if onremove}
		<div class="remove">
			<button onclick={() => onremove(id)}>
				<Icon name={'no'} height={24} />
			</button>
		</div>
	{/if}
</div>

<style>
	.cart-item {
		height: 90px;
		padding: 8px;
		background: var(--accent);
		display: flex;
		gap: 16px;
		color: var(--primary);

		.text {
			display: flex;
			flex-direction: column;
			justify-content: center;
			gap: 8px;
			height: 100%;
			width: 100%;

			.type {
				display: flex;
				align-items: center;
				gap: 4px;
			}
		}

		.numbers {
			display: flex;
			flex-direction: column;
			justify-content: center;
			gap: 8px;
			text-wrap: nowrap;
			text-align: right;

			h4 {
				font-weight: 600;
				font-size: 18px;
			}
		}

		.remove {
			display: flex;
			align-items: center;
			justify-content: center;
			padding: 0px 10px;

			button {
				display: flex;
				cursor: pointer;
			}
		}

		.image {
			height: 100%;
			aspect-ratio: 1 / 1;
			background: var(--secondary);
			cursor: pointer;
		}
	}
</style>
