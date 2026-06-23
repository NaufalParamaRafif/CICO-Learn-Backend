package handler

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/NaufalParamaRafif/CICO-Learn-Backend/domain"
	"github.com/NaufalParamaRafif/CICO-Learn-Backend/dto/request"
	"github.com/NaufalParamaRafif/CICO-Learn-Backend/internal/mapper"
	"github.com/NaufalParamaRafif/CICO-Learn-Backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase domain.UserUseCase
}

func NewUserHandler(
	router *gin.Engine,
	userUseCase domain.UserUseCase,
) {
	h := UserHandler{
		userUseCase: userUseCase,
	}

	// User
	router.GET("users/me", middleware.AuthMiddleware(userUseCase), h.GetUserData)
	router.PATCH("users/me/learning_preference", middleware.AuthMiddleware(userUseCase), h.UpdateLearningPreference)
	router.PATCH("users/me/profile", middleware.AuthMiddleware(userUseCase), h.UpdateProfile)

	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	// User - Background Image
	router.POST("/users/me/background_image", middleware.AuthMiddleware(userUseCase), h.UpdateBackgroundImage)
	router.DELETE("/users/me/background_image", middleware.AuthMiddleware(userUseCase), h.DeleteBackgroundImage)

	// User - Profile Picture
	router.POST("/users/me/profile_picture", middleware.AuthMiddleware(userUseCase), h.UpdateProfilePicture)
	router.DELETE("/users/me/profile_picture", middleware.AuthMiddleware(userUseCase), h.DeleteProfilePicture)
}

func (h UserHandler) UpdateLearningPreference(ctx *gin.Context) {
	rawUserID, exist := ctx.Get("user_id")
	if exist == false {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	userID, ok := rawUserID.(uint)
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	var req request.UpdateLearningPreferenceRequest
	var dailyReminders []domain.TimeOfDay

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}
	gin.DisableBindValidation()

	for _, val := range req.DailyReminder {
		// hour, minute, err := utils.ConvertStringToTimeOfDay(val)
		time := strings.Split(val, ":")
		if len(time) != 2 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Time format is incorrect",
			})

			return
		}

		dailyReminders = append(dailyReminders, domain.TimeOfDay{
			Hour:   time[0],
			Minute: time[1],
		})
	}

	err := h.userUseCase.Update(ctx, domain.UpdateUserInput{
		ID:            userID,
		Level:         (*domain.EnglishLevel)(req.Level),
		Language:      (*domain.Language)(req.Language),
		Target:        (*domain.Target)(req.Target),
		DailyReminder: dailyReminders,
		TimeTarget:    (*domain.TimeTarget)(req.TimeTarget),
	})
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": domain.ErrUserNotFound.Error(),
			})
		} else if errors.Is(err, domain.ErrEnglishLevelInvalid) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": domain.ErrEnglishLevelInvalid.Error(),
			})
		} else if errors.Is(err, domain.ErrLanguageInvalid) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": domain.ErrLanguageInvalid.Error(),
			})
		} else if errors.Is(err, domain.ErrTargetInvalid) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": domain.ErrTargetInvalid.Error(),
			})
		} else if errors.Is(err, domain.ErrTimeTargetInvalid) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": domain.ErrTimeTargetInvalid.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Learning preference updated successfully",
	})
}

func (h UserHandler) UpdateProfile(ctx *gin.Context) {
	rawUserID, exist := ctx.Get("user_id")
	if exist == false {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	userID, ok := rawUserID.(uint)
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	var req request.UpdateProfileRequest

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}
	gin.DisableBindValidation()

	err := h.userUseCase.Update(ctx, domain.UpdateUserInput{
		ID:        userID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Address:   req.Address,
		Gender:    (*domain.Gender)(req.Gender),
	})
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": domain.ErrUserNotFound.Error(),
			})
		} else if errors.Is(err, domain.ErrEnglishLevelInvalid) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": domain.ErrEnglishLevelInvalid.Error(),
			})
		} else if errors.Is(err, domain.ErrLanguageInvalid) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": domain.ErrLanguageInvalid.Error(),
			})
		} else if errors.Is(err, domain.ErrTargetInvalid) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": domain.ErrTargetInvalid.Error(),
			})
		} else if errors.Is(err, domain.ErrTimeTargetInvalid) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": domain.ErrTimeTargetInvalid.Error(),
			})
		} else if errors.Is(err, domain.ErrGenderInvalid) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": domain.ErrGenderInvalid.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
	})
}

