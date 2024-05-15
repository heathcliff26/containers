# tang

This image runs a tang server for Network Bound Disk Encryption (NBDE).

Usage:
```
podman run -d -p 8080:8080 -v tang-keys:/var/db/tang --name tang ghcr.io/heathcliff26/tang:latest
```
