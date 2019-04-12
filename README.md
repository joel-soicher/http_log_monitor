# HTTP Log Monitor

Implementation of an HTTP log monitoring console program

## Install and run

### Using docker:

```sh
$ docker build -t http_log_monitor .
$ docker run http_log_monitor
```

### Manual build:
Ensure Go is installed. Tested successfully with go 1.7 and go 1.11
```sh
$ export GOPATH=<the root directory of the project>
$ cd src/http_log_monitor/
$ go get -u github.com/hpcloud/tail
$ go build
$ ./http_log_monitor
```
Note: without arguments or an existing /tmp/access.log file, a panic will occur

### Command line arguments:
- file : the path to the log file (default: /tmp/access.log)
- tick : the number of seconds between two displays (default: 10s)
- alert : the number of ticks representing the period in which an alert can be triggered (default :12 -> 120s with a 10s tick)
- maxreq : the number of requests per second above which an alert will be triggered (default: 10)

example:
```sh
$ ./http_log_monitor --file="/home/user/myfolder/myfile.log" --tick=2 --alert=5 --maxreq=100
```

## Improvements
 All of these cannot be done due to lack of time...
- Add more tests!!!
- Better dependency management (dep, glide, modules, ...)
- Use a Text-UI based library for a better user experience
- I have not tested the library used for tailing the file, I don't know if it is efficient enough
- The alerting system is based on a number of ticks. It should be based on real time.
- Add a Displayer struct
- Add more checkers
- Add more tests!!!
