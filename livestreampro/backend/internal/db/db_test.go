// /home/${USER}/livestreampro/backend/internal/db/db_test.go
package db

import "testing"

func TestConnectInvalid(t *testing.T) {
	d, err := Connect("postgres://invalid")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer d.Close()

	if err := d.Ping(); err == nil {
		t.Fatalf("expected ping error for invalid connection")
	}
}
