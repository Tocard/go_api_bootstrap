package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Coins struct {
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Total int `json:"total"`
	Data  []struct {
		Hash              string `json:"hash"`
		Height            int    `json:"height"`
		Time              int    `json:"time"`
		TxCount           int    `json:"txCount"`
		Generator         string `json:"generator"`
		Previousblockhash string `json:"previousblockhash"`
		BlockDetail       struct {
			FarmerRewardAddress string `json:"farmerRewardAddress"`
			PoolTargetAddress   string `json:"poolTargetAddress"`
			Time                int    `json:"time"`
			BlockNo             int    `json:"blockNo"`
			Fees                int    `json:"fees"`
			PoolContractPuzzle  string `gorm:"default:true" json:"poolContractPuzzleHash"`
		} `json:"blockDetail"`
		Nextblockhash string `json:"nextblockhash,omitempty"`
	} `json:"data"`
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
	for p := 0; p < 300; p++ { // =4 minute d'exec
		url := fmt.Sprintf("https://chia.tt/api/chia/blockchain/block?page=%d&count=20", p)
		r, err := myClient.Get(url)
		if err != nil {
			return nil, err
		}
		err = json.NewDecoder(r.Body).Decode(&coins)
		if err != nil {
			return nil, err
		}
		for i := 0; i < len(coins.Data); i++ {
			var launcherid string
			if coins.Data[i].BlockDetail.PoolContractPuzzle != "" {
				launcherid, _ = GetFarmerFromP2SingletonPuzzleHash(coins.Data[i].BlockDetail.PoolContractPuzzle[2:])
			} else {
				launcherid, _ = GetFarmerFromRewardAdress(coins.Data[i].BlockDetail.FarmerRewardAddress)
			}
			if launcherid != "" {
				r.Body.Close()
				toreturn.Data.LastBlocks = []LastBlocks{{Height: coins.Data[i].Height, Timestamp: coins.Data[i].Time}}
				values := NewWinBlock(coins.Data[i].Time, coins.Data[i].Height, launcherid)
				json_data, err := json.Marshal(values)

				if err != nil {
					log.Fatal(err)
				}

				http.Post("https://localhost:8081/block/new_block_discord", "application/json",
					bytes.NewBuffer(json_data))
			}
		}
		r.Body.Close()
	}
	toreturn.Status = "OK"
	return &toreturn, nil
}
