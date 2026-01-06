<script lang="ts">
	import { page } from '$app/state';
</script>

<main>
	<div class="error">
		<h1>{page.status}</h1>
		<p>{page.error?.message}</p>
	</div>
</main>

<style>
	main {
		padding: 100px 100px;
		display: flex;
		justify-content: center;
		align-items: center;
		height: calc(100vh - 100px * 2);

		.error {
			display: flex;
			flex-direction: column;
			gap: 10px;
			align-items: center;

			h1 {
				font-size: 128px;
			}

			p {
				font-size: 20px;
			}
		}
	}
</style>
