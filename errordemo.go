package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// 不同类型异常处理函数
func ExcepitonTest(){
	/*
	if err != nil {
		switch err.(type) {
		case *json.SyntaxError:
			...
		case *ZeroDivisionError:
			...
		case *NullPointerError:
			...
		default:
			...
		}
	}
	*/
}

type Point struct {
	Longitude int
	Latitude int
	Distance int
	ElevationGain int
	ElevationLoss int
}

// err != nil地狱
func parse(r io.Reader) (*Point, error) {
	var p Point
	if err := binary.Read(r, binary.BigEndian, &p.Longitude); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.Latitude); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.Distance); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.ElevationGain); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.ElevationLoss); err != nil {
		return nil, err
	}

	return &p, nil
}

// 通过函数式编程，优化异常处理
func parseB(r io.Reader) (*Point, error) {
	var p Point
	var err error
	read := func(data interface{}) {
		if err != nil {
			return
		}
		err = binary.Read(r, binary.BigEndian, data)
	}

	read(&p.Longitude)
	read(&p.Latitude)
	read(&p.Distance)
	read(&p.ElevationGain)
	read(&p.ElevationLoss)

	if err != nil {
		return &p, err
	}
	return &p, nil
}

// 通过结构体，再次优化异常处理
type Reader struct {
	r   io.Reader
	err error
}

func (r *Reader) read(data interface{}) {
	if r.err == nil {
		r.err = binary.Read(r.r, binary.BigEndian, data)
	}
}

func parseC(input io.Reader) (*Point, error) {
	var p Point
	r := Reader{r: input}

	r.read(&p.Longitude)
	r.read(&p.Latitude)
	r.read(&p.Distance)
	r.read(&p.ElevationGain)
	r.read(&p.ElevationLoss)

	if r.err != nil {
		return nil, r.err
	}

	return &p, nil
}

// 向上传递错误内容
func errorWrap(err error) error {
	if err != nil {
		return fmt.Errorf("something failed: %v", err)
	}

	return nil
}

// 包装错误1
type authorizationError struct {
	operation string
	err error   // original error
}

func (e *authorizationError) Error() string {
	return fmt.Sprintf("authorization failed during %s: %v", e.operation, e.err)
}

// 包装错误2
type causer interface {
	Cause() error
}

func (e *authorizationError) Cause() error {
	return e.err
}

/*
// 错误包装3
func errDel (err error) error{
	if err != nil {
	return errors.Wrap(err, "read failed")
	}

	// Cause接口
	switch err := errors.Cause(err).(type) {
	case *MyError:
	// handle specifically
	default:
	// unknown error
}
*/

// 使用defer关键字在函数退出时关闭文件。
// Fatalf会导致系统退出os.Exit(1)
func main() {
	r, err := os.Open("a")
	if err != nil {
		log.Fatalf("error opening 'a'\n")
	}
	defer Close(r)

	r, err = os.Open("b")
	if err != nil {
		log.Fatalf("error opening 'b'\n")
	}
	defer Close(r)
}
