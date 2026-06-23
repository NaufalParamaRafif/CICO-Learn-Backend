package usecase

import (
	"context"
	"crypto/rand"
	"errors"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/NaufalParamaRafif/CICO-Learn-Backend/domain"
	"github.com/NaufalParamaRafif/CICO-Learn-Backend/internal/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepository domain.UserRepository
}

func NewUserUseCase(userRepository *domain.UserRepository) domain.UserUseCase {
	return &userUseCase{
		userRepository: *userRepository,
	}
}

func (u userUseCase) GetByID(ctx context.Context, userID uint) (*domain.User, error) {
	user, err := u.userRepository.GetByID(ctx, userID)

	if err != nil {
		return &domain.User{}, err
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return &domain.User{}, domain.ErrUserNotFound
	}

	return user, err
}

func (u userUseCase) Update(ctx context.Context, input domain.UpdateUserInput) error {
	_, err := u.userRepository.GetByID(ctx, input.ID)
	if err != nil {
		return err
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.ErrUserNotFound
	}

	if input.Level != nil && !input.Level.IsValid() {
		return domain.ErrEnglishLevelInvalid
	}
	if input.Language != nil && !input.Language.IsValid() {
		return domain.ErrLanguageInvalid
	}
	if input.Target != nil && !input.Target.IsValid() {
		return domain.ErrTargetInvalid
	}
	if input.TimeTarget != nil && !input.TimeTarget.IsValid() {
		return domain.ErrTimeTargetInvalid
	}
	if input.Gender != nil && !input.Gender.IsValid() {
		return domain.ErrGenderInvalid
	}

	_, err = u.userRepository.Update(ctx, &input)

	return err
}

func (u userUseCase) Delete(ctx context.Context, userID uint) error {
	_, err := u.userRepository.GetByID(ctx, userID)
	if err != nil {
		return err
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.ErrUserNotFound
	}

	err = u.userRepository.Delete(ctx, userID)

	return err
}

func (u userUseCase) Register(ctx context.Context, input domain.RegisterInput) error {
	_, err := u.userRepository.FindByEmail(ctx, input.Email)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return domain.ErrEmailExists
	}

	_, err = u.userRepository.FindByUsername(ctx, input.Username)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return domain.ErrUsernameTaken
	}

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// user.Password = string(password)
	input.Password = string(password)
	_, err = u.userRepository.Insert(ctx, &input)

	return err
}

func (u userUseCase) Login(ctx context.Context, input domain.LoginInput) (*domain.LoginResult, error) {
	if input.Identifier == "" {
		return &domain.LoginResult{}, domain.ErrEmptyIdentifier
	}
	if input.Password == "" {
		return &domain.LoginResult{}, domain.ErrEmptyPassword
	}

	user, err := u.userRepository.FindByEmailOrUsername(ctx, input.Identifier)
	if errors.Is(err, pgx.ErrNoRows) {
		return &domain.LoginResult{}, domain.ErrUserNotFound
	}
	if err != nil {
		return &domain.LoginResult{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return &domain.LoginResult{}, domain.ErrInvalidIdentifierPassword
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      user.ID,
		"username": user.Username,
		// "exp":      time.Now().Add(time.Minute).Unix(),
		"exp": time.Now().Add(time.Hour).Unix(), // TODO: Fix this
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return &domain.LoginResult{}, err
	}

	refreshToken := rand.Text() // TODO: SAVE REFRESH TOKEN, MAKE MORE LONG

	return &domain.LoginResult{
		RefreshToken: refreshToken,
		AccessToken:  tokenString,
	}, err
}

func (u userUseCase) UploadUserImage(ctx context.Context, userID uint, file *domain.FileInput, imageType domain.UserImageType) (fileURL string, err error) {
	// Change file name
	fileExtension := filepath.Ext(file.Filename) // TODO : fix ..
	file.Filename = uuid.NewString() + fileExtension

	// Save file
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	fileLocation := filepath.Join(dir, "files", "images", file.Filename)

	err = os.MkdirAll(filepath.Dir(fileLocation), 0755)
	if err != nil {
		return "", err
	}

	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return "", err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, file.Data); err != nil {
		return "", err
	}

	// if greater than 4mb compress
	if file.Size > 4096 {
		utils.CompressImage(fileLocation, fileLocation, 80)
	}

	// delete old profile photo if exists
	user, err := u.userRepository.GetByID(ctx, userID)
	if err != nil {
		return "", err
	}

	var input domain.UpdateUserInput

	switch imageType {
	case domain.UserImageProfile:
		if user.ProfilePicture != nil {
			fileLocation = filepath.Join(dir, "files", "images", *user.ProfilePicture)

			_ = os.Remove(fileLocation)
			// if err != nil {
			// 	return "", err
			// }
		}

		input = domain.UpdateUserInput{
			ID:             userID,
			ProfilePicture: &file.Filename,
		}
	case domain.UserImageBackground:
		if user.BackgroundImage != nil {
			fileLocation = filepath.Join(dir, "files", "images", *user.BackgroundImage)

			_ = os.Remove(fileLocation)
			// if err != nil {
			// 	return "", err
			// }
		}

		input = domain.UpdateUserInput{
			ID:              userID,
			BackgroundImage: &file.Filename,
		}
	}

	user, err = u.userRepository.Update(ctx, &input)
	if err != nil {
		return "", err
	}

	return file.Filename, err
}

func (u userUseCase) DeleteUserImage(ctx context.Context, userID uint, imageType domain.UserImageType) error {
	user, err := u.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	var inputUserUpdate domain.UpdateUserInput
	var fileLocation string

	switch imageType {
	case domain.UserImageProfile:
		inputUserUpdate = domain.UpdateUserInput{
			ID:             userID,
			ProfilePicture: new(string),
		}

		fileLocation = filepath.Join(dir, "files", "images", *user.ProfilePicture)
	case domain.UserImageBackground:
		inputUserUpdate = domain.UpdateUserInput{
			ID:              userID,
			BackgroundImage: new(string),
		}

		fileLocation = filepath.Join(dir, "files", "images", *user.BackgroundImage)
	}

	err = u.Update(ctx, inputUserUpdate)
	if err != nil {
		return err
	}

	err = os.Remove(fileLocation)
	if err != nil {
		return err
	}

	return nil
}
