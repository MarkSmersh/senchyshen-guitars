<script lang="ts">
	import Header from '$lib/components/Header.svelte';

	let { children } = $props();
</script>

<svelte:head>
	<!-- <link rel="icon" href={favicon} /> -->
</svelte:head>

<Header />
{@render children()}

<style>
	@import url('https://fonts.googleapis.com/css2?family=Lato:ital,wght@0,100;0,300;0,400;0,700;0,900;1,100;1,300;1,400;1,700;1,900&display=swap');

	:root {
		--primary: #e8d2c4;
		--secondary: #cea890;
		--accent: #413630;

		--secondary-opacity: rgba(206, 168, 144, 0.2);
	}

	:global(*) {
		padding: 0;
		margin: 0;
		border: 0;
	}

	:global(body) {
		background: var(--primary);
		color: var(--accent);
		font-family: 'Lato', sans-serif;
		font-style: normal;
	}

	:global(input) {
		background: transparent;
		border: 0;
		color: var(--accent);
	}

	:global(input:focus) {
		outline: none;
	}

	:global(input ::placeholder) {
		color: var(--secondary);
	}
</style>
