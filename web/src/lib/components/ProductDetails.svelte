<script lang="ts">
	import type { ProductDetails, ProductModel } from '$lib/api/products';
	import ProductDetail from './ProductDetail.svelte';
	import Color from './Color.svelte';
	import Icon from './Icon.svelte';

	const { product }: { product: ProductModel } = $props();

	const productDetails = $derived(product[product.type]);
</script>

{#if productDetails}
	<div class="details">
		{#each Object.keys(productDetails) as ProductDetails[] as detail}
			<div class="detail">
				<ProductDetail name={detail} />
				{#if 'color' in productDetails && detail == 'color'}
					{@const value = productDetails['color']}
					<div class="color" title={'#' + value}>
						<Color size={20} color={value} />
						<!-- <p>{productDetails['color']}</p> -->
					</div>
				{:else if 'bodyshapeId' in productDetails && detail == 'bodyshapeId'}
					{@const value = productDetails['bodyshapeId']}
					<div class="bodyshape">
						<Icon name="bodyshapes" height={20} alt />
						<p>
							ID:
							<a href={'/products/' + value}>{value}</a>
						</p>
					</div>
				{:else if 'pickups' in productDetails && detail == 'pickups'}
					{@const value = productDetails['pickups']}
					<div class="pickups">
						{#each value as p}
							<div class="pickup">
								<Icon alt name="pickups" height={20} />
								<p>
									ID:
									<a href={'/products/' + p.pickupId}>{p.pickupId}</a>
									({p.position})
								</p>
							</div>
						{/each}
					</div>
				{:else if 'power' in productDetails && detail == 'power'}
					{@const value = productDetails['power']}
					<p>
						{value} W
					</p>
				{:else}
					<p>{(productDetails as any)[detail]}</p>
				{/if}
			</div>
		{/each}
	</div>
{/if}

<style>
	.pickups {
		display: flex;
		gap: 10px;

		.pickup {
			display: flex;
			gap: 4px;
		}
	}

	.bodyshape {
		display: flex;
		align-items: center;
		gap: 4px;
	}

	.color {
		display: flex;
		gap: 4px;
		align-items: center;
	}

	.details {
		display: flex;
		flex-direction: column;
		gap: 12px;

		.detail {
			display: flex;
			justify-content: space-between;
			align-items: center;
		}

		a {
			color: var(--accent);
		}
	}
</style>
