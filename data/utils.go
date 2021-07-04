package data

func GetFees() (float64, string) {
	return 0.001, "PPLNS"
}

func GetName() string {
	return "French Farmers"
}

func GetLogo() string {
	return API_HOST + ":" + API_PORT + "/logo.svg"
}
