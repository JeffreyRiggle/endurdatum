package endurdatum

import "testing"

func TestOrRequest(t *testing.T) {
	orFilter := FilterRequest{
		Filters: []ComplexFilter{
			{
				Filters: []Filter{
					{
						Property:       "testprop",
						ComparisonType: "equals",
						Value:          "foo",
					},
				},
			},
		},
		Junction: "or",
	}

	res := isOrRequest(&orFilter)

	if !res {
		t.Error(
			"For", orFilter,
			"expected", true,
			"got", res)
	}

	andFilter := FilterRequest{
		Filters: []ComplexFilter{
			{
				Filters: []Filter{
					{
						Property:       "testprop",
						ComparisonType: "equals",
						Value:          "foo",
					},
				},
			},
		},
		Junction: "and",
	}

	res2 := isOrRequest(&andFilter)

	if res2 {
		t.Error(
			"For", andFilter,
			"expected", false,
			"got", res2)
	}
}

func TestAndRequest(t *testing.T) {
	andFilter := FilterRequest{
		Filters: []ComplexFilter{
			{
				Filters: []Filter{
					{
						Property:       "testprop",
						ComparisonType: "equals",
						Value:          "foo",
					},
				},
			},
		},
		Junction: "and",
	}

	res := isAndRequest(&andFilter)

	if !res {
		t.Error(
			"For", andFilter,
			"expected", true,
			"got", res)
	}

	orFilter := FilterRequest{
		Filters: []ComplexFilter{
			{
				Filters: []Filter{
					{
						Property:       "testprop",
						ComparisonType: "equals",
						Value:          "foo",
					},
				},
			},
		},
		Junction: "or",
	}

	res2 := isAndRequest(&orFilter)

	if res2 {
		t.Error(
			"For", orFilter,
			"expected", false,
			"got", res2)
	}
}

func TestOrFilter(t *testing.T) {
	orFilter := ComplexFilter{
		Filters: []Filter{
			{
				Property:       "testprop",
				ComparisonType: "equals",
				Value:          "foo",
			},
		},
		Junction: "or",
	}

	res := isOrFilter(orFilter)

	if !res {
		t.Error(
			"For", orFilter,
			"expected", true,
			"got", res)
	}

	andFilter := ComplexFilter{
		Filters: []Filter{
			{
				Property:       "testprop",
				ComparisonType: "equals",
				Value:          "foo",
			},
		},
		Junction: "and",
	}

	res2 := isOrFilter(andFilter)

	if res2 {
		t.Error(
			"For", andFilter,
			"expected", false,
			"got", res2)
	}
}

func TestAndFilter(t *testing.T) {
	andFilter := ComplexFilter{
		Filters: []Filter{
			{
				Property:       "testprop",
				ComparisonType: "equals",
				Value:          "foo",
			},
		},
		Junction: "and",
	}

	res := isAndFilter(andFilter)

	if !res {
		t.Error(
			"For", andFilter,
			"expected", true,
			"got", res)
	}

	orFilter := ComplexFilter{
		Filters: []Filter{
			{
				Property:       "testprop",
				ComparisonType: "equals",
				Value:          "foo",
			},
		},
		Junction: "or",
	}

	res2 := isAndFilter(orFilter)

	if res2 {
		t.Error(
			"For", orFilter,
			"expected", false,
			"got", res2)
	}
}

func TestEqualsComparison(t *testing.T) {
	eFilter := Filter{
		Property:       "Someprop",
		Value:          "somevalue",
		ComparisonType: "equals",
	}

	res := isEqualsComparision(eFilter)

	if !res {
		t.Error(
			"For", eFilter,
			"expected", true,
			"got", res)
	}

	neFilter := Filter{
		Property:       "Someprop",
		Value:          "somevalue",
		ComparisonType: "notequals",
	}

	res2 := isEqualsComparision(neFilter)

	if res2 {
		t.Error(
			"For", neFilter,
			"expected", false,
			"got", res2)
	}

	cFilter := Filter{
		Property:       "Someprop",
		Value:          "somevalue",
		ComparisonType: "contains",
	}

	res3 := isEqualsComparision(cFilter)

	if res3 {
		t.Error(
			"For", cFilter,
			"expected", false,
			"got", res3)
	}
}

func TestNotEqualsComparison(t *testing.T) {
	eFilter := Filter{
		Property:       "Someprop",
		Value:          "somevalue",
		ComparisonType: "equals",
	}

	res := isNotEqualsComparision(eFilter)

	if res {
		t.Error(
			"For", eFilter,
			"expected", false,
			"got", res)
	}

	neFilter := Filter{
		Property:       "Someprop",
		Value:          "somevalue",
		ComparisonType: "notequals",
	}

	res2 := isNotEqualsComparision(neFilter)

	if !res2 {
		t.Error(
			"For", neFilter,
			"expected", true,
			"got", res2)
	}

	cFilter := Filter{
		Property:       "Someprop",
		Value:          "somevalue",
		ComparisonType: "contains",
	}

	res3 := isNotEqualsComparision(cFilter)

	if res3 {
		t.Error(
			"For", cFilter,
			"expected", false,
			"got", res3)
	}
}

func TestContainsComparison(t *testing.T) {
	eFilter := Filter{
		Property:       "Someprop",
		Value:          "somevalue",
		ComparisonType: "equals",
	}

	res := isContainsComparision(eFilter)

	if res {
		t.Error(
			"For", eFilter,
			"expected", false,
			"got", res)
	}

	neFilter := Filter{
		Property:       "Someprop",
		Value:          "somevalue",
		ComparisonType: "notequals",
	}

	res2 := isContainsComparision(neFilter)

	if res2 {
		t.Error(
			"For", neFilter,
			"expected", false,
			"got", res2)
	}

	cFilter := Filter{
		Property:       "Someprop",
		Value:          "somevalue",
		ComparisonType: "contains",
	}

	res3 := isContainsComparision(cFilter)

	if !res3 {
		t.Error(
			"For", cFilter,
			"expected", true,
			"got", res3)
	}
}
