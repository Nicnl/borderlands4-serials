package serial

import (
	"borderlands_4_serials/b4s/codex_loader/codex"
)

func (s *Serial) FindItemType() (codex.ItemType, bool) {
	manufacturerIndex, found := s.FindIntAtPos(0)
	if !found {
		return codex.ItemType{}, false
	}

	return codex.GetItemTypeByIndex(manufacturerIndex)
}
