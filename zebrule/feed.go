package zebrule

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
)

//Feed tells the zebrule what to stream out
func (z *Zebrule) Feed(report Data) error {

	switch report.Type {
	case "WARNING":
		return z.Warning.feed(report)
	case "FATAL":
		return z.Fatal.feed(report)
	case "ERROR":
		return z.Error.feed(report)
	case "DEBUG":
		return z.Debug.feed(report)
	case "INFO":
		return z.Info.feed(report)
	case "NOTICE":
		return z.Notice.feed(report)
	default:
		return fmt.Errorf("%s is an unknown report type, please check docs", report.Type)
	}
}

func (d Destination) feed(report Data) error {
	if d.Config == nil {
		return errors.New("Logging has not been assigned to this zebrule")
	}

	if d.ID == "" {
		return errors.New("No ID given for the target destination")
	}

	output := report.Aluminum.Bytes()
	fmt.Println(string(output))

	switch d.Type {
	case reflect.TypeOf(&aws.Config{}):
		s := session.Must(session.NewSession())
		hose := firehose.New(s, (*d.Config).(*aws.Config))

		_, err := hose.PutRecord(&firehose.PutRecordInput{
			Record:             &firehose.Record{Data: append(output, byte('\n'))},
			DeliveryStreamName: aws.String(d.ID),
		})
		if err != nil {
			return fmt.Errorf("Error streaming to AWS: %v", err)
		}

	default:
		return fmt.Errorf("%s is not a valid type of config", d.Type)
	}

	return nil
}
