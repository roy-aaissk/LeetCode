type ParkingSystem struct {
	Big    int
	Medium int
	Small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		Big:    big,
		Medium: medium,
		Small:  small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.Big >= 1 {
			this.Big--
			return true
		}
	case 2:
		if this.Medium >= 1 {
			this.Medium--
			return true
		}
	case 3:
		if this.Small >= 1 {
			this.Small--
			return true
		}
	default:
		return false
	}

	return false
}
