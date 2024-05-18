package interfaces

import (
	"time"

	"github.com/bondzai/test/utils"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var (
	org         = utils.GetEnv("GO_INFLUX_ORG", "")
	user_bucket = utils.GetEnv("GO_INFLUX_BUCKET_USER", "")
)

// InfluxClientInterface defines methods required for an InfluxDB client.
type InfluxClientInterface interface {
	SetDataToInflux(data *User)
}

// InfluxClient implements the InfluxClientInterface.
type InfluxClient struct {
	client influxdb2.Client
}

// NewInfluxClient creates a new InfluxClient instance.
func NewInfluxClient(host, token string) *InfluxClient {
	return &InfluxClient{
		client: influxdb2.NewClient(host, token),
	}
}

// SetDataToInflux writes data to InfluxDB.
func (c *InfluxClient) SetDataToInflux(data *User) {
	writeAPI := c.client.WriteAPI(org, user_bucket)

	point := influxdb2.NewPointWithMeasurement("measurement").
		AddField("usage_count", data.TotalUsers).
		AddTag("other", "other").
		SetTime(time.Now())

	writeAPI.WritePoint(point)
	writeAPI.Flush()
}
