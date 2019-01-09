package net

// // ComposeError - creates an err.CompositeError object inside the request's context for collecting all the errors
// func ComposeError(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		r = r.WithContext(ctx)
// 		next.ServeHTTP(w, r)
// 	})
// }
