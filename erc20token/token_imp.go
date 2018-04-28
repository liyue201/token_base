package erc20token

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Erc20TokenImp struct {
	erc20Token *Token
	tokenAddr  string
}

func (t *Erc20TokenImp) GetAbi() string {
	return TokenABI
}

func (t *Erc20TokenImp) GetTokenAddr() string {
	return t.tokenAddr
}

func (t *Erc20TokenImp) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	return t.erc20Token.TotalSupply(opts)
}

func (t *Erc20TokenImp) BalanceOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	return t.erc20Token.BalanceOf(opts, addr)
}

func (t *Erc20TokenImp) Allowance(opts *bind.CallOpts, addr common.Address, arg1 common.Address) (*big.Int, error) {
	return t.erc20Token.Allowance(opts, addr, arg1)
}

func (t *Erc20TokenImp) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return t.erc20Token.Transfer(opts, to, value)
}

func (t *Erc20TokenImp) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return t.erc20Token.Approve(opts, spender, value)
}

func (t *Erc20TokenImp) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return t.erc20Token.TransferFrom(opts, from, to, value)
}

func NewTokenERC20(nodeAddr, tokenAddr string) (token Erc20token, err error) {
	ethCli, err := ethclient.Dial(nodeAddr)
	if err != nil {
		return
	}
	t, err := NewToken(common.HexToAddress(tokenAddr), ethCli)
	if err != nil {
		return nil, err
	}
	token = &Erc20TokenImp{erc20Token: t, tokenAddr: tokenAddr}
	return
}
