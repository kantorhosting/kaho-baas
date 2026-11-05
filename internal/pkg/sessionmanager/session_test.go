package sessionmanager

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSessionInstance_SameProjectReturnsSameStore(t *testing.T) {
	sm := NewSessionManager()

	first := sm.GetSessionInstance("project-a")
	second := sm.GetSessionInstance("project-a")
	other := sm.GetSessionInstance("project-b")

	assert.Same(t, first, second)
	assert.NotSame(t, first, other)
}

func TestGetSessionInstance_Concurrent(t *testing.T) {
	sm := NewSessionManager()

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.GetSessionInstance(fmt.Sprintf("project-%d", i%10))
		}(i)
	}
	wg.Wait()

	assert.Len(t, sm.cache, 10)
}
