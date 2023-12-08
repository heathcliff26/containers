# speedtest-exporter

This project is a prometheus exporter for speedtests implemented in go.
It runs automatically speedtests with the speedtest.net API and exports the result as prometheus metrics.
It supports native speedtest by using [speedtest-go](https://github.com/showwin/speedtest-go).

I created this project because i saw [Jeff Geerling's](https://github.com/geerlingguy) video about his speedtest setup.
When looking into it, he uses a [python based speedtest-exporter](https://github.com/MiguelNdeCarvalho/speedtest-exporter).
At this point i could have just used that as well, but i was bored and wanted to program something. Hence this project.

## Table of Contents

- [speedtest-exporter](#speedtest-exporter)
  - [Table of Contents](#table-of-contents)
  - [Container Images](#container-images)
    - [Image location](#image-location)
    - [Tags](#tags)
  - [Usage](#usage)
  - [Metrics](#metrics)
  - [Dashboard](#dashboard)

## Container Images

### Image location

| Container Registry                                                                                      | Image                                       |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------- |
| [Github Container](https://github.com/users/heathcliff26/packages/container/package/speedtest-exporter) | `ghcr.io/heathcliff26/speedtest-exporter`   |
| [Docker Hub](https://hub.docker.com/repository/docker/heathcliff26/speedtest-exporter)                  | `docker.io/heathcliff26/speedtest-exporter` |

### Tags

There are different flavors of the image:

| Tag(s)           | Describtion                                                                                                                 |
| ---------------- | --------------------------------------------------------------------------------------------------------------------------- |
| **latest, slim** | Contains only the speedtest-exporter binary and uses native golang implementation.                                          |
| **cli**          | Alpine based container that also contains the speedtest.net cli client binary. Uses the speedtest.net cli to run the tests. |

## Usage

Output of `speedtest-exporter -h`
```
Usage of speedtest-exporter:
  -cacheTime int
        Time in minutes to cache speedtest output (default 5)
  -instance string
        Label added to all metrics for identification, defaults to hostname
  -port int
        Port for the webserver, default 8080 (default 8080)
  -speedtest-path string
        Specify speedtest executable to use, defaults to internal implementation
  -v    Enable verbose output
```

Alternatively these values can be configured using enviroment variables:

| Flag              | Enviroment Variable    |
| ----------------- | ---------------------- |
| `-cacheTime`      | `SPEEDTEST_CACHE_TIME` |
| `-instance`       | `SPEEDTEST_INSTANCE`   |
| `-port`           | `SPEEDTEST_PORT`       |
| `-speedtest-path` | `SPEEDTEST_PATH`       |
| `-v`              | `SPEEDTEST_DEBUG`      |

## Metrics

The following metrics are exported:

| Metric                                   | Description                                |
| ---------------------------------------- | ------------------------------------------ |
| `speedtest_jitter_latency_milliseconds`  | Speedtest current Jitter in ms             |
| `speedtest_ping_latency_milliseconds`    | Speedtest current Ping in ms               |
| `speedtest_download_megabits_per_second` | Speedtest current Download Speed in Mbit/s |
| `speedtest_upload_megabits_per_second`   | Speedtest current Upload Speed in Mbit/s   |
| `speedtest_data_used_megabytes`          | Data used for speedtest in MB              |
| `speedtest_up`                           | Indicates if the speedtest was successful  |

## Dashboard

A ready made dashboard for the exporter can be imported from json. The json file can be found [here](dashboard/dashboard.json).

The dashboard is also published on grafana.com with the id [20115](https://grafana.com/grafana/dashboards/20115).

Here is a preview of the dashboard:
![](images/dashboard.png)
