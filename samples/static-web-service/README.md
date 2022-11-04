# Samples of static web service written in Go

## How to run the sample

```bash
cd samples/static-web-service

go build .

static-web-service.exe

# from another terminal

curl http://localhost:3000/?name=World

curl -H "Accept: application/json" http://localhost:3000/?name=World
```
