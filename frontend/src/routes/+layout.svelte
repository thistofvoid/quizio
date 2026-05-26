<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';
	import { pb } from '$lib/pocketbase';
	import { auth } from '$lib/auth.svelte';

	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>
<div>
	<div class="navbar bg-base-100 shadow-sm">
		<div class="flex-1">
			<a class="btn text-xl btn-ghost">QUIZIO</a>
		</div>
		<div class="flex-none">
			<ul class="menu menu-horizontal px-1">
				<li><a href="/">Home</a></li>
				{#if !auth.isValid}
					<li>
						<details>
							<summary>Authentificate</summary>
							<ul class="rounded-t-none bg-base-100 p-2">
								<li><a href="/signin">Sign In</a></li>
								<li><a href="/signup">Sign Up</a></li>
							</ul>
						</details>
					</li>
				{:else}
					<li><a href="/" onclick={() => pb.authStore.clear()}>Log Out</a></li>
				{/if}
			</ul>
		</div>
	</div>
	<div>
		{@render children()}
	</div>
</div>
