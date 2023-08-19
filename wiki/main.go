package main

import "github.com/gregidonut/deMMO/wiki/topLevelObjects"

func main() {
	// gathering resources need to be declared twice since we are using pointer values
	// this is to make it nil value for biomes that don't have resources e.g(wasteLand)

	var (
		WATER_RESOURCE topLevelObjects.GatheringResource = new(string)
		WOOD_RESOURCE  topLevelObjects.GatheringResource = new(string)
		OIL_RESOURCE   topLevelObjects.GatheringResource = new(string)
		ROCK_RESOURCE  topLevelObjects.GatheringResource = new(string)
	)
	*WATER_RESOURCE = "water"
	*WOOD_RESOURCE = "wood"
	*OIL_RESOURCE = "oil"
	*ROCK_RESOURCE = "rock"

	biomeInstances := [5]topLevelObjects.Biome{
		{
			Name:              "muddyPlains",
			GatheringResource: WATER_RESOURCE,
			Mobs: []topLevelObjects.Mob{
				{
					Name:       "snek",
					BossMob:    false,
					RegularMob: true,
				},
				{
					Name:       "babyWaterElemental",
					BossMob:    false,
					RegularMob: true,
				},
				{
					Name:       "Frog",
					BossMob:    false,
					RegularMob: true,
				},
			},
		},
		{
			Name:              "forest",
			GatheringResource: WOOD_RESOURCE,
			Mobs: []topLevelObjects.Mob{
				{
					Name:       "wolf",
					BossMob:    false,
					RegularMob: true,
				},
				{
					Name:       "treant",
					BossMob:    false,
					RegularMob: true,
				},
				{
					// TODO: need one more mob since requirement is 3 mobs
				},
			},
		},
		{
			Name:              "rockyPlains",
			GatheringResource: ROCK_RESOURCE,
			Mobs: []topLevelObjects.Mob{
				{
					Name:       "golem",
					BossMob:    false,
					RegularMob: true,
				},
				{
					// TODO: need two more mob since requirement is 3 mobs
				},
				{
					// TODO: need one more mob since requirement is 3 mobs
				},
			},
		},
		{
			Name:              "desert",
			GatheringResource: OIL_RESOURCE,
			Mobs: []topLevelObjects.Mob{
				{
					Name:       "scorpion",
					BossMob:    false,
					RegularMob: true,
				},
				{
					Name:       "vulture",
					BossMob:    false,
					RegularMob: true,
				},
				{
					// TODO: need one more mob since requirement is 3 mobs
				},
			},
		},
		{
			Name:              "wasteland",
			GatheringResource: nil,
			Mobs: []topLevelObjects.Mob{
				// this biome has only one mob(boss)
				{
					// TODO: think of bossmob name
					Name:       "",
					BossMob:    true,
					RegularMob: false,
				},
			},
		},
	}

	characterClassInstances := []topLevelObjects.CharacterClass{
		{
			Name: "gatherer",
		},
		{
			Name: "hunter",
		},
	}

	// throwing away memory instance wince we only need the linter to make sure
	// instantiations of topLevelObjects all follow business rules
	_ = biomeInstances
	_ = characterClassInstances
}
