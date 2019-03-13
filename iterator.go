package goflattree

type Iterator struct {
	Index  uint
	Offset uint
	Factor uint
}

func NewIterator(index uint) *Iterator {
	var iter Iterator
	iter = Iterator{
		Index:  0,
		Offset: 0,
		Factor: 0,
	}

	iter.Seek(index)
	return &iter
}

func (iter *Iterator) Seek(index uint) {
	iter.Index = index
	if IsOdd(index) {
		iter.Offset = Offset(index)
		iter.Factor = TwoPow(Depth(index) + 1)
	} else {
		iter.Offset = index / 2
		iter.Factor = 2
	}
}

func (iter *Iterator) IsLeft() bool {
	return IsEven(iter.Offset)
}

func (iter *Iterator) IsRight() bool {
	return IsOdd(iter.Offset)
}

func (iter *Iterator) Prev() uint {
	if iter.Offset == 0 {
		return iter.Index
	}
	iter.Offset--
	iter.Index -= iter.Factor
	return iter.Index
}

func (iter *Iterator) Sibling() uint {
	if iter.IsLeft() {
		iter.Next()
	}
	return iter.Prev()

}

func (iter *Iterator) Parent() uint {
	if IsOdd(iter.Offset) {
		iter.Index -= iter.Factor / 2
		iter.Offset = (iter.Offset - 1) / 2
	} else {
		iter.Index += iter.Factor / 2
		iter.Offset /= 2
	}
	iter.Factor *= 2
	return iter.Index
}

func (iter *Iterator) LeftSpan() uint {
	iter.Index = iter.Index + 1 - iter.Factor/2
	iter.Offset = iter.Index / 2
	iter.Factor = 2
	return iter.Index
}

func (iter *Iterator) RightSpan() uint {
	iter.Index = iter.Index + iter.Factor/2 - 1
	iter.Offset = iter.Index / 2
	iter.Factor = 2
	return iter.Index
}

func (iter *Iterator) LeftChild() uint {
	if iter.Factor == 2 {
		return iter.Index
	}
	iter.Factor /= 2
	iter.Index -= iter.Factor / 2
	iter.Offset *= 2
	return iter.Index
}

func (iter *Iterator) RightChild() uint {
	if iter.Factor == 2 {
		return iter.Index
	}
	iter.Factor /= 2
	iter.Index += iter.Factor / 2
	iter.Offset = 2*iter.Offset + 1
	return iter.Index
}

func (iter *Iterator) Next() uint {
	iter.Offset++
	iter.Index += iter.Factor
	return iter.Index
}

func TwoPow(n uint) uint {
	if n < 31 {
		return 1 << n
	}
	return (1 << 30) * (1 << (n - 30))
}
