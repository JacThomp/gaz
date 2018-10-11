package zebrule

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/service/firehose"
)

func (z *zebrule) Feed(report Aluminum) error {

	switch *(report.Type) {
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
		return fmt.Errorf("%s is an unknown report type, please check docs", *(report.Type))
	}
}

func (d Destination) feed(report Aluminum) error {
	if d.firehose == nil {
		return errors.New("Warning logging has not been assigned to this zebrule")
	}

	switch *(d.Type) {
	case "AWS":

		j, err := json.Marshal(report.Data)
		if err != nil {
			return err
		}
		hose := d.firehose.(*firehose.Firehose)
		_, err = hose.PutRecord(&firehose.PutRecordInput{
			Record:             &firehose.Record{Data: j},
			DeliveryStreamName: d.id,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
