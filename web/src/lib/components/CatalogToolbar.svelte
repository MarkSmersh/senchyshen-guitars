<script lang="ts">
	import type { Order, OrderBy } from '$lib/api/products';
	import { OrderByToName, OrderToName } from '$lib/dictionary';
	import DropdownList from './DropdownList.svelte';
	import Icon from './Icon.svelte';

	export interface CatalogToolbarEvent {
		query: string;
		limit: string;
		orderBy: string;
		order: string;
	}

	interface CatalogToolbarProps {
		limits: string[];
		orderBys: string[];
		orders: string[];
		onupdate: (e: CatalogToolbarEvent) => void;
	}

	const { limits, orderBys, orders, onupdate }: CatalogToolbarProps = $props();

	let limit = $state(limits[0]);
	let orderBy = $state(orderBys[0]);
	let order = $state(orders[0]);
	let q = $state('');

	function update(): boolean {
		onupdate({
			query: q,
			order: order,
			orderBy: orderBy,
			limit: limit
		});

		return true;
	}
</script>

<div class="toolbar">
	<div class="search-container">
		<p>Wyszukiwanie</p>
		<div class="search">
			<Icon name={'search'} height={16} />
			<input bind:value={q} oninput={() => update()} placeholder="Nazwa produktu..." type="text" />
		</div>
	</div>
	<div class="dropdowns">
		<DropdownList
			onupdate={() => update()}
			bind:item={limit}
			items={limits}
			title={'Limit na stronę'}
		/>
		<DropdownList
			translateWith={(v) => OrderByToName[v as OrderBy]}
			onupdate={() => update()}
			bind:item={orderBy}
			items={orderBys}
			title={'Sortować według'}
		/>
		<DropdownList
			translateWith={(v) => OrderToName[v as Order]}
			onupdate={() => update()}
			bind:item={order}
			items={orders}
			title={'Kolejność'}
		/>
	</div>
</div>

<style>
	.toolbar {
		background: var(--accent);
		padding: 20px;
		display: flex;
		align-items: center;
		gap: 20px;

		.search-container {
			width: 100%;
			gap: 5px;
			display: flex;
			flex-direction: column;
			color: var(--primary);

			.search {
				background: var(--primary);
				padding: 5px;
				display: flex;
				align-items: center;
				gap: 5px;
				height: calc(30px - 5px * 2);
			}
		}

		.dropdowns {
			display: flex;
			gap: 20px;
			text-align: right;
		}
	}
</style>
