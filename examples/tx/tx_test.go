package main

import (
	"testing"

	"github.com/gatechain/gatechainsdk/api/client"
	"github.com/gatechain/gatechainsdk/examples/testutil"
	"github.com/gatechain/gatechainsdk/gatechain/crypto/keys"
	"github.com/michaelhou11/gatedemo/examples/common"
)

var kb keys.Keybase

func init() {
	kb = testutil.ImportAllPrivateKeys()
}

func TestQueryTx(t *testing.T) {
	client := client.NewClient(common.EndPoint, common.APIToken)
	hashHexStr := "IRREVOCABLEPAY-08A79A9031ACB80A848D19ADE96FCD39FF6A9510752A36A211A4B1E029A378EC1A687DA4704012696B36FED883DC89AF"
	client.Tx.QueryTX(hashHexStr)
}

func TestBroadcastTX(t *testing.T) {
	from_addr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	client := client.NewClientWithFromMemKeyBase(from_addr, common.EndPoint, common.APIToken, kb)

	to_addr := "gt11ka7wph4uzstt06v36v9nf4h4rh8hr47l69adlsmce5shwn02rx9ay5xpl686cn5dxpkp3f"
	amount := "100000000000NANOGT"
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"

	client.Tx.SendTX(from_addr, to_addr, amount, fees, chainID, gas, kb)
}

func TestCreateTxJsonFile(t *testing.T) {
	from_addr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9"
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"
	to_addr := "gt11ka7wph4uzstt06v36v9nf4h4rh8hr47l69adlsmce5shwn02rx9ay5xpl686cn5dxpkp3f"
	amount := "100000000000NANOGT"

	client := client.NewClientWithFromMemKeyBase(from_addr, common.EndPoint, common.APIToken, kb)
	client.Tx.CreateUnsignTX(from_addr, to_addr, amount, common.UnsignTxFile, fees, chainID, gas)

}

func TestSignTxJsonFile(t *testing.T) {
	from_addr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9" //
	fees := "100000000NANOGT"
	gas := uint64(200000)
	chainID := "gate-66"
	client := client.NewClientWithFromMemKeyBase(from_addr, common.EndPoint, common.APIToken, kb)
	client.Tx.CreateSignTX(common.UnsignTxFile, common.SignTxFile, fees, chainID, gas, kb)
}

func TestBroadCastTxJsonFile(t *testing.T) {
	from_addr := "gt11380m6lv6xr9fphasqunurpfus50h6eu4vgy8cvhrzayut9xkhju2zvfmergng55pee48u9" //
	client := client.NewClientWithFromMemKeyBase(from_addr, common.EndPoint, common.APIToken, kb)
	client.Tx.BroadcastSignTx(common.SignTxFile)
}
