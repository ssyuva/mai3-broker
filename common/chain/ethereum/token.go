package ethereum

import (
	"context"
	"fmt"
	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/mcarloai/mai-v3-broker/common/chain/ethereum/erc20"
)

func (c *Client) GetTokenSymbol(ctx context.Context, address string) (string, error) {
	var opts *ethBind.CallOpts
	addr, err := HexToAddress(address)
	if err != nil {
		return "", fmt.Errorf("invalid erc20 address:%w", err)
	}
	contract, err := erc20.NewErc20(addr, c.ethCli)
	if err != nil {
		return "", fmt.Errorf("init NewErc20 failed:%w", err)
	}

	return contract.Symbol(opts)
}
