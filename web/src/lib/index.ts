import { env } from '$env/dynamic/public';

export async function request(
	url: string,
	method: 'GET' | 'POST' | 'PATCH' | 'DELETE' = 'GET',
	body: string | null = null
) {
	const res = await fetch(env.PUBLIC_SERVER + url, {
		method: method,
		headers: body
			? {
					'Content-Type': 'application/json'
				}
			: {},
		body: body
	});

	// const text = await res.json();
	//
	console.log(res);

	// createNotify({
	// 	message: text,
	// 	status: res.status
	// });

	return res;
}
