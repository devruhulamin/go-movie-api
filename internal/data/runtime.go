package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)
var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int32
// override the marshajosn to customize the json encoding
func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	doubleQuotedValue := strconv.Quote(jsonValue)
	return []byte(doubleQuotedValue), nil
}
// overrride the  unmarshal to customize the json decoding 
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error{
	unquotedJsonValue,err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	parts := strings.Split(unquotedJsonValue," ")
	if len(parts) != 2 || parts[1] != "mins"{
		return ErrInvalidRuntimeFormat
	}
	i,err := strconv.ParseInt(parts[0],10,32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	*r = Runtime(i)
	return nil
}