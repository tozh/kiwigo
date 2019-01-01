package client

import (
	"strings"
	"strconv"
)

func ResolveBulkStr(s string) (string, error) {
	if s[0] != '$' {
		// fmt.Println("111111")
		return "", errWrongFormat
	}
	pos := strings.IndexByte(s, '\r')
	// fmt.Println("Pos is ----->", pos)
	if pos == -1 {
		// fmt.Println("222222")
		return "", errWrongFormat
	}
	bulkLen, err := strconv.Atoi(s[1:pos])
	if err != nil {
		// fmt.Println(s[1:pos-1])
		// fmt.Println("333333")
		return "", errWrongFormat
	}
	newPos := strings.IndexByte(s[pos+2:],'\r' )
	// fmt.Println("newPos is ----->", newPos)

	if newPos != bulkLen {
		// fmt.Println("444444")
		return "", errWrongFormat
	}
	// fmt.Println(s[pos+2:newPos+pos+2])
	return s[pos+2:newPos+pos+2], nil
}

