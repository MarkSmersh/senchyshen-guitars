<script lang="ts">
	import Button from './Button.svelte';
	import Icon from './Icon.svelte';
	import Price from './Price.svelte';

	interface CartFormProps {
		total: number;
		phoneNumber: string;
		comment: string;
		onorder: () => void;
	}

	let {
		total,
		phoneNumber = $bindable(),
		comment = $bindable(),
		onorder = () => {}
	}: CartFormProps = $props();
</script>

<form class="cart-form" onsubmit={() => onorder()}>
	<section class="total">
		<h3>Podsumowanie</h3>
		<div class="block">
			<p>Wartość produktów</p>
			<p><Price price={total} /></p>
		</div>
	</section>
	<section class="form">
		<h3>Formularz</h3>
		<div class="block">
			<p>Numer telefonu (obowiązkowy)</p>
			<input bind:value={phoneNumber} required pattern={'(\\+?48)?([0-9]{9})'} type="tel" />
		</div>
		<div class="block">
			<p>Komentarz</p>
			<input bind:value={comment} maxlength={256} type="text" />
		</div>
	</section>
	<Button formtype={'submit'} onclick={() => {}}>
		<Icon name={'order'} alt height={24} />
		Złożyc zamówienie
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
		}

		.form {
			.block {
				flex-direction: column;
				gap: 10px;
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
