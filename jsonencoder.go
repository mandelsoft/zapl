package zapl

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mandelsoft/zapl/pool"
	"go.uber.org/zap/zapcore"
)

var _dataPool = pool.New(func() *DataEncoder {
	return &DataEncoder{object: make(map[string]interface{})}
})

func JSONEncodeObject(m zapcore.ObjectMarshaler) ([]byte, error) {
	e := _dataPool.Get()
	defer e.free()

	err := m.MarshalLogObject(e)
	if err != nil {
		return nil, err
	}
	return json.Marshal(e.object)
}

func JSONEncodeArray(m zapcore.ArrayMarshaler) ([]byte, error) {
	e := _dataPool.Get()
	defer e.free()

	err := m.MarshalLogArray(e)
	if err != nil {
		return nil, err
	}
	return json.Marshal(e.slice)
}

////////////////////////////////////////////////////////////////////////////////

type DataEncoder struct {
	namespace string
	object    map[string]interface{}
	slice     []interface{}
}

var _ zapcore.ArrayEncoder = (*DataEncoder)(nil)
var _ zapcore.ObjectEncoder = (*DataEncoder)(nil)

func NewJSONEncoder() *DataEncoder {
	return &DataEncoder{object: make(map[string]interface{})}
}

func (e *DataEncoder) clone() *DataEncoder {
	return _dataPool.Get()
}

func (e *DataEncoder) Reset() {
	clear(e.object)
	e.slice = e.slice[:0]
	e.namespace = ""
}

func (e *DataEncoder) free() {
	e.Reset()
	_dataPool.Put(e)
}

func (e *DataEncoder) Object() map[string]interface{} {
	return e.object
}

func (e *DataEncoder) Slice() []interface{} {
	return e.slice
}

// array encoder

func (e *DataEncoder) AppendBool(v bool) {
	e.slice = append(e.slice, v)
}

func (e *DataEncoder) AppendByteString(v []byte) {
	e.slice = append(e.slice, v)
}

func (e *DataEncoder) AppendComplex128(v complex128) {
	e.slice = append(e.slice, v)
}

func (e *DataEncoder) AppendComplex64(v complex64) {
	e.slice = append(e.slice, v)
}

func (e *DataEncoder) AppendFloat64(v float64) {
	e.slice = append(e.slice, v)
}

func (e *DataEncoder) AppendFloat32(v float32) {
	e.slice = append(e.slice, v)
}

func (e *DataEncoder) AppendInt(v int) {
	e.slice = append(e.slice, v)
}

func (e *DataEncoder) AppendInt64(v int64) {
	e.slice = append(e.slice, v)
}

func (e *DataEncoder) AppendInt32(v int32) {
	e.slice = append(e.slice, v)
}

func (e *DataEncoder) AppendInt16(v int16) {
	e.slice = append(e.slice, v)
}

func (e *DataEncoder) AppendInt8(v int8) {
	e.slice = append(e.slice, v)
}

func (e *DataEncoder) AppendString(v string) {
	e.slice = append(e.slice, v)
}

func (e *DataEncoder) AppendUint(v uint) {
	e.slice = append(e.slice, v)
}

func (e *DataEncoder) AppendUint64(v uint64) {
	e.slice = append(e.slice, v)
}

func (e *DataEncoder) AppendUint32(v uint32) {
	e.slice = append(e.slice, v)
}

func (e *DataEncoder) AppendUint16(v uint16) {
	e.slice = append(e.slice, v)
}

func (e *DataEncoder) AppendUint8(v uint8) {
	e.slice = append(e.slice, v)
}

func (e *DataEncoder) AppendUintptr(v uintptr) {
	e.slice = append(e.slice, fmt.Sprintf("%#x", v))
}

func (e *DataEncoder) AppendDuration(v time.Duration) {
	e.slice = append(e.slice, v)
}

func (e *DataEncoder) AppendTime(v time.Time) {
	e.slice = append(e.slice, v)
}

func (e *DataEncoder) AppendArray(marshaler zapcore.ArrayMarshaler) error {
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

func (e *DataEncoder) AppendObject(marshaler zapcore.ObjectMarshaler) error {
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

func (e *DataEncoder) AppendReflected(v interface{}) error {
	e.slice = append(e.slice, v)
	return nil
}

// object encoder

func (e *DataEncoder) AddArray(key string, marshaler zapcore.ArrayMarshaler) error {
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

func (e *DataEncoder) AddObject(key string, marshaler zapcore.ObjectMarshaler) error {
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

func (e *DataEncoder) AddBinary(key string, value []byte) {
	e.object[e.namespace+key] = value
}

func (e *DataEncoder) AddByteString(key string, value []byte) {
	e.object[e.namespace+key] = value
}

func (e *DataEncoder) AddBool(key string, value bool) {
	e.object[e.namespace+key] = value
}

func (e *DataEncoder) AddComplex128(key string, value complex128) {
	e.object[e.namespace+key] = value
}

func (e *DataEncoder) AddComplex64(key string, value complex64) {
	e.object[e.namespace+key] = value
}

func (e *DataEncoder) AddDuration(key string, value time.Duration) {
	e.object[e.namespace+key] = value
}

func (e *DataEncoder) AddFloat64(key string, value float64) {
	e.object[e.namespace+key] = value
}

func (e *DataEncoder) AddFloat32(key string, value float32) {
	e.object[e.namespace+key] = value
}

func (e *DataEncoder) AddInt(key string, value int) {
	e.object[e.namespace+key] = value
}

func (e *DataEncoder) AddInt64(key string, value int64) {
	e.object[e.namespace+key] = value
}

func (e *DataEncoder) AddInt32(key string, value int32) {
	e.object[e.namespace+key] = value
}

func (e *DataEncoder) AddInt16(key string, value int16) {
	e.object[e.namespace+key] = value
}

func (e *DataEncoder) AddInt8(key string, value int8) {
	e.object[e.namespace+key] = value
}

func (e *DataEncoder) AddString(key, value string) {
	e.object[e.namespace+key] = value
}

func (e *DataEncoder) AddTime(key string, value time.Time) {
	e.object[e.namespace+key] = value
}

func (e *DataEncoder) AddUint(key string, value uint) {
	e.object[e.namespace+key] = value
}

func (e *DataEncoder) AddUint64(key string, value uint64) {
	e.object[e.namespace+key] = value
}

func (e *DataEncoder) AddUint32(key string, value uint32) {
	e.object[e.namespace+key] = value
}

func (e *DataEncoder) AddUint16(key string, value uint16) {
	e.object[e.namespace+key] = value
}

func (e *DataEncoder) AddUint8(key string, value uint8) {
	e.object[e.namespace+key] = value
}

func (e *DataEncoder) AddUintptr(key string, value uintptr) {
	e.object[e.namespace+key] = fmt.Sprintf("%#x", value)
}

func (e *DataEncoder) AddReflected(key string, value interface{}) error {
	e.object[e.namespace+key] = value
	return nil
}

func (e *DataEncoder) OpenNamespace(key string) {
	e.namespace = e.namespace + key + "/"
}
