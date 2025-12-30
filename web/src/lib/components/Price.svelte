<script lang="ts">
	const { price }: { price: number } = $props();

	let converted = $derived(price / 100);
</script>

{converted.toFixed(2)} zł
