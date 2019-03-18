package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseRequest(t *testing.T) {
	assert.NotPanicsf(t, func() {
		ParseRequest([]byte("*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n"))
	}, "the request header is legal and won't cause panic")
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

