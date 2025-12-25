package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

// PrependMiddlewareChain prepends given middleware arguments in the order as passed to the
// function to a given http handler that will be called after the middleware chain.
func PrependMiddlewareChain(next http.Handler, middlewareArgs ...Middleware) http.Handler {
	if len(middlewareArgs) == 0 {
		return next
	}

	handler := next

	for i := len(middlewareArgs) - 1; i >= 0; i-- {
		handler = middlewareArgs[i](handler)
	}

	return handler
}
