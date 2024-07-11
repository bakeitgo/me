### gRPC

* Can use protocol buffer to define grpc service and its message format.

* Client -> Server communication

* Language agnostic

* Can be used with JSON as message format, but by default uses proto buffers.

#### Proto Buffers

* Text file with .proto ext, defines message format and services.

* Generating from this .proto file data structures (message format) for specified language using protoc compiler
```proto
// The greeter service definition.
service Greeter {
  // Sends a greeting
// The greeter service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}
```

#### Available grpc services

* Single request / Single Response
```proto
rpc SayHello(HelloRequest) returns (HelloResponse);
```
* Single Request / Stream responses
    * Client reads until server sends all messages in stream (response is ordered)
```proto
rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse);
```
* Stream Requests / Single response
    * Client write multiple messages, after writing it waits until server reads it and return response (request messages are ordered)
```proto
rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse);
```
* Stream Requests / Stream responses
    * Bidirectional communication, both of the streams operate independently, messages within streams are ordered. Server can wait for all messages until respond, or make "ping-pong", its application specific. ORder is preserved
```proto
rpc BidiHello(stream HelloRequest) returns (stream HelloResponse);
```

#### RPC life cycle

* Unary RPC (Single request / Single response)
    1. Client calls stub method, and the RPC gets verified about client metadata, possible deadline and method called.
    2. Server responds with its metadata immediately or waits for the client response (application-specific)
    3. Client sends request, Server do work then respond with metadata AND IF response is successfull, status code and optional message is returned, and optional trailing metadata
    4. If response status is OK, client gets the response which completed on the client side

* Server streaming
    * Same as unary RPC except server returns stream of messages
* Client streaming
    * Same as unary RPC except client sends stream of messages
* Bidirectional streaming
    * Therefore cleint and server are streams, application specific how the communication will act

#### Deadlines / Timeout

* Client can specify deadline, everything is independent. Server can query for time left for the rpc. If its exceeded the DEADLINE\_EXCEEDED err is raised


#### Cancellation

* Client can cancel RPC at any time

* **CHANGES ARE NOT ROLLED BACK BEFORE CANCELLATION**

#### Metadata

* Information about RPC call, in form of **key-value** pairs.

* Can be user defined, key with *grpc-* prefix is reserved by gRPC itself

* Keys are strings, Values can be binary when nended with -bin, and strings

#### Channel


* gRPC channel provides connnection to the gRPC server on specified host and port.

* Client can specify channel args to modify gRPC default behavior

* Connection establishes when creating client stub.

