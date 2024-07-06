package snowflake

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"log"
	"sync"
	"testing"
	"time"
)

const (
	numMachines            = 10
	numIDsPerMachine       = 10
	machineID        int64 = 1
	printAllID             = true
)

func TestGenerateIDConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	idMap := sync.Map{}
	node, err := NewSnowflakeNode(time.Now().Format(time.DateOnly), machineID)
	if err != nil {
		t.Fatalf("Failed to create snowflake node: %v", err)
	}
	for i := 0; i < numIDsPerMachine; i++ {
		wg.Add(1)
		go func(node *snowflake.Node) {
			defer wg.Done()
			id := GenerateID(node)
			if printAllID {
				fmt.Println(id)
			}
			if _, loaded := idMap.LoadOrStore(id, struct{}{}); loaded {
				log.Fatal("Duplicate ID found!")
			}
		}(node)
	}
	wg.Wait()
}

func TestGenerateIDDistributedConcurrency(t *testing.T) {
	wg := sync.WaitGroup{}
	idMap := sync.Map{}
	// 启动多个goroutine来模拟不同机器生成ID
	for machineID := int64(1); machineID <= numMachines; machineID++ {
		node, err := NewSnowflakeNode(time.Now().Format(time.DateOnly), machineID)
		if err != nil {
			t.Fatalf("Failed to create snowflake node: %v", err)
		}
		for i := 0; i < numIDsPerMachine; i++ {
			wg.Add(1)
			go func(node *snowflake.Node) {
				defer wg.Done()
				id := GenerateID(node)
				if printAllID {
					fmt.Println(id)
				}
				if _, loaded := idMap.LoadOrStore(id, struct{}{}); loaded {
					log.Fatal("Duplicate ID found!")
				}
			}(node)
		}
	}
	wg.Wait()
}
