package store

type Schema interface {
	CreateSchema() error
	CreateRoot() error
	HasTable(tb interface{}) bool
	Close() error
	Migrate(tables []interface{}) error
	Exec(string) error
}
