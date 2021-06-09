package data

type MiningStatPool struct {
	Fee          string `gorm:"-" json:"Fees"`
	FarmerOnline int    `gorm:"-" json:"farmer_online"`
	Logo         string `gorm:"-" json:"logo"`
	PoolName     string `gorm:"-" json:"pool_name"`
	Power        string `gorm:"-" json:"power"`
	Mode         string `gorm:"-" json:"mode"`
}

// GetMiningStatPool return structure for minig stat pool
func GetMiningStatPool() (*MiningStatPool, error) {
	toreturn := MiningStatPool{}
	toreturn.Fee = GetFees()
	toreturn.FarmerOnline, _ = GetFarmersCount()
	toreturn.Logo = GetLogo()
	toreturn.PoolName = GetName()
	toreturn.Power = GetPower()
	toreturn.Mode = GetMode()
	return &toreturn, nil
}
