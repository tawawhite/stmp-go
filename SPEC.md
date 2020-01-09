# STMP Specification

**VERSION: 1.0**

- **Note**: `text` is only used for `WebSockets`, it contains message length already, so we do not need `<LENGTH>` field.
- **Note**: if transport could split packet, the binary packets should not contains `XXX LENGTH` fields, it should be replaced by `\n`

1. Client handshake

   - binary

     ```text
     STMP<MAJOR|MINOR><HEADER LENGTH><HEADER><PAYLOAD LENGTH><PAYLOAD>
     ```

   - text
   
     ```text
     <STMP><MAJOR><MINOR>
     H<HEADER>
     M<MESSAGE>
     ```

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
     STMP<HANDSHAKE STATUS><HEADER LENGTH><HEADER><PAYLOAD LENGTH><PAYLOAD>
     ```

   - text

     ```text
     STMP<HANDSHAKE STATUS>
     H<HEADER>
     M<MESSAGE>
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
     | F |     K     | W | W | S | R |
     | I |     I     | P | H | A | S |
     | N |     N     |   |   |   | V |
     |   |     D     |   |   |   |   |
     ```

     - `FIN`: must be `1` currently
     - `KIND`: message kind
       - `0x0`: Ping message
       - `0x1`: Pong message
       - `0x2`: Request message
       - `0x3`: Notify message
       - `0x4`: Response message
       - `0x5`: Close message
     - `WP`: with payload or not, if `1`, means the packet has `PAYLOAD LENGTH` and `PAYLOAD` part
     - `WH`: with header or not, if `1`, means the packet has `HEADER LENGTH` and `HEADER` part
     - `SA`: action is string or not, if `1`, means the packet action is `METHOD LENGTH` and `METHOD` rather than `ACTION`

   - text

     ```text
     <KIND>
     ```

     - `KIND`:
         - `I`: Ping message
         - `O`: Pong message
         - `Q`: Request message
         - `N`: Notify message
         - `S`: Response message
         - `C`: Close message

4. Request/Notify message

   - binary

     ```text
     <HEAD>[MESSAGE ID](<ACTION>|<METHOD LENGTH><METHOD>)([HEADER LENGTH][HEADER])([PAYLOAD LENGTH][PAYLOAD])
     ```

   - text

     ```text
     <HEAD>[MESSAGE ID]:A<ACTION>|M<METHOD>
     H[HEADER]
     P[PAYLOAD]
     ```

   - `MESSAGE ID`: a `uint16(LE)` represents the message id, for request message only
   - `ACTION`: a `uvarint` represents method
   - `METHOD LENGTH`: a `uvarint` represents the `METHOD` length
   - `METHOD`: the string method
   - `HEADER LENGTH`: a `uvarint` represents the `HEADER` length
   - `HEADER`: the header pairs
   - `PAYLOAD LENGTH`: a `uvarint` represents the `PAYLOAD` length
   - `PAYLOAD`: any content will be explained by codec

5. Response message

   - binary

     ```text
     <HEAD><MESSAGE ID><RESPONSE STATUS>([HEADER LENGTH][HEADER])[PAYLOAD LENGTH][PAYLOAD]
     ```

   - text

     ```text
     <HEAD><MESSAGE ID>:<RESPONSE STATUS>
     H[HEADER]
     P[PAYLOAD]
     ```

   - `RESPONSE STATUS`: see `Status code`

6. Close message

   - binary

     ```text
     <HEAD><CLOSE STATUS>[PAYLOAD LENGTH][PAYLOAD]
     ```
   
   - text
   
     ```text
     <HEAD><CLOSE STATUS>
     [PAYLOAD]
     ```

   - `CLOSE STATUS`: see `Status code`

**Status code**
