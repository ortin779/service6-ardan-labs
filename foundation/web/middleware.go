package web

type MidHandler func(handler Handler) Handler

func wrapMiddleware(mws []MidHandler, handler Handler) Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		mwFunc := mws[i]
		if mwFunc != nil {
			handler = mwFunc(handler)
		}
	}
	return handler
}
