package data

type Farmer struct {
	Model
	LauncherId                       string `gorm:"-" json:"launcher_id"`
	AuthenticationPublicKey          string `gorm:"-" json:"authentication_public_key"`
	AuthenticationPublicKeyTimestamp int    `gorm:"-" json:"authentication_public_key_timestamp"`
	OwnerPublicKey                   string `gorm:"-" json:"owner_public_key"`
	TargetPuzzleHash                 string `gorm:"-" json:"target_puzzle_hash"`
	RelativeLockHeight               int    `gorm:"-" json:"relative_lock_height"`
	P2SingletonPuzzleHash            string `gorm:"-" json:"p2_singleton_puzzle_hash"`
	BlockchainHeight                 int    `gorm:"-" json:"blockcxhain_height"`
	SingletonCoinId                  string `gorm:"-" json:"singleton_coin_id"`
	Points                           int    `gorm:"-" json:"points"`
	Difficulty                       int    `gorm:"-" json:"difficulty"`
	PoolPayoutInstructions           string `gorm:"-" json:"pool_payout_instructions"`
	IsPoolMember                     bool   `gorm:"-" json:"is_pool_member"`
}

// GetFarmer get farmer from launcher_id.
func GetFarmer(LauncherId string) (*Farmer, error) {
	db := GetConn()
	defer db.Close()
	toreturn := Farmer{}
	db.Raw("SELECT * FROM farmer where launcher_id=\"" + LauncherId + "\"").Scan(&toreturn)
	errs := db.GetErrors()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return &toreturn, nil
}

// GetFarmers get all farmer
func GetTest(LauncherId string) (*Farmer, error) {
	db := GetConn()
	defer db.Close()
	toreturn := &Farmer{}
	db.Where("launcher_id = ?", LauncherId).Find(&toreturn)
	errs := db.GetErrors()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return toreturn, nil
}

// GetFarmers get all farmer
func GetFarmers() ([]*Farmer, error) {
	db := GetConn()
	defer db.Close()
	toreturn := []*Farmer{}
	db.Raw("SELECT * FROM farmer").Scan(&toreturn)
	errs := db.GetErrors()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return toreturn, nil
}

// GetFarmers get all farmer
func GetFarmersCount() (int, error) {
	db := GetConn()
	defer db.Close()
	var toreturn int
	db.Table("farmer").Count(&toreturn)
	errs := db.GetErrors()
	if len(errs) > 0 {
		return toreturn, errs[0]
	}
	return toreturn, nil
}
