package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"log"
	"sync"
	"testing"
	"time"
)

const (
	numMachines      = 10   // 机器数量
	numIDsPerMachine = 1000 // 每台机器生成的 ID 数量
	printAllID       = true // 是否打印所有生成的 ID
)

func TestGenerateIDConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	idMap := sync.Map{}
	var machineID int64 = 1
	// 创建一个 Snowflake 节点
	node, err := NewSnowflakeNode(time.Now().Format(time.DateOnly), machineID)
	if err != nil {
		t.Fatalf("Failed to create snowflake node: %v", err)
	}

	// 并发生成 ID
	for i := 0; i < numIDsPerMachine; i++ {
		wg.Add(1)
		go func(node *snowflake.Node) {
			defer wg.Done()
			id := GenerateID(node)
			if printAllID {
				t.Logf("Generated ID:%d", id)
			}
			// 检查是否有重复 ID
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
					t.Logf("MachineID:%d, %d of %d:%d", machineID, i, numIDsPerMachine, id)
				}
				if _, loaded := idMap.LoadOrStore(id, struct{}{}); loaded {
					log.Fatal("Duplicate ID found!")
				}
			}(node)
		}
	}
	wg.Wait()
}
