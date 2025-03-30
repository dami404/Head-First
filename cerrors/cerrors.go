package cerrors

import (
	"errors"
)

var ErrorDraw error = errors.New("It's draw! Vote again")

var SQLError = struct {
	TableNotFoundError func(tableName string) error
}{
	TableNotFoundError: func(tableName string) error {
		return errors.New("Table " + tableName + " not found")
	},
}
