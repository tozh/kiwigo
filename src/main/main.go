package main

import (
	"fmt"
	"time"
	"strconv"
	. "kiwigo/src/client"
	"github.com/gomodule/redigo/redis"
	"strings"
)

func main() {
	fmt.Println()
	testRedis0_9999()
	fmt.Println()
	testKiwi0_9999()
	fmt.Println()
	testRedis10000_20000()
	fmt.Println()
	testKiwi10000_20000()
	fmt.Println()
	testRedisComplex()
	fmt.Println()
	testKiwiComplex()

}

func testRedis0_9999() {
	fmt.Println("test Redis 0_9999")
	t1 := time.Now()
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		// handle error
	}
	defer c.Close()
	t2 := time.Now()
	tsset := time.Duration(0)
	tsget := time.Duration(0)
	setok := 0
	getok := 0
	equal := 0
	setfailed := 0
	getfailed := 0
	notEqual := 0
	for i:=1;i<10000;i++ {
		key := strconv.Itoa(i)
		tset:=time.Now()
		_, err1 := c.Do("SET", key, key)
		tsset += time.Since(tset)
		if err1 == nil {
			setok++
			tget := time.Now()
			v, err2 := redis.String(c.Do("GET", key))
			tsget += time.Since(tget)
			if err2 == nil {
				getok++
				if key != v {
					notEqual++
					// fmt.Println("Get Wrong Value--->", key, "<<< >>>", v)
				} else {
					equal++
				}
			} else {
				getfailed++
				// fmt.Println("Get Failed--->", i, "---->", err2)
			}
		} else {
			setfailed++
			// fmt.Println("Set Failed--->", i, "--->", err1)
		}
	}
	fmt.Println("t2---->", time.Since(t2))
	fmt.Println("t1---->", time.Since(t1))
	fmt.Println("setok", setok)
	fmt.Println("getok", getok)
	fmt.Println("equal", equal)
	fmt.Println("setfailed", setfailed)
	fmt.Println("getfailed", getfailed)
	fmt.Println("notEqual", notEqual)
	fmt.Println("time set", tsset)
	fmt.Println("time get", tsget)


}
func testRedis10000_20000() {
	fmt.Println("test Redis 10000_20000")

	t1 := time.Now()
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		// handle error
	}
	defer c.Close()
	t2 := time.Now()
	tsset := time.Duration(0)
	tsget := time.Duration(0)
	setok := 0
	getok := 0
	equal := 0
	setfailed := 0
	getfailed := 0
	notEqual := 0
	for i:=1;i<10000;i++ {
		key := strconv.Itoa(i)
		tset:=time.Now()
		_, err1 := c.Do("SET", key, key)
		tsset += time.Since(tset)
		if err1 == nil {
			setok++
			tget := time.Now()
			v, err2 := redis.String(c.Do("GET", key))
			tsget += time.Since(tget)
			if err2 == nil {
				getok++
				if key != v {
					notEqual++
					// fmt.Println("Get Wrong Value--->", key, "<<< >>>", v)
				} else {
					equal++
				}
			} else {
				getfailed++
				// fmt.Println("Get Failed--->", i, "---->", err2)
			}
		} else {
			setfailed++
			// fmt.Println("Set Failed--->", i, "--->", err1)
		}
	}
	fmt.Println("t2---->", time.Since(t2))
	fmt.Println("t1---->", time.Since(t1))
	fmt.Println("setok", setok)
	fmt.Println("getok", getok)
	fmt.Println("equal", equal)
	fmt.Println("setfailed", setfailed)
	fmt.Println("getfailed", getfailed)
	fmt.Println("notEqual", notEqual)
	fmt.Println("time set", tsset)
	fmt.Println("time get", tsget)
}
func testRedisComplex() {
	fmt.Println("test Redis complex")

	t1 := time.Now()
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		// handle error
	}
	defer c.Close()
	t2 := time.Now()
	tsset := time.Duration(0)
	tsget := time.Duration(0)
	setok := 0
	getok := 0
	equal := 0
	setfailed := 0
	getfailed := 0
	notEqual := 0
	for i:=1;i<10000;i++ {
		key := createKey(i)
		tset:=time.Now()
		_, err1 := c.Do("SET", key, key)
		tsset += time.Since(tset)
		if err1 == nil {
			setok++
			tget := time.Now()
			v, err2 := redis.String(c.Do("GET", key))
			tsget += time.Since(tget)
			if err2 == nil {
				getok++
				if key != v {
					notEqual++
					// fmt.Println("Get Wrong Value--->", key, "<<< >>>", v)
				} else {
					equal++
				}
			} else {
				getfailed++
				// fmt.Println("Get Failed--->", i, "---->", err2)
			}
		} else {
			setfailed++
			// fmt.Println("Set Failed--->", i, "--->", err1)
		}
	}
	fmt.Println("t2---->", time.Since(t2))
	fmt.Println("t1---->", time.Since(t1))
	fmt.Println("setok", setok)
	fmt.Println("getok", getok)
	fmt.Println("equal", equal)
	fmt.Println("setfailed", setfailed)
	fmt.Println("getfailed", getfailed)
	fmt.Println("notEqual", notEqual)
	fmt.Println("time set", tsset)
	fmt.Println("time get", tsget)
}

