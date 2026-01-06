<script lang="ts">
	import type { Snippet } from 'svelte';

	type ButtonStyle = 'primary' | 'secondary' | 'accent';

	interface ButtonProps {
		children: Snippet;
		type?: ButtonStyle;
		formtype?: 'submit' | 'reset' | 'button';
		onclick: () => void;
		onsubmit?: () => void;
	}

	const {
		children,
		type = 'primary',
		formtype,
		onclick,
		onsubmit = () => {}
	}: ButtonProps = $props();
</script>

<button type={formtype} {onsubmit} {onclick} class={type}>
	{@render children()}
</button>

<style>
	button {
		padding: 10px 20px;
		font-size: 16px;
		font-weight: bold;
		cursor: pointer;
		transition: 0.2s;
		color: var(--accent);
		display: flex;
		align-items: center;
		gap: 8px;
		justify-content: center;
	}

	.primary {
		background: var(--primary);
	}

	.primary:hover {
		background: var(--secondary);
	}

	.secondary {
		background: var(--secondary);
	}

	.secondary:hover {
		background: var(--primary);
	}

	.accent {
		background: var(--accent);
		color: var(--primary);
	}

	.accent:hover {
		background: var(--primary);
		color: var(--accent);
	}
</style>
