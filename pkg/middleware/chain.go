package middleware

import (
	"net/http"
)

// Middleware - функция, обрабатывающая HTTP-запросы
type Middleware func(http.Handler) http.Handler

// Chain объединяет несколько middleware в один обработчик
func Chain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		// Применяем middleware в обратном порядке
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}
