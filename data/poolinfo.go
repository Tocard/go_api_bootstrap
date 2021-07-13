package data


type PoolInfo struct {
	Status string `json:"status"`
	Data DataPool `json:"data"`
}
type DataPool struct {
	LastBlocks []LastBlocks `json:"lastBlocks"`
	PoolSpaceTiB   float64 `json:"poolSpaceTiB"`
	Farmers        int     `json:"farmers"`
	CurrentFeeType string  `json:"currentFeeType"`
	CurrentFee     float64 `json:"currentFee"`
	TotalPoints    int `json:"totalPoints"`
	PointValue 	   float64 `json:"pointValue"`
}



// GetMiningStatPool return structure for minig stat pool
func GetPoolInfo() (*PoolInfo, error) {
	toreturn := PoolInfo{}
	fees, feestype := GetFees()
	toreturn.Data.Farmers, _ = GetFarmersCount()
	toreturn.Data.CurrentFee = fees
	toreturn.Data.CurrentFeeType = feestype
	toreturn.Data.PoolSpaceTiB, _ = GetNetSpaceTotal()
	toreturn.Data.PoolSpaceTiB += LoadFileSoloPlot()
	toreturn.Data.LastBlocks = []LastBlocks{{Height: 536606, Timestamp: 1625640238}}
	toreturn.Data.TotalPoints, _ = GetTotalPoint()
	toreturn.Data.PointValue, _ = GetValuePoint()
	toreturn.Status = "OK"

	return &toreturn, nil
}


