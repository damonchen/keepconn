# message

消息以帧的形式存在

## frame

- type： uint8, 消息的类型
- length: uint16，表示payload的长度 
- payload： 消息的载体，为序列化后的值。

序列化的方法目前为：json

type: 

- 0: unknown
- 1: result
- 1: handshake
- 2: authenticate
- 3: data message
- 4: 

## handshake

client  request:

- type: 1
- payload: "0.0.1"

server response:

- type: 1
- errno:
    {
        "status": "success|fail"
        "message": "",
    }


## authenticate

client request:

- type: 2
- payload: []byte
    
    
## message send

