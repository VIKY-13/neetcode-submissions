import "slices"
import "cmp"
func carFleet(target int, position []int, speed []int) int {
	type car struct{
		position int
		speed int
	}

	cars := make([]car, len(position))
	for i,_ := range position{
		cars[i] = car{
        	position: position[i],
        	speed: speed[i],
    	}
	}

	slices.SortFunc(cars, func(i,j car) int{
		return cmp.Compare(j.position, i.position)
	})

	fleetCount := 0
	var currentFleetTime float64
	for _, car := range cars{
		time := float64(target-car.position) / float64(car.speed)

		if time > currentFleetTime {
			fleetCount++
			currentFleetTime = time
		}
	}

	return fleetCount
}
