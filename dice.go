package main

import "math/rand"

func rollXdY(X, Y int) int {
  var roll int
  for i := 0; i < X; i++ {
    roll += rand.Intn(Y) + 1
  }
  return roll
}
