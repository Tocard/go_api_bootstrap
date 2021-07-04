package data

type MiningStatPool struct {
	Status string `json:"status"`
	Data   struct {
		LastBlocks []string `json:"lastBlocks"`
		PoolStats  struct {
			PoolSpaceTiB   float64 `json:"poolSpaceTiB"`
			Farmers        int     `json:"farmers"`
			CurrentFeeType string  `json:"currentFeeType"`
			CurrentFee     float64 `json:"currentFee"`
		} `json:"poolStats"`
		Xch struct {
			Usdt float64 `json:"usdt"`
			Btc  float64 `json:"btc"`
		} `json:"xch"`
	} `json:"data"`
}

// GetMiningStatPool return structure for minig stat pool
func GetMiningStatPool() (*MiningStatPool, error) {
	toreturn := MiningStatPool{}
	fees, feestype := GetFees()
	toreturn.Data.PoolStats.Farmers, _ = GetFarmersCount()
	toreturn.Data.PoolStats.CurrentFee = fees
	toreturn.Data.PoolStats.CurrentFeeType = feestype
	toreturn.Data.PoolStats.PoolSpaceTiB = 1.0
	toreturn.Data.Xch.Btc = 0
	toreturn.Data.Xch.Usdt = 0
	toreturn.Data.LastBlocks = []string{"first", "second"}
	return &toreturn, nil
}
