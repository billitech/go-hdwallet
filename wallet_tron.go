package hdwallet

import "github.com/billitech/go-hdwallet/tron/address"

func init() {
	coins[TRON] = newTRX
}

type trx struct {
	name   string
	symbol string
	key    *Key
}

func newTRX(key *Key) Wallet {
	return &trx{
		name:   "tron",
		symbol: "xrp",
		key:    key,
	}
}

func (c *trx) GetType() uint32 {
	return c.key.Opt.CoinType
}

func (c *trx) GetName() string {
	return c.name
}

func (c *trx) GetSymbol() string {
	return c.symbol
}

func (c *trx) GetKey() *Key {
	return c.key
}

func (c *trx) GetAddress() (string, error) {
	return address.FromPublicKey(c.key.PublicECDSA).ToBase58(), nil
}
