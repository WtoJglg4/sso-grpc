package models

type User struct {
	ID       int64  `db:"id"`
	Email    string `db:"email"`
	PassHash []byte `db:"pass_hash"`
}
