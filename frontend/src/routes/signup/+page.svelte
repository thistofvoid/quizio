<script>
	import { pb } from '$lib/pocketbase';
	import { goto } from '$app/navigation';

	let email = $state('');
	let password = $state('');
	let passwordConfirm = $state('');
	let error = $state('');
	let loading = $state(false);

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

<input bind:value={email} type="email" placeholder="Email" />
<input bind:value={password} type="password" placeholder="Password" />
<input bind:value={passwordConfirm} type="password" placeholder="Confirm password" />
<button onclick={signup} disabled={loading}>
	{loading ? 'Creating account…' : 'Sign up'}
</button>
{#if error}<p>{error}</p>{/if}
