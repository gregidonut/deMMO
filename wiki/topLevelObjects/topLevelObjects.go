package topLevelObjects

import "github.com/gregidonut/deMMO/wiki/topLevelObjects/biomeTypes"

type Biome struct {
	Name              biomeTypes.BiomeType
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
