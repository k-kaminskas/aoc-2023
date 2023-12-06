package internal

import "strconv"

type Race struct {
	Time     int
	Distance int
}

func (r Race) getValidCases() (cases int) {
	for i := 1; i < r.Time; i++ {
		if i*(r.Time-i) > r.Distance {
			cases++
		}
	}
	return cases
}

// Races ----------------------------------------------------------------------------------------------------------- */

type Races []*Race

func (rs Races) Product() (product int) {
	product = 1
	for _, race := range rs {
		product *= race.getValidCases()
	}
	return product
}

func (rs Races) AggregatedProduct() int {
	timeStr, distanceStr := "", ""
	for _, race := range rs {
		timeStr += strconv.Itoa(race.Time)
		distanceStr += strconv.Itoa(race.Distance)
	}

	totalTime, _ := strconv.Atoi(timeStr)
	totalDistance, _ := strconv.Atoi(distanceStr)

	return Race{totalTime, totalDistance}.getValidCases()
}
