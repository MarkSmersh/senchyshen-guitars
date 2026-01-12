<script lang="ts">
	import { addCartItem } from '$lib/api/carts';
	import Button from '$lib/components/Button.svelte';
	import Icon from '$lib/components/Icon.svelte';
	import IconType from '$lib/components/IconType.svelte';
	import ImageSelector from '$lib/components/ImageSelector.svelte';
	import Price from '$lib/components/Price.svelte';
	import ProductDetails from '$lib/components/ProductDetails.svelte';
	import ProductType from '$lib/components/ProductType.svelte';
	import type { PageData } from './$types';

	const { data }: { data: PageData } = $props();
</script>

<main>
	<div class="selector-container">
		<ImageSelector images={data.images.sort((a, b) => a.id - b.id).map((i) => i.path)} />
	</div>
	<div class="data">
		<div class="top">
			<h1>{data.title}</h1>
			<div class="type">
				<IconType alt height={28} type={data.type} />
				<ProductType type={data.type} />
			</div>
		</div>
		<p>{data.description}</p>
		<div class="details">
			<h4>Szczegóły:</h4>
			<div class="container">
				<ProductDetails details={data[data.type]} />
			</div>
		</div>
		<div class="bottom">
			<h3><Price price={data.price} /></h3>
			<Button type={'accent'} onclick={() => addCartItem({ productId: data.id, count: 1 })}>
				<Icon height={24} name={'cart'} />
				Do koszyka
			</Button>
		</div>
	</div>
</main>

<style>
	main {
		display: flex;
		padding: 100px;
		gap: 50px;

		.selector-container {
			display: flex;
			height: calc(100vh - 150px);
		}

		.data {
			padding-top: 20px;
			display: flex;
			flex-direction: column;
			gap: 20px;
			width: 100%;

			.top {
				display: flex;
				flex-direction: column;
				gap: 5px;

				.type {
					display: flex;
					align-items: center;
					gap: 5px;
					font-size: 20px;
				}
			}

			p {
				font-weight: 600;
				white-space: pre-line;
			}

			.details {
				display: flex;
				flex-direction: column;
				gap: 10px;

				.container {
					padding: 10px;
				}
			}

			.bottom {
				display: flex;
				justify-content: space-between;
				margin-top: 30px;
				align-items: center;

				h3 {
					font-size: 24px;
				}
			}
		}
	}
</style>
