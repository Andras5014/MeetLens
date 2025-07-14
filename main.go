package MeetLens

import "MeetLens/pkg/snowflake"

func main() {
	// 建议 nodeID 0~1023，生产可用机器ID + 数据中心ID 组合
	snowflake.Init(1)
}
