package client

type status int

func (c *Client) setResult() (status, error) {
	reply := c.readBuf.String()
	if reply[0:len(reply)-1] == shared.Ok {
		return STATUS_SUCCESS, nil
	}
	return STATUS_FAIL, errSetFailed
}

func (c *Client) setQuery(key string, value string) {
	c.addQueryMultiBulkLen(3)
	c.addQueryBulkStr("set")
	c.addQueryBulkStr(key)
	c.addQueryBulkStr(value)
}
func (c *Client) Set(key string, value string) (status, error) {
	defer c.reset()

	c.setQuery(key, value)
	c.writeToServer()
	readCh := make(chan error, 1)
	defer close(readCh)
	go c.readFromServer(readCh)
	select {
	case <-c.passiveCloseCh:
		// fmt.Println("Set ----> Stop Client")
		return STATUS_UNKNOWN, errEOF
	case err := <-readCh:
		if err == nil {
			return c.setResult()
		} else {
			return STATUS_FAIL, err
		}
		//case <- time.After(c.timeOut):
		//	// fmt.Println("TimeOut")
		//	return errTimeOut
	}
}

func (c *Client) getResult() (string, error) {
	reply := c.readBuf.String()
	// fmt.Println([]byte(reply))
	tmp := reply[0 : len(reply)-1]
	// fmt.Println("---------\n",tmp,"---------")
	// fmt.Println([]byte(tmp))

	if tmp == shared.WrongTypeErr {
		return "", errWrongType
	} else if tmp == shared.NullBulk {
		return "", errNil
	}

	return ResolveBulkStr(tmp)
}

func (c *Client) getQuery(key string) {
	c.addQueryMultiBulkLen(2)
	c.addQueryBulkStr("get")
	c.addQueryBulkStr(key)
}

func (c *Client) Get(key string) (string, error) {
	defer c.reset()

	c.getQuery(key)
	c.writeToServer()
	readCh := make(chan error, 1)
	defer close(readCh)
	go c.readFromServer(readCh)
	select {
	case <-c.passiveCloseCh:
		// fmt.Println("Get ----> Stop Client")
		return "", errEOF
	case err := <-readCh:
		if err == nil {
			return c.getResult()
		} else {
			return "", err
		}
	}
}

type Clienter interface {
	Set(key string, value string) (bool, error)
	MSet(params ...string) (int64, error)
	Get(key string) (string, error)
	MGet(keys ...string) (string, error)
	SetNx(key string) (bool, error)
	MSetNx(keys ...string) (int64, error)
	SetEx(key string) (bool, error)
	Append(key string) (int64, error)
	StrLen(key string) (int64, error)
	Del(keys ...string) (int64, error)
	Select(id int64) (bool, error)
	FlushAll() (bool, error)
	RandomKey() (string, error)
	Incr(key string) (int64, error)
	Decr(key string) (int64, error)
	Exists(keys ...string) (int64, error)
}
