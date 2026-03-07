# Dev Log

## 2026-03-07

Fresh start, will try be more formal with this implementation.

Going over character basics.
Idea will be to have some basic type Actor, that represents any creature that can act in a combat scenario.
Maybe there will be something like a Combatant interface that embeds Actor.

I'll call it Character for now, but if we want to serialize it, it will look something like:

```json
{
    "hp": {
        "val": 10,
        "max": 12,
        "tmp": 1,
    },
    // Base Armor Class of the creature.
    // if omitted, assume 10
    "ac": 10,
    // Speed of the character, if omitted, assume standard 30
    "speed": 30,
    //  all abilities in order STR, DEX, CON, INT, WIS, CHA
    "abilities": [1, 2, 3, 4, 5, 6],
    // proficiencies
    "proficiencies": {
        "skills": [1, 2, 3],
        "savingThrows": [1, 2],
        "equipment" : {
            "weapons": [1, 2],
            "tools": [1, 2],
        },
    },
    // Heroic Inspiration
    "hi": 0,
    // Conditions
    // type defines what the condition is, optional duration field indicates when the condition stops
    "conditions": [
        { "type": 1 },
        { "type": 2, "duration": 2 }
    ]
}
```

Maybe I need to decide the proper scale for what this should be.
It's probably easy to make it a simple combat calculator, charater tracker, etc. 


At the moment I'm just focussed on the combat side, but interesting to note some of the other rules about exploration scenarios.

- Pace rules for travel
  - Fast gives disadvantage on perception, medium is normal, slow is advantage.
- Ties for initiative rolls
  - DM decides order of ties between monsters
  - Players decide about order when players are tied
  - DM decides if players are tied with monsters.
