# WIP
# deadmanssnitch-api

API written in Golang - https://deadmanssnitch.com/

# Start API Server

To initialize the HTTPServer API, it is required to use `DEADMANSSNITCH_APIKEY` environment variable

```bash
$ export DEADMANSSNITCH_APIKEY="XXXXPPPTTTOOOOODDDD111123"

$ go run main.go
YYYY/MM/dd 00:00:00 Starting server http://0.0.0.0:8000...
```

You can create new APIKEY access on the Deadmanssnitch Website, follow the example:
- https://deadmanssnitch.com/cases/{ID}/keys


## POST

You can create new Snitch send a POST request to `/api/snitch` 

```bash
curl http://localhost:8000/api/snitch -X POST -H '{"Content-Type: application/json' -d '{"name": "hola", "interval": "daily", "tags": ["prod", "critical"]}'
```

The response will have HTTP Status Code `(201 Created)`
