import { request } from '$lib';

interface PreviewData {
	bodyshapes: number;
	pickups: number;
}

export async function getPreview(): Promise<PreviewData | undefined> {
	const res = await request(`/api/preview`, 'GET');

	const body: PreviewData = await res.json();

	return body;
}
