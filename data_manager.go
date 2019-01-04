package endurdatum

// DataManager defines a mininal implementation required for managing data.
type DataManager interface {
	Add(id interface{}, item interface{})
	Get(id interface{}, returnType interface{}) interface{}
	Update(id interface{}, newItem interface{})
	Remove(id interface{})
	Clear()
}