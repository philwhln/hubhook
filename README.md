# HubHook

Handle webhook requests from DockerHub and reply to callback 

## Run locally (no Docker)

```
bash build/build
./app
```

## Test locally

From a separate terminal window send a request to the running Go process `app`.

```
curl -X POST http://0.0.0.0:8080 -d @example_payload.json
```

## Build for Docker

Build with Docker

```
docker build build
```

## Build with Docker for deployment

```
docker build deploy
```

