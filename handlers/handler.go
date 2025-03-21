package handlers

import "user-api/repository"

var repo repository.Repository

func InitHandlers(r repository.Repository) {
	repo = r
}
