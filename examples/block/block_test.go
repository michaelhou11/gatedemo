package main

import (
	"testing"

	"github.com/gatechain/gatechainsdk/api/client"
	"github.com/michaelhou11/gatedemo/examples/common"
)

func TestQueryCurBlock(t *testing.T) {
	client := client.NewClient(common.EndPoint, common.APIToken)
	blockHeight := uint64(0)

	client.Block.Block(blockHeight)
}

func TestQueryBlock(t *testing.T) {
	client := client.NewClient(common.EndPoint, common.APIToken)
	blockHeight := uint64(1)

	client.Block.Block(blockHeight)
}
