// /home/${USER}/livestreampro/backend/internal/auth/service.go
package auth

import "context"

// Service defines authentication operations used across the backend.
type Service interface {
	// Login authenticates a user and returns a JWT token on success.
	Login(ctx context.Context, username, password string) (string, error)
	// Logout invalidates the provided token.
	Logout(ctx context.Context, token string) error
}

// StubService is a placeholder implementation that allows other packages to compile.
type StubService struct{}

func (s *StubService) Login(ctx context.Context, username, password string) (string, error) {
	return "stub-token", nil
}

func (s *StubService) Logout(ctx context.Context, token string) error {
	return nil
}
