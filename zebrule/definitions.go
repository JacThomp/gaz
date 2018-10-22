package zebrule

import (
	"reflect"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/firehose"
)

var ch = make(chan error)
var wg sync.WaitGroup

type fire interface {
	PutRecord(input *firehose.PutRecordInput) (*firehose.PutRecordOutput, error)
}

type config interface {
	Copy(...*aws.Config) *aws.Config
}

//Aluminum is the data interface, in case you don't want to use the defaul
type Aluminum interface {
	Bytes() []byte
}

//Aluminum tells the zebrule what to do
type Data struct {
	Type     string   `required:"true"`
	Aluminum Aluminum `required:"true"`
}

//Destination is where the stuff gets sent
type Destination struct {
	Type   reflect.Type `required:"true"`
	ID     string       `required:"true"`
	Config *config      `required:"true"`
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
	Endpoints endpoint `required:"true"`
}
