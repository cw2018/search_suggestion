package treemap


type StrKey string

func (a StrKey) Compare(bKey Key) int {
	b := bKey.(StrKey)
	if string(a) > string(b) {
		return 1
	} else if string(a) < string(b) {
		return -1
	}
	return 0
}


type IntKey int

func (a IntKey) Compare(bKey Key) int {
        b := bKey.(IntKey)
        if int(a) > int(b) {
                return 1
        } else if int(a) < int(b) {
                return -1
        }
        return 0
}