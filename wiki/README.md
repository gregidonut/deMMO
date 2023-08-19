# Wiki

This directory contains the main wiki of game objects

## Current business logic

All instances are currently housed in [ main.go ](https://github.com/gregidonut/deMMO/blob/main/wiki/main.go)

### 5 main objects:

- Biome  
  We currently have 5 Biome instances
    ```go
    biomeInstances := [5]topLevelObjects.Biome{
        {
            Name:              "muddyPlains",
        },
        {
            Name:              "forest",
        },
        {
            Name:              "rockyPlains",
        },
        {
            Name:              "desert",
        },
        {
            Name:              "wasteland",
        },
    }
    ```

we plan on making wasteland a pvp zone( hadn't been clearly defined yet )

- Map(no instances yet)
- Mob  
  we currently have 9 mob instances(including *nameless* boss) which is still
  incomplete some of the biomes lack the specified number of mobs as per business
  requirements since it should be three for each biome except wasteLand
    ```go
    biomeInstances := [5]topLevelObjects.Biome{
        {
            Name:              "muddyPlains",
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
        {
            Name:              "forest",
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
    ```
- Resource  
  we currently have 4 resources for all the biomes except for `wasteLand` biome.
    ```go
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
    ```
- Character  
  we currently have 2 characterClass instances, with a
  glaring TODO of what key differences should be
  ```go
    characterClassInstances := []topLevelObjects.CharacterClass{
        {
            Name: "gatherer",
        },
        {
            Name: "hunter",
        },
    }
  ```