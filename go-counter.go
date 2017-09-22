package go_counter

func init() {

}

type GoCounter struct {
	counterMap map[string]uint
}

func (counter GoCounter) Increment(alias string) {
	if counter.counterMap[alias] == 0 {
		counter.counterMap[alias]++
	}
}

func (counter GoCounter) GetCounter(alias string) uint{
	return counter.counterMap[alias]
}
