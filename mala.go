package mala

import "database/sql"

type Beginner interface {
	Begin() (*sql.Tx, error)
}

type Mala struct {
	beginner Beginner
}
