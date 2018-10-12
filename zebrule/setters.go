package zebrule

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
)

func (d *Destination) generateDestination(c config) (*Destination, error) {

	if c == nil {
		return d, errors.New("Bad config")
	}

	switch *(d.Type) {
	case "AWS":
		s := session.Must(session.NewSession())
		d.firehose = firehose.New(s, c.(*aws.Config))
	default:
		return d, fmt.Errorf("%s is not a valid type", *(d.Type))
	}

	return d, nil
}

//SetFatal creates a new hose
func (z *Zebrule) SetFatal(typeof, target, id string) (*Zebrule, error) {
	if id == "" {
		return z, errors.New("Hose ID not supplied, hose not generated")
	}
	if typeof == "" {
		return z, errors.New("Type of hose not supplied, hose not generated")
	}

	d := NewDestination(typeof, target, id)
	d, err := d.generateDestination(*(z.Config))
	if err != nil {
		return z, fmt.Errorf("Generation of hose failed: %v", err)
	}

	z.Endpoints.Fatal = d

	return z, nil
}

//SetWarning creates a new hose
func (z *Zebrule) SetWarning(typeof, target, id string) (*Zebrule, error) {
	if id == "" {
		return z, errors.New("Hose ID not supplied, hose not generated")
	}
	if typeof == "" {
		return z, errors.New("Type of hose not supplied, hose not generated")
	}

	d := NewDestination(typeof, target, id)
	d, err := d.generateDestination(*(z.Config))
	if err != nil {
		return z, fmt.Errorf("Generation of hose failed: %v", err)
	}

	z.Endpoints.Warning = d

	return z, nil
}

//SetError creates a new hose
func (z *Zebrule) SetError(typeof, target, id string) (*Zebrule, error) {
	if id == "" {
		return z, errors.New("Hose ID not supplied, hose not generated")
	}
	if typeof == "" {
		return z, errors.New("Type of hose not supplied, hose not generated")
	}

	d := NewDestination(typeof, target, id)
	d, err := d.generateDestination(*(z.Config))
	if err != nil {
		return z, fmt.Errorf("Generation of hose failed: %v", err)
	}

	z.Endpoints.Error = d

	return z, nil
}

//SetDebug creates a new hose
func (z *Zebrule) SetDebug(typeof, target, id string) (*Zebrule, error) {
	if id == "" {
		return z, errors.New("Hose ID not supplied, hose not generated")
	}
	if typeof == "" {
		return z, errors.New("Type of hose not supplied, hose not generated")
	}

	d := NewDestination(typeof, target, id)
	d, err := d.generateDestination(*(z.Config))
	if err != nil {
		return z, fmt.Errorf("Generation of hose failed: %v", err)
	}

	z.Endpoints.Debug = d

	return z, nil
}

//SetInfo creates a new hose
func (z *Zebrule) SetInfo(typeof, target, id string) (*Zebrule, error) {
	if id == "" {
		return z, errors.New("Hose ID not supplied, hose not generated")
	}
	if typeof == "" {
		return z, errors.New("Type of hose not supplied, hose not generated")
	}

	d := NewDestination(typeof, target, id)
	d, err := d.generateDestination(*(z.Config))
	if err != nil {
		return z, fmt.Errorf("Generation of hose failed: %v", err)
	}

	z.Endpoints.Info = d

	return z, nil
}

//SetNotice creates a new hose
func (z *Zebrule) SetNotice(typeof, target, id string) (*Zebrule, error) {
	if id == "" {
		return z, errors.New("Hose ID not supplied, hose not generated")
	}
	if typeof == "" {
		return z, errors.New("Type of hose not supplied, hose not generated")
	}

	d := NewDestination(typeof, target, id)
	d, err := d.generateDestination(*(z.Config))
	if err != nil {
		return z, fmt.Errorf("Generation of hose failed: %v", err)
	}

	z.Endpoints.Notice = d

	return z, nil
}
