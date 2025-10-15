package codex

func (i *Item) Type() (ItemType, bool) {
	manufacturerIndex, found := i.FindIntAtPos(0)
	if !found {
		return ItemType{}, false
	}

	return GetItemTypeByIndex(manufacturerIndex)
}
