package rest

import (
	"fmt"
	"github.com/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("about to start test cases...")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups() // remove any mock up we might have from previous tests
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "https://api.bookstore.com/users/login",
		ReqBody: `{"email":"email@gmail.com","password":"password"}`,
		RespHTTPCode: -1,
		RespBody: `{}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid restclient response trying to login user", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups() // remove any mock up we might have from previous tests
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "https://api.bookstore.com/users/login",
		ReqBody: `{"email":"email@gmail.com","password":"password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody: `{"message":"invalid login credentials", "status":"404", "error":"not_found"}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface trying to login user", err.Message)
}

func TestLoginUserInvalidloginCredentials(t *testing.T) {
	rest.FlushMockups() // remove any mock up we might have from previous tests
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "https://api.bookstore.com/users/login",
		ReqBody: `{"email":"email@gmail.com","password":"password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody: `{"message":"invalid login credentials", "status":404, "error":"not_found"}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "invalid login credentials", err.Message)
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	rest.FlushMockups() // remove any mock up we might have from previous tests
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "https://api.bookstore.com/users/login",
		ReqBody: `{"email":"email@gmail.com","password":"password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody: `{"id":"1", "first_name":"eric", "last_name":"ha", "email":"eric.ha@gmail.com"}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error trying to unmarshal users login response", err.Message)
}

func TestLoginUserNoError(t *testing.T)	{
	rest.FlushMockups() // remove any mock up we might have from previous tests
	rest.AddMockups(&rest.Mock{
		HTTPMethod: http.MethodPost,
		URL: "https://api.bookstore.com/users/login",
		ReqBody: `{"email":"email@gmail.com","password":"password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody: `{"id":1, "first_name":"eric", "last_name":"ha", "email":"eric.ha@gmail.com"}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("email@gmail.com", "password")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "eric", user.FirstName)
	assert.EqualValues(t, "ha", user.LastName)
	assert.EqualValues(t, "eric.ha@gmail.com", user.Email)
}



