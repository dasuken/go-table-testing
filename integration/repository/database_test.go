package repository

import (
	"github.com/go-sql-driver/mysql"
	"testing"
)

func TestGenerateDatabaseQuery(t *testing.T) {
	got := generateDatabaseQuery("root", "mysql", "127.0.0.1", "3306", "todo", "?parseTime=true&loc=Japan")

	_, err := mysql.ParseDSN(got)
	if err != nil {
		t.Fatalf("cannot Parse DSN. err=%v", err)
	}
}
