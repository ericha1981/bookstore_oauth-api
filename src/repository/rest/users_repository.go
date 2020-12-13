package rest

import (
	"encoding/json"
	"github.com/ericha1981/bookstore_oauth-api/src/domain/users"
	"github.com/ericha1981/bookstore_oauth-api/src/utils/errors"
	"github.com/golang-restclient/rest"
	"time"
)

var usersRestClient = rest.RequestBuilder{
	BaseURL: "https://api.bookstore.com",
	Timeout: 100 * time.Millisecond,
}

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct {}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser (email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email: email,
		Password: password,
	}

	response := usersRestClient.Post("/users/login", request)

	if response == nil || response.Response == nil { // Timeout
		return nil, errors.NewInternalServerError("invalid restclient response trying to login user")
	}
	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface trying to login user")
		}
		return nil, &restErr
	}
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error trying to unmarshal users login response")
	}
	return &user, nil
}