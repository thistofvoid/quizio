<script>
	import { pb } from '$lib/pocketbase';
	import { goto } from '$app/navigation';

	let email = $state('');
	let password = $state('');
	let error = $state('');

	async function login() {
		try {
			await pb.collection('users').authWithPassword(email, password);
			goto('/');
		} catch {
			error = 'Invalid email or password.';
		}
	}
</script>

<div class="flex items-center justify-center p-2">
	<div class="flex w-1/2 flex-col gap-2 rounded border p-2">
		<div class="flex flex-col">
			<label for="email">Email</label>
			<input class="input" bind:value={email} name="email" type="email" />
		</div>
		<div class="flex flex-col">
			<label for="password">Password</label>
			<input class="input" bind:value={email} name="password" type="password" />
		</div>
		<button class="btn" onclick={login}>Sign in</button>
		<div>
			{#if error}<p>{error}</p>{/if}
		</div>
	</div>
</div>
