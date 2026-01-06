<script lang="ts">
	import { goto } from '$app/navigation';
	import Category from '$lib/components/Category.svelte';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
</script>

<main>
	{#each data.categories as c}
		<Category
			id={c.id}
			title={c.title}
			description={c.description}
			image={c.image ? c.image : ''}
			onclick={() => goto('/catalog?category=' + c.id)}
		/>
	{/each}
</main>

<style>
	main {
		padding: 100px;
		gap: 50px;
		display: grid;
		grid: auto / repeat(3, 1fr);
	}
</style>
