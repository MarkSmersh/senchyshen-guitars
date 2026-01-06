<script lang="ts">
	import { goto } from '$app/navigation';
	import CartItem from '$lib/components/CartItem.svelte';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();

	$inspect(data);
</script>

<main>
	<div class="order">
		<h1>{data.uuid}</h1>
		<h3>Telefon: {data.tel}</h3>
		{#if data.comment != ''}
			<h3>Komentarz: {data.comment}</h3>
		{/if}
		{#each data.items as i}
			<CartItem
				id={i.productId}
				image={i.image}
				price={i.price}
				title={i.title}
				count={i.count}
				onclick={() => goto('/products/' + i.productId)}
			/>
		{/each}
	</div>
</main>

<style>
	main {
		padding: 100px;

		.order {
			display: flex;
			flex-direction: column;
			gap: 10px;
		}
	}
</style>
