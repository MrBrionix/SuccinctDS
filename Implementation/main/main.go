package main

import (
  "test"
  //"jacobsonDS"
  //"fmt"
)

func main() {
  //test.SuccinctDSManualTest()
  //test.SuccinctDSAutomatedTest(false, true)      // showDetails, slowCheck
  //test.SuccinctRankSelectTest(false, true)       // showDetails, slowCheck
  test.JacobsonTest(false)
}