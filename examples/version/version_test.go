package main

import (
	"testing"

	"github.com/gatechain/gatechainsdk/api/client"
	"github.com/michaelhou11/gatedemo/examples/common"
)

func TestVersion(t *testing.T) {
	client := client.NewClient(common.EndPoint, common.APIToken)
	client.Version.Version()
}
