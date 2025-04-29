package sqlcerr

import (
	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
)

var (
	ErrNoRows = pgx.ErrNoRows
)

func Is(err, target error) bool {
	return err.Error() == target.Error()
}

func IsDuplicate(err error) bool {
	if err != nil {
		duplicateError, ok := err.(*pq.Error)
		if ok {
			return duplicateError.Code == "23505"
		}
	}
	return false
}
