package jsonutil

import (
	"fmt"
	"log"
	"reflect"
)

// MergeStructs merges two instances of the same struct type and returns the merged struct
func MergeStructs(dest, src interface{}) error {
    // Check if both dest and src are pointers to the same type
    destVal := reflect.ValueOf(dest).Elem()
    srcVal := reflect.ValueOf(src).Elem()

    if destVal.Kind() != reflect.Ptr || srcVal.Kind() != reflect.Ptr {
        return fmt.Errorf("both parameters must be pointers")
    }

    // // Merge src into dest
    for i := 0; i < srcVal.NumField(); i++ {
        srcField := srcVal.Field(i)
        srcFieldType := srcVal.Type().Field(i)
        destField := destVal.FieldByName(srcFieldType.Name)
        log.Printf("%v",destField.IsValid() && destField.CanSet())

        if destField.IsValid() && destField.CanSet() {
            // If srcField is a pointer, we need to check if it is nil
            if srcField.Kind() == reflect.Ptr {
                if !srcField.IsNil() {
                    destField.Set(srcField.Elem())
                }
            } else {
                zeroValue := reflect.Zero(srcField.Type()).Interface()
                if !reflect.DeepEqual(srcField.Interface(), zeroValue) {
                    destField.Set(srcField)
                }
            }
        }
    }

    return nil
}

