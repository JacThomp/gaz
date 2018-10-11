package zebrule

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/gobuffalo/uuid"
)

/*
type endpoint struct {
	Fatal   *Destination `required:"true"`
	Warning *Destination `required:"true"`
	Error   *Destination `required:"true"`
	Debug   *Destination `required:"true"`
	Info    *Destination `required:"true"`
	Notice  *Destination `required:"true"`
}
type zebrule struct {
	Stream    *streamer `required:"true"`
	Hose      *string   `required:"true"`
	Endpoints *endpoint `required:"true"`
}
type Destination struct {
	Type     string   `required:"true"`
	Target   string   `required:"true"`
	firehose *request `required:"false"`
	shutoff  *cancel  `required:"false"`
}
*/

func (d *Destination) generateDestination(c config) (*Destination, error) {

	id := uuid.Must(uuid.NewV4()).String()
	s := session.Must(session.NewSession())

	switch d.Type {
	case "AWS":
		awsConf := c.(*aws.Config)
		awsConf.Endpoint = &d.Target
		d.firehose = firehose.New(s, awsConf)
		d.id = &id
	default:
		return nil, fmt.Errorf("%s is not a valid type", d.Type)
	}

	return d, nil
}

func (z *zebrule) SetFatal()   {}
func (z *zebrule) SetWarning() {}
func (z *zebrule) SetError()   {}
func (z *zebrule) SetDebug()   {}
func (z *zebrule) SetInfo()    {}
func (z *zebrule) SetNotice()  {}
