// -*-coding utf-8-*-
//from[https://qiita.com/tfrcm/items/e2a3d7ce7ab8868e37f7]
//next "構造体"
package main

import "fmt"

func main(){
  /*
  f := func(a int, b int)(int, int){return a-b, a+b}
  fmt.Println(f(1, 2))
  func(msg string){fmt.Println(msg)}("hello golang")
  */
}


func constant(){
  const(
    sun = iota
    mon
    tue
  )
  fmt.Printf("%d, %d, %d\n", sun, mon, tue)
}

func Plus(a int, b int) (c int){
  c = a + b
  return
}

func addless(){
  p := 5
  var pa *int
  pa = &p
  fmt.Println(pa)
  fmt.Println(*pa)
  /*
  NG : fmt.Printf(pa)
  def : fmt.Print("string")
  save : fmt.Fprint(os.Stdout, "Hello world!")
  No-output : hello := fmt.Sprint("Hello world!")
  format : hello := "Hello world!"
           fmt.Printf("%s\n", hello)
           fmt.Printf("%#v\n", hello)
  space&\n : fmt.Println("Hello", "world!")
  */
}

func array(){
  var a1[5]int
  b1 := [3]int{1, 2, 3}
  c1 := [...]int{2, 4, 6, 2, 4}
}

func slice(){
  a2 := [5]int{2, 4, 6, 8, 10}
  b2 := a2[2:4] // [6, 8]
  fmt.Println(len(a2))
  fmt.Println(cap(a2))

  s1 := make([]int, 3) //[0 0 0]
  s2 := []int{1, 3, 5}
  s3 := append(s2, 8, 2, 10)
  s4 := copy(s1, s2) //s1 = [1, 3, 5]
}

func mapRoot(){
  // 値を指定しながら宣言する
  m := map[string]int{"fujimoto": 100, "arita": 200} // map[fujimoto:100 arita:200]

  // キーの存在を調べる
  v, ok := m["fujimoto"]

  fmt.Println(v)  // 100
  fmt.Println(ok) // true

  // 要素を削除する
  delete(m, "fujimoto")
  fmt.Println(m) // map[arita:200]
}

func ifor(){
  if score := 52; score > 80{
    fmt.Println(score)
  }else{
    fmt.Println("----")
  }

  for i := 0;i<10;i++{
  }
}

func rangeTest(){
  s := []int{2, 4, 6}
  for i, v := range s{ fmt.Println(i, v) }

  m := map[string]int{"hello":100, "golang":200}
  for k, l :=range m{ fmt.Println(k, l) }
}
