# Config

Config is a transport protocol on top of protobuffer for defining
extensible configuration protocols.  The protocol defines a request
and a response structure.

To define a given configuration protocol for a device there needs to
be a defined set of commands and defined responses as part of the
application firmware and the backend services.  It is advisable to
document this protocol in a place that can be referred to by all
components that depend on the given application protocol.

## Build

In order to build this package you will need to have `protobuf`
installed on your system.  If you do you also need `protoc-gen-go`
which can be obtained by running

    make deps

If you have both installed you can build and test with 

    make
	
This generates the protobuffer code and runs a simple test.  This is
meant to act as a tool for verifying changes to the protobuffer
definition.

**We intend to check in ready generated protobuffer code to this
project so that other projects can depend on these files without
having to install the tooling to generate them**.

## Documentation

### Request and Response

The Request and Response messages are defined as follows:

    message Request {
        uint32 id = 1;
        uint32 command = 2;
        repeated Value values = 3;
    }

    message Response {
        uint32 id = 1;
        uint32 command = 2;
        uint32 sequence = 3;
        uint32 responseCode = 4;
        repeated Value values = 5;
    }


#### `id` (Request, Response)

The `id` fields are used to correlate `Response` messages to their
`Request` counterparts so that when a device responds to a request it
will copy the `id` from the `Request` into the `Response`.

If the `id` is 0 this means that the `Response` was in response to any
particular `Request`.

A request can elicit multiple responses.  For instance if the response
does not fit in a single message or if a request turns on logging and
you want to want to be able to correlate it to the `Request` that
turned on logging.

#### `command` (Request, Response)

The values of `command` are used to identify which command the
`Request` represents.  This will be specific to each application.

#### `values` (Request, Response)

In order to represent values that are sent with a `Request` or
returned by a `Response` we use the `Value` message which is just a
structure that can hold int32, int64, double, string or raw
bytes.  

    message Value {
        uint32 id = 1;
        int32  int32Val = 2;
        int64  int64Val = 3;
        double doubleVal = 4;
        string stringVal = 5;
        bytes bytesVal = 6;
    }


Both `Request` and `Response` define the `values` field as
`repeatable` which means you can have zero or more values in a given
`Request` or `Response`.

Since all fields are optional in proto3 it is up to the application to
define which fields should be set for which commands.

*We have avoided using `oneof` or similar tricks to maintain
simplicity.  We could save a few bytes of RAM and packet size, but at
the cost of added complexity*.

#### `responseCode` (Response)

This field contains an application specific response code in the
`Request` message.

#### `sequence` (Response)

The sequence field is used to enumerate responses that require more
than one message.  It is recommended that the first response indicates
how many messages are expected if this is known.  For instance by
adding a 

Sequences should start with 0 (zero).


## Recommendations

### Document carefully

As mentioned above, it is recommended to document the application
protocols that make use of this transport in a common place so that
both the firmware and the backend software refers to the same
definition of the protocol.

You should document all commands complete with which values will be
set in both the request and response.  Make sure that for each
`command` you document which `responseCode` values developers should
expect back.

### Commands that should exist

There should at minimum be commands to ask the device for its current
status (uptime, battery, id of any MCUs or peripherials attached,
firmware version etc) and commands to do simple household tasks like
reboot the device etc.

### Versioning

Beware that this is a transport protocol and that the application
protocol you put on top of it should be versioned separately from the
transport protocol.

## Ideas

### Reserved commands

It might be a good idea to have some reserved commands that are
defined by default.  For instance commands to inquire for status,
firmware, serial numbers, ask the device to reboot etc.
