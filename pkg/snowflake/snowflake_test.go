package snowflake

import (
	"fmt"
	"sync"
	"testing"
)

func TestSnowflakeIDUniqueness(t *testing.T) {
	err := Init("2023-01-01", 1)
	if err != nil {
		t.Fatalf("Failed to initialize Snowflake node: %v", err)
	}

	var wg sync.WaitGroup
	idCount := 100000
	ids := make(chan int64, idCount)

	// 启动多个goroutine来生成ID
	for i := 0; i < idCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			id := GenID()
			ids <- id
		}()
	}

	wg.Wait()
	close(ids)

	idMap := make(map[int64]struct{})

	for id := range ids {
		if _, exists := idMap[id]; exists {
			t.Fatalf("Duplicate ID found: %d", id)
		}
		idMap[id] = struct{}{}
	}

	fmt.Printf("Generated %d unique IDs\n", len(idMap))
}
