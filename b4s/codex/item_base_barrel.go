package codex

import part2 "borderlands_4_serials/b4s/serial_datatypes/part"

func (i *Item) BaseBarrel() (BaseBarrelInfos, bool) {
	manufacturerIndex, found := i.FindIntAtPos(0)
	if !found {
		return BaseBarrelInfos{}, false
	}

	baseIndex := i.FindPartAtPos(0, false)
	if baseIndex == nil {
		return BaseBarrelInfos{}, false
	}

	// Barrel may be anywhere from pos 1 to {n}
	pos := 1
	for {
		part := i.FindPartAtPos(pos, false)

		if part == nil {
			break
		}

		if part.SubType == part2.SUBTYPE_NONE {
			maybeBaseBarrel := BaseBarrel{
				ManufacturerIndex: manufacturerIndex,
				BaseIndex:         baseIndex.Index,
				BarrelIndex:       part.Index,
			}

			infos, found := Barrels[maybeBaseBarrel]
			if found {
				return infos, true
			}
		}

		pos++
	}

	return BaseBarrelInfos{}, false
}
