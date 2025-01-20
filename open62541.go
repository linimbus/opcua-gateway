package main

/*
#cgo CFLAGS: -I. -std=c99
#cgo LDFLAGS: -lws2_32 -lIphlpapi

#include <stdlib.h>
#include <stdio.h>
#include "open62541.h"
#include "open62541_cgo.h"

*/
import "C"

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"unsafe"

	"github.com/astaxie/beego/logs"
)

type Client struct {
	addr    string
	cli     uintptr
	cLogger C.UA_Logger
}

type Server struct {
	sync.WaitGroup
	addr      string
	running   bool
	namespace map[string]uint32
	srv       uintptr
	cLogger   C.UA_Logger
}

type NodeInfo struct {
	NsIndex uint32
	NodeID  string
}

type NodeTree struct {
	Level    uint32
	Node     NodeInfo
	SubNodes []*NodeTree
}

type ValueType int

const (
	UA_BOOLEAN ValueType = iota
	UA_INT8
	UA_UINT8
	UA_INT16
	UA_UINT16
	UA_INT32
	UA_UINT32
	UA_INT64
	UA_UINT64
	UA_FLOAT
	UA_DOUBLE
	UA_STRING
	UA_DATETIME
	// UA_GUID
	UA_BYTESTRING
	// UA_XMLELEMENT
	// UA_NODEID
	// UA_EXPANDEDNODEID
	// UA_STATUSCODE
	// UA_QUALIFIEDNAME
	// UA_LOCALIZEDTEXT
	// UA_EXTENSIONOBJECT
	// UA_DATAVALUE
	// UA_VARIANT
	// UA_DIAGNOSTICINFO
)

type NodeValue struct {
	Type  ValueType
	Array bool
	Value interface{} /*
		if array is true, the value type is : []bool, []int8, []uint8, []int16, []uint16, []int32, []uint32, []int64, []uint64, []float, []double, []string, [][]byte
		if array is false, the value type is : bool, int8, uint8, int16, uint16, int32, uint32, int64, uint64, float, double, string, []byte
	*/
}

func (v *NodeValue) Clone() *NodeValue {
	value := &NodeValue{Type: v.Type, Array: v.Array}

	switch v.Type {
	case UA_BOOLEAN:
		if v.Array {
			value.Value = BoolListClone(v.Value.([]bool))
		} else {
			value.Value = v.Value.(bool)
		}
	case UA_INT8:
		if v.Array {
			value.Value = Int8ListClone(v.Value.([]int8))
		} else {
			value.Value = v.Value.(int8)
		}
	case UA_UINT8:
		if v.Array {
			value.Value = Uint8ListClone(v.Value.([]uint8))
		} else {
			value.Value = v.Value.(uint8)
		}
	case UA_INT16:
		if v.Array {
			value.Value = Int16ListClone(v.Value.([]int16))
		} else {
			value.Value = v.Value.(int16)
		}
	case UA_UINT16:
		if v.Array {
			value.Value = Uint16ListClone(v.Value.([]uint16))
		} else {
			value.Value = v.Value.(uint16)
		}
	case UA_INT32:
		if v.Array {
			value.Value = Int32ListClone(v.Value.([]int32))
		} else {
			value.Value = v.Value.(int32)
		}
	case UA_UINT32:
		if v.Array {
			value.Value = Uint32ListClone(v.Value.([]uint32))
		} else {
			value.Value = v.Value.(uint32)
		}
	case UA_INT64:
		if v.Array {
			value.Value = Int64ListClone(v.Value.([]int64))
		} else {
			value.Value = v.Value.(int64)
		}
	case UA_UINT64:
		if v.Array {
			value.Value = Uint64ListClone(v.Value.([]uint64))
		} else {
			value.Value = v.Value.(uint64)
		}
	case UA_DATETIME:
		if v.Array {
			value.Value = Uint64ListClone(v.Value.([]uint64))
		} else {
			value.Value = v.Value.(uint64)
		}
	case UA_FLOAT:
		if v.Array {
			value.Value = FloatListClone(v.Value.([]float32))
		} else {
			value.Value = v.Value.(float32)
		}
	case UA_DOUBLE:
		if v.Array {
			value.Value = DoubleListClone(v.Value.([]float64))
		} else {
			value.Value = v.Value.(float64)
		}
	case UA_STRING:
		if v.Array {
			value.Value = StringListClone(v.Value.([]string))
		} else {
			value.Value = v.Value.(string)
		}
	case UA_BYTESTRING:
		if v.Array {
			value.Value = ByteListClone(v.Value.([][]byte))
		} else {
			value.Value = ByteClone(v.Value.([]byte))
		}
	default:
		return nil
	}
	return value
}

func (v *NodeValue) Compare(b *NodeValue) bool {
	if v.Type != b.Type || v.Array != b.Array {
		return false
	}
	switch v.Type {
	case UA_BOOLEAN:
		if v.Array {
			return BoolListCompare(v.Value.([]bool), b.Value.([]bool))
		}
		return v.Value.(bool) == b.Value.(bool)
	case UA_INT8:
		if v.Array {
			return Int8ListCompare(v.Value.([]int8), b.Value.([]int8))
		}
		return v.Value.(int8) == b.Value.(int8)
	case UA_UINT8:
		if v.Array {
			return Uint8ListCompare(v.Value.([]uint8), b.Value.([]uint8))
		}
		return v.Value.(uint8) == b.Value.(uint8)
	case UA_INT16:
		if v.Array {
			return Int16ListCompare(v.Value.([]int16), b.Value.([]int16))
		}
		return v.Value.(int16) == b.Value.(int16)
	case UA_UINT16:
		if v.Array {
			return Uint16ListCompare(v.Value.([]uint16), b.Value.([]uint16))
		}
		return v.Value.(uint16) == b.Value.(uint16)
	case UA_INT32:
		if v.Array {
			return Int32ListCompare(v.Value.([]int32), b.Value.([]int32))
		}
		return v.Value.(int32) == b.Value.(int32)
	case UA_UINT32:
		if v.Array {
			return Uint32ListCompare(v.Value.([]uint32), b.Value.([]uint32))
		}
		return v.Value.(uint32) == b.Value.(uint32)
	case UA_INT64:
		if v.Array {
			return Int64ListCompare(v.Value.([]int64), b.Value.([]int64))
		}
		return v.Value.(int64) == b.Value.(int64)
	case UA_UINT64:
		if v.Array {
			return Uint64ListCompare(v.Value.([]uint64), b.Value.([]uint64))
		}
		return v.Value.(uint64) == b.Value.(uint64)
	case UA_DATETIME:
		if v.Array {
			return Uint64ListCompare(v.Value.([]uint64), b.Value.([]uint64))
		}
		return v.Value.(uint64) == b.Value.(uint64)
	case UA_FLOAT:
		if v.Array {
			return FloatListCompare(v.Value.([]float32), b.Value.([]float32))
		}
		return v.Value.(float32) == b.Value.(float32)
	case UA_DOUBLE:
		if v.Array {
			return DoubleListCompare(v.Value.([]float64), b.Value.([]float64))
		}
		return v.Value.(float64) == b.Value.(float64)
	case UA_STRING:
		if v.Array {
			return StringListCompare(v.Value.([]string), b.Value.([]string))
		}
		return v.Value.(string) == b.Value.(string)
	case UA_BYTESTRING:
		if v.Array {
			return ByteListCompare(v.Value.([][]byte), b.Value.([][]byte))
		}
		return bytes.Equal(v.Value.([]byte), b.Value.([]byte))
	default:
		return false
	}
}

