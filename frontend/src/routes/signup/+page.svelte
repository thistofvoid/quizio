<script>
	import { pb } from '$lib/pocketbase';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { redirect } from '@sveltejs/kit';
	import { auth } from '$lib/auth.svelte';

	let email = $state('');
	let password = $state('');
	let passwordConfirm = $state('');
	let error = $state('');
	let loading = $state(false);

	onMount(() => {
		if (auth.isValid) {
			goto('/');
		}
	});

	async function signup() {
		error = '';

		// client-side check for fast feedback (not security)
		if (password !== passwordConfirm) {
			error = 'Passwords do not match.';
			return;
		}
		if (password.length < 8) {
			error = 'Password must be at least 8 characters.';
			return;
		}

		loading = true;
		try {
			// 1. create the user record
			await pb.collection('users').create({
				email,
				password,
				passwordConfirm
			});

			// 2. log them in immediately
			await pb.collection('users').authWithPassword(email, password);

			// 3. optionally trigger a verification email
			await pb.collection('users').requestVerification(email);

			goto('/');
		} catch (err) {
			// PocketBase returns structured field errors
			// @ts-ignore
			error = err?.response?.message || 'Signup failed.';
		} finally {
			loading = false;
		}
	}
</script>

<div class="flex items-center justify-center p-2">
	<div class="flex w-1/2 flex-col gap-2 rounded border p-2">
		<div class="flex flex-col">
			<label for="email">Email</label>
			<input
				class="input"
				placeholder="test@mail.com"
				bind:value={email}
				name="email"
				type="email"
			/>
		</div>
		<div class="flex flex-col">
			<label for="password">Password</label>
			<input
				class="input"
				placeholder="Password"
				bind:value={password}
				name="password"
				type="password"
			/>
		</div>
		<div class="flex flex-col">
			<label for="confirm_password">Confirm Password</label>
			<input
				class="input"
				placeholder="Confirm Password"
				bind:value={passwordConfirm}
				name="confirm_password"
				type="password"
			/>
		</div>
		<button class="btn" onclick={signup} disabled={loading}
			>{loading ? 'Creating account…' : 'Sign up'}</button
		>
		<div>
			{#if error}<p>{error}</p>{/if}
		</div>
	</div>
</div>
