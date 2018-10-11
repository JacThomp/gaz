package zebrule

import (
	"errors"
)

//NewDestination returns a blank destination to use
func NewDestination(typeof, target string) *Destination {
	return &Destination{
		Type:     typeof,
		Target:   target,
		firehose: nil,
		id:       nil,
	}
}

//NewZebrule returns a zebrule to use
func NewZebrule(config config, fatal, warning, erro *Destination) (*zebrule, error) {

	if fatal.Target == "" && warning.Target == "" && erro.Target == "" {
		return &zebrule{}, errors.New("No endpoints given")
	}

	if config == nil {
		return &zebrule{}, errors.New("No streaming config object given")
	}

	err := errors.New("")

	if fatal.Target != "" {
		fatal, err = fatal.generateDestination(config)
		if err != nil {
			return &zebrule{}, err
		}
	}
	if warning.Target != "" {
		warning, err = warning.generateDestination(config)
		if err != nil {
			return &zebrule{}, err
		}
	}
	if erro.Target != "" {
		erro, err = erro.generateDestination(config)
		if err != nil {
			return &zebrule{}, err
		}
	}

	debug := &Destination{Type: "", Target: ""}
	info := &Destination{Type: "", Target: ""}
	notice := &Destination{Type: "", Target: ""}

	ep := endpoint{
		Fatal:   fatal,
		Warning: warning,
		Error:   erro,
		Debug:   debug,
		Info:    info,
		Notice:  notice,
	}

	zeb := &zebrule{
		Config:    &config,
		Endpoints: &ep,
	}

	return zeb, nil
}

//Must wrapper to ensure New doesn't return an error
func Must(zeb *zebrule, err error) *zebrule {
	if err != nil {
		panic(err)
	}
	return zeb
}
