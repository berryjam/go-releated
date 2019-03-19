# 本文档为redis通讯协议解析器的使用示例文档，分别介绍单元测试和命令行两种模式下的使用方法。

## Prerequisite

1. cd $GOPATH/src

2. mkdir -p github.com/berryjam

3. cd github.com/berryjam

4. git clone git@github.com:berryjam/go-releated.git

5. cd go-releated/src/redis_protocol_parser


### 单元测试模式

6. go test ./


### 命令行模式


7.1 go build redis_protocol_parser.go

7.2 解析请求:./redis_protocol_parser request $'*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n'

7.3 解析回复:./redis_protocol_parser reply $'*5\r\n:1\r\n:2\r\n:3\r\n:4\r\n$6\r\nfoobar\r\n'
