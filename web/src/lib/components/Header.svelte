<script lang="ts">
	import logo from '$lib/assets/logo.svg';
	import Icon from './Icon.svelte';
</script>

<header>
	<a href="/" class="logo">
		<img src={logo} alt="logo" />
	</a>
	<div class="icons">
		<a href="/constructor">
			<Icon height={47} name={'constructor'} />
		</a>
		<a href="/catalog">
			<Icon height={45} name={'catalog'} />
		</a>
		<a href="/categories">
			<Icon height={47} name={'categories'} />
		</a>
		<a href="/cart">
			<Icon height={43} name={'cart'} />
		</a>
	</div>
</header>

<style>
	header {
		background-color: var(--accent);
		height: calc(80px - 10px * 2);
		display: flex;
		align-items: center;
		justify-content: space-between;
		padding: 10px;
		position: fixed;
		width: calc(100vw - 20px);
		z-index: 69;
		top: 0;

		img {
			height: 100%;
		}

		a {
			background: transparent;
			border: 0;
			cursor: pointer;
			height: 45px;
			transition: 0.2s;
		}

		a:hover {
			scale: 105%;
		}

		.logo {
			height: 100%;
		}

		.icons {
			display: flex;
			gap: 8px;
		}
	}
</style>
