package queue

type Message struct {
	id       string
	key      string
	value    []byte
	errCount int
}

func (m *Message) GetId() string {
	return m.id
}
func (m *Message) SetId(id string) {
	m.id = id
}

func (m *Message) SetKey(key string) {
	m.key = key
}
func (m *Message) GetKey() string {
	return m.key
}

func (m *Message) GetValues() []byte {
	return m.value
}

func (m *Message) SetValues(value []byte) {
	m.value = value
}

func (m *Message) SetErrorCount(count int) {
	m.errCount = count
}

func (m *Message) GetErrorCount() int {
	return m.errCount
}