func (v *NodeValue) ToString() string {
	switch v.Type {
	case UA_BOOLEAN:
		if v.Array {
			return BoolListToString(v.Value.([]bool))
		}
		return BoolToString(v.Value.(bool))
	case UA_INT8:
		if v.Array {
			return Int8ListToString(v.Value.([]int8))
		}
		return fmt.Sprintf("%d", int16(v.Value.(int8)))
	case UA_UINT8:
		if v.Array {
			return Uint8ListToString(v.Value.([]uint8))
		}
		return fmt.Sprintf("%d", uint16(v.Value.(uint8)))
	case UA_INT16:
		if v.Array {
			return Int16ListToString(v.Value.([]int16))
		}
		return fmt.Sprintf("%d", v.Value.(int16))
	case UA_UINT16:
		if v.Array {
			return Uint16ListToString(v.Value.([]uint16))
		}
		return fmt.Sprintf("%d", v.Value.(uint16))
	case UA_INT32:
		if v.Array {
			return Int32ListToString(v.Value.([]int32))
		}
		return fmt.Sprintf("%d", v.Value.(int32))
	case UA_UINT32:
		if v.Array {
			return Uint32ListToString(v.Value.([]uint32))
		}
		return fmt.Sprintf("%d", v.Value.(uint32))
	case UA_INT64:
		if v.Array {
			return Int64ListToString(v.Value.([]int64))
		}
		return fmt.Sprintf("%d", v.Value.(int64))
	case UA_UINT64:
		if v.Array {
			return Uint64ListToString(v.Value.([]uint64))
		}
		return fmt.Sprintf("%d", v.Value.(uint64))
	case UA_DATETIME:
		if v.Array {
			return DateTimeListToString(v.Value.([]uint64))
		}
		return DatetimeToString(v.Value.(uint64))
	case UA_FLOAT:
		if v.Array {
			return FloatListToString(v.Value.([]float32))
		}
		return fmt.Sprintf("%0.5f", v.Value.(float32))
	case UA_DOUBLE:
		if v.Array {
			return DoubleListToString(v.Value.([]float64))
		}
		return fmt.Sprintf("%0.5f", v.Value.(float64))
	case UA_STRING:
		if v.Array {
			return StringListToString(v.Value.([]string))
		}
		return v.Value.(string)
	case UA_BYTESTRING:
		if v.Array {
			return ByteListToString(v.Value.([][]byte))
		}
		return base64.StdEncoding.EncodeToString(v.Value.([]byte))
	default:
		return ""
	}
}

