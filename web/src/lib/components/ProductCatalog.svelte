<script lang="ts">
	import type { ProductType as IProductType } from '$lib/api/products';
	import IconType from './IconType.svelte';
	import Image from './Image.svelte';
	import Price from './Price.svelte';
	import ProductType from './ProductType.svelte';

	interface PreviewProductsProps {
		id: number;
		images: string[];
		title: string;
		type: IProductType;
		price: number;
	}

	let image = $state(0);
	let container: HTMLDivElement | undefined = $state();

	function swapImage(e: MouseEvent) {
		if (container) {
			const parent = container.getBoundingClientRect();

			const perBlock = parent.width / images.length;
			const x = Math.floor((e.clientX - parent.left) / perBlock);

			image = x;
		}
	}

	const { id, images, type, price, title }: PreviewProductsProps = $props();
</script>

<!-- svelte-ignore a11y_no_static_element_interactions, a11y_mouse_events_have_key_events -->
<div class="container">
	<div
		bind:this={container}
		class="image-container"
		onmousemove={(e) => swapImage(e)}
		onmouseleave={() => (image = 0)}
	>
		<Image src={images[image]} alt={title} />
	</div>
	<div class="data">
		<a href={'/products/' + id}><h4>{title}</h4></a>
		<div class="type">
			<IconType {type} height={18} />
			<ProductType {type} />
		</div>
	</div>
	<p>
		<Price {price} />
	</p>
</div>

<style>
	.container {
		background: var(--accent);
		box-sizing: border-box;
		aspect-ratio: 9 / 13;
		padding: 10px;
		position: relative;
		display: flex;
		flex-direction: column;
		transition: 0.2s;

		.image-container {
			aspect-ratio: 1 / 1;
			background: var(--secondary);
			width: 100%;
			display: flex;
		}

		.data {
			padding: 8px 4px;
			display: flex;
			flex-direction: column;
			justify-content: space-between;
			gap: 8px;

			h4 {
				overflow-x: hidden;
				text-overflow: ellipsis;
				white-space: nowrap;
				font-size: 16px;
			}

			a {
				color: var(--primary);
				text-decoration: none;
			}

			.type {
				display: flex;
				align-items: center;
				gap: 4px;
				font-size: 14px;
				color: var(--secondary);
			}
		}

		p {
			color: var(--primary);
			display: flex;
			font-size: 18px;
			font-weight: 600;
			padding: 8px 4px;
		}
	}

	.container:hover {
		a {
			text-decoration: underline;
		}
	}

	/* .container:hover { */
	/* 	box-shadow: -4px 0px 12px rgba(0, 0, 0, 0.25); */
	/* } */
</style>
