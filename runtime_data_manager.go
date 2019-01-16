package endurdatum

import (
	"reflect"
	"strings"
)

// RuntimeDataManager manages data in the applications runtime.
// This would be useful for developer testing. This would however
// not be practical for production use since data would be lost on application shutdown.
type RuntimeDataManager struct {
	Items map[interface{}]interface{}
}

// CreateRuntimeDataManager creates a new runtime data manager instance.
func CreateRuntimeDataManager() RuntimeDataManager {
	return RuntimeDataManager{make(map[interface{}]interface{}, 0)}
}

// Add adds an item to the runtime.
func (manager RuntimeDataManager) Add(id interface{}, item interface{}) {
	manager.Items[id] = item
}

// Get gets an item based off of its id from the runtime.
func (manager RuntimeDataManager) Get(id interface{}, retVal interface{}) {
	retValType := reflect.ValueOf(retVal)
	retValVal := reflect.Indirect(retValType)

	item := reflect.ValueOf(manager.Items[id])
	retValVal.Set(reflect.Indirect(item))
}

// Update updates an item in the runtime.
func (manager RuntimeDataManager) Update(id interface{}, item interface{}) {
	manager.Items[id] = item
}

// Remove removes an item from the runtime.
func (manager RuntimeDataManager) Remove(id interface{}) {
	delete(manager.Items, id)
}

// Length lists the length of items in the runtime.
func (manager RuntimeDataManager) Length() int {
	return len(manager.Items)
}

// Filter applies a filter to the runtime items and returns a subset of the applicable items
func (manager RuntimeDataManager) Filter(filter *FilterRequest) []interface{} {
	retVal := make([]interface{}, 0)

	for _, v := range manager.Items {
		if inFilterRequest(v, filter) {
			retVal = append(retVal, v)
		}
	}

	return retVal
}

func inFilterRequest(item interface{}, filter *FilterRequest) bool {
	if filter == nil {
		return true
	}

	if isOrRequest(filter) {
		for _, k := range filter.Filters {
			if inComplexFilter(item, k) {
				return true
			}
		}

		return false
	}

	for _, k := range filter.Filters {
		if !inComplexFilter(item, k) {
			return false
		}
	}

	return true
}

func inComplexFilter(item interface{}, filter ComplexFilter) bool {
	if filter.Children != nil {
		if isOrFilter(filter) {
			for _, v := range filter.Children {
				if inComplexFilter(item, *v) {
					return true
				}
			}

			return false
		}

		for _, v := range filter.Children {
			if !inComplexFilter(item, *v) {
				return false
			}
		}

		return true
	}

	if isOrFilter(filter) {
		for _, v := range filter.Filters {
			if itemInFilter(item, v) {
				return true
			}
		}

		return false
	}

	for _, v := range filter.Filters {
		if !itemInFilter(item, v) {
			return false
		}
	}

	return true
}

func itemInFilter(item interface{}, filter Filter) bool {
	val := getItemPropertyValue(filter.Property, item)
	if isEqualsComparision(filter) {
		return strings.EqualFold(filter.Value, val)
	}

	if isContainsComparision(filter) {
		return strings.Contains(strings.ToLower(val), strings.ToLower(filter.Value))
	}

	if isNotEqualsComparision(filter) {
		return !strings.EqualFold(filter.Value, val)
	}

	return false
}

func getItemPropertyValue(propertyName string, item interface{}) string {
	s := reflect.ValueOf(item).Elem()
	t := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)

		if t.Field(i).Name == propertyName {
			return f.Interface().(string)
		}
	}

	return ""
}

// Clear removes all items from the runtime.
func (manager RuntimeDataManager) Clear() {
	for i := range manager.Items {
		delete(manager.Items, i)
	}
}
