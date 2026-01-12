<script lang="ts">
	import Checkbox from './Checkbox.svelte';

	interface CheckInputProps {
		title: string;
		checked: boolean;
	}

	let { title, checked = $bindable() }: CheckInputProps = $props();
</script>

<div class="check-input">
	<button type="button" onclick={() => (checked = !checked)}>
		<Checkbox {checked} />
	</button>
	<p>{title}</p>
</div>

<style>
	.check-input {
		display: flex;
		gap: 10px;
		align-items: center;

		button {
			display: flex;
		}
	}
</style>
