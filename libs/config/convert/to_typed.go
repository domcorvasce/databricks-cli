package convert

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/databricks/cli/libs/config"
)

func ToTyped(dst any, src config.Value) error {
	dstv := reflect.ValueOf(dst)

	// Dereference pointer if necessary
	for dstv.Kind() == reflect.Pointer {
		// If the source value is nil and the destination is a settable pointer,
		// set the destination to nil. Also see `end_to_end_test.go`.
		if dstv.CanSet() && src == config.NilValue {
			dstv.SetZero()
			return nil
		}
		if dstv.IsNil() {
			dstv.Set(reflect.New(dstv.Type().Elem()))
		}
		dstv = dstv.Elem()
	}

	// Verify that vv is settable.
	if !dstv.CanSet() {
		panic("cannot set destination value")
	}

	switch dstv.Kind() {
	case reflect.Struct:
		return toTypedStruct(dstv, src)
	case reflect.Map:
		return toTypedMap(dstv, src)
	case reflect.Slice:
		return toTypedSlice(dstv, src)
	case reflect.String:
		return toTypedString(dstv, src)
	case reflect.Bool:
		return toTypedBool(dstv, src)
	case reflect.Int, reflect.Int32, reflect.Int64:
		return toTypedInt(dstv, src)
	case reflect.Float32, reflect.Float64:
		return toTypedFloat(dstv, src)
	}

	return fmt.Errorf("unsupported type: %s", dstv.Kind())
}

func toTypedStruct(dst reflect.Value, src config.Value) error {
	switch src.Kind() {
	case config.KindMap:
		info := getStructInfo(dst.Type())
		for k, v := range src.MustMap() {
			index, ok := info.Fields[k]
			if !ok {
				// Ignore unknown fields.
				// A warning will be printed later. See PR #904.
				continue
			}

			// Create intermediate structs embedded as pointer types.
			// Code inspired by [reflect.FieldByIndex] implementation.
			f := dst
			for i, x := range index {
				if i > 0 {
					if f.Kind() == reflect.Pointer {
						if f.IsNil() {
							f.Set(reflect.New(f.Type().Elem()))
						}
						f = f.Elem()
					}
				}
				f = f.Field(x)
			}

			err := ToTyped(f.Addr().Interface(), v)
			if err != nil {
				return err
			}
		}

		// Populate field(s) for [config.Value], if any.
		if info.ValueField != nil {
			vv := dst.FieldByIndex(info.ValueField)
			vv.Set(reflect.ValueOf(src))
		}

		return nil
	case config.KindNil:
		dst.SetZero()
		return nil
	}

	return TypeError{
		value: src,
		msg:   fmt.Sprintf("expected a map, found a %s", src.Kind()),
	}
}

func toTypedMap(dst reflect.Value, src config.Value) error {
	switch src.Kind() {
	case config.KindMap:
		m := src.MustMap()

		// Always overwrite.
		dst.Set(reflect.MakeMapWithSize(dst.Type(), len(m)))
		for k, v := range m {
			kv := reflect.ValueOf(k)
			vv := reflect.New(dst.Type().Elem())
			err := ToTyped(vv.Interface(), v)
			if err != nil {
				return err
			}
			dst.SetMapIndex(kv, vv.Elem())
		}
		return nil
	case config.KindNil:
		dst.SetZero()
		return nil
	}

	return TypeError{
		value: src,
		msg:   fmt.Sprintf("expected a map, found a %s", src.Kind()),
	}
}

func toTypedSlice(dst reflect.Value, src config.Value) error {
	switch src.Kind() {
	case config.KindSequence:
		seq := src.MustSequence()

		// Always overwrite.
		dst.Set(reflect.MakeSlice(dst.Type(), len(seq), len(seq)))
		for i := range seq {
			err := ToTyped(dst.Index(i).Addr().Interface(), seq[i])
			if err != nil {
				return err
			}
		}
		return nil
	case config.KindNil:
		dst.SetZero()
		return nil
	}

	return TypeError{
		value: src,
		msg:   fmt.Sprintf("expected a sequence, found a %s", src.Kind()),
	}
}

func toTypedString(dst reflect.Value, src config.Value) error {
	switch src.Kind() {
	case config.KindString:
		dst.SetString(src.MustString())
		return nil
	case config.KindBool:
		dst.SetString(strconv.FormatBool(src.MustBool()))
		return nil
	case config.KindInt:
		dst.SetString(strconv.FormatInt(src.MustInt(), 10))
		return nil
	case config.KindFloat:
		dst.SetString(strconv.FormatFloat(src.MustFloat(), 'f', -1, 64))
		return nil
	}

	return TypeError{
		value: src,
		msg:   fmt.Sprintf("expected a string, found a %s", src.Kind()),
	}
}

func toTypedBool(dst reflect.Value, src config.Value) error {
	switch src.Kind() {
	case config.KindBool:
		dst.SetBool(src.MustBool())
		return nil
	case config.KindString:
		// See https://github.com/go-yaml/yaml/blob/f6f7691b1fdeb513f56608cd2c32c51f8194bf51/decode.go#L684-L693.
		switch src.MustString() {
		case "y", "Y", "yes", "Yes", "YES", "on", "On", "ON":
			dst.SetBool(true)
			return nil
		case "n", "N", "no", "No", "NO", "off", "Off", "OFF":
			dst.SetBool(false)
			return nil
		}
	}

	return TypeError{
		value: src,
		msg:   fmt.Sprintf("expected a boolean, found a %s", src.Kind()),
	}
}

func toTypedInt(dst reflect.Value, src config.Value) error {
	switch src.Kind() {
	case config.KindInt:
		dst.SetInt(src.MustInt())
		return nil
	case config.KindString:
		if i64, err := strconv.ParseInt(src.MustString(), 10, 64); err == nil {
			dst.SetInt(i64)
			return nil
		}
	}

	return TypeError{
		value: src,
		msg:   fmt.Sprintf("expected an int, found a %s", src.Kind()),
	}
}

func toTypedFloat(dst reflect.Value, src config.Value) error {
	switch src.Kind() {
	case config.KindFloat:
		dst.SetFloat(src.MustFloat())
		return nil
	case config.KindString:
		if f64, err := strconv.ParseFloat(src.MustString(), 64); err == nil {
			dst.SetFloat(f64)
			return nil
		}
	}

	return TypeError{
		value: src,
		msg:   fmt.Sprintf("expected a float, found a %s", src.Kind()),
	}
}
