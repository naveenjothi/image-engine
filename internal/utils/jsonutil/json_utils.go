package jsonutil

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// MergeStructs merges two instances of the same struct type and returns the merged struct
func MergeStructs(dest, src interface{}) error {
    // Check if both dest and src are pointers to the same type
    destVal := reflect.ValueOf(dest)
    srcVal := reflect.ValueOf(src)

    if destVal.Kind() != reflect.Ptr || srcVal.Kind() != reflect.Ptr {
        return fmt.Errorf("both parameters must be pointers")
    }

    destElem := destVal.Elem()
    srcElem := srcVal.Elem()

    if destElem.Type() != srcElem.Type() {
        return fmt.Errorf("both parameters must be pointers to the same type")
    }

    // Merge src into dest
    for i := 0; i < destElem.NumField(); i++ {
        destField := destElem.Field(i)
        srcField := srcElem.Field(i)

        // Check if srcField is zero value, if not, set it to destField
        if !srcField.IsZero() {
            destField.Set(srcField)
        }
    }

    return nil
}

// MergeJSONStructs unmarshals two JSON byte slices into the same struct type,
// merges them, and returns the resulting struct
func MergeJSONStructs(json1, json2 []byte, modelType interface{}) (interface{}, error) {
    // Create instances of the modelType
    model1 := reflect.New(reflect.TypeOf(modelType)).Interface()
    model2 := reflect.New(reflect.TypeOf(modelType)).Interface()

    // Unmarshal JSON into model1 and model2
    if err := json.Unmarshal(json1, model1); err != nil {
        return nil, fmt.Errorf("error unmarshaling json1: %v", err)
    }

    if err := json.Unmarshal(json2, model2); err != nil {
        return nil, fmt.Errorf("error unmarshaling json2: %v", err)
    }

    // Merge model2 into model1
    if err := MergeStructs(model1, model2); err != nil {
        return nil, err
    }

    return model1, nil
}
