<script lang="ts">
	import Button from './Button.svelte';
	import CheckInput from './CheckInput.svelte';
	import Icon from './Icon.svelte';
	import Price from './Price.svelte';

	interface CartFormProps {
		title: string;
		bodyshapePrice: number | undefined;
		pickupPrices: number[];
		colorPrice: number | undefined;
		servicePrice: number;
		publish: boolean;
		oncart: () => void;
	}

	let {
		title = $bindable(),
		publish = $bindable(),
		bodyshapePrice,
		pickupPrices,
		colorPrice,
		servicePrice,
		oncart = () => {}
	}: CartFormProps = $props();

	let general = $derived.by(() => {
		let v = 0;

		bodyshapePrice ? (v += bodyshapePrice) : null;
		colorPrice ? (v += colorPrice) : null;

		pickupPrices.forEach((p) => {
			v += p;
		});

		v += servicePrice;

		return v;
	});
</script>

<form class="cart-form" onsubmit={() => oncart()}>
	<section class="total">
		<h3>Podsumowanie</h3>
		{#if bodyshapePrice}
			<div class="block">
				<p>Kształt</p>
				<p><Price price={bodyshapePrice} /></p>
			</div>
		{/if}
		{#each pickupPrices as p}
			<div class="block">
				<p>Przetwornik</p>
				<p><Price price={p} /></p>
			</div>
		{/each}
		{#if colorPrice}
			<div class="block">
				<p>Pokolorowanie</p>
				<p><Price price={colorPrice} /></p>
			</div>
		{/if}
		<div class="block">
			<p>Usługi serwisowe</p>
			<p><Price price={servicePrice} /></p>
		</div>
		<div class="dashed"></div>
		<div class="block">
			<p>Wartość ogólna</p>
			<p><Price price={general} /></p>
		</div>
	</section>
	<section class="form">
		<h3>Informacja</h3>
		<div class="block">
			<p>Nazwa</p>
			<input bind:value={title} maxlength={256} type="text" />
		</div>
		<div class="publish">
			<CheckInput title="Chcę opublikować gitarę w katalogu" bind:checked={publish} />
		</div>
	</section>
	<Button formtype={'submit'} onclick={() => {}}>
		<Icon name={'cart'} alt height={24} />
		Do koszyka
	</Button>
</form>

<style>
	.cart-form {
		background: var(--accent);
		color: var(--primary);
		padding: 10px;
		display: flex;
		flex-direction: column;
		height: fit-content;

		section {
			padding: 10px;

			.block {
				padding: 10px;
			}
		}

		.block {
			display: flex;
		}

		.total {
			.block {
				justify-content: space-between;
			}

			.dashed {
				border-top: 1px dashed var(--primary);
			}
		}

		.form {
			.block {
				flex-direction: column;
				gap: 10px;
			}

			.publish {
				margin-top: 8px;
				margin-bottom: 8px;
			}
		}

		input {
			background: var(--primary);
			width: calc(100% - 8px * 2);
			height: 20px;
			font-size: 16px;
			padding: 8px;
		}
	}
</style>
