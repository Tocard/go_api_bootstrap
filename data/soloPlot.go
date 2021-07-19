package data

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type FloatingPartial struct {
	LauncherId  string  `gorm:"launcher_id" json:"launcher_id"`
	PartialPart float64 `gorm:"partial_part" json:"partial_part"`
}

func (f *FloatingPartial) Save() error {
	db := GetConn()
	defer db.Close()
	db = db.Save(f)
	return db.Error
}

// NewFloatingPartial returns a FloatingPartial pointer.
func NewFloatingPartial(launcherId string, partialpart float64) *FloatingPartial {
	p := &FloatingPartial{}
	p.LauncherId = launcherId
	p.PartialPart = partialpart
	return p
}

// GetFloatingPartial get FloatingPartial from launcher_id.
func GetFloatingPartial(LauncherId string) (FloatingPartial, error) {
	db := GetConn()
	defer db.Close()
	toreturn := FloatingPartial{}
	db.Model(&FloatingPartial{}).Where(
		"launcher_id = ?",
		LauncherId,
	).Find(&toreturn)
	errs := db.GetErrors()
	if len(errs) > 0 {
		return toreturn, errs[0]
	}
	return toreturn, nil
}

// UpdateFloatingPartial set FloatingPartial from launcher_id.
func UpdateFloatingPartial(launcherid string, part float64) error {
	db := GetConn()
	defer db.Close()
	db.Model(&FloatingPartial{}).Where("launcher_id = ?", launcherid).Update("partial_part", part)
	errs := db.GetErrors()
	if len(errs) > 0 {
		return errs[0]
	}
	return nil
}

func LoadFileSoloPlot() float64 {
	content, err := ioutil.ReadFile("solo_plot.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)
	Stringfloat := strings.TrimSuffix(text, "\n")
	soloNetspace, _ := strconv.ParseFloat(Stringfloat, 64)
	return soloNetspace
}
