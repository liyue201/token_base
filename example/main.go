package main

import (
	"fmt"
	"github.com/liyue201/token_base/erc20token"
	"github.com/liyue201/token_base/geth"
	"math/big"
)

func main() {
	nodeAddr := "http://192.168.8.53:8545"

	fmt.Println("============================================================")

	//rinkeby eos network
	//https://rinkeby.etherscan.io/tx/0xcc9d8fba9d1d773179646a4e8bb120d5fb933d14664ce79cb12487e268a173fc
	TestToken(nodeAddr,
		"0xD0cd0114D255DDEE071D863c77dFB63889c6cEA0",
		"0xd38f740e2360a86135423f07144fcceaa3665de5",
		2167247)

	//rinkeby yee network
	TestToken(nodeAddr,
		"0x1893bc5B5d715ff2Dda02Cb96C7F7460ba167d40",
		"0xd38f740e2360a86135423f07144fcceaa3665de5",
		2167247)

	//rinkeby yee network
	//https://rinkeby.etherscan.io/tx/0xb5916866318816be3a789348819c3f3e6381157a2901bc9675669b47d00e1e2f
	TestToken(nodeAddr,
		"0x9227E7097A8Eec5d3418daADA643991eeDA1D28C",
		"0xaf1bc95420616ce6e39d23885c990d80ca13c506",
		2124447)
}

func TestToken(nodeAddr, tokenAddr, addr string, height int64) {

	erc20, err := erc20token.NewTokenERC20(nodeAddr, tokenAddr)
	if err != nil {
		fmt.Println("[NewTokenERC20] error,", err)
		return
	}
	cli, err := geth.NewTokenClient(erc20, nodeAddr)
	if err != nil {
		fmt.Println("[NewTokenClient] error,", err)
		return
	}
	blockNum, err := cli.GetLatestBlockNum()
	if err != nil {
		fmt.Println("[GetLatestBlockNum] error,", err)
		return
	}
	fmt.Println("blockNum:", blockNum)

	totalSupply, err := cli.TotalSupply()
	if err != nil {
		fmt.Println("[TotalSupply] error,", err)
		return
	}
	fmt.Println("totalSupply:", totalSupply)

	balance, err := cli.BalanceOfToken(addr)
	if err != nil {
		fmt.Println("[BalanceOfToken] error,", err)
		return
	}
	fmt.Println("balance:", balance)

	logs, err := cli.GetAllTransferLogByAddress(big.NewInt(0), big.NewInt(height), addr)
	if err != nil {
		fmt.Println("[GetAllTransferLogByAddress] error,", err)
		return
	}

	fmt.Println("len(logs)=", len(logs))
	for _, log := range logs {
		lg, err := geth.ConvertTransferLog(log)
		if err == nil {
			fmt.Println(lg.TxHash, lg.From, lg.To, lg.Val)
		}
	}

	fmt.Println("============================================================")
}
