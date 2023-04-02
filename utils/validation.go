package utils

import (
	"plusone/backend/types"
	"reflect"
)

func UserFormValidation(data types.UserForm) bool {
	value := reflect.ValueOf(data)
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		switch field.Kind() {
		case reflect.String:
			if field.String() == "" {
				return false
			}
		}
	}
	return true
}

func EventFormValidation(data types.EventForm) bool {
	value := reflect.ValueOf(data)
	for i := 0; i < value.NumField(); i++ {
		if i != 2 && i != 3 {
			field := value.Field(i)
			switch field.Kind() {
			case reflect.String:
				if field.String() == "" {
					return false
				}
			}
		}
	}
	return true
}
