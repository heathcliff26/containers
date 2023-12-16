package promremote

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/heathcliff26/containers/apps/speedtest-exporter/pkg/collector"
	"github.com/heathcliff26/containers/apps/speedtest-exporter/pkg/speedtest"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/prometheus/prompb"
	"github.com/stretchr/testify/assert"
)

var mockSpeedtestResult = speedtest.NewSpeedtestResult(0.5, 15, 876.53, 12.34, 950.3079, "Foo Corp.", "127.0.0.1")

func NewMockRegistry() *prometheus.Registry {
	s := &speedtest.MockSpeedtest{Result: mockSpeedtestResult}
	c, _ := collector.NewCollector(0, s)
	reg := prometheus.NewRegistry()
	reg.MustRegister(c)
	return reg
}

func TestNewWriteClient(t *testing.T) {
	tMatrix := []struct {
		Name, Endpoint, Instance, Job string
		Registry                      *prometheus.Registry
		Error                         string
	}{
		{"MissingEndpoint", "", "testinstance", "testjob", prometheus.NewRegistry(), "promremote.ErrMissingEndpoint"},
		{"MissingInstance", "test-endpoint", "", "testjob", prometheus.NewRegistry(), "promremote.ErrMissingInstance"},
		{"MissingJob", "test-endpoint", "testinstance", "", prometheus.NewRegistry(), "promremote.ErrMissingJob"},
		{"MissingRegistry", "test-endpoint", "testinstance", "testjob", nil, "promremote.ErrMissingRegistry"},
	}

	for _, tCase := range tMatrix {
		t.Run(tCase.Name, func(t *testing.T) {
			c, err := NewWriteClient(tCase.Endpoint, tCase.Instance, tCase.Job, tCase.Registry)

			assert := assert.New(t)

			assert.Nil(c)
			if !assert.Error(err) {
				t.Fatal("Did not receive an error")
			}
			if !assert.Equal(tCase.Error, reflect.TypeOf(err).String()) {
				t.Fatalf("Received invalid error: %v", err)
			}
		})
	}

	t.Run("Success", func(t *testing.T) {
		c, err := NewWriteClient("test-endpoint", "testinstance", "testjob", prometheus.NewRegistry())

		assert := assert.New(t)

		assert.Nil(err)
		assert.NotEmpty(c)
	})
}

func TestClientGet(t *testing.T) {
	c, _ := NewWriteClient("test-endpoint", "test", "test", prometheus.NewRegistry())
	var cNil *Client = nil
	t.Run("Endpoint", func(t *testing.T) {
		assert := assert.New(t)

		res := c.Endpoint()
		assert.NotEmpty(res)
		res = cNil.Endpoint()
		assert.Empty(res)
	})

	t.Run("Registry", func(t *testing.T) {
		assert := assert.New(t)

		res := c.Registry()
		assert.NotEmpty(res)
		res = cNil.Registry()
		assert.Empty(res)
	})
}

func TestClientSetBasicAuth(t *testing.T) {
	c, _ := NewWriteClient("test-endpoint", "test", "test", prometheus.NewRegistry())

	tMatrix := []struct {
		Username, Password string
		Error              error
	}{
		{"testuser", "password", nil},
		{"testuser", "", ErrMissingAuthCredentials{}},
		{"", "password", ErrMissingAuthCredentials{}},
		{"", "", ErrMissingAuthCredentials{}},
	}

	assert := assert.New(t)

	for _, tCase := range tMatrix {
		err := c.SetBasicAuth(tCase.Username, tCase.Password)
		if tCase.Error == nil {
			assert.Nil(err)
			assert.Equal(tCase.Username, c.username)
			assert.Equal(tCase.Password, c.password)
		} else {
			assert.Equal(tCase.Error, err)
		}
	}
}

func TestPost(t *testing.T) {
	tMatrix := []struct {
		Name   string
		Client *Client
		Auth   bool
	}{
		{
			Name:   "WithoutAuth",
			Client: &Client{},
			Auth:   false,
		},
		{
			Name: "WithAuth",
			Client: &Client{
				username: "testuser",
				password: "password",
			},
			Auth: true,
		},
	}

	for _, tCase := range tMatrix {
		t.Run(tCase.Name, func(t *testing.T) {
			assert := assert.New(t)

			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				assert.Equal(http.MethodPost, req.Method)

				assert.Equal("snappy", req.Header.Get("Content-Encoding"))
				assert.Equal("application/x-protobuf", req.Header.Get("Content-Type"))
				assert.Equal("0.1.0", req.Header.Get("X-Prometheus-Remote-Read-Version"))
				if tCase.Auth {
					auth := req.Header.Get("Authorization")
					expectedAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(tCase.Client.username+":"+tCase.Client.password))
					assert.Equal(expectedAuth, auth)
				}

				_, _ = rw.Write([]byte("Success"))
			}))
			defer server.Close()

			tCase.Client.endpoint = server.URL
			err := tCase.Client.post([]prompb.TimeSeries{})
			assert.Nil(err)
		})
	}
}

func TestCollect(t *testing.T) {
	c, _ := NewWriteClient("testendpoint", "test", "test", NewMockRegistry())

	assert := assert.New(t)

	ts, err := c.collect()

	assert.Nil(err)
	assert.NotEmpty(ts)
}
