package snowflake

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

func NewSnowflakeNode(startTime string, machineID int64) (*snowflake.Node, error) {
	var st time.Time
	var err error
	st, err = time.Parse(time.DateOnly, startTime) // 修改时间格式
	if err != nil {
		return nil, err
	}
	// 设置时间为较早的时间戳，例如Twitter的Snowflake算法使用的时间
	snowflake.Epoch = st.UnixNano() / 1e6
	node, err := snowflake.NewNode(machineID)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func GenerateID(node *snowflake.Node) int64 {
	return node.Generate().Int64()
}
