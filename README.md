# Logrus

Implements logrus Log interface

```
package main

import (
	logr "github.com/go-log/logrus"
	"github.com/sirupsen/logrus"
)

func main {
	l := logr.WithFields(logrus.Fields{
		"foo": "bar",
	})

	l.Log("a log ling")
}
```
