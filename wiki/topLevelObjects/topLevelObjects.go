package topLevelObjects

type Biome struct {
	Name              string
	GatheringResource GatheringResource
	Mobs              []Mob
}

type Map struct {
}

type Mob struct {
	Name       string
	BossMob    bool
	RegularMob bool
}

type GatheringResource *string

type CharacterClass struct {
	Name string
}
