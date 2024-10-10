package convertthetemperature

func ConvertTemperature(celsius float64) []float64 {
	result := []float64{}
	result = append(result, celsius+273.15, celsius*1.80+32.00)
	return result
}
