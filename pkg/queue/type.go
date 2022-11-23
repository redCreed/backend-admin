package queue

type ConsumerFunc func(messager []Messager) error

type Queue interface {
	string() string
	Add(Messager) error
	Register(string, ConsumerFunc)
	Run()
	Close()
}
type Messager interface {
	SetId(string)
	GetId() string
	SetKey(string)
	GetKey() string
	SetValues([]byte)
	GetValues() []byte
	SetErrorCount(int)
	GetErrorCount() int
}
