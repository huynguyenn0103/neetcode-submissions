func destCity(paths [][]string) string {
	mapCityToHaveOutTrip := make(map[string]bool)
	for _, path := range paths{
		from, to := path[0], path[1]
		mapCityToHaveOutTrip[from]=true

		if !mapCityToHaveOutTrip[to]{
			mapCityToHaveOutTrip[to]=false
		}
	}
	for key, val := range mapCityToHaveOutTrip{
		if !val{
			return key
		}
	}
	return ""
}
