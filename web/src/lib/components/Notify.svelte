<script lang="ts">
	import type { INotify } from '$lib/notify';
	import { notify } from '$lib/notify';

	import Icon from './Icon.svelte';

	let notifications = $state<INotify[]>([]);

	notify.subscribe((n) => (notifications = n));
</script>

<div class="notify">
	{#each notifications as n}
		<!-- svelte-ignore a11y_click_events_have_key_events, a11y_no_static_element_interactions (because of reasons) -->
		<div
			onclick={() => notify.update((notify) => notify.filter((v) => v.id !== n.id))}
			class="notification"
		>
			<div class="icon">
				{#if n.status && n.status > 399}
					<!-- <Icon name="fa-skull" /> -->
					bad icon
				{:else}
					<!-- <Icon name="fa-check" /> -->
					good icon
				{/if}
			</div>
			<div class="text">
				<p>{n.message}</p>
			</div>
		</div>
	{/each}
</div>
