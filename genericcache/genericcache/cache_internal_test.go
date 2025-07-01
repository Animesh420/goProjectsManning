package genericcache

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCache_Parallel_goroutines(t *testing.T) {
	ttl := time.Duration(time.Second * 1000)

	c := New[int, string](5, ttl)

	const parallelTasks = 10
	wg := sync.WaitGroup{}
	wg.Add(parallelTasks)

	for i := 0; i < parallelTasks; i++ {
		go func(j int) {
			defer wg.Done()
			c.Upsert(4, fmt.Sprint(j))
		}(i)
	}

	wg.Wait()

}

func TestCache_Parallel(t *testing.T) {
	ttl := time.Duration(time.Second * 1000)
	c := New[int, string](5, ttl)

	t.Run("write six", func(t *testing.T) {
		t.Parallel()
		c.Upsert(6, "six")
	})

	t.Run("write keys", func(t *testing.T) {
		t.Parallel()
		c.Upsert(6, "keys")
	})
}

func TestCache_TTL(t *testing.T) {
	t.Parallel()

	c := New[string, string](5, time.Millisecond*100)
	c.Upsert("Norwegian", "Blue")

	got, found := c.Read("Norwegian")
	assert.True(t, found)
	assert.Equal(t, "Blue", got)

	time.Sleep(time.Millisecond * 200)

	got, found = c.Read("Norwegian")

	assert.False(t, found)
	assert.Equal(t, "", got)

}

// TestCache_Maxsize tests the maximum capacity feature of a cache
// It checks that update items are properly requeued as "new" item
// and that we make room by removing the most ancient item for ones.

func TestCache_Maxsize(t *testing.T){
	t.Parallel()

	// Give it a TTL long enough to survive the test
	c := New[int, int](3, time.Minute)

	c.Upsert(1, 1)
	c.Upsert(2, 2)
	c.Upsert(3, 3)

	got, found := c.Read(1)

	assert.True(t, found)
	assert.Equal(t, 1, got)

	// Update 1 which will no longer make it the olders
	c.Upsert(1, 10)

	// Adding a fourth element will discard the olders - 2 in this case
	c.Upsert(4, 4)

	// Trying to retrieve an element that should have been discarded
	got, found = c.Read(2)
	assert.False(t, found)
	assert.Equal(t, 0, got)

}