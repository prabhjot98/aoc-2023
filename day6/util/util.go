package util

type Race struct {
	Duration int
	Distance int
}

func GetDistance(raceDuration int, timeButtonHeld int) int {
	return (raceDuration - timeButtonHeld) * timeButtonHeld
}

func (r Race) BeatsRecord(timeButtonHeld int) bool {
	if GetDistance(r.Duration, timeButtonHeld) > r.Distance {
		return true
	}
	return false
}
