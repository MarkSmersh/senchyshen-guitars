import { getProducts, type ProductType } from '$lib/api/products';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ url }) => {
	const category = url.searchParams.get('category');
	const type = url.searchParams.get('type') as ProductType;

	if (category) {
		const res = await getProducts({ limit: 25, category: parseInt(category) });
		return res;
	}

	if (type) {
		const res = await getProducts({ limit: 25, types: [type] });
		return res;
	}

	const res = getProducts({ limit: 25 });
	return res;
};
