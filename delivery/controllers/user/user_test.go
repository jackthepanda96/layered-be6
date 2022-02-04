package user

import (
	"Project/playground/be6/layered/entities"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	context := e.NewContext(req, res)
	context.SetPath("/users")

	userController := New(&MockUserRepository{})
	userController.Get()(context)

	response := GetUserResponseFormat{}

	json.Unmarshal([]byte(res.Body.Bytes()), &response)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "Adlan", response.Data[0].Nama)

}

type MockUserRepository struct{}

func (m MockUserRepository) Get() ([]entities.User, error) {
	return []entities.User{
		{Nama: "Adlan", Email: "adlan@adlan.com", Password: "adlan123"},
	}, nil
}

func (m *MockUserRepository) Insert(newUser entities.User) (entities.User, error) {
	return entities.User{Nama: "Arya", Email: "arya@arya.com", Password: "arya123"}, nil
}
