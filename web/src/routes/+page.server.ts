import { getPreview } from '$lib/api/getPreview';
import { getProducts } from '$lib/api/products';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async () => {
	const preview = await getPreview();
	const res = await getProducts({ types: ['guitar'], limit: 3 });

	return {
		preview: preview,
		products: res?.products
	};
};
