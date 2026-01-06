<script lang="ts">
	import Icon from './Icon.svelte';

	const { checked }: { checked: boolean } = $props();
</script>

{#if checked}
	<Icon name="yes" height={20} />
{:else}
	<Icon name="no" height={20} />
{/if}

<style>
</style>
