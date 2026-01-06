import { getOrder } from '$lib/api/orders';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params }) => {
	const slug = params.slug;

	const order = await getOrder(slug);
	return order;
};
