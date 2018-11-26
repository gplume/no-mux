package middle

import (
	"log"
	"net/http"
	"time"
)

// Notify ...
func Notify(logger *log.Logger) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Println(">>>>>>>>>>>> Notify() before")
			defer func(begin time.Time) {
				logger.Printf("<<<<<<<<<<<< Notify() after %vs", time.Since(begin).Seconds())
			}(time.Now())
			h.ServeHTTP(w, r)
		})
	}
}
