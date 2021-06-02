package main

import (
	"bytes"
	"fmt"
)

// interfaces describe behaviour
// interfaces are implicitly implemented

// naming conventions:
// if its an interface with a single method,
// name the interface the same as the method + 'er'

// implementing with values v pointers
// method set of value is all methods with value receivers
// method set of pointer is all methods, regradless of receiver type

// best practices
// prefer many, small interfaces
// don't export interfaces for types that will be consumed
// export interfaces for types that will be used by package
// design functions and methods to receive interfaces whenever possible

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	myInt := IntCounter(0)
	var inc Incrementer = &myInt

	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Yoyoyo Mr White, Im Jessy Joy"))
	wc.Close()

	// interface type conversion
	bwc, ok := wc.(*BufferedWriterCloser)
	if ok {
		fmt.Println(bwc)
	} else {
		fmt.Println("Conversion failed")
	}

	// empty interface
	// var myObj interface{} = NewBufferedWriterCloser()
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

// implicit implementation of Writer interface
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Printf(string(data))

	return n, err
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++

	return int(*ic)
}

// embedding interfaces
type Closer interface {
	Close() error
}

// interface composed of interfaces
type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

// buffers data
func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}

		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}

	return n, nil
}

// write data to console in 8 character chunks
func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}

	return nil
}

// constructor method
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
