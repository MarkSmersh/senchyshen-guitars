import { request } from '$lib';

interface CreateOrder {
	tel: string;
	comment?: string;
}

export async function createOrder(req: CreateOrder): Promise<string> {
	const res = await request('/api/orders/', 'POST', JSON.stringify(req));
	const order = await res.text();
	return order;
}

export interface OrderItemModel {
	productId: number;
	title: string;
	description: string;
	image: string;
	price: number;
	count: number;
}

export interface OrderModel {
	uuid: string;
	tel: string;
	comment: string;
	items: OrderItemModel[];
}

export async function getOrder(uuid: string): Promise<OrderModel> {
	const res = await request('/api/orders/' + uuid);
	const order: OrderModel = await res.json();
	return order;
}
