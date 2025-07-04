// /home/${USER}/livestreampro/backend/internal/auth/service_test.go
package auth

import (
	"context"
	"testing"
)

func TestStubService(t *testing.T) {
	var _ Service = (*StubService)(nil)

	s := &StubService{}
	token, err := s.Login(context.Background(), "user", "pass")
	if err != nil {
		t.Fatalf("login error: %v", err)
	}
	if token == "" {
		t.Fatalf("expected token")
	}
	if err := s.Logout(context.Background(), token); err != nil {
		t.Fatalf("logout error: %v", err)
	}
}
