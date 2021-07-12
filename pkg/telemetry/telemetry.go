package telemetry

import (
	"fmt"
	"github.com/AkshayBhansali18/telemetry-dependency-go/pkg/constants"
	"gopkg.in/segmentio/analytics-go.v3"
)

type SegmentClient struct {
	client analytics.Client
}

func InitClient() (SegmentClient, error) {

	segmentClient := SegmentClient{}
	segmentClient.client = analytics.New(constants.WRITE_KEY)
	return segmentClient, nil

}

func (SegmentClient SegmentClient) CloseClient() error {

	err := SegmentClient.client.Close()
	if err != nil {
		return fmt.Errorf("An Error Occurred while trying to close client %v: ", err)
	}
	return nil
}

func (segmentClient SegmentClient) IdentifyUser(userId string, traits map[string]interface{}) error {

	newClient := segmentClient.client
	if len(userId) == 0 {
		return fmt.Errorf("length of userId must be greater than 0")
	}

	err := newClient.Enqueue(analytics.Identify{
		UserId: userId,
		Traits: traits,
	})

	if err != nil {
		return fmt.Errorf("An error occurred: %v", err)
	}
	return nil
}

func (segmentClient SegmentClient) PageTelemetry(userId string, pageName string, properties map[string]interface{}) error {

	if len(userId) == 0 {
		return fmt.Errorf("length of \"userId\" must be greater than 0")
	}

	if len(pageName) == 0 {
		return fmt.Errorf("length of \"pageName\" must be greater than 0")
	}

	err := segmentClient.client.Enqueue(analytics.Page{
		UserId:     userId,
		Name:       pageName,
		Properties: properties,
	})

	if err != nil {
		return fmt.Errorf("An error occurred: %v", err)
	}
	return nil
}

func (segmentClient SegmentClient) TrackTelemetry(userId string, event string, properties map[string]interface{}) error {

	if len(userId) == 0 {
		return fmt.Errorf("length of \"userId\" must be greater than 0")
	}
	if len(event) == 0 {
		return fmt.Errorf("length of \"event\" must be greater than 0")
	}
	fmt.Println(segmentClient.client)
	err := segmentClient.client.Enqueue(analytics.Track{
		UserId:     userId,
		Event:      event,
		Properties: properties,
	})

	if err != nil {
		return fmt.Errorf("An error occurred: %v", err)
	}
	return nil
}
