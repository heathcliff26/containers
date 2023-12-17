# cloudflare-dyndns

Implements the API from [Fritz!Box DynDNS Script for Cloudflare](https://github.com/1rfsNet/Fritz-Box-Cloudflare-DynDNS), but can also be used as a standalone client.

Additionally to consuming less resources and being a smaller image, it also implements POST in addition to GET requests, meaning no longer does the token need to be included in the url.

The client package can also be used as a golang API, should you want to build your application with included cloudflare dyndns capabilities.

## Table of Contents

- [cloudflare-dyndns](#cloudflare-dyndns)
  - [Table of Contents](#table-of-contents)
  - [Container Images](#container-images)
    - [Image location](#image-location)
    - [Tags](#tags)
  - [Usage](#usage)
  - [API (Server Mode)](#api-server-mode)
    - [Examples](#examples)

## Container Images

### Image location

| Container Registry                                                                                     | Image                                      |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------ |
| [Github Container](https://github.com/users/heathcliff26/packages/container/package/cloudflare-dyndns) | `ghcr.io/heathcliff26/cloudflare-dyndns`   |
| [Docker Hub](https://hub.docker.com/repository/docker/heathcliff26/cloudflare-dyndns)                  | `docker.io/heathcliff26/cloudflare-dyndns` |

### Tags

There are different flavors of the image:

| Tag(s)           | Describtion                                                                                                                                                                              |
| ---------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **latest, slim** | Contains only the golang binary                                                                                                                                                          |
| **php**          | Deprecated: Contains the original php script from [Fritz!Box DynDNS Script for Cloudflare](https://github.com/1rfsNet/Fritz-Box-Cloudflare-DynDNS) and is based on `php:apache-bookworm` |

## Usage

The binary can be run either as a server, a standalone client or in relay mode where it will call a server.

The main use case for relay mode would be when you want to restrict your cloudflare API key to a static IP.

Output of `cloudflare-dyndns -h`
```
Usage of cloudflare-dyndns:
  -config string
        Path to config file, can be empty when running in mode server
  -env
        Used together with -config, when set will expand enviroment variables in config
  -mode string
        Set what mode to run, options are "server", "client" and "relay" (default "server")
```
An example config can be found [here](configs/example-config.yaml).

## API (Server Mode)

| Parameter        | Description                                                                    |
| ---------------- | ------------------------------------------------------------------------------ |
| token (cf_key)   | Token needed for accessing cloudflare api                                      |
| domains (domain) | The domain to update, parsed from comma (,) separated string, needs at least 1 |
| ipv4             | IPv4 Address, optional, when IPv6 set                                          |
| ipv6             | IPv6 Address, optional, when IPv4 set                                          |
| proxy            | Indicate if domain should be proxied, defaults to true                         |

### Examples
Here is an example GET request:
```
https://dyndns.example.com/?token=testtoken&domains=foo.example.net,bar.example.org,example.net&ipv4=100.100.100.100&ipv6=fd00::dead&proxy=true
```
or alternatively in the format [Fritz!Box DynDNS Script for Cloudflare](https://github.com/1rfsNet/Fritz-Box-Cloudflare-DynDNS) from :
```
http://example.org/?cf_key=testtoken&domain=foo.example.net&ipv4=100.100.100.100&ipv6=fd00::dead&proxy=true
```
When using POST the format is:
```
{
  "token": "",
  "domains": [
    "foo.example.org",
    "bar.example.net"
  ],
  "ipv4": "100.100.100.100",
  "ipv6": "fd00::dead",
  "proxy": true
}
```
