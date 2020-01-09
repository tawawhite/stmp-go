# stmp-go

A lightweight real-time bidirectional framework for Golang.

**Features**

- Fast
    - Supports `Protocol Buffers`, `MsgPack`
    - Extremely tidy packet format, only 1 byte header
    - Supports `KCP`
    - Supports connection level `gzip` compression
- Easy to use
    - Supports `WebSockets` and highly optimized for it
    - Supports `text` format packet for `WebSockets`
    - Supports `JSON`
    - Supports compose different listeners in one server
    - Most concepts are the same as `gRPC`
    - `.proto` based service definition
- Secure
    - Supports `TLS`

*Supported transport layer protocols*

- `KCP`
- `TCP`
- `WebSockets`

## Install

```bash
go get -u github.com/acrazing/stmp-go
```

## Examples

You can get full examples at [examples](./examples) directory, you can run the examples with source code in root
directory, the examples list:

- [quick start](./examples/quick_start): The quick start example

    ```bash
    # run server
    make run-quick-start-server
  
    # run client
    make run-quick-start-client
    ```

- [room](./examples/room): A complex chat room service, include golang server, golang terminal client, and typescript browser client.

    ```bash
    # run server
    make run-room-server
  
    # run client
    make run-room-client
    ```

## Documentation

- `stmp`: <https://godoc.org/github.com/acrazing/stmp-go/stmp>
- `protoc-gen-stmp`: <https://godoc.org/github.com/acrazing/stmp-go/protoc-gen-stmp>

## LICENSE

```text
The MIT License (MIT)

Copyright (c) 2019-2020 acrazing

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```