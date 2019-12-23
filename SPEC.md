# STMP Specification

1. Client handshake

    - binary
    
        ```text
        <LENGTH><VERSION MAJOR><VERSION MINOR>
        <KEY>:<VALUE>
        ```

    - text

        *reuse WebSockets handshake request*
        
        - All the query will be treated as headers
        - `stmp-version`: with format `<major>.<minor>`
    
    - headers
    
        - `Content-Type`: determine the `content-type` for full connection
        - `Encoding`: the compress mode, only allow `gzip` or empty
        - `Action-Kind`: the action kind, only `string` or `varint`
    
    - `VERSION MAJOR`, `VERSION MINOR`: each use one byte as its value
        result is `uint16`

2. Server handshake

    - binary
    
        ```text
        <HANDSHAKE STATUS>[PAYLOAD LENGTH][PAYLOAD]
        ```
    
    - text
    
        ```text
        <HANDSHAKE STATUS>[MESSAGE]
        ```
      
        - *WebSockets handshake response is deprecated, because client cannot access it*
    
    - `HANDSHAKE STATUS`, first bit is used for with `PAYLOAD` or not, if with load
        the `PAYLOAD LENGTH` and `PAYLOAD` is required

        - `0x0`: OK
        - `0x1`: Authenticate failed

3. Exchange message header

    - binary
    
        ```text
        | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
        | M |    KIND   | F |  ENC  | R |
        | A |           | I |   O   | S |
        | S |           | N |   D   | V |
        | K |           |   |  ING  |   |
        ```
      
        - `MASK`: must be `1`, used for distinguishing `text` and `binary` message
        - `KIND`: message kind
            - `0b000`: `Ping message`
            - `0b001`: `Request message`
            - `0b010`: `Notify message`
            - `0b011`: `Response message`
            - `0b100`: `Following message`
            - `0b101`: `Close message`
        - `FIN`: only for `Request`, `Notify`, `Response` and `Following`, if fin, means
            that the message is end, else has more `Following` message with same `MESSAGE ID`,
            the `PAYLOAD` of the message with same `MESSAGE ID` should be concat
        - `ENCODING`: the `PAYLOAD` encoding
            - `0b00`: without payload
            - `0b01`: binary
            - `0b10`: `UTF-8` payload
            - `0b11`: `UTF-16` payload
        - `RSV`: must be `0`

    - text
    
        text message is used for `WebSockets` only, so, it only contains
        `KIND` field with one byte:
        
        - `p`: `Ping message`
        - `q`: `Request message`
        - `n`: `Notify message`
        - `s`: `Response message`
        - `f`: `Following message`, `deprecated`, for `WebSockets` is frame based
        - `c`: `Close message`
        
4. Request message

    - binary
    
        ```text
        <HEAD><MESSAGE ID>[PAYLOAD LENGTH]<ACTION>
        [PAYLOAD]
        ```
      
    - text
    
        ```text
        q<MESSAGE ID>,<ACTION>
        [PAYLOAD]
        ```
    
    - `MESSAGE ID`: a `uint16LE`
    - `PAYLOAD LENGTH`: `varuint`
    - `ACTION`: a string or `varuint`

5. Notify message

    - binary
    
        ```text
        <HEAD>[MESSAGE ID]<ACTION>
        [PAYLOAD LENGTH][PAYLOAD]
        ```
      
        - If `FIN`, `MESSAGE ID` should be omitted
        
    - text
    
        ```text
        n<ACTION>
        [PAYLOAD]
        ```

6. Response message

    - binary
    
        ```text
        <HEAD><RESPONSE STATUS><MESSAGE ID>
        [PAYLOAD LENGTH][PAYLOAD]
        ```
  
    - text
    
        ```text
        s<RESPONSE STATUS><MESSAGE ID>
        ```

    - `RESPONSE STATUS`
        - `0x0`: OK
        - `0x1`: Client error, should not retry
        - `0x2`: Server internal error, should retry
        - `0x3`: Network error, should not appears in messages, just emit to client

7. Following message

    - binary
    
        ```text
        <HEAD><MESSAGE ID><PAYLOAD LENGTH><PAYLOAD>
        ```

8. Close message

    - binary
    
        ```text
        <HEAD><CLOSE STATUS>[PAYLOAD LENGTH][PAYLOAD]
        ```
  
    - `CLOSE STATUS`
    
        - `0x0`: close normal
        - `0x1`: connection closed
        - `0x3`: server close
