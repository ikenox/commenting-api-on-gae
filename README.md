# Comment API on GAE

API server of commenting service which is run on Google App Engine Standard Environment

## About this project

This project adopts Clean Architecture and Domain Driven Design, and keeps these principles as much as possible.

Because of Clean Architecture, application logic and domain logic are independent with detail of infrastructure.  
Followings are not appeared in core of application.

- Various packages related to Google App Engine
- Technological details of web application (e.g. context.Context)

TBW:
- Application architecture overview
- Domain model overview

## Requirements

- go 1.8
- google-cloud-sdk
    - goapp
    - dev_appserver.py
- dep

## Setup

```shell
# /path/to/comment-api-on-gae/src/commenting
$ GOPATH=/path/to/comment-api-on-gae dep ensure
```

## Run

```shell
# /path/to/comment-api-on-gae/src/commenting
$ GOPATH=/path/to/comment-api-on-gae dev_appserver.py app --enable_watching_go_path --log_level=debug --datastore_path=.storage
```

## Deploy

```shell
# /path/to/comment-api-on-gae/src/commenting
$ GOPATH=/path/to/comment-api-on-gae goapp deploy app
```
