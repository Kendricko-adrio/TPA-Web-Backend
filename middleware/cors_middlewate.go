package middleware

import "net/http"

func CorsMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")https://604c4b7ea9112e6d754fa318--suspicious-engelbart-00f72e.netlify.app/
		w.Header().Set("Access-Control-Allow-Origin", "https://604c4b7ea9112e6d754fa318--suspicious-engelbart-00f72e.netlify.app")
		//w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8080")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")

		next.ServeHTTP(w, r)
	})

}
