import { request } from '$lib';
import type { Category } from './products';

export async function getCategories(): Promise<Category[] | undefined> {
	const res = await request('/api/categories/');

	if (res.status > 399) {
		const resCopy = res.clone();
		const text = await resCopy.text();
		return undefined;
	}

	const body: Category[] = await res.json();
	return body;
}
