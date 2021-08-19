package pagination

import "reflect"

// Use this if your web framework's query parser does not support complex
// struct like PaginationQuerySort. For example, for the fiber framework, use
// Bind(&PaginationQuerySort{}, fiber.Ctx.QueryParser).
func Bind(object interface{}, bindFunc func(interface{}) error) error {
	fields, values := getFieldsAndValues(reflect.ValueOf(object), reflect.TypeOf(object))
	n := reflect.New(reflect.StructOf(fields))
	if err := bindFunc(n.Interface()); err != nil {
		return err
	}
	for i := range fields {
		values[i].Set(n.Elem().Field(i))
	}
	return nil
}

func getFieldsAndValues(rv reflect.Value, rt reflect.Type) (fields []reflect.StructField, values []reflect.Value) {
	if rt.Kind() == reflect.Ptr {
		rv = rv.Elem()
		rt = rt.Elem()
	}
	if rt.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		value := rv.Field(i)
		if field.Anonymous {
			f, v := getFieldsAndValues(value, field.Type)
			fields = append(fields, f...)
			values = append(values, v...)
		} else if field.Tag.Get("query") != "" {
			fields = append(fields, field)
			values = append(values, value)
		}
	}
	return
}
