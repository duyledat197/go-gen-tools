package http

// type authHandler struct {
// 	authService auth.Service
// 	logger      kitlog.Logger
// }

// func (h *authHandler) router() chi.Router {
// 	r := chi.NewRouter()
// 	r.Post("/register", h.register)
// 	r.Post("/login", h.login)
// 	r.Post("/forgot-password", h.forgotPassword)
// 	r.Post("/reset-password", h.resetPassword)

// 	return r
// }

// func (h *authHandler) register(w http.ResponseWriter, r *http.Request) {
// 	ctx := context.Background()
// 	var req request.Register

// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		h.logger.Log("error", err)
// 		encodeError(ctx, err, w)
// 		return
// 	}

// 	u := domain.User{
// 		Email: req.Email,
// 	}

// 	err := h.authService.Register(&u, req.Password)

// 	if err != nil {
// 		encodeError(ctx, err, w)
// 		return
// 	}

// 	res := response.Register{
// 		Success: true,
// 	}

// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	if err = json.NewEncoder(w).Encode(res); err != nil {
// 		h.logger.Log("error", err)
// 		encodeError(ctx, err, w)
// 		return
// 	}
// }

// func (h *authHandler) login(w http.ResponseWriter, r *http.Request) {
// 	ctx := context.Background()

// 	var req request.Login
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		h.logger.Log("error", err)
// 		encodeError(ctx, err, w)
// 		return
// 	}
// 	user, token, err := h.authService.Login(req.Email, req.Password)
// 	// if user exist error equal null
// 	if err != nil {
// 		encodeError(ctx, err, w)
// 		return
// 	}

// 	res := response.Login{
// 		User:    user,
// 		Token:   token,
// 		Success: true,
// 	}

// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	if err = json.NewEncoder(w).Encode(res); err != nil {
// 		h.logger.Log("error", err)
// 		encodeError(ctx, err, w)
// 		return
// 	}
// }

// func (h *authHandler) forgotPassword(w http.ResponseWriter, r *http.Request) {

// }

// func (h *authHandler) resetPassword(w http.ResponseWriter, r *http.Request) {

// }
