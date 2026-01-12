<script lang="ts">
	interface DropdownListProps {
		items: string[];
		item: string;
		title: string;
		onupdate?: (item: string) => void;
		translateWith?: (value: string) => string;
	}

	let {
		items,
		title,
		item = $bindable(),
		onupdate = () => {},
		translateWith = (v) => v
	}: DropdownListProps = $props();
</script>

<div class="dropdown">
	<p>{title}</p>
	<select bind:value={item} onchange={() => onupdate(item)}>
		{#each items as i}
			<option value={i}>{translateWith(i)}</option>
		{/each}
	</select>
</div>

<style>
	.dropdown {
		text-align: inherit;
		text-wrap: nowrap;
	}

	p {
		color: var(--primary);
		margin-bottom: 5px;
	}

	select {
		background: var(--primary);
		padding: 4px 8px;
		color: var(--accent);
		font-size: 14px;
		height: 30px;
	}
</style>