func testKiwi0_9999() {
	fmt.Println("test Kiwi 0_9999")

	t3 := time.Now()
	c := TcpClient("0.0.0.0", 9988)
	defer 	c.Close()

	t4 := time.Now()
	tsset := time.Duration(0)
	tsget := time.Duration(0)
	setok := 0
	getok := 0
	equal := 0
	setfailed := 0
	getfailed := 0
	notEqual := 0
	for i:=1;i<10000;i++ {
		key := strconv.Itoa(i)
		tset:=time.Now()
		_, err1 := c.Set(key, key)

		tsset += time.Since(tset)
		if err1 == nil {
			setok++
			tget := time.Now()
			v, err2 := c.Get(key)

			tsget += time.Since(tget)
			if err2 == nil {
				getok++
				if key != v {
					notEqual++
					// fmt.Println("Get Wrong Value--->", key, "<<< >>>", v)
				} else {
					equal++
				}
			} else {
				getfailed++
				// fmt.Println("Get Failed--->", i, "---->", err2)
			}
		} else {
			setfailed++
			// fmt.Println("Set Failed--->", i, "--->", err1)
		}
	}
	fmt.Println("t4---->", time.Since(t4))
	c.Close()
	fmt.Println("t3---->", time.Since(t3))
	fmt.Println("setok", setok)
	fmt.Println("getok", getok)
	fmt.Println("equal", equal)
	fmt.Println("setfailed", setfailed)
	fmt.Println("getfailed", getfailed)
	fmt.Println("notEqual", notEqual)
	fmt.Println("time set", tsset)
	fmt.Println("time get", tsget)
}
func testKiwi0_2() {
	fmt.Println("test Kiwi 0_9999")

	t3 := time.Now()
	c := TcpClient("0.0.0.0", 9988)
	defer 	c.Close()

	t4 := time.Now()
	tsset := time.Duration(0)
	tsget := time.Duration(0)
	setok := 0
	getok := 0
	equal := 0
	setfailed := 0
	getfailed := 0
	notEqual := 0
	for i:=1;i<10;i++ {
		key := strconv.Itoa(i)
		tset:=time.Now()
		_, err1 := c.Set(key, key)

		tsset += time.Since(tset)
		if err1 == nil {
			setok++
			tget := time.Now()
			v, err2 := c.Get(key)

			tsget += time.Since(tget)
			if err2 == nil {
				getok++
				if key != v {
					notEqual++
					// fmt.Println("Get Wrong Value--->", key, "<<< >>>", v)
				} else {
					equal++
				}
			} else {
				getfailed++
				// fmt.Println("Get Failed--->", i, "---->", err2)
			}
		} else {
			setfailed++
			// fmt.Println("Set Failed--->", i, "--->", err1)
		}
	}
	fmt.Println("t4---->", time.Since(t4))
	c.Close()
	fmt.Println("t3---->", time.Since(t3))
	fmt.Println("setok", setok)
	fmt.Println("getok", getok)
	fmt.Println("equal", equal)
	fmt.Println("setfailed", setfailed)
	fmt.Println("getfailed", getfailed)
	fmt.Println("notEqual", notEqual)
	fmt.Println("time set", tsset)
	fmt.Println("time get", tsget)
}