func (v *NodeValue) FromString(values []string) error {
	if !v.Array && len(values) > 1 {
		return fmt.Errorf("ua client node value from string failed, the node value type is not array")
	}

	if !v.Array && len(values) == 0 {
		return fmt.Errorf("ua client node value from string failed, input values is empty")
	}

	switch v.Type {
	case UA_BOOLEAN:
		if v.Array {
			arrayList := v.Value.([]bool)
			length := SmallLength(len(arrayList), len(values))
			for i := 0; i < length; i++ {
				arrayList[i] = StringToBool(values[i])
			}
		} else {
			v.Value = StringToBool(values[0])
		}
	case UA_INT8:
		if v.Array {
			arrayList := v.Value.([]int8)
			length := SmallLength(len(arrayList), len(values))
			for i := 0; i < length; i++ {
				val, err := strconv.Atoi(values[i])
				if err != nil {
					return fmt.Errorf("ua client node value from string failed, strconv.int8 %s", err.Error())
				}
				arrayList[i] = int8(val)
			}
		} else {
			val, err := strconv.Atoi(values[0])
			if err != nil {
				return fmt.Errorf("ua client node value from string failed, strconv.int8 %s", err.Error())
			}
			v.Value = int8(val)
		}
	case UA_UINT8:
		if v.Array {
			arrayList := v.Value.([]uint8)
			length := SmallLength(len(arrayList), len(values))
			for i := 0; i < length; i++ {
				val, err := strconv.Atoi(values[i])
				if err != nil {
					return fmt.Errorf("ua client node value from string failed, strconv.uint8 %s", err.Error())
				}
				arrayList[i] = uint8(val)
			}
		} else {
			val, err := strconv.Atoi(values[0])
			if err != nil {
				return fmt.Errorf("ua client node value from string failed, strconv.uint8 %s", err.Error())
			}
			v.Value = uint8(val)
		}
	case UA_INT16:
		if v.Array {
			arrayList := v.Value.([]int16)
			length := SmallLength(len(arrayList), len(values))
			for i := 0; i < length; i++ {
				val, err := strconv.Atoi(values[i])
				if err != nil {
					return fmt.Errorf("ua client node value from string failed, strconv.int16 %s", err.Error())
				}
				arrayList[i] = int16(val)
			}
		} else {
			val, err := strconv.Atoi(values[0])
			if err != nil {
				return fmt.Errorf("ua client node value from string failed, strconv.int16 %s", err.Error())
			}
			v.Value = int16(val)
		}
	case UA_UINT16:
		if v.Array {
			arrayList := v.Value.([]uint16)
			length := SmallLength(len(arrayList), len(values))
			for i := 0; i < length; i++ {
				val, err := strconv.Atoi(values[i])
				if err != nil {
					return fmt.Errorf("ua client node value from string failed, strconv.uint16 %s", err.Error())
				}
				arrayList[i] = uint16(val)
			}
		} else {
			val, err := strconv.Atoi(values[0])
			if err != nil {
				return fmt.Errorf("ua client node value from string failed, strconv.uint16 %s", err.Error())
			}
			v.Value = uint16(val)
		}
	case UA_INT32:
		if v.Array {
			arrayList := v.Value.([]int32)
			length := SmallLength(len(arrayList), len(values))
			for i := 0; i < length; i++ {
				val, err := strconv.Atoi(values[i])
				if err != nil {
					return fmt.Errorf("ua client node value from string failed, strconv.int32 %s", err.Error())
				}
				arrayList[i] = int32(val)
			}
		} else {
			val, err := strconv.Atoi(values[0])
			if err != nil {
				return fmt.Errorf("ua client node value from string failed, strconv.int32 %s", err.Error())
			}
			v.Value = int32(val)
		}
	case UA_UINT32:
		if v.Array {
			arrayList := v.Value.([]uint32)
			length := SmallLength(len(arrayList), len(values))
			for i := 0; i < length; i++ {
				val, err := strconv.Atoi(values[i])
				if err != nil {
					return fmt.Errorf("ua client node value from string failed, strconv.uint32 %s", err.Error())
				}
				arrayList[i] = uint32(val)
			}
		} else {
			val, err := strconv.Atoi(values[0])
			if err != nil {
				return fmt.Errorf("ua client node value from string failed, strconv.uint32 %s", err.Error())
			}
			v.Value = uint32(val)
		}
	case UA_INT64:
		if v.Array {
			arrayList := v.Value.([]int64)
			length := SmallLength(len(arrayList), len(values))
			for i := 0; i < length; i++ {
				val, err := strconv.Atoi(values[i])
				if err != nil {
					return fmt.Errorf("ua client node value from string failed, strconv.int64 %s", err.Error())
				}
				arrayList[i] = int64(val)
			}
		} else {
			val, err := strconv.Atoi(values[0])
			if err != nil {
				return fmt.Errorf("ua client node value from string failed, strconv.int64 %s", err.Error())
			}
			v.Value = int64(val)
		}
	case UA_UINT64:
		if v.Array {
			arrayList := v.Value.([]uint64)
			length := SmallLength(len(arrayList), len(values))
			for i := 0; i < length; i++ {
				val, err := strconv.Atoi(values[i])
				if err != nil {
					return fmt.Errorf("ua client node value from string failed, strconv.uint64 %s", err.Error())
				}
				arrayList[i] = uint64(val)
			}
		} else {
			val, err := strconv.Atoi(values[0])
			if err != nil {
				return fmt.Errorf("ua client node value from string failed, strconv.uint64 %s", err.Error())
			}
			v.Value = uint64(val)
		}
	case UA_DATETIME:
		if v.Array {
			arrayList := v.Value.([]uint64)
			length := SmallLength(len(arrayList), len(values))
			for i := 0; i < length; i++ {
				val, err := strconv.Atoi(values[i])
				if err != nil {
					return fmt.Errorf("ua client node value from string failed, strconv.datetime %s", err.Error())
				}
				arrayList[i] = uint64(val)
			}
		} else {
			val, err := strconv.Atoi(values[0])
			if err != nil {
				return fmt.Errorf("ua client node value from string failed, strconv.datetime %s", err.Error())
			}
			v.Value = uint64(val)
		}
	case UA_FLOAT:
		if v.Array {
			arrayList := v.Value.([]float32)
			length := SmallLength(len(arrayList), len(values))
			for i := 0; i < length; i++ {
				val, err := strconv.ParseFloat(values[i], 32)
				if err != nil {
					return fmt.Errorf("ua client node value from string failed, strconv.float32 %s", err.Error())
				}
				arrayList[i] = float32(val)
			}
		} else {
			val, err := strconv.ParseFloat(values[0], 32)
			if err != nil {
				return fmt.Errorf("ua client node value from string failed, strconv.float32 %s", err.Error())
			}
			v.Value = float32(val)
		}
	case UA_DOUBLE:
		if v.Array {
			arrayList := v.Value.([]float64)
			length := SmallLength(len(arrayList), len(values))
			for i := 0; i < length; i++ {
				val, err := strconv.ParseFloat(values[i], 64)
				if err != nil {
					return fmt.Errorf("ua client node value from string failed, strconv.float64 %s", err.Error())
				}
				arrayList[i] = val
			}
		} else {
			val, err := strconv.ParseFloat(values[0], 64)
			if err != nil {
				return fmt.Errorf("ua client node value from string failed, strconv.float64 %s", err.Error())
			}
			v.Value = val
		}
	case UA_STRING:
		if v.Array {
			arrayList := v.Value.([]string)
			length := SmallLength(len(arrayList), len(values))
			for i := 0; i < length; i++ {
				arrayList[i] = values[i]
			}
		} else {
			v.Value = values[0]
		}
	default:
		return fmt.Errorf("ua client node value from string failed, type = %d not support", uint32(v.Type))
	}
	return nil
}

func (n NodeInfo) ToString() string {
	return fmt.Sprintf("ns=%d,s=%s", n.NsIndex, n.NodeID)
}

func (n NodeInfo) Compare(b NodeInfo) bool {
	return n.NsIndex == b.NsIndex && n.NodeID == b.NodeID
}

/*

typedef enum {
    UA_LOGLEVEL_TRACE = 0,
    UA_LOGLEVEL_DEBUG,
    UA_LOGLEVEL_INFO,
    UA_LOGLEVEL_WARNING,
    UA_LOGLEVEL_ERROR,
    UA_LOGLEVEL_FATAL
} UA_LogLevel;

typedef enum {
    UA_LOGCATEGORY_NETWORK = 0,
    UA_LOGCATEGORY_SECURECHANNEL,
    UA_LOGCATEGORY_SESSION,
    UA_LOGCATEGORY_SERVER,
    UA_LOGCATEGORY_CLIENT,
    UA_LOGCATEGORY_USERLAND,
    UA_LOGCATEGORY_SECURITYPOLICY
} UA_LogCategory;

*/

