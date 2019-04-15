# HTTP Log Monitor

Implementation of an HTTP log monitoring console program.

note: There's a misinterpretation in the subject for the default file that could be either /tmp/access.log or /var/log/access.log.
I have chosen /tmp/access.log as default.

## Install and run

### Using docker:

```sh
$ docker build -t http_log_monitor .
$ docker run http_log_monitor
```

Command line arguments can be passed through docker in the run command. See below.


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

examples:

```sh
$ docker run http_log_monitor --file="/home/user/myfolder/myfile.log" --tick=2 --alert=5 --maxreq=100
```

```sh
$ ./http_log_monitor --file="/home/user/myfolder/myfile.log" --tick=2 --alert=5 --maxreq=100
```

## Tests
For manual testing, I have a simple application for writing file to the log file that I change according to my needs.
It is located in src/tester and you can launch it simply by running, assuming you are in that folder:
```sh
$ go run main.go
```

Go tests can be launched by this command, assuming you are in the src/http_log_monitor/ folder:
```sh
$ go test
```

## Improvements
 All of these cannot be done due to lack of time...
- Add more tests!!!
- Test with real data (a lot of sample files can be found on the web)
- Make different packages.
- Better error management, logs the potential errors
- Better dependency management (dep, glide, modules, ...)
- Use a Text-UI based library for a better user experience
- I have not tested the library used for tailing the file, I don't know if it is efficient enough
- The alerting system is based on a number of ticks. It should be based on real time.
- Decouple computation and display, because the DisplayString methods do more than display by now.
- The displayer should be abstract, in order to redirect to different outputs (console, file, network, ...)
- Add many more checkers, add a way to externally customize them and selecting the ones you want to activate.
- Benchmark to identify potential bottlenecks, and optimize these parts
- Read other log formats and improve the parsing
- Add more tests!!!
