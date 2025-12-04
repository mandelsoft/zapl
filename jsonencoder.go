package zapl

import (
	"encoding/json"
	"fmt"
	"time"

	"go.uber.org/zap/zapcore"
)

type JSONEncoder struct {
	namespace string
	object    map[string]interface{}
	slice     []interface{}
}

var _ zapcore.ArrayEncoder = (*JSONEncoder)(nil)
var _ zapcore.ObjectEncoder = (*JSONEncoder)(nil)

func NewJSONObjectEncoder() *JSONEncoder {
	return &JSONEncoder{object: make(map[string]interface{})}
}

func (e *JSONEncoder) Bytes() ([]byte, error) {
	if e.object != nil {
		return json.Marshal(e.object)
	} else {
		return json.Marshal(e.slice)
	}
}

func (e *JSONEncoder) String() (string, error) {
	b, err := e.Bytes()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// array encoder

func (e *JSONEncoder) AppendBool(v bool) {
	e.slice = append(e.slice, v)
}

func (e *JSONEncoder) AppendByteString(v []byte) {
	e.slice = append(e.slice, v)
}

func (e *JSONEncoder) AppendComplex128(v complex128) {
	e.slice = append(e.slice, v)
}

func (e *JSONEncoder) AppendComplex64(v complex64) {
	e.slice = append(e.slice, v)
}

func (e *JSONEncoder) AppendFloat64(v float64) {
	e.slice = append(e.slice, v)
}

func (e *JSONEncoder) AppendFloat32(v float32) {
	e.slice = append(e.slice, v)
}

func (e *JSONEncoder) AppendInt(v int) {
	e.slice = append(e.slice, v)
}

func (e *JSONEncoder) AppendInt64(v int64) {
	e.slice = append(e.slice, v)
}

func (e *JSONEncoder) AppendInt32(v int32) {
	e.slice = append(e.slice, v)
}

func (e *JSONEncoder) AppendInt16(v int16) {
	e.slice = append(e.slice, v)
}

func (e *JSONEncoder) AppendInt8(v int8) {
	e.slice = append(e.slice, v)
}

func (e *JSONEncoder) AppendString(v string) {
	e.slice = append(e.slice, v)
}

func (e *JSONEncoder) AppendUint(v uint) {
	e.slice = append(e.slice, v)
}

func (e *JSONEncoder) AppendUint64(v uint64) {
	e.slice = append(e.slice, v)
}

func (e *JSONEncoder) AppendUint32(v uint32) {
	e.slice = append(e.slice, v)
}

func (e *JSONEncoder) AppendUint16(v uint16) {
	e.slice = append(e.slice, v)
}

func (e *JSONEncoder) AppendUint8(v uint8) {
	e.slice = append(e.slice, v)
}

func (e *JSONEncoder) AppendUintptr(v uintptr) {
	e.slice = append(e.slice, fmt.Sprintf("%#x", v))
}

func (e *JSONEncoder) AppendDuration(v time.Duration) {
	e.slice = append(e.slice, v)
}

func (e *JSONEncoder) AppendTime(v time.Time) {
	e.slice = append(e.slice, v)
}

func (e *JSONEncoder) AppendArray(marshaler zapcore.ArrayMarshaler) error {
	old := e.slice
	e.slice = nil
	err := marshaler.MarshalLogArray(e)
	if err != nil {
		return err
	}
	old = append(old, e.slice)
	e.slice = old
	return nil
}

func (e *JSONEncoder) AppendObject(marshaler zapcore.ObjectMarshaler) error {
	old := e.object
	ns := e.namespace
	e.object = map[string]interface{}{}

	err := marshaler.MarshalLogObject(e)
	if err != nil {
		return err
	}
	e.slice = append(e.slice, e.object)
	e.object = old
	e.namespace = ns
	return nil
}

func (e *JSONEncoder) AppendReflected(v interface{}) error {
	e.slice = append(e.slice, v)
	return nil
}

// object encoder

func (e *JSONEncoder) AddArray(key string, marshaler zapcore.ArrayMarshaler) error {
	old := e.slice
	e.slice = nil
	err := marshaler.MarshalLogArray(e)
	if err != nil {
		return err
	}
	e.object[e.namespace+key] = e.slice
	e.slice = old
	return nil
}

func (e *JSONEncoder) AddObject(key string, marshaler zapcore.ObjectMarshaler) error {
	ns := e.namespace
	old := e.object
	e.object = map[string]interface{}{}
	err := marshaler.MarshalLogObject(e)
	if err != nil {
		return err
	}
	old[e.namespace+key] = e.object
	e.object = old
	e.namespace = ns
	return nil
}

func (e *JSONEncoder) AddBinary(key string, value []byte) {
	e.object[e.namespace+key] = value
}

func (e *JSONEncoder) AddByteString(key string, value []byte) {
	e.object[e.namespace+key] = value
}

func (e *JSONEncoder) AddBool(key string, value bool) {
	e.object[e.namespace+key] = value
}

func (e *JSONEncoder) AddComplex128(key string, value complex128) {
	e.object[e.namespace+key] = value
}

func (e *JSONEncoder) AddComplex64(key string, value complex64) {
	e.object[e.namespace+key] = value
}

func (e *JSONEncoder) AddDuration(key string, value time.Duration) {
	e.object[e.namespace+key] = value
}

func (e *JSONEncoder) AddFloat64(key string, value float64) {
	e.object[e.namespace+key] = value
}

func (e *JSONEncoder) AddFloat32(key string, value float32) {
	e.object[e.namespace+key] = value
}

func (e *JSONEncoder) AddInt(key string, value int) {
	e.object[e.namespace+key] = value
}

func (e *JSONEncoder) AddInt64(key string, value int64) {
	e.object[e.namespace+key] = value
}

func (e *JSONEncoder) AddInt32(key string, value int32) {
	e.object[e.namespace+key] = value
}

func (e *JSONEncoder) AddInt16(key string, value int16) {
	e.object[e.namespace+key] = value
}

func (e *JSONEncoder) AddInt8(key string, value int8) {
	e.object[e.namespace+key] = value
}

func (e *JSONEncoder) AddString(key, value string) {
	e.object[e.namespace+key] = value
}

func (e *JSONEncoder) AddTime(key string, value time.Time) {
	e.object[e.namespace+key] = value
}

func (e *JSONEncoder) AddUint(key string, value uint) {
	e.object[e.namespace+key] = value
}

func (e *JSONEncoder) AddUint64(key string, value uint64) {
	e.object[e.namespace+key] = value
}

func (e *JSONEncoder) AddUint32(key string, value uint32) {
	e.object[e.namespace+key] = value
}

func (e *JSONEncoder) AddUint16(key string, value uint16) {
	e.object[e.namespace+key] = value
}

func (e *JSONEncoder) AddUint8(key string, value uint8) {
	e.object[e.namespace+key] = value
}

func (e *JSONEncoder) AddUintptr(key string, value uintptr) {
	e.object[e.namespace+key] = fmt.Sprintf("%#x", value)
}

func (e *JSONEncoder) AddReflected(key string, value interface{}) error {
	e.object[e.namespace+key] = value
	return nil
}

func (e *JSONEncoder) OpenNamespace(key string) {
	e.namespace = e.namespace + key + "/"
}
