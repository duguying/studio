# sentry

[![Build Status](https://travis-ci.org/gin-contrib/sentry.svg?branch=master)](https://travis-ci.org/gin-contrib/sentry)
[![Go Report Card](https://goreportcard.com/badge/github.com/gin-contrib/sentry)](https://goreportcard.com/report/github.com/gin-contrib/sentry)
[![GoDoc](https://godoc.org/github.com/gin-contrib/sentry?status.svg)](https://godoc.org/github.com/gin-contrib/sentry)
[![Join the chat at https://gitter.im/gin-gonic/gin](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/gin-gonic/gin)

Middleware to integrate with [sentry](https://getsentry.com/) crash reporting.  Middleware version of `raven.RecoveryHandler()`.

## Example

See the [example](example/main.go)

[embedmd]:# (example/main.go go)
```go
package main

import (
	"github.com/getsentry/raven-go"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/gin"
)

func init() {
	raven.SetDSN("https://<key>:<secret>@app.getsentry.com/<project>")
}

func main() {
	r := gin.Default()
	r.Use(sentry.Recovery(raven.DefaultClient, false))
	// only send crash reporting
	// r.Use(sentry.Recovery(raven.DefaultClient, true))
	r.Run(":8080")
}
```
