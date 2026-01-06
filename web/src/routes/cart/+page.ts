import { getCart } from '$lib/api/carts';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({}) => {
	const res = await getCart();
	return { items: res };
};
