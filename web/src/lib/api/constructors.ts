import { request } from '$lib';

export interface Pickup {
	pickupId: number;
	position: string;
}

export interface ConstructorCreate {
	title?: string;
	bodyshapeId: number;
	pickups: Pickup[];
	color?: string;
	publish?: boolean;
}

export async function constructGuitar(params: ConstructorCreate) {
	const res = request('/api/constructors/', 'POST', JSON.stringify(params));
}
