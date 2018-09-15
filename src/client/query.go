package client

import (
	"strconv"
)

func (c *Client) addQuery(s string) {
	_, err := c.writer.WriteString(s)
	if err != nil {
		// fmt.Println(err)
		return
	}
}

func (c *Client) addQueryIntWithPrefix(i int, prefix byte) {
	if prefix == '*' && i < SHARED_BULKHDR_LEN {
		c.addQuery(shared.MultiBulkHDR[i])
	} else if prefix == '$' && i < SHARED_BULKHDR_LEN {
		//// fmt.Println("--------->",shared.BulkHDR[i])
		c.addQuery(shared.BulkHDR[i])
	} else {
		buf := Buffer{}
		buf.WriteByte(prefix)
		if 0 < i && i < SHARED_INTEGERS {
			buf.WriteString(shared.Integers[i])
		} else {
			buf.WriteString(strconv.Itoa(i))
		}
		buf.WriteByte('\r')
		buf.WriteByte('\n')
		c.addQuery(buf.String())
	}
}

func (c *Client) addQueryInt(i int) {
	if i == 0 {
		c.addQuery(shared.Zero)
	} else if i == 1 {
		c.addQuery(shared.One)
	} else {
		c.addQueryIntWithPrefix(i, ':')
	}
}

func (c *Client) addQueryMultiBulkLen(length int) {
	c.addQueryIntWithPrefix(length, '*')
}

func (c *Client) addQueryBulkLenOfStr(s string) {
	c.addQueryIntWithPrefix(len(s), '$')
}

func (c *Client) addQueryBulkStr(s string) {
	if s == "" {
		c.addQuery(shared.NullBulk)
	} else {
		//// fmt.Println("addQueryBulkStr", s)
		c.addQueryBulkLenOfStr(s)
		c.addQuery(s)
		c.addQuery(shared.Crlf)
	}
}

func (c *Client) addQueryBulkInt(i int) {
	c.addQueryBulkStr(strconv.Itoa(i))
}
