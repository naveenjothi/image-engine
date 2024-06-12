package jsonutil

import (
	"fmt"
	"reflect"
)

// MergeStructs merges two instances of the same struct type and returns the merged struct
func MergeStructs(dest, src interface{}) error {
    destVal := reflect.ValueOf(dest).Elem()
    srcVal := reflect.ValueOf(src).Elem()

    if destVal.Kind() != reflect.Struct || srcVal.Kind() != reflect.Struct {
        return fmt.Errorf("both dest and src must be structs")
    }

    for i := 0; i < srcVal.NumField(); i++ {
        srcField := srcVal.Field(i)
        srcFieldType := srcVal.Type().Field(i)
        destField := destVal.FieldByName(srcFieldType.Name)

        if !destField.IsValid() || !destField.CanSet() {
            continue
        }

        // Handle pointer fields in src
        if srcField.Kind() == reflect.Ptr {
            if !srcField.IsNil() {
                if destField.Type() == srcField.Elem().Type() {
                    destField.Set(srcField.Elem())
                } else {
                    return fmt.Errorf("cannot assign value of type %s to field %s of type %s",
                        srcField.Elem().Type(), srcFieldType.Name, destField.Type())
                }
            }
        } else {
            zeroValue := reflect.Zero(srcField.Type()).Interface()
            if !reflect.DeepEqual(srcField.Interface(), zeroValue) {
                if destField.Type() == srcField.Type() {
                    destField.Set(srcField)
                } else {
                    return fmt.Errorf("cannot assign value of type %s to field %s of type %s",
                        srcField.Type(), srcFieldType.Name, destField.Type())
                }
            }
        }
    }
    return nil
}

