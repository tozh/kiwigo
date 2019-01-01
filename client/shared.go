package client

import (
	"strconv"
)

type SharedObjects struct {
	Crlf           string // "\r\n"
	NullBulk       string // "$-1\r\n"
	EmptyBulk      string // "$0\r\n"
	NullMultiBulk  string // "*-1\r\n"
	EmptyMultiBulk string // "*0\r\n"
	Zero           string // ":0\r\n"
	One            string // ":1\r\n"
	NegOne         string // ":-1\r\n"
	Ok             string // "+OK\r\n"
	Err            string // "-ERR\r\n"
	NoAuthErr      string // "-NOAUTH Authentication required.\r\n"
	OOMErr         string // "-OOM command not allowed when used memory > 'maxmemory'.\r\n"
	LoadingErr     string // "-LOADING Redis is loading the dataset in memory\r\n"
	SyntaxErr      string // "-ERR syntax error\r\n"
	WrongTypeErr   string // "-WRONGTYPE Operation against a key holding the wrong kind of value\r\n"
	Integers       [SHARED_INTEGERS]string
	MultiBulkHDR   [SHARED_BULKHDR_LEN]string // "*<value>\r\n"
	BulkHDR        [SHARED_BULKHDR_LEN]string // "$<value>\r\n"
}

func CreateShared() *SharedObjects {
	so := SharedObjects{
		Crlf:           "\r\n",
		NullBulk:       "$-1\r\n",
		EmptyBulk:      "$0\r\n",
		NullMultiBulk:  "*-1\r\n",
		EmptyMultiBulk:  "*0\r\n",
		Zero:           ":0\r\n",
		One:            ":1\r\n",
		NegOne:         ":-1\r\n",
		Ok:             "+OK\r\n",
		Err:            "-ERR\r\n",
		NoAuthErr:      "-NOAUTH Authentication required.\r\n",
		OOMErr:         "-OOM command not allowed when used memory > 'maxmemory'.\r\n",
		LoadingErr:     "-LOADING Redis is loading the dataset in memory\r\n",
		SyntaxErr:      "-ERR syntax error\r\n",
		WrongTypeErr:   "-WRONGTYPE Operation against a key holding the wrong kind of value\r\n",
		Integers:       [SHARED_INTEGERS]string{},
		MultiBulkHDR:   [SHARED_BULKHDR_LEN]string{}, // "*<value>\r\n"
		BulkHDR:        [SHARED_BULKHDR_LEN]string{}, // "$<value>\r\n"
	}
	for i := 0; i < SHARED_INTEGERS; i++ {
		so.Integers[i] = strconv.Itoa(i)
	}
	buf := Buffer{}
	for i := 0; i < SHARED_BULKHDR_LEN; i++ {
		buf.WriteByte('*')
		numStr := strconv.Itoa(i)
		buf.WriteString(numStr)
		buf.WriteString("\r\n")
		so.MultiBulkHDR[i] = buf.String()
		buf.Reset()
	}
	for i := 0; i < SHARED_BULKHDR_LEN; i++ {
		buf.WriteByte('$')
		numStr := strconv.Itoa(i)
		buf.WriteString(numStr)
		buf.WriteString("\r\n")
		so.BulkHDR[i] = buf.String()
		buf.Reset()
	}
	return &so
}

var shared = CreateShared()
