package main

import (
	"fmt"
	"testing"
)

func TestCreateDatabase(t *testing.T) {
	create_database("testing")
	db_exists_query := fmt.Sprintf("select exists(SELECT datname FROM pg_catalog.pg_database WHERE lower(datname) = lower('dbname'));")

}
