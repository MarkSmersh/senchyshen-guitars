<script lang="ts">
	import Image from './Image.svelte';

	interface CategoryProps {
		id: number;
		title: string;
		description: string;
		image: string;
		onclick?: () => void;
	}

	const { id, image, title, description, onclick = () => {} }: CategoryProps = $props();
</script>

<div class="category">
	<a href={'/catalog?category=' + id}><h3>{title}</h3></a>
	<p>{description}</p>
	<div class="circle">
		<div class="inner-circle">
			<Image src={image} alt={title} />
		</div>
	</div>
</div>

<style>
	.category {
		display: flex;
		flex-direction: column;
		background: var(--accent);
		color: var(--primary);
		gap: 10px;
		aspect-ratio: 2 / 1;
		height: calc(100% - 20px * 2);
		padding: 20px;
		box-shadow: 10px 10px 0px 0px var(--secondary);
		justify-content: center;
		position: relative;
		overflow: hidden;

		a {
			color: var(--primary);
		}

		p {
			color: var(--secondary);
		}

		.circle {
			position: absolute;
			background: var(--secondary);
			aspect-ratio: 1 / 1;
			border-radius: 50%;
			height: calc(100% - 10px * 2);
			right: -70px;
			bottom: -50px;
			overflow: hidden;
			padding: 10px;
			transition: 0.2s;

			.inner-circle {
				width: 100%;
				height: 100%;
				border-radius: 50%;
				overflow: hidden;
			}
		}
	}

	.category:hover {
		.circle {
			scale: 110%;
		}
	}
</style>
