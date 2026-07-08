package solution3894

func trafficSignal(timer int) string {
	switch {
	case timer == 0:
		return "Green"
	case timer == 30:
		return "Orange"
	case 30 < timer && timer <= 90:
		return "Red"
	default:
		return "Invalid"
	}
}
