package main


type Character struct {
  STR int
  DEX int
  CON int
  INT int
  WIS int
  CHA int

  HP int
}

const MinAS = 8

func (c *Character) isOk() bool {
  return c.STR > MinAS || c.DEX > MinAS || c.CON > MinAS || c.INT > MinAS || c.WIS > MinAS || c.CHA > MinAS
}

func (c *Character) rollStats() {
  for ok := true; ok; ok = !c.isOk() {
    c.STR = rollXdY(3, 6)
    c.DEX = rollXdY(3, 6)
    c.CON = rollXdY(3, 6)
    c.INT = rollXdY(3, 6)
    c.WIS = rollXdY(3, 6)
    c.CHA = rollXdY(3, 6)
  }
}

func (c *Character) rollHP_0lvl() {
  c.HP = rollXdY(1, 4) + modOf[c.CON]
  if c.HP < 1 {
    c.HP = 1
  }
}

var modOf map[int]int = map[int]int{
  3:  -3,
  4:  -2,
  5:  -2,
  6:  -1,
  7:  -1,
  8:  -1,
  9:  0,
  10: 0,
  11: 0,
  12: 0,
  13: 1,
  14: 1,
  15: 1,
  16: 2,
  17: 2,
  18: 3,
}

func genChar_0lvl() (newChar Character) {
  newChar.rollStats()
  newChar.rollHP_0lvl()
  return
}

