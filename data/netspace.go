package data

import (
	"fmt"
	"time"
	"chia_api/redis"
	"chia_api/utils"
)

// GetNetSpaceByLauncherId calculate size from partial share
func GetNetSpaceByLauncherId(LauncherId string) (float64, error) {
	db := GetConn()
	defer db.Close()

	redisValue := redis.GetFromToRedis(0, "netspace:"+LauncherId)
	redisNetspace := utils.StringToFloat(redisValue)
	if redisNetspace != 0.0 {
		return redisNetspace, nil
	}

	toreturn := []*Partial{}
	t := time.Now()
	timeToCheck := int64(43200)
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
		size = (float64(count) / (float64(timeToCheck) * ((5 / averageDifficulty) / 43200 / 106364865085.00))) / 1099511627776
	}
	fmt.Printf("db %f\n", size)
	errs := db.GetErrors()
	if len(errs) > 0 {
		return 0, errs[0]
	}

	redis.WriteToRedis(0, "netspace:"+LauncherId, utils.FloatToString(size))

	return size, nil
}

// GetNetSpaceTotal calculate size from partial
func GetNetSpaceTotal() (float64, error) {
	db := GetConn()
	defer db.Close()

	redisValue := redis.GetFromToRedis(0, "netspacetotal")
	redisNetspace := utils.StringToFloat(redisValue)
	if redisNetspace != 0.0 {
		return redisNetspace, nil
	}

	partial := []*Partial{}
	t := time.Now()
	timeToCheck := int64(21600)
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
	}
	errs := db.GetErrors()
	if len(errs) > 0 {
		return 0, errs[0]
	}

	redis.WriteToRedis(0, "netspacetotal", utils.FloatToString(size))

	return size, nil
}
