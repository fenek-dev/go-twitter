package services

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/fenek-dev/go-twitter/src/auth/internal/lib"
	sl "github.com/fenek-dev/go-twitter/src/common"
	"github.com/fenek-dev/go-twitter/src/common/mappers"
	"github.com/fenek-dev/go-twitter/src/common/models"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func (s *Services) RegisterNewUser(ctx context.Context, username string, pass string) (string, error) {
	const op = "auth.service.RegisterNewUser"
	ctx, span := s.tracer.Start(ctx, op)
	defer span.End()

	log := s.log.With(
		slog.String("op", op),
		slog.String("username", username),
	)

	log.Info("registering user")

	_, span = s.tracer.Start(ctx, op+".GenerateFromPassword")
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	span.End()
	if err != nil {
		log.Error("failed to generate password hash", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	user, err := s.pg.SaveUser(ctx, username, passHash)
	if err != nil {
		log.Error("failed to save user", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	_, span = s.tracer.Start(ctx, op+".NewToken")
	token, err := lib.NewToken(user, s.secret, s.tokenTTL)
	span.End()
	if err != nil {
		s.log.Error("failed to generate token", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}

func (s *Services) Login(
	ctx context.Context,
	username string,
	password string,
) (string, error) {
	const op = "auth.service.Login"
	ctx, span := s.tracer.Start(ctx, op)
	defer span.End()

	log := s.log.With(
		slog.String("op", op),
		slog.String("username", username),
	)

	log.Info("attempting to login user")

	user, err := s.pg.User(ctx, username)
	if err != nil {
		// if errors.Is(err, storage.ErrUserNotFound) {
		// 	s.log.Warn("user not found", sl.Err(err))

		// 	return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		// }

		s.log.Error("failed to get user", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	_, span = s.tracer.Start(ctx, op+".CompareHashAndPassword")
	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		s.log.Info("invalid credentials", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
	}
	span.End()

	log.Info("user logged in successfully")

	// Создаём токен авторизации
	_, span = s.tracer.Start(ctx, op+".NewToken")
	token, err := lib.NewToken(user, s.secret, s.tokenTTL)
	span.End()
	if err != nil {
		s.log.Error("failed to generate token", sl.Err(err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}

func (s *Services) Verify(
	ctx context.Context,
	token string,
) (*models.User, error) {
	const op = "auth.service.Verify"
	ctx, span := s.tracer.Start(ctx, op)
	defer span.End()

	log := s.log.With(
		slog.String("op", op),
	)

	// Создаём токен авторизации

	_, span = s.tracer.Start(ctx, op+".GetFromToken")
	claims, err := lib.GetFromToken(token, s.secret)
	span.End()
	if err != nil {
		log.Error("failed to verify token", sl.Err(err))

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	user := mappers.ClaimsToUserModel(claims)

	return user, nil
}
