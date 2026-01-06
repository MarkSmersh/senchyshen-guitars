import { request } from '$lib';
import type { ProductType } from './products';

export interface CartItemModel {
	id: number;
	count: number;
	total: number;
	price: number;
	title: string;
	description: string;
	image: string;
	type: ProductType;
}

export async function getCart(): Promise<CartItemModel[]> {
	const res = await request('/api/carts/');
	const body: CartItemModel[] = await res.json();
	return body;
}

export async function clearCart() {
	const res = await request('/api/carts/', 'PUT');
}

export async function removeCartItem(id: number) {
	const res = await request('/api/carts/' + id, 'DELETE');
}

interface AddCartItem {
	productId: number;
	count: number;
}

export async function addCartItem(item: AddCartItem) {
	const res = await request('/api/carts/', 'POST', JSON.stringify(item));
}