//export UA_Logger_golang
func UA_Logger_golang(level C.uint32_t, category C.uint32_t, message *C.char) {
	var categoryStr string
	switch category {
	case C.UA_LOGCATEGORY_NETWORK:
		categoryStr = "NETWORK"
	case C.UA_LOGCATEGORY_SECURECHANNEL:
		categoryStr = "SECURECHANNEL"
	case C.UA_LOGCATEGORY_SESSION:
		categoryStr = "SESSION"
	case C.UA_LOGCATEGORY_SERVER:
		categoryStr = "SERVER"
	case C.UA_LOGCATEGORY_CLIENT:
		categoryStr = "CLIENT"
	case C.UA_LOGCATEGORY_USERLAND:
		categoryStr = "USERLAND"
	case C.UA_LOGCATEGORY_SECURITYPOLICY:
		categoryStr = "SECURITYPOLICY"
	default:
		categoryStr = "UNKNOWN"
	}

	if level <= 1 {
		logs.Debug("OPCUA category: %s, message: %s", categoryStr, C.GoString(message))
	} else if level == 2 {
		logs.Info("OPCUA category: %s, message: %s", categoryStr, C.GoString(message))
	} else if level == 3 {
		logs.Warn("OPCUA category: %s, message: %s", categoryStr, C.GoString(message))
	} else {
		logs.Error("OPCUA category: %s, message: %s", categoryStr, C.GoString(message))
	}
}

func NewEmptyNodeValue() *NodeValue {
	return &NodeValue{Type: UA_STRING, Value: string(""), Array: false}
}

func UA_VariantToArrayValue(variant *C.UA_Variant) (*NodeValue, error) {
	length := uint32(variant.arrayLength)

	var uaType C.UA_UInt32
	retval := C.UA_VariantType(variant, &uaType)
	if retval != C.UA_STATUSCODE_GOOD {
		return nil, fmt.Errorf("ua client covert variant to node value failed, type = %d", uint32(uaType))
	}

	var arrayValue interface{}
	var valueType ValueType

	switch uaType {
	case C.UA_TYPES_BOOLEAN:
		{
			arrayList := make([]bool, length)
			for i := 0; i < int(length); i++ {
				arrayList[i] = bool(C.UA_VariantValueBoolean(variant, C.int(i)))
			}
			arrayValue = arrayList
			valueType = UA_BOOLEAN
		}
	case C.UA_TYPES_SBYTE:
		{
			arrayList := make([]int8, length)
			for i := 0; i < int(length); i++ {
				arrayList[i] = int8(C.UA_VariantValueInt8(variant, C.int(i)))
			}
			arrayValue = arrayList
			valueType = UA_INT8
		}
	case C.UA_TYPES_BYTE:
		{
			arrayList := make([]uint8, length)
			for i := 0; i < int(length); i++ {
				arrayList[i] = uint8(C.UA_VariantValueUint8(variant, C.int(i)))
			}
			arrayValue = arrayList
			valueType = UA_UINT8
		}
	case C.UA_TYPES_INT16:
		{
			arrayList := make([]int16, length)
			for i := 0; i < int(length); i++ {
				arrayList[i] = int16(C.UA_VariantValueInt16(variant, C.int(i)))
			}
			arrayValue = arrayList
			valueType = UA_INT16
		}
	case C.UA_TYPES_UINT16:
		{
			arrayList := make([]uint16, length)
			for i := 0; i < int(length); i++ {
				arrayList[i] = uint16(C.UA_VariantValueUint16(variant, C.int(i)))
			}
			arrayValue = arrayList
			valueType = UA_UINT16
		}
	case C.UA_TYPES_INT32:
		{
			arrayList := make([]int32, length)
			for i := 0; i < int(length); i++ {
				arrayList[i] = int32(C.UA_VariantValueInt32(variant, C.int(i)))
			}
			arrayValue = arrayList
			valueType = UA_INT32
		}
	case C.UA_TYPES_UINT32:
		{
			arrayList := make([]uint32, length)
			for i := 0; i < int(length); i++ {
				arrayList[i] = uint32(C.UA_VariantValueUint32(variant, C.int(i)))
			}
			arrayValue = arrayList
			valueType = UA_UINT32
		}
	case C.UA_TYPES_INT64:
		{
			arrayList := make([]int64, length)
			for i := 0; i < int(length); i++ {
				arrayList[i] = int64(C.UA_VariantValueInt64(variant, C.int(i)))
			}
			arrayValue = arrayList
			valueType = UA_INT64
		}
	case C.UA_TYPES_DATETIME:
		{
			arrayList := make([]uint64, length)
			for i := 0; i < int(length); i++ {
				arrayList[i] = uint64(C.UA_VariantValueUint64(variant, C.int(i)))
			}
			arrayValue = arrayList
			valueType = UA_DATETIME
		}
	case C.UA_TYPES_UINT64:
		{
			arrayList := make([]uint64, length)
			for i := 0; i < int(length); i++ {
				arrayList[i] = uint64(C.UA_VariantValueUint64(variant, C.int(i)))
			}
			arrayValue = arrayList
			valueType = UA_UINT64
		}
	case C.UA_TYPES_FLOAT:
		{
			arrayList := make([]float32, length)
			for i := 0; i < int(length); i++ {
				arrayList[i] = float32(C.UA_VariantValueFloat(variant, C.int(i)))
			}
			arrayValue = arrayList
			valueType = UA_FLOAT
		}
	case C.UA_TYPES_DOUBLE:
		{
			arrayList := make([]float64, length)
			for i := 0; i < int(length); i++ {
				arrayList[i] = float64(C.UA_VariantValueDouble(variant, C.int(i)))
			}
			arrayValue = arrayList
			valueType = UA_DOUBLE
		}
	case C.UA_TYPES_STRING:
		{
			arrayList := make([]string, length)
			for i := 0; i < int(length); i++ {
				cString := C.UA_VariantValueString(variant, C.int(i))
				goBytes := C.GoBytes(unsafe.Pointer(cString.data), C.int(cString.length))
				arrayList[i] = string(goBytes)
			}
			arrayValue = arrayList
			valueType = UA_STRING
		}
	case C.UA_TYPES_BYTESTRING:
		{
			arrayList := make([][]byte, length)
			for i := 0; i < int(length); i++ {
				cString := C.UA_VariantValueByteString(variant, C.int(i))
				goBytes := C.GoBytes(unsafe.Pointer(cString.data), C.int(cString.length))
				arrayList[i] = goBytes
			}
			arrayValue = arrayList
			valueType = UA_BYTESTRING
		}
	default:
		return nil, fmt.Errorf("ua client covert variant to node value failed, type = %d", uint32(uaType))
	}
	return &NodeValue{Type: valueType, Array: true, Value: arrayValue}, nil
}

