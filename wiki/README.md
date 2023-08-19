# Wiki

This directory contains the main wiki of game objects

## Current business logic

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
- Character

```go
    
```