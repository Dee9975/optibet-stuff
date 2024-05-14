package controllers

import "webnams-backend/store"

type UserController struct {
	store store.UserStore
}

func NewUserController(store store.UserStore) *UserController {
	return &UserController{store}
}
