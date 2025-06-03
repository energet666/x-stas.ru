<script lang="ts">
	//const host = window.location.hostname;
	// const apiHost = "http://" + host + ":8080"; //для отладки или если статика и бэкенд на разных хостах
	const apiHost = "";

	window.addEventListener("unload", () => {
		console.log("unload");
		navigator.sendBeacon(apiHost + "/api/shutdown");
	});

	let count = $state(0);

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
</main>

<style lang="postcss">
</style>
