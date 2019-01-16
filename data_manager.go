package endurdatum

// DataManager defines a mininal implementation required for managing data.
type DataManager interface {
	Add(id interface{}, item interface{})
	Get(id interface{}, retVal interface{})
	Update(id interface{}, newItem interface{})
	Remove(id interface{})
	Filter(filter *FilterRequest) []interface{}
	Length() int
	Clear()
}
