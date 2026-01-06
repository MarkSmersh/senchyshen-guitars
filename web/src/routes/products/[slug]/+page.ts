import { getProduct } from '$lib/api/products';
import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params }) => {
	const slug = params.slug;

	const res = await getProduct(parseInt(slug));

	if (!res) {
		error(404, 'Nie znajdziono');
	}

	return res;
};
