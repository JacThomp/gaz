package zebrule

import (
	"errors"
)

//SetWarning creates a new hose
func (z *Zebrule) SetWarning(id string, c config) (*Zebrule, error) {
	if id == "" {
		return z, errors.New("Hose ID not supplied, hose not generated")
	}

	if c == nil {
		return z, errors.New("Config File not provided, hose not generated")
	}

	d := NewDestination(id, c)

	z.Endpoints.Warning = d

	return z, nil
}

//SetError creates a new hose
func (z *Zebrule) SetError(id string, c config) (*Zebrule, error) {
	if id == "" {
		return z, errors.New("Hose ID not supplied, hose not generated")
	}

	if c == nil {
		return z, errors.New("Config File not provided, hose not generated")
	}

	d := NewDestination(id, c)

	z.Endpoints.Error = d

	return z, nil
}

//SetDebug creates a new hose
func (z *Zebrule) SetDebug(id string, c config) (*Zebrule, error) {
	if id == "" {
		return z, errors.New("Hose ID not supplied, hose not generated")
	}

	if c == nil {
		return z, errors.New("Config File not provided, hose not generated")
	}

	d := NewDestination(id, c)

	z.Endpoints.Debug = d

	return z, nil
}

//SetInfo creates a new hose
func (z *Zebrule) SetInfo(id string, c config) (*Zebrule, error) {
	if id == "" {
		return z, errors.New("Hose ID not supplied, hose not generated")
	}

	if c == nil {
		return z, errors.New("Config File not provided, hose not generated")
	}

	d := NewDestination(id, c)

	z.Endpoints.Info = d

	return z, nil
}

//SetNotice creates a new hose
func (z *Zebrule) SetNotice(id string, c config) (*Zebrule, error) {
	if id == "" {
		return z, errors.New("Hose ID not supplied, hose not generated")
	}

	if c == nil {
		return z, errors.New("Config File not provided, hose not generated")
	}

	d := NewDestination(id, c)

	z.Endpoints.Notice = d

	return z, nil
}
