package data

// TableName overrides the table name used by User to `profiles`
func (Farmer) TableName() string {
	return "farmer"
}

type Farmer struct {
	LauncherId              string  `gorm:"launcher_id" json:"launcher_id"`
	P2SingletonPuzzleHash   string  `gorm:"p2_singleton_puzzle_hash" json:"p2_singleton_puzzle_hash"`
	DelayTime               int64   `gorm:"delay_time" json:"delay_time"`
	DelayPuzzleHash         string  `gorm:"delay_puzzle_hash" json:"delay_puzzle_hash"`
	AuthenticationPublicKey string  `gorm:"authentication_public_key" json:"authentication_public_key"`
	SingletonTip            []byte  `gorm:"singleton_tip" json:"singleton_tip"`
	SingletonTipState       []byte  `gorm:"singleton_tip_state" json:"singleton_tip_state"`
	Points                  int     `gorm:"points" json:"points"`
	Difficulty              int     `gorm:"difficulty" json:"difficulty"`
	PayoutInstructions      string  `gorm:"payout_instructions" json:"payout_instructions"`
	IsPoolMember            bool    `gorm:"is_pool_member" json:"is_pool_member"`
	FarmerNetSpace          float64 `gorm:"farmer_netspace" json:"farmer_netspace"`
	LastSeen                int64   `gorm:"farmer_lastseen" json:"farmer_lastseen"`
}

// GetFarmer get farmer from launcher_id.
func GetFarmer(LauncherId string) (*Farmer, error) {
	db := GetConn()
	defer db.Close()
	toreturn := Farmer{}
	db.Raw("SELECT * FROM farmer where launcher_id=\"" + LauncherId + "\"").Scan(&toreturn)
	errs := db.GetErrors()
	Netspace, _ := GetNetSpaceByLauncherId(LauncherId)
	toreturn.FarmerNetSpace = Netspace
	toreturn.LastSeen, _ = GetLastSeen(LauncherId)

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
	for _, element := range toreturn {
		element.FarmerNetSpace, _ = GetNetSpaceByLauncherId(element.LauncherId)
	}
	errs := db.GetErrors()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return toreturn, nil
}

// GetFarmers top farmer
func GetTopFarmers() ([]*Farmer, error) {
	db := GetConn()
	defer db.Close()
	toreturn := []*Farmer{}

	db.Raw("SELECT * FROM farmer ORDER BY points DESC LIMIT 10").Scan(&toreturn)
	for _, element := range toreturn {
		element.FarmerNetSpace, _ = GetNetSpaceByLauncherId(element.LauncherId)
	}
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
	db.Table("farmer").Where("points > ?", "0").Count(&toreturn)
	errs := db.GetErrors()
	if len(errs) > 0 {
		return toreturn, errs[0]
	}
	return toreturn, nil
}
