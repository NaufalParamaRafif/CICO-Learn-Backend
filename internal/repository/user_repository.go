package repository

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/NaufalParamaRafif/CICO-Learn-Backend/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) domain.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetByID(ctx context.Context, userID uint) (*domain.User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, address, gender, profile_picture, background_image, level, language, target, daily_reminder, time_target, password, created_at, updated_at
		FROM users
		where id = $1
	`

	var user domain.User
	var dailyReminders []time.Time

	err := r.db.QueryRow(
		ctx, query, userID,
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Address,
		&user.Gender,
		&user.ProfilePicture,
		&user.BackgroundImage,
		&user.Level,
		&user.Language,
		&user.Target,
		&dailyReminders,
		&user.TimeTarget,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return &domain.User{}, err
	}

	user.DailyReminder = make([]domain.TimeOfDay, 0, len(dailyReminders))

	for _, time := range dailyReminders {
		user.DailyReminder = append(user.DailyReminder, domain.TimeOfDay{
			Hour:   fmt.Sprintf("%02d", time.Hour()),
			Minute: fmt.Sprintf("%02d", time.Minute()),
		})
	}

	return &user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, address, gender, profile_picture, background_image, level, language, target, daily_reminder, time_target, password, created_at, updated_at
		FROM users
		where email = $1
	`

	var user domain.User
	var dailyReminders []time.Time

	err := r.db.QueryRow(
		ctx, query, email,
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Address,
		&user.Gender,
		&user.ProfilePicture,
		&user.BackgroundImage,
		&user.Level,
		&user.Language,
		&user.Target,
		&dailyReminders,
		&user.TimeTarget,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return &domain.User{}, err
	}

	user.DailyReminder = make([]domain.TimeOfDay, 0, len(dailyReminders))

	for _, time := range dailyReminders {
		user.DailyReminder = append(user.DailyReminder, domain.TimeOfDay{
			Hour:   fmt.Sprintf("%02d", time.Hour()),
			Minute: fmt.Sprintf("%02d", time.Minute()),
		})
	}

	return &user, nil
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, address, gender, profile_picture, background_image, level, language, target, daily_reminder, time_target, password, created_at, updated_at
		FROM users
		WHERE
			username = $1
	`

	var user domain.User
	var dailyReminders []time.Time

	err := r.db.QueryRow(
		ctx,
		query,
		username,
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Address,
		&user.Gender,
		&user.ProfilePicture,
		&user.BackgroundImage,
		&user.Level,
		&user.Language,
		&user.Target,
		&dailyReminders,
		&user.TimeTarget,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return &domain.User{}, err
	}

	user.DailyReminder = make([]domain.TimeOfDay, 0, len(dailyReminders))

	for _, time := range dailyReminders {
		user.DailyReminder = append(user.DailyReminder, domain.TimeOfDay{
			Hour:   fmt.Sprintf("%02d", time.Hour()),
			Minute: fmt.Sprintf("%02d", time.Minute()),
		})
	}

	return &user, err
}

func (r *UserRepository) FindByEmailOrUsername(ctx context.Context, identifier string) (*domain.User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, address, gender, profile_picture, background_image, level, language, target, daily_reminder, time_target, password, created_at, updated_at
		FROM users
		where 
			email = $1
			OR
			username = $1
	`

	var user domain.User
	var dailyReminders []time.Time

	err := r.db.QueryRow(
		ctx, query, identifier,
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Address,
		&user.Gender,
		&user.ProfilePicture,
		&user.BackgroundImage,
		&user.Level,
		&user.Language,
		&user.Target,
		&dailyReminders,
		&user.TimeTarget,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return &domain.User{}, err
	}

	user.DailyReminder = make([]domain.TimeOfDay, 0, len(dailyReminders))

	for _, time := range dailyReminders {
		user.DailyReminder = append(user.DailyReminder, domain.TimeOfDay{
			Hour:   fmt.Sprintf("%02d", time.Hour()),
			Minute: fmt.Sprintf("%02d", time.Minute()),
		})
	}

	return &user, nil
}

