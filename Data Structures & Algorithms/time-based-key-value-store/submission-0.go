type TimeMap struct {
	TimeValueMap map[string][]TimeValue
}

type TimeValue struct {
	time  int
	value string
}

func Constructor() TimeMap {
	return TimeMap{
		TimeValueMap: make(map[string][]TimeValue),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	timeValues, _ := this.TimeValueMap[key]
	timeValues = append(timeValues, TimeValue{
		time:  timestamp,
		value: value,
	})
	this.TimeValueMap[key] = timeValues
}

func (this *TimeMap) Get(key string, timestamp int) string {
	timeValues, exists := this.TimeValueMap[key]
	if !exists {
		return ""
	}

	low, high := 0, len(timeValues)-1
	var candidate string

	for low <= high {
		mid := low + (high-low)/2

		if timeValues[mid].time <= timestamp {
			candidate = timeValues[mid].value
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return candidate
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */