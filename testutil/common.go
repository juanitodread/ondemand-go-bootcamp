package testutil

import (
	"reflect"
	"testing"
)

func AssertEquals(t *testing.T, expectedResult interface{}, actualResult interface{}) {
	if expectedResult != actualResult {
		t.Errorf("Error: Expected [%+v], Got [%+v]", expectedResult, actualResult)
	}
}

func AssertNil(t *testing.T, actualValue interface{}) {
	if actualValue != nil && !reflect.ValueOf(actualValue).IsNil() {
		t.Errorf("Error: Expected to be nil, Got [%+v]", actualValue)
	}
}

func AssertNotNil(t *testing.T, actualValue interface{}) {
	if actualValue == nil {
		t.Errorf("Error: Expected to be nil, Got [%+v]", actualValue)
	}
}
