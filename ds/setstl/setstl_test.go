package setstl

import (
	"fmt"
	"testing"

	"github.com/dengjiawen8955/gostl/ds/options"
)

//功能测试
// s1.Exists(1): true
// s1.Exists(3): false
// s1.Exists(1): false
// is: [2]
// 交集: [2]
// 并集=[]interface {}{1, 2, 3}
// 差集=[]interface {}{1}
func TestSet(t *testing.T) {
	s1 := New[int]()
	s1.Add(1)
	s1.Add(2)
	fmt.Printf("s1.Exists(1): %v\n", s1.Exists(1))
	fmt.Printf("s1.Exists(3): %v\n", s1.Exists(3))
	s1.Del(1)
	fmt.Printf("s1.Exists(1): %v\n", s1.Exists(1))
	is := s1.All()
	fmt.Printf("is: %v\n", is)
	s1.Add(1)
	s2 := New[int]()
	s2.Add(2)
	s2.Add(3)
	//交集
	s3 := s1.Inter(s2)
	s3s := s3.All()
	fmt.Printf("交集: %v\n", s3s)
	//并集
	s4 := s1.Union(s2)
	s4s := s4.All()
	fmt.Printf("并集=%#v\n", s4s)
	//差集
	s5 := s1.Diff(s2)
	s5s := s5.All()
	fmt.Printf("差集=%#v\n", s5s)
}

func TestWithSync(t *testing.T) {
	s1 := New[int](options.WithSync())
	go func() {
		for i := 0; i < 100000; i++ {
			s1.Add(1)
		}
	}()
	for i := 0; i < 100000; i++ {
		s1.Exists(1)
	}
	fmt.Printf("%s\n", "OK")
}

//会发生panic=
func TestWithoutSync(t *testing.T) {
	s1 := New[int]()
	go func() {
		for i := 0; i < 100000; i++ {
			s1.Add(1)
		}
	}()
	for i := 0; i < 100000; i++ {
		s1.Exists(1)
	}
	fmt.Printf("%s\n", "OK")
}
