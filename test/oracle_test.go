package oracle_test

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-oci8"
)

func TestOracleDriver(t *testing.T) {
	_, err := sql.Open("oci8", "")
	if err != nil {
		t.Fatalf("driver is not found: %s", err)
	}
}
