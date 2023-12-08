package collector

import (
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

var fakeSpeedtestResultSuccess = NewSpeedtestResult(0.5, 15, 876.53, 12.34, 950.3079, "Foo Corp.", "127.0.0.1")

type FakeSpeedtest struct {
	callback func()
	fail     bool
}

func (s *FakeSpeedtest) Speedtest() *SpeedtestResult {
	if s.callback != nil {
		s.callback()
	}
	if s.fail {
		return NewFailedSpeedtestResult()
	}
	return fakeSpeedtestResultSuccess
}

func TestNewCollector(t *testing.T) {
	s := &FakeSpeedtest{}
	expectedCollector := &Collector{
		cacheTime: 5,
		instance:  "foo",
		speedtest: s,
	}

	actualCollector, err := NewCollector(5, "foo", s)
	if err != nil {
		t.Fatalf("Could not create new Collector: %v", err)
	}

	assert := assert.New(t)

	assert.Equal(expectedCollector, actualCollector)

	_, err = NewCollector(0, "", nil)
	assert.Equal(NoSpeedtestError{}, err)
}

func TestSetNextSpeedtestTime(t *testing.T) {
	c, err := NewCollector(5, "foo", &FakeSpeedtest{})
	if err != nil {
		t.Fatalf("Could not create new Collector: %v", err)
	}

	c.nextSpeedtest = time.Now()
	c.setNextSpeedtestTime()

	assert.GreaterOrEqual(t, c.nextSpeedtest, time.Now())
}

func TestFirstSpeedtestRun(t *testing.T) {
	c, err := NewCollector(5, "foo", &FakeSpeedtest{})
	if err != nil {
		t.Fatalf("Could not create new Collector: %v", err)
	}
	result := c.getSpeedtestResult()

	assert := assert.New(t)

	assert.Equal(fakeSpeedtestResultSuccess, result)
	assert.Equal(result, c.lastResult)
	assert.NotEmpty(c.nextSpeedtest)
}

func TestResultFromCache(t *testing.T) {
	speedtestRan := false
	s := &FakeSpeedtest{
		callback: func() {
			speedtestRan = true
		},
	}

	c, err := NewCollector(5, "foo", s)
	if err != nil {
		t.Fatalf("Could not create new Collector: %v", err)
	}
	c.lastResult = NewFailedSpeedtestResult()
	c.nextSpeedtest = time.Now().Add(time.Hour)

	result := c.getSpeedtestResult()

	assert := assert.New(t)

	assert.NotNil(result)
	assert.Equal(NewFailedSpeedtestResult(), result)
	if speedtestRan {
		t.Error("Speedtest has been called")
	}
}

func TestRunSpeedtestWhenCacheEmpty(t *testing.T) {
	speedtestRan := false
	s := &FakeSpeedtest{
		callback: func() {
			speedtestRan = true
		},
	}

	c, err := NewCollector(5, "foo", s)
	if err != nil {
		t.Fatalf("Could not create new Collector: %v", err)
	}
	c.lastResult = nil
	c.nextSpeedtest = time.Now().Add(time.Hour)

	result := c.getSpeedtestResult()

	assert := assert.New(t)

	assert.NotEmpty(result)
	assert.Equal(fakeSpeedtestResultSuccess, result)
	assert.Equal(result, c.lastResult)
	if !speedtestRan {
		t.Error("Speedtest was not called")
	}
}

func TestRunSpeedtestWhenCacheExpired(t *testing.T) {
	speedtestRan := false
	s := &FakeSpeedtest{
		callback: func() {
			speedtestRan = true
		},
	}

	c, err := NewCollector(5, "foo", s)
	if err != nil {
		t.Fatalf("Could not create new Collector: %v", err)
	}
	c.lastResult = NewFailedSpeedtestResult()
	c.nextSpeedtest = time.Now().Add(time.Hour * -1)

	result := c.getSpeedtestResult()

	assert := assert.New(t)

	assert.NotEmpty(result)
	assert.NotEqual(NewFailedSpeedtestResult(), result)
	assert.Equal(fakeSpeedtestResultSuccess, result)
	assert.Equal(result, c.lastResult)
	if !speedtestRan {
		t.Error("Speedtest was not called")
	}
}

func TestSpeedtestIsNotRunConcurrently(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	i := 0
	sleeping := make(chan bool, 1)
	s := &FakeSpeedtest{
		callback: func() {
			i++
			sleeping <- true
			time.Sleep(10 * time.Second)
		},
	}

	c, err := NewCollector(5, "foo", s)
	if err != nil {
		t.Fatalf("Could not create new Collector: %v", err)
	}

	assert := assert.New(t)

	var result1, result2 *SpeedtestResult
	go func() {
		result1 = c.getSpeedtestResult()
	}()
	<-sleeping
	assert.Nil(c.lastResult)
	result2 = c.getSpeedtestResult()

	assert.NotNil(result1)
	assert.Equal(result1, result2)
	assert.Equal(result1, c.lastResult)
	assert.Equal(1, i)
}

func TestCollect(t *testing.T) {
	s := &FakeSpeedtest{}
	c, err := NewCollector(0, "foo", s)
	if err != nil {
		t.Fatalf("Could not create new Collector: %v", err)
	}

	t.Run("Success", func(t *testing.T) {
		ch := make(chan prometheus.Metric, 1)
		go c.Collect(ch)

		actualLabelValues := []string{c.instance, fakeSpeedtestResultSuccess.clientIp, fakeSpeedtestResultSuccess.clientIsp}

		actualMetric := <-ch
		expectedMetric := prometheus.MustNewConstMetric(jitterLatencyDesc, prometheus.GaugeValue, fakeSpeedtestResultSuccess.JitterLatency(), actualLabelValues...)
		assert.Equal(t, expectedMetric, actualMetric)

		actualMetric = <-ch
		expectedMetric = prometheus.MustNewConstMetric(pingDesc, prometheus.GaugeValue, fakeSpeedtestResultSuccess.Ping(), actualLabelValues...)
		assert.Equal(t, expectedMetric, actualMetric)

		actualMetric = <-ch
		expectedMetric = prometheus.MustNewConstMetric(downloadSpeedDesc, prometheus.GaugeValue, fakeSpeedtestResultSuccess.DownloadSpeed(), actualLabelValues...)
		assert.Equal(t, expectedMetric, actualMetric)

		actualMetric = <-ch
		expectedMetric = prometheus.MustNewConstMetric(uploadSpeedDesc, prometheus.GaugeValue, fakeSpeedtestResultSuccess.UploadSpeed(), actualLabelValues...)
		assert.Equal(t, expectedMetric, actualMetric)

		actualMetric = <-ch
		expectedMetric = prometheus.MustNewConstMetric(dataUsedDesc, prometheus.GaugeValue, fakeSpeedtestResultSuccess.DataUsed(), actualLabelValues...)
		assert.Equal(t, expectedMetric, actualMetric)

		actualMetric = <-ch
		expectedMetric = prometheus.MustNewConstMetric(upDesc, prometheus.GaugeValue, 1, c.instance)
		assert.Equal(t, expectedMetric, actualMetric)
	})

	t.Run("Failure", func(t *testing.T) {
		ch := make(chan prometheus.Metric, 1)
		s.fail = true
		go c.Collect(ch)
		actualMetric := <-ch
		expectedMetric := prometheus.MustNewConstMetric(upDesc, prometheus.GaugeValue, 0, c.instance)
		assert.Equal(t, expectedMetric, actualMetric)
	})
}
