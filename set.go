package goset

type Set interface {
	Add(n interface{}) bool
	Remove(n interface{}) bool
	Contains(n interface{}) bool
	Length() int
	ToSlice() []interface{}
}

func NewSet() Set {
	return newMapSet()
}
