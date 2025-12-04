package zapl_test

import (
	"github.com/mandelsoft/goutils/maputils"
	"github.com/mandelsoft/zapl"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap/zapcore"
)

var _ = Describe("JSONEncoder", func() {

	var enc *zapl.JSONEncoder

	ar := []interface{}{int(1), uint(2), int32(3), uint32(4), int64(5), uint64(6), byte(7), int8(8), uint8(9), int16(10), uint16(11), int32(12), "13", float32(14.1), float64(15.1), true}

	obj := map[string]interface{}{
		"a": 1,
		"b": true,
	}

	Context("arrays", func() {
		BeforeEach(func() {
			enc = &zapl.JSONEncoder{}
		})

		It("flat", func() {
			enc.AppendArray(TestArray(ar))
			Expect(enc.String()).To(Equal(`[[1,2,3,4,5,6,7,8,9,10,11,12,"13",14.1,15.1,true]]`))
		})
	})

	Context("objects", func() {
		BeforeEach(func() {
			enc = zapl.NewJSONObjectEncoder()
		})

		It("flat", func() {
			enc.AddObject("obj", TestObject(obj))
			Expect(enc.String()).To(Equal(`{"obj":{"a":1,"b":true}}`))
		})

		It("nested", func() {
			enc.AddObject("obj", TestObject(map[string]interface{}{"nested": obj}))
			Expect(enc.String()).To(Equal(`{"obj":{"nested":{"a":1,"b":true}}}`))
		})
	})
})

// TestArray is a simple ArrayMarshaler for testing
type TestArray []interface{}

func (a TestArray) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	for _, v := range a {
		switch val := v.(type) {
		case string:
			enc.AppendString(val)
		case int:
			enc.AppendInt(val)
		case int8:
			enc.AppendInt8(val)
		case int16:
			enc.AppendInt16(val)
		case int32:
			enc.AppendInt32(val)
		case int64:
			enc.AppendInt64(val)

		case uint:
			enc.AppendUint(val)
		case uint8:
			enc.AppendUint8(val)
		case uint16:
			enc.AppendUint16(val)
		case uint32:
			enc.AppendUint32(val)
		case uint64:
			enc.AppendUint64(val)

		case bool:
			enc.AppendBool(val)
		case float32:
			enc.AppendFloat32(val)
		case float64:
			enc.AppendFloat64(val)
		default:
			enc.AppendReflected(val)
		}
	}
	return nil
}

type TestObject map[string]interface{}

func (a TestObject) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	for _, k := range maputils.OrderedKeys(a) {
		switch val := a[k].(type) {
		case string:
			enc.AddString(k, val)
		case int:
			enc.AddInt(k, val)
		case int8:
			enc.AddInt8(k, val)
		case int16:
			enc.AddInt16(k, val)
		case int32:
			enc.AddInt32(k, val)
		case int64:
			enc.AddInt64(k, val)

		case uint:
			enc.AddUint(k, val)
		case uint8:
			enc.AddUint8(k, val)
		case uint16:
			enc.AddUint16(k, val)
		case uint32:
			enc.AddUint32(k, val)
		case uint64:
			enc.AddUint64(k, val)

		case bool:
			enc.AddBool(k, val)
		case float32:
			enc.AddFloat32(k, val)
		case float64:
			enc.AddFloat64(k, val)

		case map[string]interface{}:
			enc.AddObject(k, TestObject(val))
		case []interface{}:
			enc.AddArray(k, TestArray(val))

		default:
			enc.AddReflected(k, val)
		}
	}
	return nil
}
