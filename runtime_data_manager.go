package endurdatum

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
func (manager RuntimeDataManager) Get(id interface{}) interface{} {
	return manager.Items[id]
}

// Update updates an item in the runtime.
func (manager RuntimeDataManager) Update(id interface{}, item interface{}) {
	manager.Items[id] = item
}

// Remove removes an item from the runtime.
func (manager RuntimeDataManager) Remove(id interface{}) {
	delete(manager.Items, id)
}

// Clear removes all items from the runtime.
func (manager *RuntimeDataManager) Clear() {
	manager.Items = make(map[interface{}]interface{}, 0)
}
