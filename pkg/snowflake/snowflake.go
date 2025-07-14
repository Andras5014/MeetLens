package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"log"
)

// node 是全局单例
var node *snowflake.Node

// Init 初始化 Snowflake 节点
func Init(nodeID int64) {
	var err error
	node, err = snowflake.NewNode(nodeID)
	if err != nil {
		log.Fatalf("failed to initialize snowflake node: %v", err)
	}
}

// GenerateID 生成分布式唯一 ID
func GenerateID() int64 {
	if node == nil {
		log.Fatal("snowflake node is not initialized")
	}
	return node.Generate().Int64()
}

// GenerateStringID 返回字符串形式
func GenerateStringID() string {
	if node == nil {
		log.Fatal("snowflake node is not initialized")
	}
	return node.Generate().String()
}
