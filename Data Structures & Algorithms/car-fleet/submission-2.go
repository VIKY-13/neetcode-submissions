import "cmp"
import "slices"

func carFleet(target int, position []int, speed []int) int {
	type car struct{
		pos float64
		speed float64
	}

	cars := make([]car, len(position))
	for i := range position{
		cars[i] = car{
			pos: float64(position[i]),
			speed: float64(speed[i]),
		}
	}

	slices.SortFunc(cars, func(i, j car) int{
		return cmp.Compare(j.pos, i.pos)
	})

	fleetCount := 0
	var currentTime float64
	for _, car := range cars{
		time := (float64(target) - car.pos) / car.speed
		if time > currentTime{
			fleetCount++
			currentTime = time
		}
	}

	return fleetCount
}
