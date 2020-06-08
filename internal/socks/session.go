package socks

// ClientIndentity for client really id map
type ClientIndentity interface {
	// ID2Key id to key
	ID2Key(id string) (string, error)
	// Key2ID key to id
	Key2ID(key string) (string, error)
}

// DefaultClientIdentity default
type DefaultClientIdentity struct {
}

// ID2Key id to key
func (d *DefaultClientIdentity) ID2Key(id string) (string, error) {
	return id, nil
}

// Key2ID key to id
func (d *DefaultClientIdentity) Key2ID(key string) (string, error) {
	return key, nil
}

// Session session
type Session struct {
	conn Conn
	ID   string
}

// // SessionManager session manager
// type SessionManager struct {
// 	mutex sync.Mutex
// 	// using client id for key
// 	sessions map[string]*Session

// 	// clientIdentity key map
// 	clientIndentity ClientIndentity
// }

// func (m *SessionManager) Get(key string) (*Session, error) {
// 	id, err := m.clientIndentity.Key2ID(key)
// 	if err != nil {
// 		return nil, err
// 	}

// 	m.mutex.Lock()
// 	defer m.mutex.Unlock()
// 	return m.sessions[id], nil
// }

// func (m *SessionManager) Add(id string, session *Session) {
// 	m.mutex.Lock()
// 	defer m.mutex.Unlock()
// 	m.sessions[id] = session
// }
