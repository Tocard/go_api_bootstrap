package data

import (
	"fmt"
	"time"
)

type Partial struct {
	Model
	LauncherId string `gorm:"-" json:"launcher_id"`
	Timestamp  int    `gorm:"-" json:"timestamp"`
	Difficulty int64  `gorm:"-" json:"difficulty"`
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

// GetNetSpaceByLauncherId calculate size from partial share
func GetNetSpaceByLauncherId(LauncherId string) (float64, error) {
	db := GetConn()
	defer db.Close()
	toreturn := []*Partial{}
	t := time.Now()
	timeToCheck := int64(86400)
	query := fmt.Sprintf("SELECT * FROM partial where launcher_id=\"%s\" AND timestamp >=%d", LauncherId, t.Unix()-timeToCheck)
	db.Raw(query).Scan(&toreturn)
	count := int64(len(toreturn))
	averageDifficulty := float64(0)
	for _, element := range toreturn {
		averageDifficulty += float64(element.Difficulty)
	}
	size := float64(0)
	if count > 0 {
		averageDifficulty = averageDifficulty / float64(count)
		size = float64(count) / (float64(timeToCheck) * ((10 / averageDifficulty) / 86400.00 / 106364865085.00))
		debug1 := fmt.Sprintf("farmspace = %s, launcher_id %s diffifculty= %f, timetocheck= %d", lenReadable(int(size), 2, true), LauncherId, averageDifficulty, timeToCheck)
		fmt.Println(debug1)

	}
	errs := db.GetErrors()
	if len(errs) > 0 {
		return 0, errs[0]
	}
	return size, nil
}

// GetNetSpaceTotal calculate size from partial
func GetNetSpaceTotal() (float64, error) {
	db := GetConn()
	defer db.Close()
	partial := []*Partial{}
	t := time.Now()
	timeToCheck := int64(86400)
	query := fmt.Sprintf("SELECT * FROM partial where timestamp >=%d", t.Unix()-timeToCheck)
	db.Raw(query).Scan(&partial)
	size := float64(0)
	launcherIds := []string{}
	for _, element := range partial {
		launcherIds = append(launcherIds, element.LauncherId)
	}
	launcherIds = unique(launcherIds)
	for _, launcherId := range launcherIds {
		sizeTmp, _ := GetNetSpaceByLauncherId(launcherId)
		size += sizeTmp
		debug1 := fmt.Sprintf("total = %s", lenReadable(int(size), 2, true))
		fmt.Println(debug1)
	}
	errs := db.GetErrors()
	if len(errs) > 0 {
		return 0, errs[0]
	}
	return size, nil
}
