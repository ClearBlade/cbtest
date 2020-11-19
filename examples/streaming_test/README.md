Streaming test showcases a test against a ClearBlade system with stream services.

### Structure

- `extra/`: The system that gets imported by the test.

- `message.go`: Contains the definition and generation of messages that we send.

- `streaming_test.go`: The actual test.

### How to run

Run with default flags:

```
go test -v ./examples/streaming_test/
```

Run with 50 devices:

```
go test -v ./examples/streaming_test/ -parallel 50 -args -devices 50
```

Run with 50 devices, and 3 instances of the streaming service:

```
go test -v ./examples/streaming_test/ -parallel 50 -args -devices 50 -instances 3
```

Run for 20 seconds with 50 devices, and 3 instances of the streaming service:

```
go test -v ./examples/streaming_test/ -parallel 50 -args -duration 20s -devices 50 -instances 3
```
