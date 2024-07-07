package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

// NewSnowflakeNode 新建一个 Snowflake 节点
func NewSnowflakeNode(startTime string, machineID int64) (*sf.Node, error) {
	var st time.Time
	var err error

	// 解析开始时间
	st, err = time.Parse(time.DateOnly, startTime)
	if err != nil {
		return nil, err
	}

	sf.Epoch = st.UnixMilli()
	// 新建一个节点
	node, err := sf.NewNode(machineID)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// GenerateID 生成唯一 ID
func GenerateID(node *sf.Node) int64 {
	return node.Generate().Int64()
}
