package internal

func Map[TSource any, TResult any](s []TSource, mfn func(s TSource) TResult) []TResult {
	var result []TResult

	for _, v := range s {
		result = append(result, mfn(v))
	}

	return result
}
