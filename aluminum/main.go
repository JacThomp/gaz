package aluminum

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

//Bytes byte-fies the struct yay
func (a Aluminum) Bytes() []byte {

	cvals := reflect.ValueOf(a)
	ctype := cvals.Type()

	retval := start

	for i := 0; i < cvals.NumField(); i++ {
		v := cvals.Field(i)
		t := ctype.Field(i)

		if v.Interface() == reflect.Zero(t.Type).Interface() {
			if t.Tag.Get("required") == "true" {
				retval = append(retval, []byte(t.Name)...)
				retval = append(retval, br...)
				switch t.Type {
				case reflect.TypeOf(time.Now()):
					retval = append(retval, []byte(time.Now().UTC().String())...)
				default:
					retval = append(retval, unkown...)
				}
				retval = append(retval, se...)
			}
		} else {
			retval = append(retval, []byte(t.Name)...)
			retval = append(retval, br...)

			switch v.Kind() {
			case reflect.String:
				retval = append(retval, []byte(v.String())...)
			case reflect.Bool:
				retval = append(retval, []byte(strconv.FormatBool(v.Bool()))...)
			case reflect.Struct:
				retval = append(retval, []byte(v.Interface().(time.Time).UTC().String())...)
			case reflect.Ptr:
				resp, _ := json.Marshal(v.Interface().(*http.Response))
				retval = append(retval, []byte(resp)...)
			}

			retval = append(retval, se...)
		}
	}

	return append(retval[:len(retval)-3], end...)
}
