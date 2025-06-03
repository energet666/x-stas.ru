<script lang="ts">
	const host = window.location.hostname;
	const apiHost = "http://" + host + ":8080"; //на этом порту гошный сервер

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

<main class="flex flex-col items-center justify-center h-screen">
	<h1 class="text-8xl text-primary-500">x-stas.ru</h1>
	<h2 class="text-2xl font-mono text-surface-contrast-500/50">
		Количество посещений: {count}
	</h2>
</main>

<style lang="postcss">
</style>
