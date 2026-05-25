package handler

import (
	"auth-service/internal/models"
	"auth-service/internal/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)
// GetAllUsers godoc
//
//	@Summary		Get all users
//	@Description	Only for super admins
//	@Tags			admin
//	@Security		BearerAuth
//	@Produce		json
//	@Success		200	{array}		models.UserResponse
//	@Failure		401	{string}	string
//	@Failure		403	{string}	string
//	@Failure		500	{string}	string
//	@Router			/admin/users [get]
func GetAllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := service.GetAllUsers()

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}
// DeleteUser godoc
//
//	@Summary		Delete user
//	@Description	Delete user by ID
//	@Tags			admin
//	@Security		BearerAuth
//	@Produce		plain
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{string}	string	"User deleted"
//	@Failure		401	{string}	string
//	@Failure		403	{string}	string
//	@Failure		500	{string}	string
//	@Router			/admin/users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	userID := vars["id"]

	err := service.DeleteUser(userID)

if err != nil {

	http.Error(
		w,
		err.Error(),
		http.StatusInternalServerError,
	)

	return
}

	adminID := r.Context().Value("user_id").(string)

	service.CreateAuditLog(models.AuditLog{
		UserID: adminID,
		Action: "Deleted user " + userID,
	})

	w.Write([]byte("User deleted"))
}