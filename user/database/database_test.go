package database

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	db := New()

	fmt.Println(db)
}
