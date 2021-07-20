package data

import (
	"encoding/json"
	"net/http"
	"time"
)

type Coins struct {
	Txns []struct {
		ID        string `json:"id"`
		Coin      string `json:"coin"`
		Amount    int    `json:"amount"`
		From      string `json:"from"`
		Spent     bool   `json:"spent"`
		To        string `json:"to"`
		Timestamp int    `json:"timestamp"`
		Block     struct {
			ID         string `json:"id"`
			Height     int    `json:"height"`
			HeaderHash string `json:"header_hash"`
		} `json:"block"`
	} `json:"txns"`
}

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
	coins := Coins{}
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get("https://xchscan.com/api/txns?limit=1&offset=0")
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&coins)
	if err != nil {
		return nil, err
	}
	toreturn.Data.LastBlocks = []LastBlocks{{Height: coins.Txns[0].Block.Height, Timestamp: coins.Txns[0].Timestamp}}
	toreturn.Status = "OK"
	return &toreturn, nil
}
