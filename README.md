# A very simple logger for OCI Functions

Start the syslogger
```
$ go run .
```

Register the logger with OCI FN application
```
$ fn update app <app_name> --syslog-url tcp://<address>:<port>
```