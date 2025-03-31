// Code generated by "enumer -type=ScanType -output scantype_enumer_generated.go"; DO NOT EDIT.

package executions

import (
	"fmt"
	"strings"
)

const _ScanTypeName = "ConcreteExecutionTypeCurrentExecutionType"

var _ScanTypeIndex = [...]uint8{0, 21, 41}

const _ScanTypeLowerName = "concreteexecutiontypecurrentexecutiontype"

func (i ScanType) String() string {
	if i < 0 || i >= ScanType(len(_ScanTypeIndex)-1) {
		return fmt.Sprintf("ScanType(%d)", i)
	}
	return _ScanTypeName[_ScanTypeIndex[i]:_ScanTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ScanTypeNoOp() {
	var x [1]struct{}
	_ = x[ConcreteExecutionType-(0)]
	_ = x[CurrentExecutionType-(1)]
}

var _ScanTypeValues = []ScanType{ConcreteExecutionType, CurrentExecutionType}

var _ScanTypeNameToValueMap = map[string]ScanType{
	_ScanTypeName[0:21]:       ConcreteExecutionType,
	_ScanTypeLowerName[0:21]:  ConcreteExecutionType,
	_ScanTypeName[21:41]:      CurrentExecutionType,
	_ScanTypeLowerName[21:41]: CurrentExecutionType,
}

var _ScanTypeNames = []string{
	_ScanTypeName[0:21],
	_ScanTypeName[21:41],
}

// ScanTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ScanTypeString(s string) (ScanType, error) {
	if val, ok := _ScanTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ScanTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ScanType values", s)
}

// ScanTypeValues returns all values of the enum
func ScanTypeValues() []ScanType {
	return _ScanTypeValues
}

// ScanTypeStrings returns a slice of all String values of the enum
func ScanTypeStrings() []string {
	strs := make([]string, len(_ScanTypeNames))
	copy(strs, _ScanTypeNames)
	return strs
}

// IsAScanType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ScanType) IsAScanType() bool {
	for _, v := range _ScanTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
