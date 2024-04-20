package goset

type mapset map[interface{}]struct{}

func newMapSet() *mapset {
	set := make(mapset)
	return &set
}

func (set *mapset) Add(n interface{}) bool {
	if _, ok := (*set)[n]; ok {
		return false
	}

	(*set)[n] = struct{}{}
	return true
}

func (set *mapset) Remove(n interface{}) (ret bool) {
	ret = false
	if _, ok := (*set)[n]; ok {
		delete((*set), n)
		ret = true
	}
	return
}

func (set *mapset) Contains(n interface{}) bool {
	if _, ok := (*set)[n]; ok {
		return true
	}
	return false
}

func (set *mapset) Length() int {
	return len(*set)
}

func (set *mapset) ToSlice() (slice []interface{}) {
	slice = *new([]interface{})
	for key := range *set {
		slice = append(slice, key)
	}
	return
}
