package http

// NewHTTPHandler ...
// func NewHTTPHandler(
// 	userService user.Service,
// 	authService auth.Service,
// 	logger kitlog.Logger,
// ) http.Handler {
// 	r := chi.NewRouter()
// 	cors := cors.New(cors.Options{
// 		AllowedOrigins:   []string{"*"},
// 		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
// 		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
// 		AllowCredentials: true,
// 	})
// 	r.Use(cors.Handler)
// 	r.Route("/v1", func(r chi.Router) {
// 		userH := userHandler{userService, logger}
// 		authH := authHandler{authService, logger}
// 		r.Mount("/users", userH.router())
// 		r.Mount("/auth", authH.router())
// 		// r.With(middleware.Authentication(userService)).Mount("/wallets", walletH.router())
// 	})

// 	return r
// }

// func encodeError(_ context.Context, err error, w http.ResponseWriter) {
// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	switch err {
// 	case domain.ErrUnknowUser:
// 		w.WriteHeader(http.StatusNotFound)
// 		json.NewEncoder(w).Encode(map[string]interface{}{
// 			"success": false,
// 			"message": err.Error(),
// 		})
// 		break
// 	default:
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(map[string]interface{}{
// 			"success": false,
// 			"message": err.Error(),
// 		})
// 	}
// }
