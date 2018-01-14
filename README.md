# logr

## So far
`go run example/main.go` just writes a bunch of statically defined lines, and then reads them out again

### TODO:
- [ ] Make Thread-safe
- [ ] Have channels for reader and writers (so we can have one reader and multiple writers) (maybe local network support???)
- [ ] Support multiple kinds of writers (with a LogWriter interface (define `Write` method))
- [ ] More run examples
- [ ] Tests! (thread-safety, fill up buffer)
- [ ] Change empty system read behaviour from just returning an error, to waiting for logs to be written
- [ ] Handle logs that aren't just strings. Some sort of composite obj.
