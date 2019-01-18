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

	var retVal TestItem
	ok := manager.Get(0, &retVal)

	if !ok {
		t.Error(
			"For", ok,
			"expected", true,
			"got", ok)
	}

	if retVal.Id != 1 {
		t.Error(
			"For", retVal.Id,
			"expected", 1,
			"got", retVal.Id)
	}
}

func TestFailGetItem(t *testing.T) {
	var manager = CreateRuntimeDataManager()

	var retVal TestItem
	ok := manager.Get(0, &retVal)

	if ok {
		t.Error(
			"For", ok,
			"expected", false,
			"got", ok)
	}
}

func TestUpdateItem(t *testing.T) {
	var manager = CreateRuntimeDataManager()

	var item = TestItem{1, "This is a test"}
	manager.Add(0, item)
	var itemup = TestItem{2, "This is a different test"}
	manager.Update(0, itemup)

	var retVal TestItem
	ok := manager.Get(0, &retVal)

	if !ok {
		t.Error(
			"For", ok,
			"expected", true,
			"got", ok)
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
	var manager DataManager = CreateRuntimeDataManager()
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

func TestFilterItemEquals(t *testing.T) {
	var manager DataManager = CreateRuntimeDataManager()
	manager.Add(0, TestItem{1, "some test"})
	manager.Add(1, TestItem{2, "some other test"})
	manager.Add(2, TestItem{3, "testing test test"})
	manager.Add(3, TestItem{4, "all testers test tests"})

	filter := FilterRequest{
		Filters: []ComplexFilter{
			{
				Filters: []Filter{
					{
						Property:       "Value",
						ComparisonType: "equals",
						Value:          "some test",
					},
				},
			},
		},
		Junction: "and",
	}

	retVal := manager.Filter(&filter)

	if len(retVal) != 1 {
		t.Error(
			"For", manager,
			"expected", 1,
			"got", len(retVal))
	}

	castType, ok := retVal[0].(TestItem)

	if !ok {
		t.Error(
			"For", retVal[0],
			"expected TestItem",
			"got", castType)
	}

	if castType.Id != 1 {
		t.Error(
			"For", castType,
			"expected", 1,
			"got", castType.Id)
	}
}

func TestFilterItemNotEquals(t *testing.T) {
	var manager DataManager = CreateRuntimeDataManager()
	manager.Add(0, TestItem{1, "some test"})
	manager.Add(1, TestItem{2, "some other test"})
	manager.Add(2, TestItem{3, "testing test test"})
	manager.Add(3, TestItem{4, "all testers test tests"})

	filter := FilterRequest{
		Filters: []ComplexFilter{
			{
				Filters: []Filter{
					{
						Property:       "Value",
						ComparisonType: "notequals",
						Value:          "some test",
					},
				},
			},
		},
		Junction: "and",
	}

	retVal := manager.Filter(&filter)

	if len(retVal) != 3 {
		t.Error(
			"For", manager,
			"expected", 3,
			"got", len(retVal))
	}

	castType, ok := retVal[0].(TestItem)

	if !ok {
		t.Error(
			"For", retVal[0],
			"expected TestItem",
			"got", castType)
	}
}

func TestFilterItemContains(t *testing.T) {
	var manager DataManager = CreateRuntimeDataManager()
	manager.Add(0, TestItem{1, "some test"})
	manager.Add(1, TestItem{2, "some other test"})
	manager.Add(2, TestItem{3, "testing test test"})
	manager.Add(3, TestItem{4, "all testers test tests"})

	filter := FilterRequest{
		Filters: []ComplexFilter{
			{
				Filters: []Filter{
					{
						Property:       "Value",
						ComparisonType: "contains",
						Value:          "test",
					},
				},
			},
		},
		Junction: "and",
	}

	retVal := manager.Filter(&filter)

	if len(retVal) != 4 {
		t.Error(
			"For", manager,
			"expected", 4,
			"got", len(retVal))
	}

	castType, ok := retVal[0].(TestItem)

	if !ok {
		t.Error(
			"For", retVal[0],
			"expected TestItem",
			"got", castType)
	}
}

func TestFilterItemOrInComplexFilter(t *testing.T) {
	var manager DataManager = CreateRuntimeDataManager()
	manager.Add(0, TestItem{1, "some test"})
	manager.Add(1, TestItem{2, "some other test"})
	manager.Add(2, TestItem{3, "testing test test"})
	manager.Add(3, TestItem{4, "all testers test tests"})

	filter := FilterRequest{
		Filters: []ComplexFilter{
			{
				Filters: []Filter{
					{
						Property:       "Value",
						ComparisonType: "equals",
						Value:          "some test",
					},
					{
						Property:       "Value",
						ComparisonType: "equals",
						Value:          "some other test",
					},
				},
				Junction: "or",
			},
		},
		Junction: "and",
	}

	retVal := manager.Filter(&filter)

	if len(retVal) != 2 {
		t.Error(
			"For", manager,
			"expected", 2,
			"got", len(retVal))
	}

	castType, ok := retVal[0].(TestItem)

	if !ok {
		t.Error(
			"For", retVal[0],
			"expected TestItem",
			"got", castType)
	}
}

func TestFilterItemOrInRootFilter(t *testing.T) {
	var manager DataManager = CreateRuntimeDataManager()
	manager.Add(0, TestItem{1, "some test"})
	manager.Add(1, TestItem{2, "some other test"})
	manager.Add(2, TestItem{3, "testing test test"})
	manager.Add(3, TestItem{4, "all testers test tests"})

	filter := FilterRequest{
		Filters: []ComplexFilter{
			{
				Filters: []Filter{
					{
						Property:       "Value",
						ComparisonType: "equals",
						Value:          "some test",
					},
				},
			},
			{
				Filters: []Filter{
					{
						Property:       "Value",
						ComparisonType: "equals",
						Value:          "some other test",
					},
				},
			},
		},
		Junction: "or",
	}

	retVal := manager.Filter(&filter)

	if len(retVal) != 2 {
		t.Error(
			"For", manager,
			"expected", 2,
			"got", len(retVal))
	}

	castType, ok := retVal[0].(TestItem)

	if !ok {
		t.Error(
			"For", retVal[0],
			"expected TestItem",
			"got", castType)
	}
}

func TestFilterItemAndInComplexFilter(t *testing.T) {
	var manager DataManager = CreateRuntimeDataManager()
	manager.Add(0, TestItem{1, "some test"})
	manager.Add(1, TestItem{2, "some other test"})
	manager.Add(2, TestItem{3, "testing test test"})
	manager.Add(3, TestItem{4, "all testers test tests"})

	filter := FilterRequest{
		Filters: []ComplexFilter{
			{
				Filters: []Filter{
					{
						Property:       "Value",
						ComparisonType: "contains",
						Value:          "some",
					},
					{
						Property:       "Value",
						ComparisonType: "contains",
						Value:          "test",
					},
				},
				Junction: "and",
			},
		},
		Junction: "and",
	}

	retVal := manager.Filter(&filter)

	if len(retVal) != 2 {
		t.Error(
			"For", manager,
			"expected", 2,
			"got", len(retVal))
	}

	castType, ok := retVal[0].(TestItem)

	if !ok {
		t.Error(
			"For", retVal[0],
			"expected TestItem",
			"got", castType)
	}
}

func TestFilterItemAndInRootFilter(t *testing.T) {
	var manager DataManager = CreateRuntimeDataManager()
	manager.Add(0, TestItem{1, "some test"})
	manager.Add(1, TestItem{2, "some other test"})
	manager.Add(2, TestItem{3, "testing test test"})
	manager.Add(3, TestItem{4, "all testers test tests"})

	filter := FilterRequest{
		Filters: []ComplexFilter{
			{
				Filters: []Filter{
					{
						Property:       "Value",
						ComparisonType: "contains",
						Value:          "some",
					},
				},
			},
			{
				Filters: []Filter{
					{
						Property:       "Value",
						ComparisonType: "contains",
						Value:          "test",
					},
				},
			},
		},
		Junction: "and",
	}

	retVal := manager.Filter(&filter)

	if len(retVal) != 2 {
		t.Error(
			"For", manager,
			"expected", 2,
			"got", len(retVal))
	}

	castType, ok := retVal[0].(TestItem)

	if !ok {
		t.Error(
			"For", retVal[0],
			"expected TestItem",
			"got", castType)
	}
}

func TestFilterComplex(t *testing.T) {
	var manager DataManager = CreateRuntimeDataManager()
	manager.Add(0, TestItem{1, "some test"})
	manager.Add(1, TestItem{2, "some other test"})
	manager.Add(2, TestItem{3, "testing test test"})
	manager.Add(3, TestItem{4, "all testers test tests"})

	filter := FilterRequest{
		Filters: []ComplexFilter{
			{
				Filters: []Filter{
					{
						Property:       "Value",
						ComparisonType: "contains",
						Value:          "some",
					},
					{
						Property:       "Value",
						ComparisonType: "contains",
						Value:          "test",
					},
				},
				Junction: "and",
			},
			{
				Filters: []Filter{
					{
						Property:       "Value",
						ComparisonType: "equals",
						Value:          "all testers test tests",
					},
				},
			},
		},
		Junction: "or",
	}

	retVal := manager.Filter(&filter)

	if len(retVal) != 3 {
		t.Error(
			"For", manager,
			"expected", 3,
			"got", len(retVal))
	}

	castType, ok := retVal[0].(TestItem)

	if !ok {
		t.Error(
			"For", retVal[0],
			"expected TestItem",
			"got", castType)
	}
}
