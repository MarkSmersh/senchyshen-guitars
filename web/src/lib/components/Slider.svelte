<script lang="ts">
	import Price from './Price.svelte';

	interface SliderProps {
		from: number;
		to: number;
		title?: string;
		onupdate: (min: number, max: number) => void;
	}

	let { from, to, title, onupdate }: SliderProps = $props();

	let min = $derived(from);
	let max = $derived(to);
	let leftHolding = $state(false);
	let rightHolding = $state(false);

	let left: HTMLDivElement | undefined = $state();
	let right: HTMLDivElement | undefined = $state();
	let slider: HTMLDivElement | undefined = $state();

	let seed = $state(0);

	function move(e: MouseEvent) {
		if (slider && left && right) {
			const parent = slider.getBoundingClientRect();

			const end = Math.ceil(parent.right - parent.left - right.clientWidth - 1);
			const range = right.offsetLeft - left.offsetLeft - right.clientWidth;
			const x = e.clientX - parent.left - 13;

			// counting from the right side of left anchor and from the left side of the right anchor
			const size = end - left.offsetWidth;
			const rightAnchor = end - right.offsetLeft;
			const leftAnchor = left.offsetLeft;
			const perPixel = (to - from) / size;

			min = from + perPixel * leftAnchor;
			max = to - perPixel * rightAnchor;

			if (leftHolding) {
				if (left.offsetLeft < 0 || x < 0) {
					left.style.left = 0 + 'px';
				} else if (range < 0 || x + right.clientWidth > right.offsetLeft) {
					left.style.left = right.offsetLeft - right.clientWidth + 'px';
				} else {
					left.style.left = x + 'px';
				}

				update();
			}

			if (rightHolding) {
				if (right.offsetLeft > end || x > end) {
					right.style.left = end + 'px';
					return;
				} else if (range < 0 || x - right.clientWidth < left.offsetLeft) {
					right.style.left = left.offsetLeft + left.clientWidth + 'px';
				} else {
					right.style.left = x + 'px';
				}

				update();
			}
		}
	}

	// the seed represents the last request to send a data to a parent component.
	function update() {
		const currentSeed = Math.round(Math.random() * 10000000);
		seed = currentSeed;
		setTimeout(() => {
			if (seed == currentSeed) {
				onupdate(min, max);
			}
		}, 700);
	}

	function clearHoldings() {
		leftHolding = false;
		rightHolding = false;
	}

	// setInterval(() => (sleep = false), 350);
</script>

<!-- svelte-ignore a11y_click_events_have_key_events, a11y_no_static_element_interactions (because of reasons) -->
<div
	onmouseup={() => clearHoldings()}
	onmousemove={(e) => {
		move(e);
	}}
	onmouseleave={() => clearHoldings()}
	class="container"
>
	{title}
	<div class="core">
		<div class="slider" bind:this={slider}>
			<div
				onmousedown={() => (leftHolding = true)}
				bind:this={left}
				class="pivot"
				id="pivot-1"
			></div>
			<div class="line"></div>
			<div
				onmousedown={() => (rightHolding = true)}
				bind:this={right}
				class="pivot"
				id="pivot-2"
			></div>
		</div>
		<div class="legend">
			<p>
				<Price price={min} />
			</p>
			<p>
				<Price price={max} />
			</p>
		</div>
	</div>
</div>

<style>
	.container {
		display: flex;
		flex-direction: column;
		gap: 8px;

		.core {
			background: var(--accent);
			display: flex;
			flex-direction: column;
			padding: 15px;
			gap: 4px;
			position: relative;

			.slider {
				height: 25px;
				position: relative;
				align-items: center;
				display: flex;
				/* border: 1px solid red; */
				width: 100%;

				.pivot {
					background: var(--secondary);
					height: 25px;
					width: 25px;
					position: absolute;
					cursor: e-resize;
				}

				#pivot-2 {
					right: 0;
				}

				.line {
					background: var(--primary);
					height: 5px;
					width: 100%;
				}
			}

			.legend {
				display: flex;
				width: 100%;
				justify-content: space-between;
				color: var(--primary);
			}
		}
	}
</style>
