package store

type Db interface {
	Sys() Sys
	Schema() Schema
}
