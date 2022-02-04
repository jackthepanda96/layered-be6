package user

import (
	"Project/playground/be6/layered/configs"
	"Project/playground/be6/layered/entities"
	"Project/playground/be6/layered/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestInsert(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	repo := New(db)

	t.Run("Success Create User", func(t *testing.T) {
		mockUser := entities.User{Nama: "Steven", Email: "steven@steven.com", Password: "steven123"}
		res, err := repo.Insert(mockUser)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(res.ID))
	})

	t.Run("Fail Create User", func(t *testing.T) {
		mockUser := entities.User{Model: gorm.Model{ID: 1}, Nama: "Steven", Email: "steven@steven.com", Password: "steven123"}
		_, err := repo.Insert(mockUser)
		assert.NotNil(t, err)
	})
}

func TestGet(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	repo := New(db)
	t.Run("Success Get User", func(t *testing.T) {
		mockUser := entities.User{Nama: "Steven", Email: "steven@steven.com", Password: "steven123"}
		res, err := repo.Insert(mockUser)
		assert.Nil(t, err)
		assert.Equal(t, 1, int(res.ID))
	})
}
