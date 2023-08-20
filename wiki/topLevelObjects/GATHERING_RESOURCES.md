# Gathering Resources

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