func UA_VariantToSingleValue(variant *C.UA_Variant) (*NodeValue, error) {
	var uaType C.UA_UInt32
	retval := C.UA_VariantType(variant, &uaType)
	if retval != C.UA_STATUSCODE_GOOD {
		return nil, fmt.Errorf("ua client covert variant to node value failed, type = %d", uint32(uaType))
	}

	var value interface{}
	var valueType ValueType

	switch uaType {
	case C.UA_TYPES_BOOLEAN:
		value = bool(C.UA_VariantValueBoolean(variant, 0))
		valueType = UA_BOOLEAN
	case C.UA_TYPES_SBYTE:
		value = int8(C.UA_VariantValueInt8(variant, 0))
		valueType = UA_INT8
	case C.UA_TYPES_BYTE:
		value = uint8(C.UA_VariantValueUint8(variant, 0))
		valueType = UA_UINT8
	case C.UA_TYPES_INT16:
		value = int16(C.UA_VariantValueInt16(variant, 0))
		valueType = UA_INT16
	case C.UA_TYPES_UINT16:
		value = uint16(C.UA_VariantValueUint16(variant, 0))
		valueType = UA_UINT16
	case C.UA_TYPES_INT32:
		value = int32(C.UA_VariantValueInt32(variant, 0))
		valueType = UA_INT32
	case C.UA_TYPES_UINT32:
		value = uint32(C.UA_VariantValueUint32(variant, 0))
		valueType = UA_UINT32
	case C.UA_TYPES_INT64:
		value = int64(C.UA_VariantValueInt64(variant, 0))
		valueType = UA_INT64
	case C.UA_TYPES_DATETIME:
		value = uint64(C.UA_VariantValueUint64(variant, 0))
		valueType = UA_DATETIME
	case C.UA_TYPES_UINT64:
		value = uint64(C.UA_VariantValueUint64(variant, 0))
		valueType = UA_UINT64
	case C.UA_TYPES_FLOAT:
		value = float32(C.UA_VariantValueFloat(variant, 0))
		valueType = UA_FLOAT
	case C.UA_TYPES_DOUBLE:
		value = float64(C.UA_VariantValueDouble(variant, 0))
		valueType = UA_DOUBLE
	case C.UA_TYPES_STRING:
		cString := C.UA_VariantValueString(variant, 0)
		value = string(C.GoBytes(unsafe.Pointer(cString.data), C.int(cString.length)))
		valueType = UA_STRING
	case C.UA_TYPES_BYTESTRING:
		cString := C.UA_VariantValueByteString(variant, 0)
		value = C.GoBytes(unsafe.Pointer(cString.data), C.int(cString.length))
		valueType = UA_BYTESTRING
	default:
		return nil, fmt.Errorf("ua client covert variant to node value failed, type = %d", uint32(uaType))
	}

	return &NodeValue{Type: valueType, Value: value}, nil
}

func UA_VariantGolangValue(variant *C.UA_Variant) (*NodeValue, error) {
	if C.UA_Variant_isScalar(variant) {
		return UA_VariantToSingleValue(variant)
	}
	return UA_VariantToArrayValue(variant)
}

func UA_VariantFromSingleValue(value NodeValue, variant *C.UA_Variant) error {
	var retval C.UA_StatusCode

	switch value.Type {
	case UA_BOOLEAN:
		retval = C.UA_VariantScalarValueBoolean(variant, C.UA_Boolean(value.Value.(bool)))
	case UA_INT8:
		retval = C.UA_VariantScalarValueInt8(variant, C.UA_SByte(value.Value.(int8)))
	case UA_UINT8:
		retval = C.UA_VariantScalarValueUint8(variant, C.UA_Byte(value.Value.(uint8)))
	case UA_INT16:
		retval = C.UA_VariantScalarValueInt16(variant, C.UA_Int16(value.Value.(int16)))
	case UA_UINT16:
		retval = C.UA_VariantScalarValueUint16(variant, C.UA_UInt16(value.Value.(uint16)))
	case UA_INT32:
		retval = C.UA_VariantScalarValueInt32(variant, C.UA_Int32(value.Value.(int32)))
	case UA_UINT32:
		retval = C.UA_VariantScalarValueUint32(variant, C.UA_UInt32(value.Value.(uint32)))
	case UA_INT64:
		retval = C.UA_VariantScalarValueInt64(variant, C.UA_Int64(value.Value.(int64)))
	case UA_DATETIME:
		retval = C.UA_VariantScalarValueDateTime(variant, C.UA_DateTime(value.Value.(uint64)))
	case UA_UINT64:
		retval = C.UA_VariantScalarValueUint64(variant, C.UA_UInt64(value.Value.(uint64)))
	case UA_FLOAT:
		retval = C.UA_VariantScalarValueFloat(variant, C.UA_Float(value.Value.(float32)))
	case UA_DOUBLE:
		retval = C.UA_VariantScalarValueDouble(variant, C.UA_Double(value.Value.(float64)))
	case UA_STRING:
		cString := C.CString(value.Value.(string))
		defer C.free(unsafe.Pointer(cString))
		retval = C.UA_VariantScalarValueString(variant, cString)
	case UA_BYTESTRING:
		cString := C.CBytes(value.Value.([]byte))
		defer C.free(cString)

		retval = C.UA_VariantScalarValueByteString(variant, cString, C.size_t(len(value.Value.([]byte))))
	default:
		return fmt.Errorf("ua client covert node value to variant failed, type = %d not support", uint32(value.Type))
	}

	if retval != C.UA_STATUSCODE_GOOD {
		return fmt.Errorf("ua client covert node value to variant failed, type = %d, retval = 0x%x", uint32(value.Type), retval)
	}

	return nil
}

