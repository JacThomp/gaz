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

	z.Warning = d

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

	z.Error = d

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

	z.Debug = d

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

	z.Info = d

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

	z.Notice = d

	return z, nil
}
