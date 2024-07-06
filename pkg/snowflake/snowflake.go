package snowflake

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

// NewSnowflakeNode 新建一个 Snowflake 节点
func NewSnowflakeNode(startTime string, machineID int64) (*snowflake.Node, error) {
	var st time.Time
	var err error

	// 解析开始时间
	st, err = time.Parse(time.DateOnly, startTime)
	if err != nil {
		return nil, err
	}

	snowflake.Epoch = st.UnixMilli()
	// 新建一个节点
	node, err := snowflake.NewNode(machineID)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// GenerateID 生成唯一 ID
func GenerateID(node *snowflake.Node) int64 {
	return node.Generate().Int64()
}
