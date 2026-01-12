<script lang="ts">
	import type { ProductType as IProductType, ProductModel } from '$lib/api/products';
	import Icon from './Icon.svelte';
	import IconType from './IconType.svelte';
	import Image from './Image.svelte';
	import Price from './Price.svelte';
	import ProductDetails from './ProductDetails.svelte';
	import ProductType from './ProductType.svelte';

	interface CartItemProps {
		id: number;
		image: string;
		title: string;
		type?: IProductType;
		count?: number;
		price: number;
		alt?: boolean;
		details?: ProductModel[ProductModel['type']];
		onremove?: (id: number) => void;
		onadd?: (id: number) => void;
		onclick?: () => void;
	}

	const {
		id,
		title,
		price,
		type,
		image,
		count,
		alt = false,
		onremove,
		onadd,
		details,
		onclick = () => {}
	}: CartItemProps = $props();

	let detailsOn = $state(false);

	$inspect(details);
</script>

<div class="cart-item" class:alt>
	<div class="preview">
		<button onclick={() => onclick()} class="image">
			<Image src={image} alt={title} />
		</button>
		<div class="text">
			<h3>{title}</h3>
			{#if type}
				<div class="type">
					<IconType {type} height={20} {alt} />
					<ProductType {type} />
				</div>
			{/if}
		</div>
		<div class="numbers">
			{#if count}
				<p>{count} szt.</p>
			{/if}
			<h4><Price {price} /></h4>
		</div>
		<div class="actions">
			{#if onremove}
				<div class="action">
					<button onclick={() => onremove(id)}>
						<Icon name={'no'} height={24} />
					</button>
				</div>
			{/if}
			{#if onadd}
				<div class="action">
					<button onclick={() => onadd(id)}>
						<Icon name={'plus'} height={24} />
					</button>
				</div>
			{/if}
			{#if details}
				<div class="action">
					<button onclick={() => (detailsOn = !detailsOn)}>
						<Icon name={'down'} height={24} />
					</button>
				</div>
			{/if}
		</div>
	</div>
	{#if detailsOn}
		<div class="details">
			<ProductDetails {details} />
		</div>
	{/if}
</div>

<style>
	.cart-item {
		padding: 8px;
		background: var(--accent);
		display: flex;
		flex-direction: column;
		gap: 16px;
		color: var(--primary);

		.preview {
			gap: 10px;
			display: flex;
			height: 90px;
			justify-content: space-between;
			width: 100%;

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

			.actions {
				display: flex;
				flex-direction: column;
				gap: 10px;
				justify-content: center;

				.action {
					display: flex;
					align-items: center;
					justify-content: center;
					padding: 0px 10px;

					button {
						display: flex;
						cursor: pointer;
					}
				}
			}

			.image {
				height: 100%;
				aspect-ratio: 1 / 1;
				background: var(--secondary);
				cursor: pointer;
			}
		}

		.details {
			padding: 5px;
		}
	}

	.alt {
		background: var(--secondary);
		color: var(--accent);

		.preview {
			.image {
				background: var(--primary);
			}
		}
	}
</style>
