package web

func wrapMiddleware(handler Handler, middlewares ...Middleware) Handler {

	for _, m := range middlewares {
		if m != nil {
			handler = m(handler)
		}
	}

	return handler
}
