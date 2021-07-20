package data

type WinBlock struct {
	Height     int    `json:"height"`
	Timestamp  int    `json:"timestamp"`
	LauncherId string `json:"string"`
}

// NewWinBlock returns a WinBlock pointer.
func NewWinBlock(timestamp, height int, launcherId string) *WinBlock {
	u := &WinBlock{}
	u.LauncherId = launcherId
	u.Timestamp = timestamp
	u.Height = height

	return u
}
