<script lang="ts">
	import { goto, invalidateAll } from '$app/navigation';
	import { clearCart, getCart, removeCartItem } from '$lib/api/carts';
	import { createOrder } from '$lib/api/orders';
	import Button from '$lib/components/Button.svelte';
	import CartForm from '$lib/components/CartForm.svelte';
	import CartItem from '$lib/components/CartItem.svelte';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();

	let items = $derived(data.items);
	let total = $derived(data.items.map((i) => i.total).reduce((prev, a) => prev + a, 0));
	let phoneNumber = $state('');
	let comment = $state('');

	async function clear() {
		await clearCart();
		await updateItems();
	}

	async function updateItems() {
		items = await getCart();
		await invalidateAll();
	}

	async function removeItem(id: number) {
		await removeCartItem(id);
		await updateItems();
	}

	async function order() {
		if (data.items.length <= 0) {
			// Will I rewrite the API you ask me...? Well, that's an amazing question.
			// I mean this a great question....
			alert('Koszyk jest pusty');

			return;
		}

		const order = await createOrder({ tel: phoneNumber, comment: comment });

		goto('/orders/' + order);
	}
</script>

<main>
	<div class="cart">
		<div class="toolbar">
			<h1>Koszyk</h1>
			<Button type={'accent'} onclick={() => clear()}>Wyczyść</Button>
		</div>
		<div class="items">
			{#if items.length > 0}
				{#each items as i}
					<CartItem
						id={i.id}
						count={i.count}
						image={i.image}
						type={i.type}
						price={i.price}
						title={i.title}
						onremove={(id) => removeItem(id)}
					/>
				{/each}
			{:else}
				Koszyk jest pusty
			{/if}
		</div>
	</div>
	<CartForm {total} bind:phoneNumber bind:comment onorder={() => order()} />
</main>

<style>
	main {
		padding: 150px 100px;
		display: grid;
		grid: auto / 2fr 1fr;
		gap: 20px;

		.cart {
			display: flex;
			flex-direction: column;
			gap: 20px;

			.toolbar {
				display: flex;
				justify-content: space-between;
			}

			.items {
				padding: 10px;
				display: flex;
				flex-direction: column;
				gap: 10px;
			}
		}
	}
</style>
