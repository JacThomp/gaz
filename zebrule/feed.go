package zebrule

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/firehose"
)

//Feed tells the zebrule what to stream out
func (z *Zebrule) Feed(report Data) error {

	switch report.Type {
	case "WARNING":
		return z.Endpoints.Warning.feed(report)
	case "FATAL":
		return z.Endpoints.Fatal.feed(report)
	case "ERROR":
		return z.Endpoints.Error.feed(report)
	case "DEBUG":
		return z.Endpoints.Debug.feed(report)
	case "INFO":
		return z.Endpoints.Info.feed(report)
	case "NOTICE":
		return z.Endpoints.Notice.feed(report)
	default:
		return fmt.Errorf("%s is an unknown report type, please check docs", report.Type)
	}
}

func (d Destination) feed(report Data) error {
	if d.firehose == nil {
		return fmt.Errorf("Logging has not been assigned to this zebrule")
	}

	output := report.Aluminum.Bytes()
	fmt.Println(string(output))

	switch d.Type {
	case "AWS":

		hose := d.firehose.(*firehose.Firehose)

		d.mute.Lock()
		_, err := hose.PutRecord(&firehose.PutRecordInput{
			Record:             &firehose.Record{Data: append(output, byte('\n'))},
			DeliveryStreamName: aws.String(d.ID),
		})
		d.mute.Unlock()
		if err != nil {
			return err
		}
	}

	return nil
}
