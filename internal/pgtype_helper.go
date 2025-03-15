package internal

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func ToPgTypeUUID(id *uuid.UUID) pgtype.UUID {
	r := pgtype.UUID{Valid: false}

	if id != nil {
		r = pgtype.UUID{Valid: true, Bytes: *id}
	}

	return r
}
