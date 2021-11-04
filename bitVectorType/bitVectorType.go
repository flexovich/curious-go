package main

import "fmt"
import "bytes"

type BitVector struct {
	words []uint64
}

func (v *BitVector) Has (x int) bool {
	word, bit := x/64, uint(x%64)
	return len(v.words) > x && v.words[word]&(1<<bit) != 0
}

func (v *BitVector) Add (x int) {
	word, bit := x/64, uint(x%64)
	for len(v.words) <= word {
		v.words = append(v.words, 0)
	}
	v.words[word] |= 1 << bit
}

func (v *BitVector) AddAll(arr ...int) {
	for _, value := range arr {
		v.Add(value)
	}
}

func (v *BitVector) UnionWith(b *BitVector) {
	for i, bword := range b.words {
		if i < len(v.words) {
			v.words[i] |= bword
		} else {
			v.words = append(v.words, bword)
		}
	}
}

func (v *BitVector) IntersectWith(b *BitVector) {
	for i, bword := range b.words {
		if i < len(v.words) {
			v.words[i] &= bword
		} else { continue }
	}
}

func (v *BitVector) DifferenceWith(b *BitVector) {
	for i, bword := range b.words {
		if i < len(v.words) {
			v.words[i] &^= bword
		} else { continue }
	}
}

func (v *BitVector) SymmDifferenceWith(b *BitVector) {
        for i, bword := range b.words {
                if i < len(v.words) {
                        v.words[i] ^= bword
                } else { continue }
        }
}

func (v *BitVector) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range v.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (v *BitVector) Len() int {
	l := 0
	for _, word := range v.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<j) != 0 {
				l++
			}
		}
	}
	return l
}

func (v *BitVector) Remove(x int) {
	word, bit := x/64, x%64
	v.words[word] &^= 1 << bit
}

func (v *BitVector) Clear() {
	v.words = []uint64{}
}

func (v *BitVector) Copy() *BitVector {
	c := new(BitVector)
	c.words = make([]uint64, len(v.words))
	copy(c.words, v.words)
	return c
}

func (v *BitVector) Elems() []uint64 {
	var arr []uint64
	for i, word := range v.words {
		for j := 0; j < 64; j++ {
			if word&(1 << uint(j)) != 0 {
				arr = append(arr, uint64(i*64+j))
			}
		}
	}
	return arr
}

func main() {
	v := new(BitVector)
    v1 := &v
    
    v.AddAll(1, 2, 3)

    fmt.Println(*v)
// 	v.Add(0) why Add() doesn't create {1} set in this case? the output is {0}

	// v.AddAll(128, 97, 46)
	// fmt.Println(v)

	// arr := v.Elems()

	// fmt.Println(arr)
}
