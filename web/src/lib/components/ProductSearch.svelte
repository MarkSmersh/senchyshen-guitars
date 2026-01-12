<script lang="ts">
	import { getProduct, getProducts, type ProductModel, type ProductType } from '$lib/api/products';
	import { scale } from 'svelte/transition';
	import Icon from './Icon.svelte';
	import ProductItem from './ProductItem.svelte';
	import { goto } from '$app/navigation';

	interface ProductSearch {
		type: ProductType;
		onadd: (id: number) => void;
	}

	let { type, onadd }: ProductSearch = $props();

	let products: ProductModel[] = $state([]);

	let query = $state('');
	let active = $state(false);

	async function search() {
		const res = await getProducts({ limit: 100, query: query, types: [type] });

		if (res) {
			products = res.products;
		}
	}

	async function addDetails(id: number) {
		const i = products.findIndex((p) => p.id == id);

		if (!products[i][products[i].type]) {
			const product = await getProduct(id);
			product ? (products[i] = product) : null;
		}
	}

	search();
</script>

<!-- svelte-ignore a11y_no_noninteractive_tabindex, a11y_no_static_element_interactions, a11y_click_events_have_key_events -->
<div tabindex="0" class="search">
	<div class="query">
		<Icon name={'search'} height={20} alt />
		<input
			type="text"
			bind:value={query}
			oninput={() => search()}
			onfocus={() => (active = true)}
		/>
		{#if active}
			<button transition:scale={{ duration: 200 }} onclick={() => (active = false)}>
				<Icon name="no" height={20} />
			</button>
		{/if}
	</div>
	{#if active}
		<div class="container">
			<div class="products">
				{#each products as p, i}
					<!-- svelte-ignore a11y_no_static_element_interactions, a11y_mouse_events_have_key_events -->
					<div
						class="product"
						onmouseover={() => addDetails(p.id)}
						ontouchmove={() => addDetails(p.id)}
					>
						<ProductItem
							id={p.id}
							type={p.type}
							details={p[p.type]}
							image={p.images && p.images.length > 0 ? p.images[0].path : ''}
							price={p.price}
							title={p.title}
							onclick={() => goto('/products/' + p.id)}
							{onadd}
							alt
						/>
					</div>
				{/each}
			</div>
		</div>
	{/if}
</div>

<style>
	.search {
		display: flex;
		flex-direction: column;
		gap: 10px;

		.query {
			display: flex;
			gap: 8px;
			background: var(--accent);
			padding: 8px;

			input {
				width: 100%;
				color: var(--primary);
			}

			button {
				display: flex;
				cursor: pointer;
			}
		}

		.container {
			padding: 10px;
			background: var(--accent);
			max-height: 340px;

			.products {
				display: flex;
				max-height: 340px;
				flex-direction: column;
				gap: 10px;
				overflow: scroll;
			}
		}
	}
</style>
