# STMP Specification

**VERSION: 1.0**

**Note**: `text` is only used for `WebSockets`, it contains message length already, so we do not need `<LENGTH>` field.

1. Client handshake

   - binary

     ```text
     STMP<VERSION MAJOR><VERSION MINOR><LENGTH>
     <KEY>:<VALUE>
     ...
     ```

   - text

     _reuse WebSockets handshake request_

     - All the query will be treated as headers
     - `stmp-version`: with format `<major>.<minor>`

   - `KEY`

     - `Accept`: the accepted `Content-Type` types, see [RFC7231#5.3.2](https://tools.ietf.org/html/rfc7231#section-5.3.2).
     - `Accept-Encoding`: the accepted `Encoding` types, see [RFC7231#5.3.4](https://tools.ietf.org/html/rfc7231#section-5.3.2).
     - `Accept-Packet-Format`: the accepted `Packet-Format` types, see `Server handshake` section.

   - `VERSION MAJOR`, `VERSION MINOR`: each use one byte as its value result is `uint16`

2. Server handshake

   - binary

     ```text
     STMP<HANDSHAKE STATUS><LENGTH>
     <KEY>:<VALUE>
     ...

     [MESSAGE]
     ```

   - text

     ```text
     STMP<HANDSHAKE STATUS>
     <KEY>:<VALUE>
     ...

     [MESSAGE]
     ```

     - _WebSockets handshake response is deprecated, because client cannot access it_

   - `HANDSHAKE STATUS`: _If `HANDSHAKE STATUS` is not `OK`, both client and server should close the connection._

   - `KEY`

     - `Content-Type`: the content marshall protocol, negotiated by `Accept`, this field must be set, if server cannot parse.
       all of the types specified by `Accept`, server will response a `0x4` error code, and close connection.
     - `Encoding`: the exchange full connection encoding, negotiated by `Accept-Encoding`, if not set, means the data is plain packet.
     - `Packet-Format`: the protocol serialize format, supports `text` and `binary` for `websocket`, else supports `binary` only.

   - `MESSAGE`: a `UTF-8` encoded string

3. Exchange message header

   - binary

     ```text
     | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
     | F |     K     | H |     R     |
     | I |     I     | E |     S     |
     | N |     N     | A |     V     |
     |   |     D     | D |           |
     ```

     - `FIN`: only for `Request`, `Notify`, `Response` and `Following`, if fin, means
       that the message is end, else has more `Following` message with same `MESSAGE ID`,
       the `PAYLOAD` of the message with same `MESSAGE ID` will be merged.
     - `KIND`: message kind
       - `0x0`: Ping message
       - `0x1`: Request message
       - `0x2`: Notify message
       - `0x3`: Response message
       - `0x4`: Following message
       - `0x5`: Close message
     - `HEAD`: the message is head only or with payload, if is `0`, the `PAYLOAD` and `PAYLOAD LENGTH` field will be omitted.
     - `RSV`: must be `0` currently

   - text

     text message is used for `WebSockets` only, so, it only contains
     `KIND` field with one byte:

     - `P`: Ping message
     - `Q`: Request message
     - `N`: Notify message
     - `S`: Response message
     - `C`: Close message

4. Request message

   - binary

     ```text
     <HEAD><MESSAGE ID><ACTION>[PAYLOAD LENGTH][PAYLOAD]
     ```

   - text

     ```text
     Q<MESSAGE ID>:<ACTION>
     [PAYLOAD]
     ```

   - `MESSAGE ID`: a `int16le` represents the message id, used for connecting `Response` and `Request`, `Following` and not `FIN` `Notify` messages.
      if the message from server, it should be a negative value, else it should be a positive value.
   - `ACTION`: a `varint` represents method
   - `PAYLOAD LENGTH`: a `varint` represents the `PAYLOAD` length
   - `PAYLOAD`: any content will be explained by codec

5. Notify message

   - binary

     ```text
     <HEAD>[MESSAGE ID]<ACTION>[PAYLOAD LENGTH][PAYLOAD]
     ```

     - If `FIN`, `MESSAGE ID` should be omitted

   - text

     ```text
     N<ACTION>
     [PAYLOAD]
     ```

6. Response message

   - binary

     ```text
     <HEAD><RESPONSE STATUS><MESSAGE ID>[PAYLOAD LENGTH][PAYLOAD]
     ```

   - text

     ```text
     S<MESSAGE ID>:<RESPONSE STATUS>
     [PAYLOAD]
     ```

   - `RESPONSE STATUS`: see `Status code`

7. Following message

   - binary

     ```text
     <HEAD><MESSAGE ID>[PAYLOAD LENGTH][PAYLOAD]
     ```

8. Close message

   - binary

     ```text
     <HEAD><CLOSE STATUS>[PAYLOAD LENGTH][PAYLOAD]
     ```

**Status code**

- `0x00`: Ok, 200

_Network error_

- `0x01`: Network error
- `0x02`: Protocol error
- `0x03`: Unsupported protocol version
- `0x04`: Unsupported `Content-Type`
- `0x05`: Unsupported `Format`

_Client side error_

- `0x20`: Bad request
- `0x21`: Unauthorized
- `0x22`: Not found
- `0x23`: Request timeout
- `0x24`: Request entity too large
- `0x25`: Too many requests
- `0x26`: Client closed

_Server side error_

- `0x40`: Internal server error
- `0x41`: Server shutdown
