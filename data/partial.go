package data

import (
	"fmt"
)

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by Partial to `partial`
func (Partial) TableName() string {
	return "partial"
}

type Partial struct {
	LauncherId string `gorm:"launcher_id" json:"launcher_id"`
	Timestamp  int64  `gorm:"timestamp" json:"timestamp"`
	Difficulty int64  `gorm:"difficulty" json:"difficulty"`
}

type SolotPlot struct {
	LauncherId string `gorm:"-" json:"launcher_id"`
	Point      int    `gorm:"-" json:"pointPerHour"`
}

// GetPartials get all partial
func GetPartials() ([]*Partial, error) {
	db := GetConn()
	defer db.Close()
	toreturn := []*Partial{}
	db.Raw("SELECT * FROM partial").Scan(&toreturn)
	errs := db.GetErrors()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return toreturn, nil
}

// Last seen from launcher_id.
func GetLastSeen(LauncherId string) (int64, error) {
	db := GetConn()
	defer db.Close()
	var toreturn int64
	fmt.Println()
	db.Raw("SELECT timestamp FROM partial where launcher_id=\"" + LauncherId + "\" ORDER BY timestamp DESC LIMIT 1").Row().Scan(&toreturn)

	errs := db.GetErrors()
	if len(errs) > 0 {
		return 0, errs[0]
	}
	return toreturn, nil
}

// GetPartial get partial from launcher_id.
func GetPartial(LauncherId string) ([]*Partial, error) {
	db := GetConn()
	defer db.Close()
	toreturn := []*Partial{}
	fmt.Println()
	db.Raw("SELECT * FROM partial where launcher_id=\"" + LauncherId + "\"").Scan(&toreturn)
	errs := db.GetErrors()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return toreturn, nil
}

// GetTotalPoint return total of points
func GetTotalPoint() (int, error) {
	db := GetConn()
	defer db.Close()
	var points int
	db.Raw("SELECT SUM(points) FROM farmer").Row().Scan(&points)
	errs := db.GetErrors()
	if len(errs) > 0 {
		return 0, errs[0]
	}
	return points, nil
}

// GetPoints Value
func GetValuePoint() (float64, error) {
	var points int
	points, _ = GetTotalPoint()
	var value float64
	value = float64(1750000000000 / points)

	return value, nil
}

// NewPArtial returns a Admin pointer.
func NewPArtial(launcherId string, timestamp int64, difficulty int64) *Partial {
	p := &Partial{}
	p.LauncherId = launcherId
	p.Timestamp = timestamp
	p.Difficulty = difficulty
	return p
}

// AddSoloPartial calculate size from partial
func (p *Partial) AddSoloPartial() error {
	db := GetConn()
	defer db.Close()
	db = db.Save(p)
	return db.Error
}
