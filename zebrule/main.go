package zebrule

import (
	"errors"
	"reflect"
)

//NewDestination returns a blank destination to use
func NewDestination(id string, conf config) Destination {
	return Destination{
		ID:     id,
		Type:   reflect.TypeOf(conf),
		Config: &conf,
	}
}

//NewZebrule returns a zebrule to use
func NewZebrule(fatal, erro, warning Destination) (*Zebrule, error) {

	if fatal.ID == "" && warning.ID == "" && erro.ID == "" {
		return nil, errors.New("No Endpoints given")
	}

	debug := NewDestination("", nil)
	info := NewDestination("", nil)
	notice := NewDestination("", nil)

	zeb := &Zebrule{
		Fatal:   fatal,
		Warning: warning,
		Error:   erro,
		Debug:   debug,
		Info:    info,
		Notice:  notice,
	}

	return zeb, nil
}

//Must wrapper to ensure New doesn't return an error
func Must(zeb *Zebrule, err error) *Zebrule {
	if err != nil {
		panic(err)
	}
	return zeb
}
