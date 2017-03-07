package controllers

import (
	"github.com/itsshompi/kuubit-backend/models"
)

//Models for JSON resources
type (
	//UserResource is for Post - /user/register
	UserResource struct {
		Data models.User `json:"data"`
	}

	//LoginResource is for Post - /user/login
	LoginResource struct {
		Data LoginModel `json:"data"`
	}

	//AuthUserResource is  a response for authorized user Post - /user/login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}

	//LoginModel is a model for authentication
	LoginModel struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	//AuthUserModel is a model for authorized user with access token
	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}
)
