# logr

## Run Instructions
1. Deps are committed, but if there are issues with them `glide install`
2. `make` to build examples, then cd into each folder to run them
3. `make test` to run all tests

#### NOTE: the `cli-service` usage example is kind of a broker/pub-sub setup
1. `./cli-service --mode=serve` in one terminal tab to run the broker/log server
2. `./cli-service --mode=write --priority=[1|2|3] --msg="This is a log message"` in another terminal tab to write to logging system
3. `./cli-service --mode=read` in another terminal tab to read logs from the system
4. Be sure to watch the server tab for fun colorful system event logs