func testKiwi10000_20000() {
	fmt.Println("test Kiwi 10000_20000")

	t3 := time.Now()
	c := TcpClient("0.0.0.0", 9988)
	defer 	c.Close()


	t4 := time.Now()
	tsset := time.Duration(0)
	tsget := time.Duration(0)
	setok := 0
	getok := 0
	equal := 0
	setfailed := 0
	getfailed := 0
	notEqual := 0
	for i:=1;i<10000;i++ {
		key := strconv.Itoa(i)
		tset:=time.Now()
		_, err1 := c.Set(key, key)

		tsset += time.Since(tset)
		if err1 == nil {
			setok++
			tget := time.Now()
			v, err2 := c.Get(key)

			tsget += time.Since(tget)
			if err2 == nil {
				getok++
				if key != v {
					notEqual++
					// fmt.Println("Get Wrong Value--->", key, "<<< >>>", v)
				} else {
					equal++
				}
			} else {
				getfailed++
				// fmt.Println("Get Failed--->", i, "---->", err2)
			}
		} else {
			setfailed++
			// fmt.Println("Set Failed--->", i, "--->", err1)
		}
	}
	fmt.Println("t4---->", time.Since(t4))
	c.Close()
	fmt.Println("t3---->", time.Since(t3))
	fmt.Println("setok", setok)
	fmt.Println("getok", getok)
	fmt.Println("equal", equal)
	fmt.Println("setfailed", setfailed)
	fmt.Println("getfailed", getfailed)
	fmt.Println("notEqual", notEqual)
	fmt.Println("time set", tsset)
	fmt.Println("time get", tsget)
}

func testKiwiComplex() {
	fmt.Println("test Kiwi complex")

	t3 := time.Now()
	c := TcpClient("0.0.0.0", 9988)
	defer 	c.Close()


	t4 := time.Now()
	tsset := time.Duration(0)
	tsget := time.Duration(0)
	setok := 0
	getok := 0
	equal := 0
	setfailed := 0
	getfailed := 0
	notEqual := 0
	for i:=1;i<10000;i++ {
		key := createKey(i)
		tset:=time.Now()
		_, err1 := c.Set(key, key)

		tsset += time.Since(tset)
		if err1 == nil {
			setok++
			tget := time.Now()
			v, err2 := c.Get(key)

			tsget += time.Since(tget)
			if err2 == nil {
				getok++
				if key != v {
					notEqual++
					// fmt.Println("Get Wrong Value--->", key, "<<< >>>", v)
				} else {
					equal++
				}
			} else {
				getfailed++
				// fmt.Println("Get Failed--->", i, "---->", err2)
			}
		} else {
			setfailed++
			// fmt.Println("Set Failed--->", i, "--->", err1)
		}
	}
	fmt.Println("t4---->", time.Since(t4))
	c.Close()
	fmt.Println("t3---->", time.Since(t3))
	fmt.Println("setok", setok)
	fmt.Println("getok", getok)
	fmt.Println("equal", equal)
	fmt.Println("setfailed", setfailed)
	fmt.Println("getfailed", getfailed)
	fmt.Println("notEqual", notEqual)
	fmt.Println("time set", tsset)
	fmt.Println("time get", tsget)
}

func createKey(i int) string {
	buf := strings.Builder{}
	buf.WriteString("this is a test for int <")
	buf.WriteString(strconv.Itoa(i))
	buf.WriteString(">, do you like it?")
	return buf.String()
}