import { writable } from 'svelte/store';
import { v4 } from 'uuid';

export const notify = writable<INotify[]>([]);

export function createNotify(body: Omit<INotify, 'id'>) {
	const id = v4();

	notify.update((n) => [
		...n,
		{
			id: id,
			message: body.message,
			status: body.status
		}
	]);

	setTimeout(() => {
		notify.update((n) => n.filter((v) => v.id !== id));
	}, 5000);
}

export interface INotify {
	id: string;
	status?: number;
	message: string;
}
