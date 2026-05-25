package handler

import (
	_ "auth-service/internal/models"
	"auth-service/internal/service"
	"encoding/json"
	"net/http"
)
// GetMe godoc
//
//	@Summary		Get current user
//	@Description	Get information about authorized user
//	@Tags			users
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{object}	models.UserResponse
//	@Failure		401	{string}	string
//	@Failure		500	{string}	string
//	@Router			/me [get]
func GetMe(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("user_id").(string)

	user, err := service.GetUserByID(userID)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}