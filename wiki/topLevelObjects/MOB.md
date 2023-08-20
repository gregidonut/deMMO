# Mob

we currently have 9 mob instances(including *nameless* boss) which is still
incomplete some of the biomes lack the specified number of mobs as per business
requirements since it should be three for each biome except wasteLand  
per last meeting, there should be 17 mobs which is pretty ambitious IMO

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