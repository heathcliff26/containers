# keepalived

Using keepalived from inside a container. By default it will use the configuration from `/etc/keepalived/keepalived.conf`.

## Example command

```
sudo podman run -v ./keepalived.conf:/etc/keepalived/keepalived.conf -d --net host --cap-add=NET_ADMIN --cap-add=NET_RAW ghcr.io/heathcliff26/keepalived:latest
```
1. The container needs to be run as root to enable it to edit the interfaces
2. It needs to be run with `--net host` to be able to see the host interfaces instead of a pseudo-interface inside the container.
3. It needs the capabilities `NET_ADMIN` and `NET_RAW` to edit the network interface and receive updates from other instances of keepalived on the same network.

## keepalived scripts

The image has the `keepalived_script` user created for script execution.

Since the process will be executed inside a container, it is not possible to check for anything happening on the host out of the box.
As such the best way would be to check for open ports using `/usr/bin/nc -z 127.0.0.1 <port>`, since the container is connected to the host network.
Add `-u` if you want to check for udp port instead of the default tcp.

Note: For some reason keepalived needs the script to be an actual bash script, so make sure to wrap the command into one.
