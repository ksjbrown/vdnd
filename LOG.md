# Dev Log

## 2026-03-31

Rethinking of types and things for simplification:

- Almost all types will be declarative data structs
- Logic will be performed in a select few types, probably Encounters.
- Actions will be used to signify Character intent
- Encounters will evaluate the consequences of those actions.

Does it make sense to wrap everything in a function?
Could be a LOT of ways for data to be modified.
Might be better to hold everything behind interfaces.
Will keep primary things behind interfaces, and keep as much as possible behavioural.

## 2026-03-20

`Action` will be an interface.
This has the method `.Apply(Encounter)`.

Another thought, i think i'll not include the actual chance based stuff in the action itself.
Or at least, dont have the DieFormula perform the roll.
We might need hooks to modify roll result, or react somehow to rolls.
E.g. natural 20 on attack roll is critical hit.

Before really looking into the combat lifecycle hooks, maybe take a high level look at what actualy happens.

- Combat Begins
- Initiative Calculated
- Get next acting Character
  - process incoming actions.
  - wait until an end turn signal is received.
- check for combat end conditions
  - maybe enemies surrender after certain amount of damage.

Combat context should define very fundamental operations.
Since it tracks all active effects, modifiers, etc., it can apply these at the relevant times, or at least ensure lifecycle callbacks are run.




## 2026-03-19

Will try implement an attack action system.


## 2026-03-17

Thinking more about the combat loop and callbacks.
Will probably have something like a Combat type, containing all relevant info.
Combat can be sent to clients, and they must decide what action to take.
They return an action type, and then a separate object containing the action relevant data.

Combat Lifecycle hooks:
- `OnInitiativeRoll(c *Combat)`
- `OnCombatStart(c *Combat)`
- `OnCombatEnd(c *Combat)`
- `OnTurnStart(c *Combat)`
- `OnTurnEnd(c *Combat)`
- `OnCharacterTurnStart(c *Combat, ch *Character)`
- `OnCharacterTurnEnd(c *Combat, ch *Character)`
- `OnCharacterMoveStart(c *Combat, ch *Character, m *MoveAction)`

Regarding AI usage, I woukd like to define a structure for dialogue, exploration, dialogue choices.
Warhammer 40k Rogue Trader system will work well here I think.
AI can generatively create all scenes.
To begin, I'll probably hava the AI generate responses dynamically, it will probably be easier.

Will probably develop this Director in Python, so we can natively use the cpp llama bindings.
Director must request LLM responses to be in JSON, and we check output.

Conceptually, how to split up Dialogues?

A `Dialogue` presents some text, describing what is happening, or what is going on.
`DialogueBranch` options define dialogue options or actions a character may take, and which dialogue ID this leads to.
A dialogue branch most often leads to another dialogue.
Dialogues can also lead to combat encounters, end dialogue, etc.

Open Question, how do we handle the free exploration part? 
I think it must be all handled thru dialogues, maybe we keep it to one-shot campaigns.
This is also how its done in Real DnD.

Dialogues must also exist in some greater context, maybe `Scenario`.
This defines the location, the characters present, initial dialogue, etc.

This shall be a hybrid design design of strict database json types, and conceptual descriptions ans characteristics of characters, places, etc.

Some types that might be useful here:

```python
class Location:
    description: str
    

class Scenario:
    location: Location


class Dialog:

    # character who speaks the dialogue, or is the topic during narrator discussion.
    speaker: Npc | None

    # character who is being spoken to
    spoken_to: Character | Npc | None

class DialogueBranch:

    # the dialogue that lead to this dialogue option
    parent: Dialogue | None

    # the text to display when presenting this option
    text: str

    # conditions that must be met to have this option visible
    conditions = []

    target = Dialogue | SkillCheck | AbilityCheck | Combat

```

What I'm also thinking a lot about now is a Godot frontend.
Cartoony style sprite characters on an isometric grid.
Use cover mechanics from Warhammer game rules.

Visibility mechanics might be very cool!
Dark light shading.
Path tracing might need to he calculated on the server...

But, designing and implementing a 3D grid based world might be pretty cool.
Elevation in particular, what that means for visibility.

## 2026-03-15

Will just try to implement the a full character representation, send this data via API, and present this in our angular app.
Refactoring of various character components can come later, separation of Character from NPC, etc.

Good resource for this is just simply trying to recreate a character sheet.
Better to structure all required data, then we can move onto more functional components.

Also just had a thought, what could make this a lot cooler is using LLM to act as a GM.
Combat rules, inventory, etc. will be still stored as actual code logic.
But the creative parts, dialogues, etc. can definitely be generated.

Will implement this as a DM Creative Director type separate component.
Will probably use webhooks or similar to query responses from a separate server.

## 2026-03-08

Playing around with new Angular.
Component based stuff seems so far pretty similar to Spring Boot.
Will keep at it.

## 2026-03-08

Will try to nail down scope, and develop in a way that allows things to be added in the future.

First, will just setup a basic character creation, upload, and display.

Things that need to be shown:

- Character
  - Name
  - Level
  - Race
  - Classes
  - Origins
  - Ability Scores 
  - Proficiencies
    - Equipment
    - Saving Throws
    - 

I think I will set up in parallel the Angular app.
So I can test and use the api in parallel.

Should I use a monorepo for this?

- I think we will, will keep each project in its own directory, `server/` and `client/`. 

Docker compose will come at some point, tho I think for now I'll run everything locally on debug servers.

## 2026-03-07

Fresh start, will try be more formal with this implementation.

Going over character basics.
Idea will be to have some basic type Actor, that represents any creature that can act in a combat scenario.
Maybe there will be something like a Combatant interface that embeds Actor.

I'll call it Character for now, but if we want to serialize it, it will look something like:

```json Example Character json representation
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
