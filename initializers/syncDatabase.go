package initializers

import (
	"github.com/NaufalParamaRafif/CICO-Learn-Backend/models"
)

func SyncDatabase()  {
	DB.AutoMigrate(&models.User{})
}