package simple_checker

type bound struct {
	start byte
	end   byte
}

func (b *bound) String() string {
	return string(b.start) + "-" + string(b.end)
}
