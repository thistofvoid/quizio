import { pb } from './pocketbase';

export const auth = $state({
    user: pb.authStore.record,
    isValid: pb.authStore.isValid
});

pb.authStore.onChange(() => {
    auth.user = pb.authStore.record;
    auth.isValid = pb.authStore.isValid;
});