package telemetrydest

import "gopkg.in/segmentio/analytics-go.v3"

type TelemetryDestination interface {
	Track(client analytics.Client,userId string, event string, properties map[string]interface{}) error
	Identify(client analytics.Client, userId string, traits map[string]interface{}) error
	Page(client analytics.Client, userId string, name string, properties map[string]interface{}) error
}
