package generics

import (
	"testing"
)

func TestSetIfNotNilStructFieldValue(t *testing.T) {
	type User struct {
		FirstName string
	}
	var (
		user        User
		sourceValue = "Satoshi"
	)
	SetIfNotNil(&sourceValue, &user.FirstName)
	if user.FirstName != sourceValue {
		t.Errorf("Expected targetValue to be 'Satoshi', but got %s", user.FirstName)
	}
}

func TestSetIfNotNilUser(t *testing.T) {
	type User struct {
		FirstName *string
		LastName  *string
	}

	tests := []struct {
		Source   *User
		Target   *User
		Expected User
	}{
		{
			Source:   &User{FirstName: ToPointer("John"), LastName: ToPointer("Doe")},
			Target:   &User{},
			Expected: User{FirstName: ToPointer("John"), LastName: ToPointer("Doe")},
		},
		{
			Source:   nil,
			Target:   &User{FirstName: ToPointer("Jane"), LastName: ToPointer("Smith")},
			Expected: User{FirstName: ToPointer("Jane"), LastName: ToPointer("Smith")},
		},
		{
			Source:   &User{FirstName: nil, LastName: ToPointer("Doe")},
			Target:   &User{FirstName: ToPointer("John")},
			Expected: User{FirstName: ToPointer("John"), LastName: ToPointer("Doe")},
		},
		{
			Source:   &User{FirstName: ToPointer("Alice"), LastName: nil},
			Target:   &User{LastName: ToPointer("Smith")},
			Expected: User{FirstName: ToPointer("Alice"), LastName: ToPointer("Smith")},
		},
	}

	for _, test := range tests {
		if test.Source == nil {
			continue
		}

		SetIfNotNil(test.Source.FirstName, test.Target.FirstName)
		SetIfNotNil(test.Source.LastName, test.Target.LastName)

		tFirstName := DereferenceOrDefault(test.Target.FirstName)
		if test.Target.FirstName == nil {
			tFirstName = DereferenceOrDefault(test.Source.FirstName)
		}
		eFirstName := DereferenceOrDefault(test.Expected.FirstName)
		// Compare FirstName
		if tFirstName != eFirstName {
			t.Errorf("Expected FirstName to be '%s', but got '%s'", tFirstName, eFirstName)
		}

		tLastName := DereferenceOrDefault(test.Source.LastName)
		if test.Source.LastName == nil {
			tLastName = DereferenceOrDefault(test.Target.LastName)
		}
		eLastName := DereferenceOrDefault(test.Expected.LastName)
		// Compare LastName
		if tLastName != eLastName {
			t.Errorf("Expected LastName to be '%s', but got '%s'", tLastName, eLastName)
		}
	}
}

func TestDereferenceOrDefault(t *testing.T) {
	// Test case for *int
	var intPtr *int
	intDefault := 0
	resultInt := DereferenceOrDefault(intPtr)
	if resultInt != intDefault {
		t.Errorf("Expected %d, but got %d for *int", intDefault, resultInt)
	}

	intValue := 100
	intPtr = &intValue
	resultInt = DereferenceOrDefault(intPtr)
	if resultInt != intValue {
		t.Errorf("Expected %d, but got %d for *int", intValue, resultInt)
	}

	// Test case for *string
	var (
		strPtr     *string
		strDefault string = ""
	)
	resultStr := DereferenceOrDefault(strPtr)
	if resultStr != strDefault {
		t.Errorf("Expected %q, but got %q for *string", strDefault, resultStr)
	}

	var (
		strValue      = "hello"
		strPtrDefault = &strValue
	)
	resultStr = DereferenceOrDefault(strPtrDefault)
	if resultStr != strValue {
		t.Errorf("Expected %q, but got %q for *string", strValue, resultStr)
	}

	// Test case for *float64
	var (
		floatPtr     *float64
		floatDefault = 0.0
	)
	resultFloat := DereferenceOrDefault(floatPtr)
	if resultFloat != floatDefault {
		t.Errorf("Expected %f, but got %f for *float64", floatDefault, resultFloat)
	}

	floatValue := 2.71828
	floatPtr = &floatValue
	resultFloat = DereferenceOrDefault(floatPtr)
	if resultFloat != floatValue {
		t.Errorf("Expected %f, but got %f for *float64", floatValue, resultFloat)
	}

	// Test case for custom type
	type CustomStruct struct {
		Field string
	}

	var customPtr *CustomStruct
	var customDefault = CustomStruct{Field: ""}
	resultCustom := DereferenceOrDefault(customPtr)
	if resultCustom != customDefault {
		t.Errorf("Expected %+v, but got %+v for *CustomStruct", customDefault, resultCustom)
	}

	customValue := CustomStruct{Field: "hello"}
	customPtr = &customValue
	resultCustom = DereferenceOrDefault(customPtr)
	if resultCustom != customValue {
		t.Errorf("Expected %+v, but got %+v for *CustomStruct", customValue, resultCustom)
	}
}
