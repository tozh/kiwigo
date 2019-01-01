package client


func (c *Client) setResult() (int, error) {
	reply := c.readBuf.String()
	if reply[0:len(reply)-1] == shared.One {
		return 1, nil
	}
	return 0, errSetFailed
}

func (c *Client) setQuery(key string, value string, params ...string) {
	c.addQueryMultiBulkLen(3)
	c.addQueryBulkStr("set")
	c.addQueryBulkStr(key)
	for _, str := range params {
		c.addQueryBulkStr(str)
	}
	c.addQueryBulkStr(value)
}
func (c *Client) Set(key string, value string, params ...string) (int, error) {
	defer c.reset()
	c.setQuery(key, value)
	c.writeToServer()
	readCh := make(chan error, 1)
	defer close(readCh)
	go c.readFromServer(readCh)
	select {
	case <-c.passiveCloseCh:
		// fmt.Println("Set ----> Stop Client")
		return 0, errEOF
	case err := <-readCh:
		if err == nil {
			return c.setResult()
		} else {
			return 0, err
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

func (c *Client) SetNx(key string, value string) (int, error) {
	return c.Set(key, value, []string{"NX"}...)
}

func (c *Client) SetEx(key string, value string) (int, error) {
	return c.Set(key, value, []string{"EX"}...)
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
