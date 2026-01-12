<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import { getProducts, type Order, type OrderBy, type ProductType } from '$lib/api/products';
	import CatalogToolbar, { type CatalogToolbarEvent } from '$lib/components/CatalogToolbar.svelte';
	import Checklist, { type ChecklistItem } from '$lib/components/Checklist.svelte';
	import ProductCatalog from '$lib/components/ProductCatalog.svelte';
	import Slider from '$lib/components/Slider.svelte';
	import { ProductTypeToName } from '$lib/dictionary.js';

	const { data } = $props();

	let search = $state(data);

	let types: ChecklistItem[] = $state([]);
	let categories: ChecklistItem[] = $state([]);

	let minPrice = $state(0);
	let maxPrice = $state(0);

	let query = $state('');
	let limit = $state(25);
	let page = $state(0);
	let orderBy = $state('');
	let order = $state('');

	let limits = [25, 50, 100].map((l) => `${l}`);
	let orders: Order[] = ['asc', 'desc'];
	let orderBys: OrderBy[] = ['title', 'price', 'createdAt'];
	let products = $derived(search.products);

	function updateTypes() {
		console.log('updateTypes', search);

		const newTypes: ChecklistItem[] = [];

		search?.types?.forEach((t) => {
			const oldType = types.find((type) => type.title == t);

			if (oldType) {
				newTypes.push(oldType);
			} else {
				newTypes.push({ title: t, checked: true });
			}
		});

		types = newTypes;
	}

	function updateCategories() {
		const newCategories: ChecklistItem[] = [];

		search?.categories?.forEach((t) => {
			const oldCategory = categories.find((type) => type.title == t.title);

			if (oldCategory) {
				newCategories.push(oldCategory);
			} else {
				newCategories.push({ title: t.title, checked: false });
			}
		});

		categories = newCategories;
	}

	function updatePrice(min: number, max: number) {
		minPrice = min;
		maxPrice = max;
		updateProducts();
	}

	async function updateProducts(byCategories: boolean = false) {
		let typesMap = types.map((t) => {
			if (t.checked) {
				return t.title;
			}
		});

		let category = search?.categories?.find(
			(cat) => cat.title == categories.find((c) => c.checked)?.title
		);

		console.log('category', category);

		const res = await getProducts({
			limit: limit,
			page: page,
			orderBy: orderBy as any,
			order: order as any,
			priceMax: maxPrice,
			priceMin: minPrice,
			query: query,
			types: types.length > 0 && !byCategories ? (typesMap as any) : null,
			category: category ? category.id : undefined
		});

		// removing the length condition is destroying the universe itself
		if (res) {
			console.log('pre search', search);
			search = res;
			console.log('post search', search);

			// if (res.types && res.types.length > 0) {
			// 	console.log('UPDATE TYPES');
			// 	updateTypes();
			// }
			//
			// if (res.categories && res.categories.length > 0) {
			// 	updateCategories();
			// }

			await invalidateAll();

			if (res.products.length > 0) {
				updateCategories();
				updateTypes();
			}

			await invalidateAll();
		}
	}

	function updateToolbar(data: CatalogToolbarEvent) {
		query = data.query;
		orderBy = data.orderBy;
		order = data.order;
		limit = parseInt(data.limit);

		updateProducts();
	}

	updateTypes();
	updateCategories();

	$inspect(categories);

	// I would rewrite this piece of shit
</script>

<main>
	<div class="filter-container">
		<h1>Filtracja</h1>
		<div class="filters">
			<Checklist
				bind:items={categories}
				onupdate={() => updateProducts()}
				single
				title="Kategoria"
			/>
			<Checklist
				translateWith={(v) => ProductTypeToName(v as ProductType)}
				bind:items={types}
				title="Typy produktu"
				onupdate={() => updateProducts()}
			/>
			<Slider
				title={'Cena'}
				from={search.priceMin ? search.priceMin : 0}
				to={search.priceMax ? search.priceMax : 0}
				onupdate={(min, max) => updatePrice(min, max)}
			/>
		</div>
	</div>
	<div class="catalog-container">
		<h1>Katalog</h1>
		<div class="catalog">
			<div class="header">
				<p>Znajdziono {products ? products.length : 0} produktów</p>
				<CatalogToolbar {limits} {orderBys} {orders} onupdate={(e) => updateToolbar(e)} />
			</div>
			<div class="products">
				{#if search.products}
					{#each search.products as p}
						<ProductCatalog
							type={p.type}
							title={p.title}
							id={p.id}
							images={p.images.toSorted((a, b) => a.id - b.id).map((i) => i.path)}
							price={p.price}
						/>
					{/each}
				{/if}
			</div>
		</div>
	</div>
</main>

<style>
	main {
		padding: 100px 20px;
		display: grid;
		gap: 20px;
		position: relative;
		grid: auto / 1fr 3fr;

		h1 {
			font-size: 24px;
		}

		.filter-container {
			/* width: 80%; */
			display: flex;
			flex-direction: column;
			gap: 15px;

			.filters {
				display: flex;
				flex-direction: column;
				gap: 10px;
			}
		}

		.catalog-container {
			display: flex;
			flex-direction: column;
			gap: 15px;

			.catalog {
				.header {
					display: flex;
					flex-direction: column;
					gap: 8px;
				}

				.products {
					display: grid;
					grid: auto / repeat(4, minmax(0, 1fr));
					padding: 10px;
					gap: 10px;
				}
			}
		}
	}
</style>
