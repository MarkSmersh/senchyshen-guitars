import { request } from '$lib';

export type ProductType = 'guitar' | 'pickup' | 'bodyshape' | 'amplifier' | 'crafted';

interface ProductSearch {
	page?: number;
	limit?: number;
	types?: ProductType[];
	priceMin?: number;
	priceMax?: number;
	category?: number;
	id?: number;
	// title, price, createdAt
	orderBy?: 'title' | 'price' | 'createdAt';
	// desc | asc
	order?: 'desc' | 'asc';
	query?: string;
}

export async function getProducts(p: ProductSearch): Promise<ProductModel[]> {
	const res = await request('/api/products', 'POST', JSON.stringify(p));

	const body: ProductModel[] = await res.json();

	return body;
}

export interface GuitarModel {
	stringsCount: number;
	bodyshape: string;
	color: string;
	pickupsCount: number;
}

export interface PickupModel {
	type: string;
}

export interface AmplifierModel {
	power: number;
}

export interface Pickup {
	pickupId: number;
	position: string;
}

export interface CraftedModel {
	bodyshapeId: number;
	color: string;
	pickups: Pickup[];
}

export interface BodyshapeModel {
	color: string;
	material: string;
	height: number;
	width: number;
	thickness: number;
	pickupsCount: number;
}

// Guess — change to match your Go type if needed
export interface ImageModel {
	id: number;
	path: string;
}

export interface ProductModel {
	id: number;
	title: string;
	description: string;
	type: ProductType;
	price: number;

	category?: number; // json:"category,omitempty"
	categoryTitle?: string; // json:"categoryTitle,omitempty"

	createdAt: number;
	images: ImageModel[];

	guitar?: GuitarModel;
	pickup?: PickupModel;
	bodyshape?: BodyshapeModel;
	amplifier?: AmplifierModel;
	crafted?: CraftedModel;
}
