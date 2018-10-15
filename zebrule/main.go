package zebrule

import (
	"errors"
	"sync"
)

//NewDestination returns a blank destination to use
func NewDestination(typeof, id, seperator string) Destination {
	return Destination{
		Type:      typeof,
		Seperator: seperator,
		firehose:  nil,
		ID:        id,
		mute:      &sync.Mutex{},
	}
}

//NewZebrule returns a zebrule to use
func NewZebrule(config config, fatal, erro, warning Destination) (*Zebrule, error) {

	null := Destination{}

	if fatal == null {
		fatal = NewDestination("", "", "")
	}
	if erro == null {
		erro = NewDestination("", "", "")
	}
	if warning == null {
		warning = NewDestination("", "", "")
	}

	if fatal.ID == "" && warning.ID == "" && erro.ID == "" {
		return nil, errors.New("No Endpoints given")
	}

	if config == nil {
		return nil, errors.New("No streaming config object given")
	}

	err := errors.New("")

	if fatal.ID != "" {
		fatal, err = fatal.generateDestination(config)
		if err != nil {
			return nil, err
		}
	}
	if warning.ID != "" {
		warning, err = warning.generateDestination(config)
		if err != nil {
			return nil, err
		}
	}
	if erro.ID != "" {
		erro, err = erro.generateDestination(config)
		if err != nil {
			return nil, err
		}
	}

	debug := NewDestination("", "", "")
	info := NewDestination("", "", "")
	notice := NewDestination("", "", "")

	ep := endpoint{
		Fatal:   fatal,
		Warning: warning,
		Error:   erro,
		Debug:   debug,
		Info:    info,
		Notice:  notice,
	}

	zeb := &Zebrule{
		Config:    &config,
		Endpoints: ep,
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
