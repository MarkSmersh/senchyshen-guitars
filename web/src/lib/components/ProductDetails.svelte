<script lang="ts">
	import type { ProductDetails, ProductModel } from '$lib/api/products';
	import ProductDetail from './ProductDetail.svelte';
	import Color from './Color.svelte';
	import Icon from './Icon.svelte';

	const { details }: { details: ProductModel[ProductModel['type']] } = $props();
</script>

{#if details}
	<div class="details">
		{#each Object.keys(details) as ProductDetails[] as detail}
			<div class="detail">
				<ProductDetail name={detail} />
				{#if 'color' in details && detail == 'color'}
					{@const value = details['color']}
					<div class="color" title={'#' + value}>
						<Color size={20} color={value} />
						<!-- <p>{details['color']}</p> -->
					</div>
				{:else if 'bodyshapeId' in details && detail == 'bodyshapeId'}
					{@const value = details['bodyshapeId']}
					<div class="bodyshape">
						<Icon name="bodyshapes" height={20} alt />
						<p>
							ID:
							<a href={'/products/' + value}>{value}</a>
						</p>
					</div>
				{:else if 'pickups' in details && detail == 'pickups'}
					{@const value = details['pickups']}
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
				{:else if 'power' in details && detail == 'power'}
					{@const value = details['power']}
					<p>
						{value} W
					</p>
				{:else}
					<p>{(details as any)[detail]}</p>
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
