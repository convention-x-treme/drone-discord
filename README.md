# Discord Plugin for Drone CI

[![Discord Logo](images/discord-logo.png "Discord")](https://discord.com)

Drone plugin for sending message to Discord channel using Webhook.

Credits goes to [appleboy](https://github.com/appleboy) for his previous releases.

Webhooks are a low-effort way to post messages to channels in Discord. They do not require a bot user or authentication to use. See more [api document information](https://discord.com/developers/docs/resources/webhook).

Sending discord message using a binary, docker or [Drone CI](http://docs.drone.io/).

## Features

- [x] Send Multiple Messages
- [x] Send Multiple Files

## Usage

There are three ways to send notification.

### Usage from binary

```bash
drone-discord \
  --webhook-id xxxx \
  --webhook-token xxxx \
  --message "Test Message"
```

### Usage from docker

```bash
docker run --rm \
  -e WEBHOOK_ID=xxxxxxx \
  -e WEBHOOK_TOKEN=xxxxxxx \
  -e WAIT=false \
  -e TTS=false \
  -e USERNAME=test \
  -e AVATAR_URL=http://example.com/xxxx.png \
  -e MESSAGE=test \
  convention-x-treme/drone-discord
```

### Usage from drone ci

#### Send Notification

Execute from the working directory:

```sh
docker run --rm \
  -e WEBHOOK_ID=xxxxxxx \
  -e WEBHOOK_TOKEN=xxxxxxx \
  -e WAIT=false \
  -e TTS=false \
  -e USERNAME=test \
  -e AVATAR_URL=http://example.com/xxxx.png \
  -e MESSAGE=test \
  -e DRONE_REPO_OWNER=convention-x-treme \
  -e DRONE_REPO_NAME=drone-discord \
  -e DRONE_COMMIT_SHA=e5e82b5eb3737205c25955dcc3dcacc839b7be52 \
  -e DRONE_COMMIT_BRANCH=master \
  -e DRONE_COMMIT_AUTHOR=author \
  -e DRONE_COMMIT_AUTHOR_EMAIL=author@example.com \
  -e DRONE_COMMIT_MESSAGE=Test_Your_Commit \
  -e DRONE_BUILD_NUMBER=1 \
  -e DRONE_BUILD_STATUS=success \
  -e DRONE_BUILD_LINK=http://github.com/convention-x-treme/drone-discord \
  -e DRONE_JOB_STARTED=1477550550 \
  -e DRONE_JOB_FINISHED=1477550750 \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  conventionxtreme/drone-discord
```

You can get more [information](DOCS.md) about how to use this plugin in drone.

## Build or Download a binary

The pre-compiled binaries can be downloaded from [release page](https://github.com/convention-x-treme/drone-discord/releases). Support the following OS type.

- Windows amd64/386
- Linux arm/amd64/386
- Darwin amd64/386

With `Go` installed

```sh
go get -u -v github.com/convention-x-treme/drone-discord
```

or build the binary with the following command:

```sh
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

go test -cover ./...

go build -v -a -tags netgo -o release/linux/amd64/drone-discord
```
