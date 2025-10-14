package codex

type ItemType struct {
	Manufacturer string `json:"manufacturer"`
	Type         string `json:"type"`
	Serial       string `json:"serial"`
	Index        uint32 `json:"-"` // Extracted from Serial
}

var (
	itemTypes        []ItemType
	itemTypesByIndex = make(map[uint32]*ItemType)
)

func GetItemTypeByIndex(index uint32) (ItemType, bool) {
	itemType, found := itemTypesByIndex[index]
	if !found {
		return ItemType{}, false
	}
	return *itemType, found
}

func PushItemType(itemType ItemType) {
	itemTypes = append(itemTypes, itemType)
	itemTypesByIndex[itemType.Index] = &itemType
}
