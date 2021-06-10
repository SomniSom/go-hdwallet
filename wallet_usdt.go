package hdwallet

func init() {
	coins[USDT] = newUSDT
}

type usdt struct {
	*eth
}

func newUSDT(key *Key) Wallet {
	token := newBTC(key).(*eth)
	token.name = "Tether"
	token.symbol = "USDT"
	token.contract = "0xdAC17F958D2ee523a2206206994597C13D831ec7"
	//token.key.Opt.Params = &USDTParams

	return &usdt{eth: token}
}
