package handler

import (
	"github.com/jackc/pgx/v5"
)

type Handlers struct {
	Conn *pgx.Conn
}
