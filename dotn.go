package dotn

import (
	"reflect"
	"strconv"
	"strings"
)

// Set a value following the dot notation path
func Set(str, value string, root interface{}) error {
	keys := strings.Split(str, ".")

	v := reflect.ValueOf(root)
	if v.Kind() != reflect.Ptr {
		return ErrInterfaceNotPointer
	}

	if !v.Elem().CanAddr() {
		return ErrInterfaceNotAddressable
	}

	return set(keys, value, v, "json") // TODO: Custom tag
}

func set(keys []string, value string, v reflect.Value, tag string) error {
	var key string
	var rest []string
	if len(keys) > 0 {
		key, rest = keys[0], keys[1:]
	}

	switch v.Kind() {
	case reflect.Interface:
		return set(keys, value, reflect.ValueOf(v.Interface()), tag)
	case reflect.Ptr:
		if v.IsNil() {
			return nil
		}

		return set(keys, value, reflect.Indirect(v), tag)
	case reflect.Array, reflect.Slice:
		// Convert key to integer
		i, err := strconv.Atoi(key)
		if err != nil {
			return err
		}

		// Check bounds
		if i > v.Len() {
			return ErrIndexOutOfRange
		}

		return set(rest, value, v.Index(i), tag)
	case reflect.Map:
		iter := v.MapRange()
		for iter.Next() {
			k := iter.Key()
			if k.Kind() != reflect.String {
				// Skip on non-strings
				continue
			}

			if k.String() == key {
				return set(rest, value, iter.Value(), tag)
			}
		}

		return &KeyNotFoundError{key}
	case reflect.Struct:
		t := v.Type()

		var name string
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)

			// Get field by tag or struct name
			t := f.Tag.Get(tag)
			if t == key {
				name = f.Name
				break
			}
		}

		if name == "" {
			return &KeyNotFoundError{key}
		}

		return set(rest, value, v.FieldByName(name), tag)
	case reflect.String:
		if len(keys) == 0 {
			if v.CanSet() {
				v.SetString(value)
			}

			return nil
		}

		return ErrInterfaceTerminated
	}

	return nil
}
