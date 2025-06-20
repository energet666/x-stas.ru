<script lang="ts">
	import { Power, Save } from "@lucide/svelte";
	const host = window.location.hostname;
	const protocol = window.location.protocol;
	const apiHost = protocol + "//" + host + ":8080"; //для отладки или если статика и бэкенд на разных хостах
	// const apiHost = "http://127.0.0.1:8080";

	window.addEventListener("unload", () => {
		console.log("unload");
		navigator.sendBeacon(apiHost + "/api/shutdown");
	});

	let count = $state(0);

	let powerIsActive = $state(false);

	type Msg = {
		Count: number;
	};
	const fetchCount = async () => {
		await fetch(apiHost + "/api/count")
			.then((response) => response.json())
			.then((data: Msg) => {
				console.log("count: ", data.Count);
				count = data.Count;
			})
			.catch((error) => console.log(error));
	};
	fetchCount();
</script>

<main
	class="flex flex-col items-center justify-center h-screen select-none myrainbow-surface-50-950 myrainbow-w-5 rounded-2xl"
>
	<div class=" p-3 rounded-2xl text-center">
		<h1 class="text-primary-500 sm:text-9xl text-6xl">x-stas.ru</h1>
		<div class="text-2xl font-mono text-secondary-200-800">
			<p class="">Loads: {count}</p>
		</div>
	</div>
	<div
		class="my-block p-3 bg-primary-500 rounded-full flex gap-4 text-white/80"
	>
		<button
			class="my-button p-2 rounded-full inline-block {powerIsActive
				? 'my-button-active'
				: ''}"
			onclick={() => {
				powerIsActive = !powerIsActive;
			}}
		>
			<Power />
		</button>
		<div class="my-button p-2 rounded-full inline-block">
			<Save />
		</div>
		<!-- <input
			type="checkbox"
			class="my-checkbox"
		/> -->
	</div>
</main>

<style lang="postcss">
	.my-button {
		cursor: pointer;
		/* border: solid 2px transparent; */
		transition: all 0.1s ease-in-out;
		background-image: linear-gradient(to left top, #00000030, #ffffff30);
		/* background-repeat: repeat; */
		background-size: 125% 125%;
		background-position: center;
		border: solid 3px transparent;
		/* background-clip: border-box; */
		box-shadow:
			4px 4px 7px 0px rgb(0 0 0 / 25%),
			-4px -4px 7px 0px rgb(255 255 255 / 25%);
	}
	.my-block {
		box-shadow:
			inset 4px 4px 7px 0px rgb(0 0 0 / 25%),
			inset -4px -4px 7px 0px rgb(255 255 255 / 25%);
	}
	.my-button-active {
		box-shadow:
			inset 2px 1px 0px 0px rgb(0 0 0 / 0%),
			inset -3px -3px 5px 0px rgb(255 255 255 / 26%),
			inset 3px 3px 8px -3px #0000007d;
		background-image: none;
		color: greenyellow;
		border-color: greenyellow;
		background-clip: padding-box;
	}
</style>
