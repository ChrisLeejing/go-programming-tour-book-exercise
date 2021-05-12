package validator

import (
	val "github.com/go-playground/validator/v10"
	"reflect"
	"sync"
)

type CustomValidator struct {
	Once     sync.Once
	Validate *val.Validate
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{}
}

// implements:
// type StructValidator interface {}
func (v *CustomValidator) ValidateStruct(obj interface{}) error {
	if KindOfData(obj) == reflect.Struct {
		if err := v.Validate.Struct(obj); err != nil {
			return err
		}
	}

	return nil
}

// Engine() interface{}
func (v *CustomValidator) Engine() interface{} {
	v.lazyInit()
	return v.Validate
}

func (v *CustomValidator) lazyInit() {
	// Once is an object that will perform exactly one action.
	v.Once.Do(func() {
		// validate := val.New() // bug: nil pointer dereference
		v.Validate = val.New() // debug.
		// SetTagName allows for changing of the default tag name of 'validate'
		// validate.SetTagName("binding") // bug: nil pointer dereference
		v.Validate.SetTagName("binding")
	})
}

func KindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	kind := value.Kind()
	if kind == reflect.Ptr {
		kind = value.Elem().Kind()
	}
	return kind
}