func (r *UserRepository) Insert(ctx context.Context, input *domain.RegisterInput) (*domain.User, error) {
	columns := []string{}
	placeholders := []string{}
	args := []any{}

	columns = append(columns, "username", "email", "password")
	placeholders = append(placeholders, "$1", "$2", "$3")
	args = append(args, input.Username, input.Email, input.Password)

	if input.FirstName != nil {
		columns = append(columns, "first_name")
		placeholders = append(placeholders, "$4")
		args = append(args, *input.FirstName)
	}

	if input.LastName != nil {
		columns = append(columns, "last_name")
		placeholders = append(placeholders, "$5")
		args = append(args, *input.LastName)
	}

	query := fmt.Sprintf(
		`
			INSERT INTO users(%s)
			VALUES(%s)
			RETURNING id, first_name, last_name, username, email, password
		`,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	var user domain.User

	err := r.db.QueryRow(
		ctx,
		query,
		args...,
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return &domain.User{}, err
	}

	return &user, err
}

func (r *UserRepository) Update(ctx context.Context, input *domain.UpdateUserInput) (*domain.User, error) {
	updates := []string{}
	args := []any{}
	argPosition := 1

	if input.FirstName != nil {
		updates = append(
			updates,
			fmt.Sprintf("first_name = $%d", argPosition),
		)
		args = append(args, input.FirstName)
		argPosition++
	}

	if input.LastName != nil {
		updates = append(
			updates,
			fmt.Sprintf("last_name = $%d", argPosition),
		)
		args = append(args, input.LastName)
		argPosition++
	}

	if input.Address != nil {
		updates = append(
			updates,
			fmt.Sprintf("address = $%d", argPosition),
		)
		args = append(args, input.Address)
		argPosition++
	}

	if input.Gender != nil {
		updates = append(
			updates,
			fmt.Sprintf("gender = $%d", argPosition),
		)
		args = append(args, input.Gender)
		argPosition++
	}

	if input.ProfilePicture != nil {
		updates = append(
			updates,
			fmt.Sprintf("profile_picture = $%d", argPosition),
		)
		args = append(args, input.ProfilePicture)
		argPosition++
	}

	if input.BackgroundImage != nil {
		updates = append(
			updates,
			fmt.Sprintf("background_image = $%d", argPosition),
		)
		args = append(args, input.BackgroundImage)
		argPosition++
	}

	if input.Level != nil {
		updates = append(
			updates,
			fmt.Sprintf("level = $%d", argPosition),
		)
		args = append(args, input.Level)
		argPosition++
	}

	if input.Language != nil {
		updates = append(
			updates,
			fmt.Sprintf("language = $%d", argPosition),
		)
		args = append(args, input.Language)
		argPosition++
	}

	if input.Target != nil {
		updates = append(
			updates,
			fmt.Sprintf("target = $%d", argPosition),
		)
		args = append(args, input.Target)
		argPosition++
	}

	if input.DailyReminder != nil {
		updates = append(
			updates,
			fmt.Sprintf("daily_reminder = $%d", argPosition),
		)

		var timeStrings []string
		for _, time := range input.DailyReminder {
			timeStrings = append(timeStrings, time.Hour+":"+time.Minute)
		}

		args = append(args, timeStrings)
		argPosition++
	}

	if input.TimeTarget != nil {
		updates = append(
			updates,
			fmt.Sprintf("time_target = $%d", argPosition),
		)
		args = append(args, input.TimeTarget)
		argPosition++
	}

	updates = append(
		updates,
		fmt.Sprintf("updated_at = $%d", argPosition),
	)
	args = append(args, time.Now())
	argPosition++

	query := fmt.Sprintf(`
		UPDATE users
		SET %s
		WHERE id = $%d
		RETURNING id, first_name, last_name, username, email, address, gender, profile_picture, background_image, level, language, target, daily_reminder, time_target, password, created_at, updated_at

	`,
		strings.Join(updates, ", "),
		argPosition,
	)
	args = append(args, input.ID)

	var user domain.User

	var dailyReminders []time.Time

	err := r.db.QueryRow(
		ctx,
		query,
		args...,
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Address,
		&user.Gender,
		&user.ProfilePicture,
		&user.BackgroundImage,
		&user.Level,
		&user.Language,
		&user.Target,
		&dailyReminders,
		&user.TimeTarget,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return &domain.User{}, err
	}

	user.DailyReminder = make([]domain.TimeOfDay, 0, len(dailyReminders))

	for _, time := range dailyReminders {
		user.DailyReminder = append(user.DailyReminder, domain.TimeOfDay{
			Hour:   fmt.Sprintf("%02d", time.Hour()),
			Minute: fmt.Sprintf("%02d", time.Minute()),
		})
	}

	return &user, err
}

func (r *UserRepository) Delete(ctx context.Context, userID uint) error {
	query := `
		DELETE FROM users 
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query, userID)

	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.ErrUnsupported
	}

	return err
}
