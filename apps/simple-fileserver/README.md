# simple-fileserver

This container contains a simple HTTP Fileserver that serves static files and does nothing else.

It is implemented with golang's `http.FileServer`.

**Important Notice**

*Running the the fileserver outside of the container is not recommended, as it is possible to get the server to serve any file using dots.*
*Please ensure that there is no sensitive content inside your webroot.*

*Using simple-fileserver is done at your own risk.*

## Usage

### CLI Args
```
$ podman run ghcr.io/heathcliff26/simple-fileserver:latest -h
Usage of simple-fileserver:
  -cert string
        SFILESERVER_CERT: SSL certificate to use, needs key as well. Default is no ssl.
  -key string
        SFILESERVER_KEY: SSL private key to use, needs cert as well. Default is no ssl.
  -log
        SFILESERVER_LOG: Enable logging requests
  -no-index
        SFILESERVER_NO_INDEX: Do not serve an index for directories, return index.html or 404 instead
  -port int
        SFILESERVER_PORT: Specify port for the fileserver to listen on (default 8080)
  -webroot string
        SFILESERVER_WEBROOT: Required, root directory to serve files from
```

### Using the image
```
$ podman run -d -p 8080:8080 -v /path/to/content:/webroot ghcr.io/heathcliff26/simple-fileserver:latest
```

### Image location

| Container Registry                                                                                      | Image                                       |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------- |
| [Github Container](https://github.com/users/heathcliff26/packages/container/package/simple-fileserver) | `ghcr.io/heathcliff26/simple-fileserver`   |
| [Docker Hub](https://hub.docker.com/repository/docker/heathcliff26/simple-fileserver)                  | `docker.io/heathcliff26/simple-fileserver` |
