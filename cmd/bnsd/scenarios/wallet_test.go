package scenarios

import (
	"testing"

	"github.com/iov-one/weave/coin"
	"github.com/iov-one/weave/x/cash"

	"github.com/iov-one/weave/cmd/bnsd/client"
	"github.com/stretchr/testify/require"
)

func TestSendTokensWithoutFee(t *testing.T) {
	emilia := client.GenPrivateKey()
	aNonce := client.NewNonce(bnsClient, alice.PublicKey().Address())

	walletResp, err := bnsClient.GetWallet(alice.PublicKey().Address())
	require.NoError(t, err)
	require.NotNil(t, walletResp, "address not found")
	require.NotEmpty(t, walletResp.Wallet.Coins)

	heights := make([]int64, len(walletResp.Wallet.Coins))
	for i, cc := range walletResp.Wallet.Coins {
		// send a coin from Alice to Emilia
		c := coin.Coin{
			Ticker:     cc.Ticker,
			Fractional: 0,
			Whole:      1,
		}

		seq, err := aNonce.Next()
		require.NoError(t, err)
		tx := client.BuildSendTx(alice.PublicKey().Address(), emilia.PublicKey().Address(), c, "test tx without fee")
		require.NoError(t, client.SignTx(tx, alice, chainID, seq))
		resp := bnsClient.BroadcastTx(tx)
		require.NoError(t, resp.IsError())
		heights[i] = resp.Response.Height
		delayForRateLimits()
	}
	walletResp, err = bnsClient.GetWallet(emilia.PublicKey().Address())
	require.NoError(t, err)
	t.Log("message", "done", "height", heights, "coins", walletResp.Wallet.Coins)
}

func TestSendTokenWithFee(t *testing.T) {
	emilia := client.GenPrivateKey()
	aNonce := client.NewNonce(bnsClient, alice.PublicKey().Address())

	walletResp, err := bnsClient.GetWallet(alice.PublicKey().Address())
	require.NoError(t, err)
	require.NotNil(t, walletResp, "address not found")
	require.NotEmpty(t, walletResp.Wallet.Coins)

	heights := make([]int64, len(walletResp.Wallet.Coins))
	for i, c := range walletResp.Wallet.Coins {
		// send a coin from Alice to Emilia
		cc := coin.Coin{
			Ticker:     c.Ticker,
			Fractional: 0,
			Whole:      1,
		}

		seq, err := aNonce.Next()
		require.NoError(t, err)
		tx := client.BuildSendTx(alice.PublicKey().Address(), emilia.PublicKey().Address(), cc, "test tx with fee")
		tx.Fees = &cash.FeeInfo{
			Payer: alice.PublicKey().Address(),
			Fees: &coin.Coin{
				Ticker:     c.Ticker,
				Fractional: 1,
				Whole:      0,
			},
		}
		require.NoError(t, client.SignTx(tx, alice, chainID, seq))
		resp := bnsClient.BroadcastTx(tx)
		require.NoError(t, resp.IsError())
		heights[i] = resp.Response.Height
		delayForRateLimits()
	}
	walletResp, err = bnsClient.GetWallet(emilia.PublicKey().Address())
	require.NoError(t, err)
	t.Log("message", "done", "height", heights, "coins", walletResp.Wallet.Coins)
}

func TestQueryCurrencies(t *testing.T) {
	l, err := bnsClient.Currencies()
	require.NoError(t, err)
	require.NotNil(t, l, "no currencies found")
	require.True(t, len(l.Currencies) > 0)
}