func (h UserHandler) GetUserData(ctx *gin.Context) {
	rawUserID, exist := ctx.Get("user_id")
	if exist == false {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	userID, ok := rawUserID.(uint)
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	user, err := h.userUseCase.GetByID(ctx, userID)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": domain.ErrUserNotFound.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}

		return
	}

	response := mapper.ToUserDataResponse(user)

	ctx.JSON(http.StatusOK, response)
}

func (h UserHandler) UpdateBackgroundImage(ctx *gin.Context) {
	rawUserID, exist := ctx.Get("user_id")
	if exist == false {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	userID, ok := rawUserID.(uint)
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}
	gin.DisableBindValidation()

	file, err := fileHeader.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to open file",
		})
		return
	}
	defer file.Close()

	buffer := make([]byte, 512) // read just 512 kb
	if _, err := file.Read(buffer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read file",
		})
		return
	}

	detectedContentType := http.DetectContentType(buffer)
	if detectedContentType != "image/jpeg" && detectedContentType != "image/png" {
		ctx.JSON(http.StatusUnsupportedMediaType, gin.H{
			"error": "Image type must be jpg, jpeg, or png",
		})
		return
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	fileSize := float32(fileHeader.Size) / float32(1024*1024)

	fileInput := domain.FileInput{
		Filename:    fileHeader.Filename,
		Size:        fileSize,
		ContentType: detectedContentType,
		Data:        file,
	}

	fileURL, err := h.userUseCase.UploadUserImage(ctx, userID, &fileInput, domain.UserImageBackground)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Background image uploaded successfully",
		"file_url": fileURL,
	})
}

func (h UserHandler) DeleteBackgroundImage(ctx *gin.Context) {
	rawUserID, exist := ctx.Get("user_id")
	if exist == false {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	userID, ok := rawUserID.(uint)
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	err := h.userUseCase.DeleteUserImage(ctx, userID, domain.UserImageBackground)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": domain.ErrUserNotFound.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Background image deleted successfully",
	})
}

func (h UserHandler) UpdateProfilePicture(ctx *gin.Context) {
	rawUserID, exist := ctx.Get("user_id")
	if exist == false {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	userID, ok := rawUserID.(uint)
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	fileHeader, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	gin.DisableBindValidation()

	file, err := fileHeader.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to open file",
		})
		return
	}
	defer file.Close()

	buffer := make([]byte, 512) // read just 512 kb
	if _, err := file.Read(buffer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read file",
		})
		return
	}

	detectedContentType := http.DetectContentType(buffer)
	if detectedContentType != "image/jpeg" && detectedContentType != "image/png" {
		ctx.JSON(http.StatusUnsupportedMediaType, gin.H{
			"error": "Image type must be jpg, jpeg, or png",
		})
		return
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	fileSize := float32(fileHeader.Size) / float32(1024*1024)

	fileInput := domain.FileInput{
		Filename:    fileHeader.Filename,
		Size:        fileSize,
		ContentType: detectedContentType,
		Data:        file,
	}

	fileURL, err := h.userUseCase.UploadUserImage(ctx, userID, &fileInput, domain.UserImageProfile)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Profile picture uploaded successfully",
		"file_url": fileURL,
	})

}

func (h UserHandler) DeleteProfilePicture(ctx *gin.Context) {
	rawUserID, exist := ctx.Get("user_id")
	if exist == false {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
	userID, ok := rawUserID.(uint)
	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	err := h.userUseCase.DeleteUserImage(ctx, userID, domain.UserImageProfile)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": domain.ErrUserNotFound.Error(),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Profile picture deleted successfully",
	})

}
