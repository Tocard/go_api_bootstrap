package data

type SolotPlot struct {
	LauncherId string  `gorm:"-" json:"launcher_id"`
	Point      float64 `gorm:"-" json:"pointPerHour"`
}

// GetTotalPoint return total of points
func GetTotalPoint() (float64, error) {
	db := GetConn()
	defer db.Close()
	var points float64
	db.Raw("SELECT SUM(points) FROM farmer").Row().Scan(&points)
	errs := db.GetErrors()
	if len(errs) > 0 {
		return 0, errs[0]
	}
	return points, nil
}

// GetValuePoint return point's Value
func GetValuePoint() (float64, error) {
	points, _ := GetTotalPoint()
	value := 1750000000000.0 / points
	return value, nil
}

// UpdateFarmerPoint update points of farmer
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
