package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	sl "github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/common/models"
	"github.com/fenek-dev/go-twitter/src/sso/internal/lib"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type UserStorage interface {
	SaveUser(
		ctx context.Context,
		email string,
		passHash []byte,
	) (uid int64, err error)
	User(ctx context.Context, email string) (models.User, error)
}

type Auth struct {
	log         *slog.Logger
	userStorage UserStorage
	tokenTTL    time.Duration
}

func New(
	log *slog.Logger,
	userStorage UserStorage,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		userStorage: userStorage,
		log:         log,
		tokenTTL:    tokenTTL,
	}
}

// RegisterNewUser registers new user in the system and returns user ID.
// If user with given username already exists, returns error.
func (a *Auth) RegisterNewUser(ctx context.Context, email string, pass string) (int64, error) {
	const op = "Auth.RegisterNewUser"

	log := a.log.With(
		slog.String("op", op),
		slog.String("email", email),
	)

	log.Info("registering user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Error("failed to generate password hash", sl.Err(err))

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := a.userStorage.SaveUser(ctx, email, passHash)
	if err != nil {
		log.Error("failed to save user", sl.Err(err))

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (a *Auth) Login(
	ctx context.Context,
	email string,
	password string,
	secret string,
) (string, error) {
	const op = "Auth.Login"

	log := a.log.With(
		slog.String("op", op),
		slog.String("username", email),
	)

	log.Info("attempting to login user")

	user, err := a.userStorage.User(ctx, email)
	if err != nil {
		// if errors.Is(err, storage.ErrUserNotFound) {
		// 	a.log.Warn("user not found", sl.Err(err))

		// 	return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		// }

		a.log.Error("failed to get user", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		a.log.Info("invalid credentials", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}

	log.Info("user logged in successfully")

	// Создаём токен авторизации
	token, err := lib.NewToken(user, secret, a.tokenTTL)
	if err != nil {
		a.log.Error("failed to generate token", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}
