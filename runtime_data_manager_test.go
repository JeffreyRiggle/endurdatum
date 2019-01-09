package endurdatum

import "testing"

type TestItem struct {
	Id    int64
	Value string
}

func TestAddItem(t *testing.T) {
	var manager = CreateRuntimeDataManager()

	var item = TestItem{1, "This is a test"}
	manager.Add(0, item)

	if len(manager.Items) != 1 {
		t.Error(
			"For", manager,
			"expected", 1,
			"got", len(manager.Items))
	}

	castItem, ok := manager.Items[0].(TestItem)
	if !ok {
		t.Error(
			"For", manager.Items[0],
			"expected TestItem",
			"got error")
	}

	if castItem.Id != 1 {
		t.Error(
			"For", castItem.Id,
			"expected", 1,
			"got", castItem.Id)
	}
}

func TestGetItem(t *testing.T) {
	var manager = CreateRuntimeDataManager()

	var item = TestItem{1, "This is a test"}
	manager.Add(0, item)

	retVal, ok := manager.Get(0).(TestItem)

	if !ok {
		t.Error(
			"For", manager.Items[0],
			"expected TestItem",
			"got error")
	}

	if retVal.Id != 1 {
		t.Error(
			"For", retVal.Id,
			"expected", 1,
			"got", retVal.Id)
	}
}

func TestUpdateItem(t *testing.T) {
	var manager = CreateRuntimeDataManager()

	var item = TestItem{1, "This is a test"}
	manager.Add(0, item)
	var itemup = TestItem{2, "This is a different test"}
	manager.Update(0, itemup)

	retVal, ok := manager.Get(0).(TestItem)

	if !ok {
		t.Error(
			"For", manager.Items[0],
			"expected TestItem",
			"got error")
	}

	if retVal.Value != "This is a different test" {
		t.Error(
			"For", retVal.Value,
			"expected ", "This is a different test",
			"got", retVal.Value)
	}
}

func TestRemoveItem(t *testing.T) {
	var manager = CreateRuntimeDataManager()

	var item = TestItem{1, "This is a test"}
	manager.Add(0, item)
	manager.Remove(0)

	if len(manager.Items) != 0 {
		t.Error(
			"For", manager,
			"expected", 0,
			"got", len(manager.Items))
	}
}

func TestClear(t *testing.T) {
	var manager = CreateRuntimeDataManager()

	var item = TestItem{1, "This is a test"}
	manager.Add(0, item)
	var item2 = TestItem{2, "This is a test2"}
	manager.Add(1, item2)

	manager.Clear()
	if len(manager.Items) != 0 {
		t.Error(
			"For", manager,
			"expected", 0,
			"got", len(manager.Items))
	}
}

func TestLength(t *testing.T) {
	var manager = CreateRuntimeDataManager()
	manager.Add(0, TestItem{1, "some test"})
	len1 := manager.Length()
	if len1 != 1 {
		t.Error(
			"For", manager,
			"expected", 1,
			"got", len1)
	}

	manager.Add(1, TestItem{2, "some test"})
	len2 := manager.Length()
	if len2 != 2 {
		t.Error(
			"For", manager,
			"expected", 2,
			"got", len2)
	}

	manager.Add(2, TestItem{3, "some test"})
	len3 := manager.Length()
	if len3 != 3 {
		t.Error(
			"For", manager,
			"expected", 3,
			"got", len3)
	}
}
