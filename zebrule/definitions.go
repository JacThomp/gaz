package zebrule

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/firehose"
)

type fire interface {
	PutRecord(input *firehose.PutRecordInput) (*firehose.PutRecordOutput, error)
}

type config interface {
	Copy(...*aws.Config) *aws.Config
}

//Aluminum tells the zebrule what to do
type Aluminum struct {
	Type *string `required:"true"`
	Data *string `required:"true"`
}

//Destination is where the stuff gets sent
type Destination struct {
	Type     string  `required:"true"`
	Target   string  `required:"true"`
	firehose fire    `required:"false"`
	id       *string `required:"false"`
}

type endpoint struct {
	Fatal   *Destination `required:"true"`
	Warning *Destination `required:"true"`
	Error   *Destination `required:"true"`
	Debug   *Destination `required:"true"`
	Info    *Destination `required:"true"`
	Notice  *Destination `required:"true"`
}

type zebrule struct {
	Config    *config   `required:"true"`
	Endpoints *endpoint `required:"true"`
}
