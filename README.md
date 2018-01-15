# logr

## So far
`go run example/main.go` just writes a bunch of statically defined lines, and then reads them out again

### TODO:
- [x] Make Thread-safe
- [ ] Connects writers to reader somehow
- [x] Support multiple kinds of writers (with a LogWriter interface (define `Write` method))
- [ ] More run examples
- [ ] Tests! (thread-safety, fill up buffer)
- [x] Decorate log messages with Such as date, time, priority, thread name, class name, etc..
