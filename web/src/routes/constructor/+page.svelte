<script lang="ts">
	import { goto } from '$app/navigation';
	import { constructGuitar } from '$lib/api/constructors';
	import { getProduct, type Position, type ProductModel } from '$lib/api/products';
	import Button from '$lib/components/Button.svelte';
	import ConstructorForm from '$lib/components/ConstructorForm.svelte';
	import DropdownList from '$lib/components/DropdownList.svelte';
	import Icon from '$lib/components/Icon.svelte';
	import ProductItem from '$lib/components/ProductItem.svelte';
	import ProductSearch from '$lib/components/ProductSearch.svelte';

	interface Pickup extends ProductModel {
		pos: Position;
		index: number;
	}

	const positions: Position[] = ['top', 'middle', 'bottom'];

	let bodyshape: ProductModel | undefined = $state();

	let pickups: Pickup[] = $state([]);

	let title = $state('');

	let color = $derived(bodyshape ? '#' + bodyshape.bodyshape?.color : '');
	let isDiffentColor = $derived(bodyshape ? '#' + bodyshape.bodyshape?.color != color : false);

	let publish = $state(false);

	let conflicted = $derived.by(() => {
		let v = '';

		const poses = pickups.map((p) => p.pos).sort();

		for (let i = 0; i < poses.length; i++) {
			if (v == poses[i]) {
				return true;
			}

			v = poses[i];
		}

		return false;
	});

	async function selectBodyshape(id: number) {
		bodyshape = await getProduct(id);
	}

	async function addPickup(id: number, pos: Pickup['pos']) {
		const pickup = await getProduct(id);
		const i = Math.round(Math.random() * 1000000);
		pickup ? pickups.push({ pos: pos, index: i, ...pickup }) : null;
	}

	function updatePickup(index: number, pos: Position) {
		pickups[index].pos = pos;
	}

	function removePickup(index: number) {
		pickups = pickups.filter((_, pi) => pi != index);
	}

	function reset() {
		bodyshape = undefined;
		pickups = [];
		color = '';
	}

	async function toCart() {
		if (!bodyshape) {
			alert('Musi być wybrany kstałt');
			return;
		}

		if (pickups.length < 1) {
			alert('Minimalna ilość przetworników');
			return;
		}

		await constructGuitar({
			title: title,
			color: color.replace('#', ''),
			pickups: pickups.map((p) => {
				return { pickupId: p.id, position: p.pos };
			}),
			bodyshapeId: bodyshape.id,
			publish: publish
		});

		reset();
	}
</script>

<main>
	<div class="constructor">
		<div class="toolbar">
			<h1>Konstructor</h1>
			<Button type={'accent'} onclick={() => reset()}>Resetuj</Button>
		</div>
		<div class="parts">
			<div class="part">
				<h2>
					<Icon height={28} name={'bodyshapes'} alt />
					Kształt
				</h2>
				{#if bodyshape}
					{@const b = bodyshape}
					<ProductItem
						type={b.type}
						price={b.price}
						title={b.title}
						image={b.images && b.images.length > 0 ? b.images[0].path : ''}
						id={b.id}
						onclick={() => goto('/products/' + b.id)}
						onremove={() => (bodyshape = undefined)}
					/>
				{/if}
				{#if !bodyshape}
					<ProductSearch onadd={(id) => selectBodyshape(id)} type="bodyshape" />
				{/if}
			</div>
			<div class="part">
				<h2>
					<Icon height={28} name={'pickups'} alt />
					Przetworniky
				</h2>
				{#if conflicted}
					<p class="conflict">Nie może być więcej jednego przetwonika dla jednej pozycji.</p>
				{/if}
				{#each pickups as p, index (p.index)}
					<div class="item">
						<div class="pickup">
							<ProductItem
								type={p.type}
								price={p.price}
								title={p.title}
								image={p.images && p.images.length > 0 ? p.images[0].path : ''}
								id={p.id}
								onclick={() => goto('/products/' + p.id)}
								onremove={() => removePickup(index)}
							/>
						</div>
						<div class="pos">
							<DropdownList
								title="Pozycja"
								item={positions[0]}
								items={positions}
								onupdate={(pos) => updatePickup(index, pos as Position)}
							/>
						</div>
					</div>
				{/each}
				{#if pickups.length < 3}
					<ProductSearch onadd={(id) => addPickup(id, 'top')} type="pickup" />
				{/if}
			</div>
			<div class="part">
				<h2>
					<Icon height={28} name={'brush'} alt />
					Kolor
				</h2>
				<div class="color">
					<h3>Kolor:</h3>
					<input type="color" bind:value={color} />
					{#if color != '' && isDiffentColor}
						<button onclick={() => (color = bodyshape ? '#' + bodyshape.bodyshape?.color : '')}>
							<Icon height={24} name="no" />
						</button>
					{/if}
				</div>
			</div>
		</div>
	</div>
	<ConstructorForm
		bind:publish
		bind:title
		oncart={() => toCart()}
		colorPrice={isDiffentColor ? 200 * 100 : undefined}
		bodyshapePrice={bodyshape?.price}
		servicePrice={1000 * 100}
		pickupPrices={pickups.map((p) => p.price)}
	/>
</main>

<style>
	main {
		padding: 150px 100px;
		grid: auto-flow / 2fr 1fr;
		display: grid;
		gap: 20px;

		.constructor {
			display: flex;
			flex-direction: column;
			gap: 10px;

			.toolbar {
				display: flex;
				justify-content: space-between;
			}

			.parts {
				display: flex;
				flex-direction: column;
				gap: 20px;
				padding: 10px;

				.part {
					display: flex;
					flex-direction: column;
					gap: 10px;

					.color {
						display: flex;
						gap: 10px;
						align-items: center;
						background: var(--accent);
						padding: 10px;
						color: var(--primary);

						input {
							width: 100%;
							border: 0;
						}

						button {
							display: flex;
							cursor: pointer;
						}
					}

					h2 {
						display: flex;
						align-items: center;
						gap: 8px;
					}

					.conflict {
						color: #aa0000;
					}

					.item {
						display: flex;
						background: var(--accent);
						align-items: center;

						.pickup {
							width: 100%;
						}

						.pos {
							text-align: right;
							padding: 10px;
						}
					}
				}
			}
		}
	}
</style>
