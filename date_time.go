package msgpack

import (
	"reflect"
	"time"

	"github.com/go-openapi/strfmt"
)

var dateTimeExtId int8 = 2

func init() {
	dateTimeType := reflect.TypeOf((*strfmt.DateTime)(nil)).Elem()
	registerExt(dateTimeExtId, dateTimeType, encodeDateTimeValue, decodeDateTimeValue)
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

func (e *Encoder) EncodeDateTime(tm strfmt.DateTime) error {
	return e.EncodeTime(time.Time(tm))
}

func (d *Decoder) DecodeDateTime() (strfmt.DateTime, error) {
	result, err := d.DecodeTime()
	if err != nil {
		return strfmt.DateTime{}, err
	}
	return strfmt.DateTime(result), nil
}
