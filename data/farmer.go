package data

// TableName overrides the table name used by User to `profiles`
func (Farmer) TableName() string {
	return "farmer"
}

type Farmer struct {
	LauncherId                       string `gorm:"launcher_id" json:"launcher_id"`
	AuthenticationPublicKey          string `gorm:"authentication_public_key" json:"authentication_public_key"`
	AuthenticationPublicKeyTimestamp int    `gorm:"authentication_public_key_timestamp" json:"authentication_public_key_timestamp"`
	OwnerPublicKey                   string `gorm:"owner_public_key" json:"owner_public_key"`
	TargetPuzzleHash                 string `gorm:"target_puzzle_hash" json:"target_puzzle_hash"`
	RelativeLockHeight               int    `gorm:"relative_lock_height" json:"relative_lock_height"`
	P2SingletonPuzzleHash            string `gorm:"p2_singleton_puzzle_hash" json:"p2_singleton_puzzle_hash"`
	BlockchainHeight                 int    `gorm:"blockcxhain_height" json:"blockcxhain_height"`
	SingletonCoinId                  string `gorm:"singleton_coin_id" json:"singleton_coin_id"`
	Points                           int    `gorm:"points" json:"points"`
	Difficulty                       int    `gorm:"difficulty" json:"difficulty"`
	PoolPayoutInstructions           string `gorm:"pool_payout_instructions" json:"pool_payout_instructions"`
	IsPoolMember                     bool   `gorm:"is_pool_member" json:"is_pool_member"`
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

// UpdateFarmerPoint update points of user
func UpdateFarmerPoint(farmer *Farmer) error {
	db := GetConn()
	defer db.Close()
	db.Model(&Farmer{}).Where("launcher_id = ?", farmer.LauncherId).Update("points", farmer.Points)
	errs := db.GetErrors()
	if len(errs) > 0 {
		return errs[0]
	}
	return nil
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
