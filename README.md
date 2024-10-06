# aria2

```bash
# Start aria2
docker run --rm --name aria2 --net=host ghcr.io/chneau/aria2
# build for linux/amd64 linux/arm64 linux/386 linux/arm/v6 linux/arm/v7 linux/ppc64le linux/s390x
```

Open `localhost:3000` with `username`: "" and `password`: "look on the console"  
Then go to `AriaNg Settings` > `RPC (localhost:3000)` > `Aria2 RPC Secret Token`: "look on the console"

# env vars

- `PORT` The port the web interface will be served on, default to `3000`
- `EXTERNAL_PORT` The port the aria2 server will be available on, replace the default value on the web interface, default to the value of `PORT`
- `SECRET` The secret token for the aria2 server, default to a `random 16 characters` string
- `ARIA_PORT` The port the aria2c server will run on, default to `6800`
- `ARIA_DIR` The directory where the downloads will be saved, default to `/data`
- `USERNAME` The username for the web interface, default to `""` (empty)
- `PASSWORD` The password for the web interface, default to the value of `SECRET`

# behavior

- `Basic auth` for the web interface is on if whether `USERNAME` or `PASSWORD` are set
- Setting the Aria2c RPC Secret Token in the web interface is set if there is `basic auth`
