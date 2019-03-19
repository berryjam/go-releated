package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseRequest(t *testing.T) {
	assert.NotPanicsf(t, func() {
		ParseRequest([]byte("*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n"))
	}, "the request header is legal and won't cause panic")
	assert.Equal(t, "*3\n$3\nSET\n$5\nmykey\n$7\nmyvalue\n", ParseRequest([]byte("*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n")))
}

func TestParseRequest2(t *testing.T) {
	assert.Panicsf(t, func() {
		ParseRequest([]byte("*3\r\n$3\r\nSETNX\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n"))
	}, "参数:'SETNX'字节数量错误，应该为:3，实际为:5")
}

func TestParseRequest3(t *testing.T) {
	assert.Panicsf(t, func() {
		ParseRequest([]byte("-3\r\n$3\r\nSETNX\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n"))
	}, "第一个字节必须为*")
}

func TestParseRequest4(t *testing.T) {
	assert.Panicsf(t, func() {
		ParseRequest([]byte("*3\r$3\r\nSETNX\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n"))
	}, `请求必须以\r\n作为分割符`)
}

func TestParseRequest5(t *testing.T) {
	assert.Panicsf(t, func() {
		ParseRequest([]byte("*2\r\n$3\r\nSETNX\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n"))
	}, "参数数量不对，应该为:5，实际为:7")
}

func TestParseReply(t *testing.T) {
	assert.NotPanicsf(t, func() {
		ParseReply([]byte("+OK\r\n"))
	}, "the request header is legal and won't cause panic")
	assert.Equal(t, "OK", ParseReply([]byte("+OK\r\n")))
}

func TestParseErrReply(t *testing.T) {
	assert.NotPanicsf(t, func() {
		ParseReply([]byte("-ERR unknown command 'foobar'\r\n"))
	}, "the request header is legal and won't cause panic")
	assert.Equal(t, `ERR unknown command 'foobar'`, ParseReply([]byte("-ERR unknown command 'foobar'\r\n")))
}

func TestParseIntegerReply(t *testing.T) {
	assert.NotPanicsf(t, func() {
		ParseReply([]byte(":1000\r\n"))
	}, "the request header is legal and won't cause panic")
	assert.Equal(t, "1000", ParseReply([]byte(":1000\r\n")))
}

func TestParseBulkReply(t *testing.T) {
	assert.NotPanicsf(t, func() {
		parseBulkReply([]byte("$6\r\nfoobar\r\n"))
	}, "the request header is legal and won't cause panic")
	assert.Equal(t, "foobar", ParseReply([]byte("$6\r\nfoobar\r\n")))
}

func TestParseBulkReply1(t *testing.T) {
	assert.Equal(t, "", ParseReply([]byte("$-1")))
}

func TestParseMultiBulkReply(t *testing.T) {
	assert.Equal(t, "*5\n1\n2\n3\n4\nfoobar\n", ParseReply([]byte("*5\r\n:1\r\n:2\r\n:3\r\n:4\r\n$6\r\nfoobar\r\n")))
}