func UA_VariantFromArrayValue(value NodeValue, variant *C.UA_Variant) error {
	var retval C.UA_StatusCode
	var cArray C.ArrayValue
	var uaType C.uint32_t

	switch value.Type {
	case UA_BOOLEAN:
		{
			uaType = C.UA_TYPES_BOOLEAN
			arrayList := value.Value.([]bool)
			retval = C.UA_ArrayValueInit(&cArray, uaType)
			if retval != C.UA_STATUSCODE_GOOD {
				goto failed
			}
			for _, value := range arrayList {
				retval = C.UA_ArrayValueAppendBoolean(&cArray, C.UA_Boolean(value))
				if retval != C.UA_STATUSCODE_GOOD {
					goto failed
				}
			}
		}
	case UA_INT8:
		{
			uaType = C.UA_TYPES_SBYTE
			arrayList := value.Value.([]int8)
			retval = C.UA_ArrayValueInit(&cArray, uaType)
			if retval != C.UA_STATUSCODE_GOOD {
				goto failed
			}
			for _, value := range arrayList {
				retval = C.UA_ArrayValueAppendInt8(&cArray, C.UA_SByte(value))
				if retval != C.UA_STATUSCODE_GOOD {
					goto failed
				}
			}
		}
	case UA_UINT8:
		{
			uaType = C.UA_TYPES_BYTE
			arrayList := value.Value.([]uint8)
			retval = C.UA_ArrayValueInit(&cArray, uaType)
			if retval != C.UA_STATUSCODE_GOOD {
				goto failed
			}
			for _, value := range arrayList {
				retval = C.UA_ArrayValueAppendUint8(&cArray, C.UA_Byte(value))
				if retval != C.UA_STATUSCODE_GOOD {
					goto failed
				}
			}
		}
	case UA_INT16:
		{
			uaType = C.UA_TYPES_INT16
			arrayList := value.Value.([]int16)
			retval = C.UA_ArrayValueInit(&cArray, uaType)
			if retval != C.UA_STATUSCODE_GOOD {
				goto failed
			}
			for _, value := range arrayList {
				retval = C.UA_ArrayValueAppendInt16(&cArray, C.UA_Int16(value))
				if retval != C.UA_STATUSCODE_GOOD {
					goto failed
				}
			}
		}
	case UA_UINT16:
		{
			uaType = C.UA_TYPES_UINT16
			arrayList := value.Value.([]uint16)
			retval = C.UA_ArrayValueInit(&cArray, uaType)
			if retval != C.UA_STATUSCODE_GOOD {
				goto failed
			}
			for _, value := range arrayList {
				retval = C.UA_ArrayValueAppendUint16(&cArray, C.UA_UInt16(value))
				if retval != C.UA_STATUSCODE_GOOD {
					goto failed
				}
			}
		}
	case UA_INT32:
		{
			uaType = C.UA_TYPES_INT32
			arrayList := value.Value.([]int32)
			retval = C.UA_ArrayValueInit(&cArray, uaType)
			if retval != C.UA_STATUSCODE_GOOD {
				goto failed
			}
			for _, value := range arrayList {
				retval = C.UA_ArrayValueAppendInt32(&cArray, C.UA_Int32(value))
				if retval != C.UA_STATUSCODE_GOOD {
					goto failed
				}
			}
		}
	case UA_UINT32:
		{
			uaType = C.UA_TYPES_UINT32
			arrayList := value.Value.([]uint32)
			retval = C.UA_ArrayValueInit(&cArray, uaType)
			if retval != C.UA_STATUSCODE_GOOD {
				goto failed
			}
			for _, value := range arrayList {
				retval = C.UA_ArrayValueAppendUint32(&cArray, C.UA_UInt32(value))
				if retval != C.UA_STATUSCODE_GOOD {
					goto failed
				}
			}
		}
	case UA_INT64:
		{
			uaType = C.UA_TYPES_INT64
			arrayList := value.Value.([]int64)
			retval = C.UA_ArrayValueInit(&cArray, uaType)
			if retval != C.UA_STATUSCODE_GOOD {
				goto failed
			}
			for _, value := range arrayList {
				retval = C.UA_ArrayValueAppendInt64(&cArray, C.UA_Int64(value))
				if retval != C.UA_STATUSCODE_GOOD {
					goto failed
				}
			}
		}
	case UA_DATETIME:
		{
			uaType = C.UA_TYPES_DATETIME
			arrayList := value.Value.([]uint64)
			retval = C.UA_ArrayValueInit(&cArray, uaType)
			if retval != C.UA_STATUSCODE_GOOD {
				goto failed
			}
			for _, value := range arrayList {
				retval = C.UA_ArrayValueAppendDateTime(&cArray, C.UA_DateTime(value))
				if retval != C.UA_STATUSCODE_GOOD {
					goto failed
				}
			}
		}
	case UA_UINT64:
		{
			uaType = C.UA_TYPES_UINT64
			arrayList := value.Value.([]uint64)
			retval = C.UA_ArrayValueInit(&cArray, uaType)
			if retval != C.UA_STATUSCODE_GOOD {
				goto failed
			}
			for _, value := range arrayList {
				retval = C.UA_ArrayValueAppendUint64(&cArray, C.UA_UInt64(value))
				if retval != C.UA_STATUSCODE_GOOD {
					goto failed
				}
			}
		}
	case UA_FLOAT:
		{
			uaType = C.UA_TYPES_FLOAT
			arrayList := value.Value.([]float32)
			retval = C.UA_ArrayValueInit(&cArray, uaType)
			if retval != C.UA_STATUSCODE_GOOD {
				goto failed
			}
			for _, value := range arrayList {
				retval = C.UA_ArrayValueAppendFloat(&cArray, C.UA_Float(value))
				if retval != C.UA_STATUSCODE_GOOD {
					goto failed
				}
			}
		}
	case UA_DOUBLE:
		{
			uaType = C.UA_TYPES_DOUBLE
			arrayList := value.Value.([]float64)
			retval = C.UA_ArrayValueInit(&cArray, uaType)
			if retval != C.UA_STATUSCODE_GOOD {
				goto failed
			}
			for _, value := range arrayList {
				retval = C.UA_ArrayValueAppendDouble(&cArray, C.UA_Double(value))
				if retval != C.UA_STATUSCODE_GOOD {
					goto failed
				}
			}
		}
	case UA_STRING:
		{
			uaType = C.UA_TYPES_STRING
			arrayList := value.Value.([]string)
			retval = C.UA_ArrayValueInit(&cArray, uaType)
			if retval != C.UA_STATUSCODE_GOOD {
				goto failed
			}
			for _, value := range arrayList {
				cString := C.CString(value)
				defer C.free(unsafe.Pointer(cString))
				retval = C.UA_ArrayValueAppendString(&cArray, cString)
				if retval != C.UA_STATUSCODE_GOOD {
					goto failed
				}
			}
		}
	case UA_BYTESTRING:
		{
			uaType = C.UA_TYPES_BYTESTRING
			arrayList := value.Value.([][]byte)
			retval = C.UA_ArrayValueInit(&cArray, uaType)
			if retval != C.UA_STATUSCODE_GOOD {
				goto failed
			}
			for _, value := range arrayList {
				cString := C.CBytes(value)
				defer C.free(cString)
				retval = C.UA_ArrayValueAppendByteString(&cArray, cString, C.size_t(len(value)))
				if retval != C.UA_STATUSCODE_GOOD {
					goto failed
				}
			}
		}
	default:
		return fmt.Errorf("ua client covert node array value to variant failed, type = %d not support", uint32(value.Type))
	}

	C.UA_VariantArrayValue(variant, &cArray, uaType)
	return nil

failed:
	return fmt.Errorf("ua client covert node array value to variant failed, type = %d, retval = 0x%x", uint32(value.Type), retval)
}

