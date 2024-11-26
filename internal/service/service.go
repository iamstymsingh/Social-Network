package service

import (
	"database/sql"
	"github.com/hako/branca"
)

type Service struct {
	db    *sql.DB
	codec *branca.Branca
}
