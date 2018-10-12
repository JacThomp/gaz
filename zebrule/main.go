package zebrule

import (
	"errors"
)

//NewDestination returns a blank destination to use
func NewDestination(typeof, target, id string) *Destination {
	return &Destination{
		Type:     &typeof,
		Target:   &target,
		firehose: nil,
		ID:       &id,
	}
}

//NewZebrule returns a zebrule to use
func NewZebrule(config config, fatal, erro, warning *Destination) (*Zebrule, error) {

	if *(fatal.Target) == "" && *(warning.Target) == "" && *(erro.Target) == "" {
		return nil, errors.New("No endpoints given")
	}

	if config == nil {
		return nil, errors.New("No streaming config object given")
	}

	err := errors.New("")

	if *(fatal.Target) != "" {
		fatal, err = fatal.generateDestination(config)
		if err != nil {
			return nil, err
		}
	}
	if *(warning.Target) != "" {
		warning, err = warning.generateDestination(config)
		if err != nil {
			return nil, err
		}
	}
	if *(erro.Target) != "" {
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
		Endpoints: &ep,
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
