export interface Entity {
  id: number
  hp: {
    val: number
    max: number
    tmp?: number
  }
  ac: number
  speed: number
  abilities: [
    str: number,
    dex: number,
    con: number,
    int: number,
    wis: number,
    cha: number,
  ]
  proficiencies?: {
    skills?: number[]
    savingThrows?: number[]
    equipment?: {
      weapons?: number[]
      tools?: number[]
    }
  }
  hi: number
  conditions?: {
    type: number
    duration?: number
  }[]
}
