package main

import (
  //"test"
  "benchmark"
)

func main() {
  //test.SuccinctDSManualTest()
  //test.SuccinctDSAutomatedTest(false, true)      // showDetails, slowCheck
  //test.SuccinctRankSelectTest(false, true)       // showDetails, slowCheck
  //test.JacobsonTest(false)                       // showDetails
  //benchmark.GenerateQuery()
  benchmark.SuccinctRankSelectBenchmark()
  //benchmark.JacobsonBenchmark()
  //benchmark.JacobsonSlowBenchmark()
}
