<script lang="ts">
	import { goto } from '$app/navigation';
	import leguitar from '$lib/assets/leguitar_crop.webp';
	import Button from '$lib/components/Button.svelte';
	import Icon from '$lib/components/Icon.svelte';
	import PreviewProduct from '$lib/components/PreviewProduct.svelte';
	import type { PageData } from './$types';

	const { data }: { data: PageData } = $props();
</script>

<!-- <svelte:document onscroll={(e) => onScroll(e)} /> -->
<main>
	<section id="sec-1">
		<div>
			<h1>Oferujemy Ci prawdziwy <span>DYI</span></h1>
			<h3>To Ty decydujesz, czego chcesz, a nie rynek</h3>
		</div>
		<img src={leguitar} class="guitar" alt="guitar" />
	</section>
	<section id="sec-2">
		<div class="border">
			<h1>My mamy...</h1>
		</div>
		<div class="content">
			<div class="preview">
				<Icon height={250} name="bodyshapes" />
				<p>{data.preview?.bodyshapes}</p>
				<h4>Kształtów</h4>
			</div>
			<div class="preview">
				<Icon height={250} name="pickups" />
				<p>{data.preview?.pickups}</p>
				<h4>Przetworników</h4>
			</div>
			<div class="preview">
				<Icon height={250} name="brush" />
				<p>16M</p>
				<h4>Kolorów</h4>
			</div>
		</div>
		<div class="border">
			<Button onclick={() => goto('/constructor')}>
				<Icon name="constructor" height={25} />
				Zdubuj sam!
			</Button>
		</div>
	</section>
	<section id="sec-3">
		<h3>Również rekomendujemy</h3>
		{#if data.products}
			<div class="products">
				{#each data.products as p}
					<PreviewProduct
						id={p.id}
						type={p.type}
						title={p.title}
						price={p.price}
						images={p.images.map((i) => i.path)}
					/>
				{/each}
			</div>
		{/if}
		<div class="catalog">
			<p>Lub wybierz z</p>
			<Button onclick={() => goto('/catalog')} type="accent">
				<Icon height={24} name="catalog" />
				Katalogu
			</Button>
		</div>
	</section>
</main>

<style>
	main {
		scroll-snap-type: y mandatory;
		height: 100vh;
		overflow: scroll;
		position: relative;

		section {
			height: 100vh;
			scroll-snap-align: center;
		}

		#sec-1 {
			padding: 0px 100px;

			h1 {
				color: var(--accent);
				font-size: 36px;

				span {
					background: var(--accent);
					color: var(--primary);
					padding: 4px 8px;
				}
			}

			h3 {
				font-size: 24px;
			}

			display: flex;
			align-items: center;

			div {
				display: flex;
				flex-direction: column;
				gap: 10px;
				z-index: 1;
			}

			.guitar {
				position: absolute;
				height: 80vh;
				bottom: 0;
				right: 0;
			}
		}

		#sec-2 {
			background-color: var(--secondary);
			padding-top: 80px;
			height: calc(100vh - 80px);

			display: flex;
			flex-direction: column;
			justify-content: space-between;

			.border {
				display: flex;
				background: var(--accent);
				height: 60px;
				color: var(--primary);
				align-items: center;
				justify-content: center;
			}

			.content {
				background: var(--secondary);
				display: flex;
				align-items: center;
				justify-content: space-between;
				padding: 0px 150px;
			}
		}

		.preview {
			display: flex;
			align-items: center;
			justify-content: center;
			position: relative;
			flex-direction: column;
			gap: 10px;

			h4 {
				font-size: 26px;
				color: var(--primary);
				text-shadow: 0px 0px 4px rgba(0, 0, 0, 0.2);
			}

			p {
				position: absolute;
				height: 80px;
				padding: 8px 8px;
				margin-bottom: calc(26px + 10px);
				z-index: 10;
				font-size: 56px;
				background-color: var(--secondary-opacity);
				border-radius: 56px;
				aspect-ratio: 1 / 1;
				display: flex;
				align-items: center;
				justify-content: center;
				font-weight: bold;
				color: var(--primary);
				text-shadow: -2px 2px 8px rgba(0, 0, 0, 0.25);
			}
		}

		#sec-3 {
			height: calc(100vh - 200px);
			padding: 100px 100px;
			display: flex;
			flex-direction: column;
			gap: 40px;

			h3 {
				font-size: 28px;
				text-align: center;
			}

			.products {
				display: flex;
				justify-content: center;
				gap: 100px;
			}

			.catalog {
				display: flex;
				gap: 10px;
				align-items: center;
				font-size: 18px;
				justify-content: center;
			}
		}
	}
</style>