func UA_VariantClangValue(value NodeValue, variant *C.UA_Variant) error {
	if value.Array {
		return UA_VariantFromArrayValue(value, variant)
	}
	return UA_VariantFromSingleValue(value, variant)
}

func UA_NodeTreeExpand(cNodeTree *C.NodeTree) *NodeTree {
	subNodes := make([]*NodeTree, 0)
	node := C.UA_NodeTree_head(cNodeTree)
	for {
		if unsafe.Pointer(node) == nil {
			break
		}
		subNodes = append(subNodes, UA_NodeTreeExpand(node))
		node = C.UA_NodeTree_next(node)
	}
	return &NodeTree{
		Level: uint32(cNodeTree.level),
		Node: NodeInfo{
			NsIndex: uint32(cNodeTree.index),
			NodeID:  C.GoString(cNodeTree.nodeID)},
		SubNodes: subNodes}
}

// UA_Client //
func NewClient(addr string) (*Client, error) {
	client := C.UA_Client_new()
	if client == nil {
		return nil, errors.New("ua client create failed")
	}

	goClient := &Client{addr: addr, cli: uintptr(unsafe.Pointer(client))}
	C.UA_Logger_init(&goClient.cLogger, C.UA_Logger_golang, C.UA_LoggerWrapper, nil)

	cConfig := C.UA_Client_getConfig(client)
	cConfig.logger = goClient.cLogger
	C.UA_ClientConfig_setDefault(cConfig)

	cStr := C.CString(addr)
	defer C.free(unsafe.Pointer(cStr))

	retval := C.UA_Client_connect(client, cStr)
	if retval != C.UA_STATUSCODE_GOOD {
		C.UA_Client_delete(client)
		return nil, fmt.Errorf("ua client connect failed, retval = 0x%x", uint32(retval))
	}

	return goClient, nil
}

func (c *Client) Close() {
	client := (*C.UA_Client)(unsafe.Pointer(c.cli))
	C.UA_Client_disconnect(client)
	C.UA_Client_delete(client)
}

func (c *Client) CheckState() {
	client := (*C.UA_Client)(unsafe.Pointer(c.cli))

	var retval C.UA_StatusCode
	var channelStatus C.UA_SecureChannelState
	C.UA_Client_getState(client, &channelStatus, nil, &retval)
	if retval != C.UA_STATUSCODE_GOOD || channelStatus == C.UA_SECURECHANNELSTATE_CLOSED {
		logs.Error("ua client get connect status %d, channel status %d to reconnect", uint32(retval), uint32(channelStatus))

		C.UA_Client_disconnect(client)

		cStr := C.CString(c.addr)
		defer C.free(unsafe.Pointer(cStr))

		retval := C.UA_Client_connect(client, cStr)
		if retval != C.UA_STATUSCODE_GOOD {
			logs.Warning("ua client reconnect failed, retval = 0x%x", uint32(retval))
		}
	}
}

func (c *Client) ReadNode(node NodeInfo) (*NodeValue, error) {
	cID := C.CString(node.NodeID)
	defer C.free(unsafe.Pointer(cID))

	client := (*C.UA_Client)(unsafe.Pointer(c.cli))

	var variant C.UA_Variant
	retval := C.UA_Client_readValueAttribute(client, C.UA_NODEID_STRING(C.UA_UInt16(node.NsIndex), cID), &variant)
	if retval != C.UA_STATUSCODE_GOOD {
		return nil, fmt.Errorf("ua client read value failed, retval = 0x%x", uint32(retval))
	}
	defer C.UA_Variant_clear(&variant)

	return UA_VariantGolangValue(&variant)
}

func (c *Client) ReadNodes(nodes []NodeInfo) ([]*NodeValue, error) {
	cReadValueIDs := C.UA_ReadValueID_alloc(C.int(len(nodes)))
	if cReadValueIDs == nil {
		return nil, errors.New("ua client alloc read value ids failed, point is nil")
	}

	defer C.UA_ReadValueID_free(cReadValueIDs)

	cNodeIDs := make([]unsafe.Pointer, 0)

	for i, node := range nodes {
		cID := C.CString(node.NodeID)
		cNodeIDs = append(cNodeIDs, unsafe.Pointer(cID))
		C.UA_ReadValueID_string(cReadValueIDs, C.int(i), C.UA_UInt16(node.NsIndex), cID, C.UA_ATTRIBUTEID_VALUE)
	}

	defer func() {
		for _, cID := range cNodeIDs {
			C.free(cID)
		}
	}()

	var request C.UA_ReadRequest
	C.UA_ReadRequest_init(&request)

	request.nodesToReadSize = C.size_t(len(nodes))
	request.nodesToRead = cReadValueIDs

	response := C.UA_Client_Service_read((*C.UA_Client)(unsafe.Pointer(c.cli)), request)
	defer C.UA_ReadResponse_clear(&response)

	if C.UA_STATUSCODE_GOOD != response.responseHeader.serviceResult {
		return nil, fmt.Errorf("ua client read nodes failed, retval = 0x%x", uint32(response.responseHeader.serviceResult))
	}

	nodeValues := make([]*NodeValue, 0)

	for i := C.size_t(0); i < response.resultsSize; i++ {
		variant := C.UA_ReadResponse_variant(&response, C.int(i))
		nodeValue, err := UA_VariantGolangValue(variant)
		if err != nil {
			nodeValues = append(nodeValues, NewEmptyNodeValue())
		} else {
			nodeValues = append(nodeValues, nodeValue)
		}
	}

	return nodeValues, nil
}

