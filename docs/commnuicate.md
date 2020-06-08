# 通讯

server：控制机器的那台服务器
client: 需要被控制的那台机器

## 主要由以下几个步骤组成

1. 建立连接(client->server)
2. 发送消息(server)
3. 接受消息(client)
4. 消费消息(client)
5. 回复消息消费(client->server)
6. 接受回复的消息(server)

## 建立连接

- tcp connected
- handshake
- authenticate

### handshake

client send:

- magic
- client version

server response:

- success:
- fail:
    - unknown client
    - version not support

### authenticate

client send:

- client id
- authenticate info:
    - username
    - key

server response:

- success:
    - waiting
    - deny
    - accepted
- fail:
    - not support authenticate

## 发送消息

- generate frame message
- serialize frame message
- using tcp to send

### frame message

## 接受消息

- using tcp to receive 
- deserialize frame message
- generate frame message

## 消费消息

- work of the message with the type

## 回复消息

- generate frame message
- serialize frame message
- using tcp to send

## 接受回复的消息

- using tcp to receive 
- deserialize frame message
- generate frame message
