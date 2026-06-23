package domain

import (
	"context"
	"time"
)

type TimeOfDay struct {
	Hour   string
	Minute string
}

type User struct {
	ID              uint         `db:"id"` // TODO KALO UDAH DAPET HAK MU PERTAHANKAN HAK MU
	FirstName       *string      `db:"first_name"`
	LastName        *string      `db:"last_name"`
	Username        string       `db:"username"`
	Email           string       `db:"email"`
	Password        string       `db:"password"`
	Address         *string      `db:"address"`
	Gender          Gender       `db:"gender"`
	ProfilePicture  *string      `db:"profile_picture"`
	BackgroundImage *string      `db:"background_image"`
	Level           EnglishLevel `db:"level"`
	Language        Language     `db:"language"`
	Target          Target       `db:"target"`
	DailyReminder   []TimeOfDay  `db:"daily_reminder"`
	TimeTarget      *TimeTarget  `db:"time_target"`
	CreatedAt       time.Time    `db:"created_at"`
	UpdatedAt       *time.Time   `db:"updated_at"`
}

type UserRepository interface {
	GetByID(ctx context.Context, userID uint) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindByEmailOrUsername(ctx context.Context, identifier string) (*User, error)
	Insert(ctx context.Context, user *RegisterInput) (*User, error)
	Update(ctx context.Context, user *UpdateUserInput) (*User, error)
	Delete(ctx context.Context, userID uint) error
}

type UserUseCase interface {
	GetByID(ctx context.Context, userID uint) (*User, error)
	Update(ctx context.Context, input UpdateUserInput) error
	Delete(ctx context.Context, userID uint) error
	Register(ctx context.Context, input RegisterInput) error
	Login(ctx context.Context, input LoginInput) (*LoginResult, error)
	UploadUserImage(ctx context.Context, userID uint, file *FileInput, imageType UserImageType) (fileURL string, err error)
	DeleteUserImage(ctx context.Context, userID uint, imageType UserImageType) error
}
