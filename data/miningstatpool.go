package data

type MiningStatPool struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type LastBlocks struct {
	Height    int `json:"height"`
	Timestamp int `json:"timestamp"`
}
type PoolStats struct {
	PoolSpaceTiB   float64 `json:"poolSpaceTiB"`
	Farmers        int     `json:"farmers"`
	CurrentFeeType string  `json:"currentFeeType"`
	CurrentFee     float64 `json:"currentFee"`
}
type Data struct {
	LastBlocks []LastBlocks `json:"lastBlocks"`
	PoolStats  PoolStats    `json:"poolStats"`
}

// GetMiningStatPool return structure for minig stat pool
func GetMiningStatPool() (*MiningStatPool, error) {
	toreturn := MiningStatPool{}
	fees, feestype := GetFees()
	toreturn.Data.PoolStats.Farmers, _ = GetFarmersCount()
	toreturn.Data.PoolStats.CurrentFee = fees
	toreturn.Data.PoolStats.CurrentFeeType = feestype
	toreturn.Data.PoolStats.PoolSpaceTiB, _ = GetNetSpaceTotal()
	toreturn.Data.PoolStats.PoolSpaceTiB += LoadFileSoloPlot()
	toreturn.Data.LastBlocks = []LastBlocks{{Height: 536606, Timestamp: 1625640238}}
	toreturn.Status = "OK"
	return &toreturn, nil
}
