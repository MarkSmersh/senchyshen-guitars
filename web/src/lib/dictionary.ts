import type { Order, OrderBy, ProductDetails, ProductType } from './api/products';

export const OrderToName: Record<Order, string> = {
	asc: 'Rosnąca',
	desc: 'Malejąca'
};

export const OrderByToName: Record<OrderBy, string> = {
	title: 'Nazwa',
	price: 'Cena',
	createdAt: 'Data stworzenia'
};

export function ProductTypeToName(type: ProductType): string {
	switch (type) {
		case 'guitar':
			return 'Gitara';
		case 'pickup':
			return 'Przetwornik';
		case 'bodyshape':
			return 'Kstałt';
		case 'amplifier':
			return 'Wzmacniacz';
		case 'crafted':
			return 'Na zamówienie';
	}
}

export const ProductDetailToName: Record<ProductDetails, string> = {
	type: 'Typ',
	height: 'Wysokość',
	width: 'Szerokość',
	bodyshape: 'Kstałt',
	bodyshapeId: 'Identyfikator kstałtu',
	color: 'Kolor',
	material: 'Materiał',
	pickups: 'Przetworniki',
	pickupsCount: 'Ilość przetworników',
	power: 'Potężność',
	stringsCount: 'Ilość strun',
	thickness: 'Grubość'
};

// export function ProductDetailtToName(detail: string): string {
//
// }
