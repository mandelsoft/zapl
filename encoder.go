package zapl

import (
	"encoding/json"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/mandelsoft/logging"
	"github.com/mandelsoft/zapl/pool"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

var _encPool = pool.New(func() *Encoder {
	return &Encoder{}
})

var _bufferPool = buffer.NewPool()
var _fieldPool = pool.New(func() []any {
	return []any{}
})

type cache struct {
	lock    sync.Mutex
	ctx     logging.Context
	creator LoggerFactory
	cache   map[string]logging.Logger
}

func (c *cache) Get(name string) logging.Logger {
	c.lock.Lock()
	defer c.lock.Unlock()
	l := c.cache[name]
	if l == nil {
		l = c.creator(c.ctx, name)
		c.cache[name] = l
	}
	return l
}

type LoggerFactory func(ctx logging.Context, name string) logging.Logger

type Encoder struct {
	cache *cache

	fields []any
	enc    DataEncoder
	buf    *buffer.Buffer
}

var _ zapcore.Encoder = (*Encoder)(nil)

// StaticLogging configures the static logger creation.
func StaticLogging(ctx logging.Context, name string) logging.Logger {
	return ctx.Logger(logging.NewRealm(name))
}

// DynamicLogging configured the dynamic logger creation.
func DynamicLogging(ctx logging.Context, name string) logging.Logger {
	return logging.DynamicLogger(ctx, logging.NewRealm(name))
}

func NewEncoder(ctx logging.Context, creator ...LoggerFactory) zapcore.Encoder {
	c := StaticLogging
	for _, o := range creator {
		if o != nil {
			c = o
			break
		}

	}
	return &Encoder{cache: &cache{ctx: ctx, creator: c, cache: map[string]logging.Logger{}}}
}

func (e *Encoder) Clone() zapcore.Encoder {
	return e.clone()
}

func (e *Encoder) clone() *Encoder {
	clone := _encPool.Get()

	clone.cache = e.cache
	clone.fields = append(_fieldPool.Get(), e.fields...)
	clone.buf = _bufferPool.Get()
	return clone
}

func (e *Encoder) free() {
	e.buf = nil
	_fieldPool.Put(e.fields[:0])
	_encPool.Put(e)
}

func realmNameForLoggerName(name string) string {
	if name == "" {
		return "zap"
	}

	return "zap/" + strings.ReplaceAll(name, ".", "/")
}

func (e *Encoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	e = e.clone()

	r := realmNameForLoggerName(entry.LoggerName)
	e.fields = nil

	l := e.cache.Get(r)
	for _, f := range fields {
		f.AddTo(e)
	}
	if entry.Caller.Defined {
		if entry.Caller.File != "" {
			e.fields = append(e.fields, "file", entry.Caller.File+":"+strconv.Itoa(entry.Caller.Line))
		}
		if entry.Caller.Function != "" {
			e.fields = append(e.fields, "function", entry.Caller.Function)
		}
	}
	switch entry.Level {
	case zapcore.DebugLevel:
		l.Debug(entry.Message, e.fields...)
	case zapcore.InfoLevel:
		l.Info(entry.Message, e.fields...)
	case zapcore.WarnLevel:
		l.Warn(entry.Message, e.fields...)
	case zapcore.ErrorLevel:
		l.Error(entry.Message, e.fields...)
	case zapcore.FatalLevel, zapcore.DPanicLevel, zapcore.PanicLevel:
		l.Error(entry.Message, e.fields...)
	}
	buf := e.buf
	e.free()
	return buf, nil
}

////////////////////////////////////////////////////////////////////////////////

func (e *Encoder) AddArray(key string, marshaler zapcore.ArrayMarshaler) error {
	if err := marshaler.MarshalLogArray(&e.enc); err != nil {
		return err
	}
	s, err := json.Marshal(e.enc.slice)
	if err != nil {
		return err
	}
	e.fields = append(e.fields, logging.KeyValue(e.enc.namespace+key, string(s)))
	return nil
}

func (e *Encoder) AddObject(key string, marshaler zapcore.ObjectMarshaler) error {
	if err := marshaler.MarshalLogObject(&e.enc); err != nil {
		return err
	}
	s, err := json.Marshal(e.enc.object)
	if err != nil {
		return err
	}
	return e.AddReflected(key, s)
}

func (e *Encoder) AddBinary(key string, value []byte) {
	e.AddReflected(key, value)
}

func (e *Encoder) AddByteString(key string, value []byte) {
	e.AddReflected(key, value)
}

func (e *Encoder) AddBool(key string, value bool) {
	e.AddReflected(key, value)
}

func (e *Encoder) AddComplex128(key string, value complex128) {
	e.AddReflected(key, value)
}

func (e *Encoder) AddComplex64(key string, value complex64) {
	e.AddReflected(key, value)
}

func (e *Encoder) AddDuration(key string, value time.Duration) {
	e.AddReflected(key, value)
}

func (e *Encoder) AddFloat64(key string, value float64) {
	e.AddReflected(key, value)
}

func (e *Encoder) AddFloat32(key string, value float32) {
	e.AddReflected(key, value)
}

func (e *Encoder) AddInt(key string, value int) {
	e.AddReflected(key, value)
}

func (e *Encoder) AddInt64(key string, value int64) {
	e.AddReflected(key, value)
}

func (e *Encoder) AddInt32(key string, value int32) {
	e.AddReflected(key, value)
}

func (e *Encoder) AddInt16(key string, value int16) {
	e.AddReflected(key, value)
}

func (e *Encoder) AddInt8(key string, value int8) {
	e.AddReflected(key, value)
}

func (e *Encoder) AddString(key, value string) {
	e.AddReflected(key, value)
}

func (e *Encoder) AddTime(key string, value time.Time) {
	e.AddReflected(key, value)
}

func (e *Encoder) AddUint(key string, value uint) {
	e.AddReflected(key, value)
}

func (e *Encoder) AddUint64(key string, value uint64) {
	e.AddReflected(key, value)
}

func (e *Encoder) AddUint32(key string, value uint32) {
	e.AddReflected(key, value)
}

func (e *Encoder) AddUint16(key string, value uint16) {
	e.AddReflected(key, value)
}

func (e *Encoder) AddUint8(key string, value uint8) {
	e.AddReflected(key, value)
}

func (e *Encoder) AddUintptr(key string, value uintptr) {
	// TODO implement me
	panic("implement me")
}

func (e *Encoder) AddReflected(key string, value interface{}) error {
	e.fields = append(e.fields, logging.KeyValue(e.enc.namespace+key, value))
	return nil
}

func (e *Encoder) OpenNamespace(key string) {
	e.enc.OpenNamespace(key)
}
