package data

import (
	"chia_api/utils"
)

type PoolAdmin struct {
	Model
	LauncherId  string `gorm:"launcher_di" json:"launcher_id"`
	IsPoolAdmin bool   `gorm:"is_pool_admin" json:"is_pool_admin"`
	IsPoolModo  bool   `gorm:"is_pool_modo" json:"is_pool_modo"`
	Password    string `gorm:"password" json:"password"`
}

func (f *PoolAdmin) Save() error {
	db := GetConn()
	defer db.Close()
	db = db.Save(f)
	return db.Error
}

// NewAdmin returns a Admin pointer.
func NewAdmin(launcherId string) *PoolAdmin {
	u := &PoolAdmin{}
	u.LauncherId = launcherId
	u.IsPoolAdmin = true
	u.IsPoolModo = true
	// encrypt password:
	/*	password := make([]byte, 15)  On utiliseras ça ensuite, histoire de pas store un password en clair
		for i := 0; i < 15; i++ {
			n := rand.Intn(58) + 32
			password = append(password, byte(n))
		}
		pwd, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		u.Password = string(pwd)*/
	u.Password, _ = utils.GenerateRandowStringByLenght(20)
	return u
}
