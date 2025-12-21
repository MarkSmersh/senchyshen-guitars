import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params }) => {
	const text = await (await fetch('http://localhost:1488/api/healthy')).text();

	return { text: text };
};
