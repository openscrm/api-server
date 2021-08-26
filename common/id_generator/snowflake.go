package id_generator

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

var Node *snowflake.Node

func SetupIDGenerator() {
	var err error
	if Node, err = snowflake.NewNode(1); err != nil {
		panic(err)
	}
}

func ID() int64 {
	return Node.Generate().Int64()
}

func StringID() string {
	return fmt.Sprintf("%d", Node.Generate().Int64())
}
