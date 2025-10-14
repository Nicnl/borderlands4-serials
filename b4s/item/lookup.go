package item

// Kind is a unified enum for manufacturers, item types, and characters
// Add all relevant values here

type Kind uint32

const (
	Unknown Kind = iota
	Jacobs
	Tediore
	Torgue
	Maliwan
	Daedalus
	Vladof
	Order
	Ripper
	COV
	Hyperion
	Atlas
	// Item types
	Pistol
	Shotgun
	SMG
	Sniper
	AssaultRifle
	HeavyWeapon
	Grenade
	Shield
	Repkit
	ClassMod
	Enhancer
	// Characters
	Vex
	Amon
	Rafa
	Harlowe
)

func (k Kind) String() string {
	switch k {
	case Jacobs:
		return "Jacobs"
	case Tediore:
		return "Tediore"
	case Torgue:
		return "Torgue"
	case Maliwan:
		return "Maliwan"
	case Daedalus:
		return "Daedalus"
	case Vladof:
		return "Vladof"
	case Order:
		return "Order"
	case Ripper:
		return "Ripper"
	case COV:
		return "COV"
	case Hyperion:
		return "Hyperion"
	case Atlas:
		return "Atlas"
	case Pistol:
		return "Pistol"
	case Shotgun:
		return "Shotgun"
	case SMG:
		return "SMG"
	case Sniper:
		return "Sniper"
	case AssaultRifle:
		return "Assault Rifle"
	case HeavyWeapon:
		return "Heavy Weapon"
	case Grenade:
		return "Grenade"
	case Shield:
		return "Shield"
	case Repkit:
		return "Repkit"
	case ClassMod:
		return "Class Mod"
	case Enhancer:
		return "Enhancer"
	case Vex:
		return "Vex"
	case Amon:
		return "Amon"
	case Rafa:
		return "Rafa"
	case Harlowe:
		return "Harlowe"
	default:
		return "Unknown"
	}
}

// key is a key for two Kind values: always [Manufacturer/Character, ItemType]
type key struct {
	First, Second Kind
}

// idMap and reverseIDMap are private
var idMap = map[key]uint32{
	{Daedalus, Pistol}: 2,
	{Jacobs, Pistol}: 3,
	{Order, Pistol}: 4,
	{Tediore, Pistol}: 5,
	{Torgue, Pistol}: 6,
	{Ripper, Shotgun}: 7,
	{Daedalus, Shotgun}: 8,
	{Jacobs, Shotgun}: 9,
	{Maliwan, Shotgun}: 10,
	{Tediore, Shotgun}: 11,
	{Torgue, Shotgun}: 12,
	{Daedalus, AssaultRifle}: 13,
	{Tediore, AssaultRifle}: 14,
	{Order, AssaultRifle}: 15,
	{Vladof, Sniper}: 16,
	{Torgue, AssaultRifle}: 17,
	{Vladof, AssaultRifle}: 18,
	{Ripper, SMG}: 19,
	{Daedalus, SMG}: 20,
	{Maliwan, SMG}: 21,
	{Vladof, SMG}: 22,
	{Ripper, Sniper}: 23,
	{Jacobs, Sniper}: 24,
	{Maliwan, Sniper}: 25,
	{Order, Sniper}: 26,
	{Jacobs, AssaultRifle}: 27,
	{Vex, ClassMod}: 254,
	{Amon, ClassMod}: 255,
	{Rafa, ClassMod}: 256,
	{Harlowe, ClassMod}: 259,
	{Torgue, Repkit}: 261,
	{Maliwan, Grenade}: 263,
	{Hyperion, Enhancer}: 264,
	{Jacobs, Grenade}: 267,
	{Jacobs, Enhancer}: 268,
	{Daedalus, Grenade}: 270,
	{Maliwan, Enhancer}: 271,
	{Order, Grenade}: 272,
	{Ripper, HeavyWeapon}: 275,
	{Ripper, Repkit}: 274,
	{Daedalus, Repkit}: 277,
	{Maliwan, Shield}: 279, // Shield Energy
	{Order, Enhancer}: 281,
	{Vladof, HeavyWeapon}: 282,
	{Vladof, Shield}: 283, // Shield Armor
	{Atlas, Enhancer}: 284,
	{Order, Repkit}: 285,
	{COV, Enhancer}: 286,
	{Tediore, Shield}: 287, // Shield Armor
	{Maliwan, HeavyWeapon}: 289,
	{Tediore, Enhancer}: 292,
	{Order, Shield}: 293, // Shield Energy
	{Ripper, Enhancer}: 296,
	{Daedalus, Enhancer}: 299,
	{Ripper, Shield}: 300, // Shield Energy
	{Torgue, Enhancer}: 303,
	{Jacobs, Shield}: 306, // Shield Armor
	{Vladof, Enhancer}: 310,
	{Daedalus, Shield}: 312, // Shield Energy
	{Torgue, Shield}: 321, // Shield Armor
	// ... add more as needed ...
}

var reverseIDMap = map[uint32]key{}

func init() {
	for k, v := range idMap {
		reverseIDMap[v] = k
	}
}

// LookupID returns the ID for a [Manufacturer/Character, ItemType] combination
func GetItemTypeID(first, second Kind) (uint32, bool) {
	id, ok := idMap[key{first, second}]
	return id, ok
}

// LookupEnums returns the [Manufacturer/Character, ItemType] for a given ID
func GetKindEnums(id uint32) (Kind, Kind, bool) {
	if k, ok := reverseIDMap[id]; ok {
		return k.First, k.Second, true
	}
	return Unknown, Unknown, false
}
