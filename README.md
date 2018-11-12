# Logrus

A [logrus][] implementation of [`log.Logger`][Logger].
For more on the interface and other implementations, see [the log README][log].

```go
package main

import (
	logr "github.com/go-log/logrus"
	"github.com/sirupsen/logrus"
)

func main {
	l := logr.WithFields(logrus.Fields{
		"foo": "bar",
	})

	l.Log("a log line")
}
```

You can also use [`print.New`][print-New] to convert logrus loggers to `log.Logger`.
For an example, see [here][logrus-print].

[log]: https://github.com/go-log/log/blob/master/README.md
[Logger]: https://godoc.org/github.com/go-log/log#Logger
[logrus]: https://github.com/sirupsen/logrus
[logrus-print]: https://github.com/go-log/log/blob/master/README.md#example
[print-New]: https://godoc.org/github.com/go-log/log/print#New
