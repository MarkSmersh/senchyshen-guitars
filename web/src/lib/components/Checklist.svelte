<script lang="ts">
	import type { Component } from 'svelte';
	import Checkbox from './Checkbox.svelte';

	export interface ChecklistItem {
		title: string;
		checked: boolean;
	}

	interface ChecklistProps {
		items: ChecklistItem[];
		single?: boolean;
		title?: string;
		onupdate: () => void;
		translateWith?: (v: string) => string;
	}

	let {
		items = $bindable(),
		title,
		onupdate,
		single = false,
		translateWith = (v) => v
	}: ChecklistProps = $props();

	function update(i: ChecklistItem) {
		if (single) {
			items.forEach((item) => (item.title == i.title ? null : (item.checked = false)));
		}

		i.checked = !i.checked;

		onupdate();
	}
</script>

<div class="checklist">
	{title}
	<div class="items">
		{#each items as i}
			<button class="item" onclick={() => update(i)}>
				{#if single}
					{#if i.checked}
						<Checkbox checked={i.checked} />
					{/if}
				{:else}
					<Checkbox checked={i.checked} />
				{/if}
				<!-- <Checkbox checked={i.checked} /> -->
				{translateWith(i.title)}
			</button>
		{/each}
	</div>
</div>

<style>
	.checklist {
		display: flex;
		flex-direction: column;
		gap: 8px;

		.items {
			display: flex;
			flex-direction: column;
			padding: 10px;
			gap: 4px;
			background: var(--accent);

			.item {
				display: flex;
				padding: 4px;
				gap: 8px;
				align-items: center;
				background: var(--accent);
				color: var(--primary);
				transition: 0.2s;
				cursor: pointer;
			}

			.item:hover {
				background: var(--secondary);
			}
		}
	}
</style>
