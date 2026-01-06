import { request } from '$lib';

export type ProductType = 'guitar' | 'pickup' | 'bodyshape' | 'amplifier' | 'crafted';

export interface ProductSearch {
	page?: number;
	limit?: number;
	types?: ProductType[];
	priceMin?: number;
	priceMax?: number;
	category?: number;
	id?: number;
	// title, price, createdAt
	orderBy?: OrderBy;
	// desc | asc
	order?: 'desc' | 'asc';
	query?: string;
}

export type OrderBy = 'title' | 'price' | 'createdAt';
export type Order = 'desc' | 'asc';

export interface ProductSearchRes {
	products: ProductModel[];
	categories: Category[];
	types?: string[];
	priceMin: number;
	priceMax: number;
}

export async function getProducts(p: ProductSearch): Promise<ProductSearchRes | undefined> {
	const res = await request('/api/products/', 'POST', JSON.stringify(p));

	if (res.status > 399) {
		const resCopy = res.clone();
		const text = await resCopy.text();
		console.log(text);
		return undefined;
	}

	const body: ProductSearchRes = await res.json();

	return body;
}

export async function getProduct(id: number): Promise<ProductModel | undefined> {
	const res = await request('/api/products/' + id);

	if (res.status > 399) {
		const resCopy = res.clone();
		const text = await resCopy.text();
		console.log(text);
		return undefined;
	}

	const body: ProductModel = await res.json();

	return body;
}

export interface Category {
	id: number;
	title: string;
	description: string;
	image?: string;
}

export type ProductDetails =
	| keyof GuitarModel
	| keyof PickupModel
	| keyof AmplifierModel
	| keyof CraftedModel
	| keyof BodyshapeModel;

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
