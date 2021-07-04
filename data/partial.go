package data

import "fmt"

type Partial struct {
	Model
	LauncherId string `gorm:"-" json:"launcher_id"`
	Timestamp  int    `gorm:"-" json:"timestamp"`
	Difficulty int    `gorm:"-" json:"difficulty"`
}

// getPartial get all partial
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

// GetSizePartial calculate size from partial
func GetSizePartial(LauncherId string) ([]*Partial, error) {
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
