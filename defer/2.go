import "github.com/happypancake/nanobus/guid"

func (m *Map) Add(correlationId guid.Guid, s *Context) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if _, ok := m.items[correlationId]; ok {
		return ErrAlreadyExists
	}

	m.items[correlationId] = s
	return nil
}