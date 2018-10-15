package zebrule

import (
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/firehose"
)

type fire interface {
	PutRecord(input *firehose.PutRecordInput) (*firehose.PutRecordOutput, error)
}

type config interface {
	Copy(...*aws.Config) *aws.Config
}

type data interface {
	String() []byte
}

//Aluminum tells the zebrule what to do
type Aluminum struct {
	Type string `required:"true"`
	Data data   `required:"true"`
}

//Destination is where the stuff gets sent
type Destination struct {
	Type      string      `required:"true"`
	Seperator string      `required:"true"`
	firehose  fire        `required:"false"`
	ID        string      `required:"false"`
	mute      *sync.Mutex `required:"false"`
}

type endpoint struct {
	Fatal   Destination `required:"true"`
	Warning Destination `required:"true"`
	Error   Destination `required:"true"`
	Debug   Destination `required:"true"`
	Info    Destination `required:"true"`
	Notice  Destination `required:"true"`
}

//Zebrule is a poor struct, brought into this world only to eat aluminum and stream logs
type Zebrule struct {
	Config    *config  `required:"true"`
	Endpoints endpoint `required:"true"`
}
