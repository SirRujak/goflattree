package goflattree

import (
	"errors"
	"math/bits"
)

func Index(depth, offset uint) uint {
	return (offset << (depth + 1)) | ((1 << depth) - 1)
}

func Depth(i uint) uint {
	return uint(bits.TrailingZeros(^i))
}

func Offset(i uint) uint {
	var depth uint
	depth = Depth(i)
	if IsEven(i) {
		return i / 2
	}
	return i >> (depth + 1)
}

func Parent(i uint) uint {
	var depth uint
	depth = Depth(i)
	return Index(depth+1, Offset(i)>>1)
}

func Sibling(i uint) uint {
	var depth uint
	depth = Depth(i)
	return Index(depth, Offset(i)^1)
}

func Uncle(i uint) uint {
	var depth uint
	return Index(depth+1, Offset(Parent(i))^1)
}

func Children(i uint) *[]uint {
	var tempResponse []uint
	var depth uint
	depth = Depth(i)
	if IsEven(i) {
		return nil
	} else if depth == 0 {
		tempResponse = []uint{i, i}
		return &tempResponse
	}
	var offset uint
	offset = Offset(i) * 2
	tempResponse = []uint{Index(depth-1, offset), Index(depth-1, offset+1)}
	return &tempResponse
}

func LeftChild(i uint) *uint {
	var depth uint
	depth = Depth(i)

	var tempResponse uint
	if IsEven(i) {
		return nil
	} else if depth == 0 {
		tempResponse = i
		return &tempResponse
	}
	tempResponse = Index(depth-1, Offset(i)<<1)
	return &tempResponse
}

func RightChild(i uint) *uint {
	var depth uint
	depth = Depth(i)

	var tempResponse uint
	if IsEven(i) {
		return nil
	} else if depth == 0 {
		tempResponse = i
		return &tempResponse
	}
	tempResponse = Index(depth-1, (Offset(i)<<1)+1)
	return &tempResponse
}

func RightSpan(i uint) uint {
	var depth uint
	depth = Depth(i)

	if depth == 0 {
		return i
	}
	return (Offset(i)+1)*(2<<depth) - 2

}

func LeftSpan(i uint) uint {
	var depth uint
	depth = Depth(i)
	if depth == 0 {
		return i
	}
	return Offset(i) * (2 << depth)
}

func Spans(i uint) (uint, uint) {
	return LeftSpan(i), RightSpan(i)
}

func Count(i uint) uint {
	var depth uint
	depth = Depth(i)
	return (2 << depth) - 1
}

func FullRoots(i uint, nodes *[]uint) error {
	if !IsEven(i) {
		return errors.New("you can only look up roots for depth 0 blocks")
	}
	if nodes == nil {
		return errors.New("you must pass a constructed slice into the nodes variable")
	}
	var tmp uint
	tmp = i >> 1
	var offset uint
	offset = 0
	var factor uint
	factor = 1

	for {
		if tmp == 0 {
			break
		}
		for factor*2 <= tmp {
			factor *= 2
		}
		*nodes = append(*nodes, offset+factor-1)
		offset += 2 * factor
		tmp -= factor
		factor = 1
	}
	return nil
}

func IsEven(num uint) bool {
	return (num & 1) == 0
}

func IsOdd(num uint) bool {
	return (num & 1) != 0
}
