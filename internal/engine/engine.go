package engine

import (
	"database/sql"

	"github.com/hako/branca"
	gonanoid "github.com/matoous/go-nanoid"
)

// Engine is the blog engine
type Engine struct {
	codec *branca.Branca
	db    *sql.DB
}

//NewID creates a new Unique Id
func (*Engine) NewID() (string, error) {
	return gonanoid.Nanoid()
}

// New creates a new engine.Engine
func New(d *sql.DB, key string) *Engine {
	codec := branca.NewBranca(key)
	return &Engine{
		db:    d,
		codec: codec,
	}
}
