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

```go
    
```