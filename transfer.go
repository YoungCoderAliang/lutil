package lutil

import (
	"container/list"
	"fmt"
	"io"
	"strings"
)

func Reader2String(r io.Reader) (string, error) {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return buf.String(), nil
}

func Reader2Byte(r io.Reader) []byte {
	length := 0
	blist := list.New()
	b := make([]byte, 100)
	for {
		n, err := r.Read(b)
		if n > 0 {
			tp := make([]byte, n)
			copy(tp[0:n], b[0:n])
			blist.PushBack(tp)
			length += n
		}
		// 最后几个字节，会同时返回EOF
		if err == io.EOF {
			break
		}
	}
	rt := make([]byte, length)
	copied := 0
	for i := blist.Front(); i != nil; i = i.Next() {
		tp := i.Value.([]byte)
		copy(rt[copied:(copied+len(tp))], tp[0:])
		copied += len(tp)
	}
	return rt
}
