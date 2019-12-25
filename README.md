# log

Simple log for Go.

# Import
```go
import "github.com/eviot/log"
```

# Examples

### log.Debug(err interface{})
```go
res, err := checkSome()
if log.Debug(err) {
    return
}
```

### log.Debug(err interface{}, i ...interface{})
```go
res, err := checkID(id)
if log.Debug(err, id, userEmail) {
    return
}
```

### log.Info(i ...interface{})
```go
func logMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Info("incoming new request from", r.UserAgent())
        next.ServeHTTP(w, r)
    })
}
```

### log.Infof(format string, i ...interface{})
```go
func logMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Infof("incoming new request from %s", r.UserAgent())
        next.ServeHTTP(w, r)
    })
}
```

### log.Pretty(i interface{})
```go
type T struct {
    Foo string
    Bar int
    Map map[string]bool
}
t := T{
    Foo: "foo",
    Bar: 789,
    Map: map[string]bool{"flag": true, "monday": false},
}
log.Pretty(t)
```
Output:
```
[INF] 10.11.09 23:00:00.000 [prog.go:17] {
	"Foo": "foo",
	"Bar": 789,
	"Map": {
		"flag": true,
		"monday": false
	}
}
```
