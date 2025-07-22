// Type of a successful async result
type Success<T> = {
	data: T;
	error: null;
};

// Type of a failed async result
type Failure<E> = {
	data: null;
	error: E;
};

type Result<T, E = Error> = Success<T> | Failure<E>;

/**  Utility funtion that handles error instead of try-catch */
export async function tryCatch<T, E = Error>(
	promise: Promise<T>,
): Promise<Result<T, E>> {
	try {
		const data = await promise;
		return {data, error: null};
	} catch (error) {
		return {data: null, error: error as E};
	}
}
