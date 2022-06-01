package main

import "fmt"

func main() {
  var (
    x SuccinctDS
    s string
  )

  fmt.Println("Inserire la sequenza di pantesi ben bilanciata su cui eseguire le operazioni:")
  fmt.Scan(&s)
  x.Build(s)

  var q int
  fmt.Println("Inserire il numero di query:")
  fmt.Scan(&q)

  for i := 0; i < q; i++ {
    var ind, tipo int

    fmt.Println("Inserire il tipo di query (0 = FindClose, 1 = FindOpen, 2 = LeftEnclose, 3 = FindEnclose):")
    fmt.Scan(&tipo)

    fmt.Println("Inserire la posizione su cui eseguire la query:")
    fmt.Scan(&ind)

    fmt.Print("Il risultato della query è: ")
    switch tipo {
      case 0:
        fmt.Println(x.FindClose(ind))
      case 1:
        fmt.Println(x.FindOpen(ind))
      case 2:
        fmt.Println(x.LeftEnclose(ind))
      case 3:
        fmt.Println(x.FindEnclose(ind))
    }
  }
}