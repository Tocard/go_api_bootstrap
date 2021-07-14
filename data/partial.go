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

// GetLastSeen Last seen from launcher_id.
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
	db.Raw("SELECT * FROM partial where launcher_id=\"" + LauncherId + "\"").Scan(&toreturn)
	errs := db.GetErrors()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return toreturn, nil
}

// NewPartial returns a NewPartial pointer.
func NewPartial(launcherId string, timestamp int64, difficulty int64) *Partial {
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
