package msgpack

import (
	"reflect"
	"time"

	"github.com/go-openapi/strfmt"
)

func init() {
	dateTimeType := reflect.TypeOf((*strfmt.DateTime)(nil)).Elem()
	registerExt(timeExtId, dateTimeType, encodeDateTimeValue, decodeDateTimeValue)
}

func encodeDateTimeValue(e *Encoder, v reflect.Value) error {
	tm := v.Interface().(strfmt.DateTime)
	b := e.encodeTime(time.Time(tm))
	return e.write(b)
}

func decodeDateTimeValue(d *Decoder, v reflect.Value) error {
	tm, err := d.DecodeDateTime()
	if err != nil {
		return err
	}
	v.Set(reflect.ValueOf(tm))
	return nil
}
