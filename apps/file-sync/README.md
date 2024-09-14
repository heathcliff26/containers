# file-sync

Example `/config/config.yaml`:
```
---
target_dir: /path/to/remote/target/dir
hosts:
  - user@example.com
known_hosts: |
  example.com ssh-ed25519 asdfghjkl
post_hook: "echo posthook"
ssh_key: /path/to/private_ssh.key
```

Example kubernetes pod:
```
---
apiVersion: v1
kind: Pod
metadata:
  name: file-sync
  labels:
    app: file-sync
spec:
  containers:
  - name: file-sync
    image: ghcr.io/heathcliff26/file-sync:latest
    volumeMounts:
      - name: config
        mountPath: /config
      - name: ssh-key
        mountPath: /ssh-key
      - name: files
        mountPath: /files
  volumes:
    - name: config
      configMap:
        name: file-sync-config
    - name: ssh-key
      secret:
        secretName: file-sync-key
        defaultMode: 0600
    - name: files
      projected:
        defaultMode: 0644
        sources:
          - secret:
              name: example-cert
              items:
                - key: tls-combined.pem
                  path: example.pem
```
