package mapper

import (
	"github.com/NaufalParamaRafif/CICO-Learn-Backend/domain"
	"github.com/NaufalParamaRafif/CICO-Learn-Backend/dto/response"
	"github.com/NaufalParamaRafif/CICO-Learn-Backend/internal/utils"
)

func ToUserDataResponse(user *domain.User) response.UserDataResponse {
	var dailyReminder []string

	for _, time := range user.DailyReminder {
		dailyReminder = append(dailyReminder, (string(time.Hour) + ":" + string(time.Minute)))
	}

	return response.UserDataResponse{
		ID:                 user.ID,
		FirstName:          utils.StringFromPtr(user.FirstName),
		LastName:           utils.StringFromPtr(user.LastName),
		Username:           user.Username,
		Email:              user.Email,
		Address:            utils.StringFromPtr(user.Address),
		Gender:             string(user.Gender),
		ProfilePictureURL:  utils.StringFromPtr(user.ProfilePicture),
		BackgroundImageURL: utils.StringFromPtr(user.BackgroundImage),
		Level:              string(user.Level),
		Language:           string(user.Language),
		Target:             string(user.Target),
		DailyReminder:      dailyReminder,
		TimeTarget:         utils.StringFromPtr((*string)(user.TimeTarget)),
	}
}
