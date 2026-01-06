import { getCategories } from '$lib/api/categories';
import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load: PageLoad = async () => {
	const res = await getCategories();

	if (!res) {
		error(404, 'Nie znajdziono');
	}

	return { categories: res };
};
