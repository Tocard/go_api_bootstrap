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
	} `json:"data"`
}

// GetMiningStatPool return structure for minig stat pool
func GetMiningStatPool() (*MiningStatPool, error) {
	toreturn := MiningStatPool{}
	fees, feestype := GetFees()
	toreturn.Data.PoolStats.Farmers, _ = GetFarmersCount()
	toreturn.Data.PoolStats.CurrentFee = fees
	toreturn.Data.PoolStats.CurrentFeeType = feestype
	toreturn.Data.PoolStats.PoolSpaceTiB, _ = GetNetSpaceTotal()
	toreturn.Data.PoolStats.PoolSpaceTiB *= 0.0000000000009094947
	toreturn.Data.LastBlocks = []string{}
	toreturn.Status = "OK"
	return &toreturn, nil
}
