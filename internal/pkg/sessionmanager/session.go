package sessionmanager

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

// SessionManager struct
type SessionManager struct {
	cache map[string]*session.Store
}

// NewSessionManager untuk membuat instance SessionManager
func NewSessionManager() *SessionManager {
	return &SessionManager{
		cache: make(map[string]*session.Store),
	}
}

// GetSessionInstance mengembalikan instance session berdasarkan Project ID
func (s *SessionManager) GetSessionInstance(projectID string) *session.Store {
	// Cek apakah sudah ada session untuk project ini
	if sess, exists := s.cache[projectID]; exists {
		return sess
	}

	// Jika belum ada, buat session baru
	newSession := session.New(session.Config{
		KeyLookup:  fmt.Sprintf("cookie:kaho_session_%s", projectID),
		Expiration: 24 * time.Hour,
	})

	// Simpan ke cache
	s.cache[projectID] = newSession

	return newSession
}