func (c *Client) BrowseNode() ([]*NodeTree, error) {
	cNodeTree := C.UA_NodeTree_root_init()
	if cNodeTree == nil {
		return nil, errors.New("ua client browse node failed, point is nil")
	}

	defer C.UA_NodeTree_clear(cNodeTree)

	retval := C.UA_Browse_nodeTree((*C.UA_Client)(unsafe.Pointer(c.cli)), cNodeTree)
	if retval != C.UA_STATUSCODE_GOOD {
		return nil, fmt.Errorf("ua client browse node failed, retval = 0x%x", uint32(retval))
	}

	return UA_NodeTreeExpand(cNodeTree).SubNodes, nil
}

func (c *Client) WriteNode(node NodeInfo, value NodeValue) error {
	cID := C.CString(node.NodeID)
	defer C.free(unsafe.Pointer(cID))

	client := (*C.UA_Client)(unsafe.Pointer(c.cli))

	var variant C.UA_Variant
	err := UA_VariantClangValue(value, &variant)
	if err != nil {
		return err
	}
	defer C.UA_Variant_clear(&variant)

	retval := C.UA_Client_writeValueAttribute(client, C.UA_NODEID_STRING(C.UA_UInt16(node.NsIndex), cID), &variant)
	if retval != C.UA_STATUSCODE_GOOD {
		return fmt.Errorf("ua client write value failed, retval = 0x%x", uint32(retval))
	}

	return nil
}

// UA_Server //
func NewServer(addr string, port int) (*Server, error) {
	if err := ListenTest(addr, port); err != nil {
		return nil, err
	}

	server := C.UA_Server_new()
	if server == nil {
		return nil, errors.New("ua server create failed")
	}

	cAddr := C.CString(addr)
	defer C.free(unsafe.Pointer(cAddr))

	goServer := &Server{addr: addr, srv: uintptr(unsafe.Pointer(server)), namespace: make(map[string]uint32)}
	C.UA_Logger_init(&goServer.cLogger, C.UA_Logger_golang, C.UA_LoggerWrapper, nil)

	cConfig := C.UA_Server_getConfig(server)
	cConfig.logger = goServer.cLogger

	C.UA_ServerConfig_setMinimal(cConfig, C.UA_UInt16(port), nil)
	C.UA_String_clear(&cConfig.customHostname)
	cConfig.customHostname = C.UA_String_fromChars(cAddr)

	retval := C.UA_Server_run_startup(server)
	if retval != C.UA_STATUSCODE_GOOD {
		C.UA_Server_delete(server)
		return nil, fmt.Errorf("ua server startup failed, retval = 0x%x", uint32(retval))
	}

	goServer.running = true
	goServer.Add(1)
	go goServer.serverRunningTask()

	logs.Info("ua server startup success, addr: %s, port: %d", addr, port)
	return goServer, nil
}

func (s *Server) serverRunningTask() {
	s.Done()

	server := (*C.UA_Server)(unsafe.Pointer(s.srv))
	for s.running {
		C.UA_Server_run_iterate(server, true)
	}
}

func (s *Server) CheckState() bool {
	return s.running
}

func (s *Server) AddNameSpace(name string) (uint32, error) {
	index, ok := s.namespace[name]
	if ok {
		return index, nil
	}
	server := (*C.UA_Server)(unsafe.Pointer(s.srv))

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cIndex := C.UA_Server_addNamespace(server, cName)
	retval := C.UA_ServerAddObject(server, cIndex, cName)
	if retval != C.UA_STATUSCODE_GOOD {
		return 0, fmt.Errorf("ua server add name space failed, retval = 0x%x", uint32(retval))
	}
	s.namespace[name] = uint32(cIndex)
	logs.Info("ua server add name space success, name: %s, index: %d", name, uint32(cIndex))
	return uint32(cIndex), nil
}

func (s *Server) AddNode(parent, current NodeInfo, name string, value NodeValue) error {
	server := (*C.UA_Server)(unsafe.Pointer(s.srv))

	cParentID := C.CString(parent.NodeID)
	defer C.free(unsafe.Pointer(cParentID))

	cCurrentID := C.CString(current.NodeID)
	defer C.free(unsafe.Pointer(cCurrentID))

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var variant C.UA_Variant
	err := UA_VariantClangValue(value, &variant)
	if err != nil {
		logs.Error("ua server add node failed, convert value to variant failed, error: %s", err.Error())
		return err
	}
	defer C.UA_Variant_clear(&variant)

	retval := C.UA_ServerAddVariable(server,
		C.UA_UInt16(parent.NsIndex), cParentID,
		C.UA_UInt16(current.NsIndex), cCurrentID,
		cName, &variant)

	if retval != C.UA_STATUSCODE_GOOD {
		return fmt.Errorf("ua server add node failed, retval = 0x%x", uint32(retval))
	}
	return nil
}

func (s *Server) ReadNode(node NodeInfo) (*NodeValue, error) {
	server := (*C.UA_Server)(unsafe.Pointer(s.srv))

	cID := C.CString(node.NodeID)
	defer C.free(unsafe.Pointer(cID))

	var variant C.UA_Variant
	retval := C.UA_Server_readValue(server, C.UA_NODEID_STRING(C.UA_UInt16(node.NsIndex), cID), &variant)
	if retval != C.UA_STATUSCODE_GOOD {
		return nil, fmt.Errorf("ua server read value failed, retval = 0x%x", uint32(retval))
	}
	defer C.UA_Variant_clear(&variant)

	return UA_VariantGolangValue(&variant)
}

func (s *Server) WriteNode(node NodeInfo, value NodeValue) error {
	server := (*C.UA_Server)(unsafe.Pointer(s.srv))

	cID := C.CString(node.NodeID)
	defer C.free(unsafe.Pointer(cID))

	var variant C.UA_Variant
	err := UA_VariantClangValue(value, &variant)
	if err != nil {
		logs.Error("ua server add node failed, convert value to variant failed, error: %s", err.Error())
		return err
	}
	defer C.UA_Variant_clear(&variant)

	retval := C.UA_Server_writeValue(server, C.UA_NODEID_STRING(C.UA_UInt16(node.NsIndex), cID), variant)
	if retval != C.UA_STATUSCODE_GOOD {
		return fmt.Errorf("ua server write value failed, retval = 0x%x", uint32(retval))
	}

	return nil
}

func (s *Server) Close() {
	server := (*C.UA_Server)(unsafe.Pointer(s.srv))

	s.running = false
	s.Wait()

	C.UA_Server_run_shutdown(server)
	C.UA_Server_delete(server)
}
