package common

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/itchyny/base58-go"
	"math/big"
	"os"
	"reflect"
)

func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return nil == err
}

func ByteSliceReverse(s []byte) {
	last := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[last-i] = s[last-i], s[i]
	}
	return
}

func ScriptHashToAddress(scriptHash string) (addr string, err error) {
	scriptHashLen := len(scriptHash)
	if 0 == scriptHashLen {
		return "", errors.New("script hash is empty")
	}

	//script hash is a 40 character string, will be considered as an address string if not.
	if 40 != scriptHashLen {
		return scriptHash, nil
	}

	scriptSlice, err := hex.DecodeString(scriptHash)
	if nil != err {
		return
	}

	ByteSliceReverse(scriptSlice)

	scriptSlice = append([]byte{23}, scriptSlice...)

	x := sha256.Sum256(scriptSlice)
	x = sha256.Sum256(x[:])

	xSlice := append(scriptSlice, x[0:4]...)

	bi := new(big.Int).SetBytes(xSlice).String()
	addrBase58, err := base58.BitcoinEncoding.Encode([]byte(bi))

	addr = string(addrBase58)
	return
}

//!!never use this method in http request data convert
func Fmt2Str(i interface{}) string {
	return fmt.Sprint(i)
}

func StringSlice2InterfaceSlice(data []string) (ret []interface{}) {
	ret = []interface{}{}

	for _, v := range data {
		ret = append(ret, v)
	}

	return
}

func JsonPointer2Struct(in interface{}, out interface{}) (err error) {
	data, err := json.Marshal(in)
	if nil != err {
		return
	}

	err = json.Unmarshal(data, out)

	if nil != err {
		return
	}

	return
}

func StructDataMerge(target interface{}, source interface{}, defaultValue interface{}) {
	tType := reflect.TypeOf(target)
	sType := reflect.TypeOf(source)

	if reflect.Ptr != tType.Kind() || tType != sType {
		return
	}

	defaultRV := *getDefaultValue(defaultValue, tType)

	targetRV := reflect.ValueOf(target).Elem()
	sourceRV := reflect.ValueOf(source).Elem()

	for i := 0; i < sourceRV.NumField(); i++ {
		sourceValue := sourceRV.Field(i)
		targetValue := targetRV.Field(i)
		if !targetValue.CanSet() {
			continue
		}

		if eq, _ := compareReflectValue(defaultRV.Field(i), sourceValue); true == eq {
			continue
		}

		if sourceValue.IsValid() {
			targetRV.Field(i).Set(sourceValue)
		}
	}

	return
}

func getDefaultValue(def interface{}, theType reflect.Type) (v *reflect.Value) {
	if nil == def {
		return
	}

	defType := reflect.TypeOf(def)

	if theType != defType {
		return
	}

	ret := reflect.ValueOf(def).Elem()

	return &ret
}

func IsIntKind(v interface{}) bool {
	vKind := reflect.TypeOf(v).Kind()

	return vKind == reflect.Int ||
		vKind == reflect.Int8 ||
		vKind == reflect.Int16 ||
		vKind == reflect.Int32 ||
		vKind == reflect.Int64

}

func IsUintKind(v interface{}) bool {
	vKind := reflect.TypeOf(v).Kind()

	return vKind == reflect.Uint ||
		vKind == reflect.Uint8 ||
		vKind == reflect.Uint16 ||
		vKind == reflect.Uint32 ||
		vKind == reflect.Uint64
}

//begin: code under this line based on https://golang.org/src/text/template/funcs.go
type VariableKind int

const (
	InvalidKind VariableKind = iota
	BoolKind
	ComplexKind
	IntKind
	FloatKind
	StringKind
	UintKind
)

func InterfaceKindPtrCompatible(v interface{}) VariableKind {
	vType := reflect.TypeOf(v)
	if reflect.Ptr == vType.Kind() {
		switch vPtr := v.(type) {
		case *int, *int8, *int16, *int32, *int64:
			return IntKind

		case *uint, *uint8, *uint16, *uint32, *uint64:
			return UintKind

		case *string:
			return StringKind

		case *bool:
			return BoolKind

		case *float32, *float64:
			return FloatKind

		case *complex64, *complex128:
			return ComplexKind

		case *interface{}:
			return InterfaceKindPtrCompatible(*vPtr)

		default:
			return InvalidKind
		}
	}
	return BasicKind(reflect.TypeOf(v))
}

func BasicKind(v reflect.Type) VariableKind {
	switch v.Kind() {
	case reflect.Bool:
		return BoolKind
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return IntKind
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return UintKind
	case reflect.Float32, reflect.Float64:
		return FloatKind
	case reflect.Complex64, reflect.Complex128:
		return ComplexKind
	case reflect.String:
		return StringKind
	}
	return InvalidKind
}

func compareReflectValue(v1 reflect.Value, v2 reflect.Value) (eq bool, err error) {
	v1Type := v1.Type()
	v2Type := v1.Type()
	if !v1Type.Comparable() || !v2Type.Comparable() {
		return false, &GatewayError{Code: GW_ERR_DATA_TYPE}
	}

	k1 := BasicKind(v1Type)
	k2 := BasicKind(v2Type)

	if v1Type != v2Type {
		switch {
		case IntKind == k1 && UintKind == k2:
			eq = v1.Int() >= 0 && uint64(v1.Int()) == v2.Uint()
		case k1 == UintKind && k2 == IntKind:
			eq = v2.Int() >= 0 && v1.Uint() == uint64(v2.Int())

		default:
			return false, &GatewayError{Code: GW_ERR_DATA_TYPE}
		}
	} else {
		switch k1 {
		case BoolKind:
			eq = v1.Bool() == v2.Bool()
		case ComplexKind:
			eq = v1.Complex() == v2.Complex()
		case FloatKind:
			eq = v1.Float() == v2.Float()
		case IntKind:
			eq = v1.Int() == v2.Int()
		case StringKind:
			eq = v1.String() == v2.String()
		case UintKind:
			eq = v1.Uint() == v2.Uint()
		default:
			return false, &GatewayError{Code: GW_ERR_DATA_TYPE}
		}
	}

	return
}
