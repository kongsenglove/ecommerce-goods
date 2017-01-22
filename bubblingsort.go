package main

import(
  "fmt"
)

type SortInterface interface{
  sort()
}

type Sortor struct{
  name string
}

func (sortor Sortor) sort(array []int){
  arraylenth := len(array)
  for i := 0;  i< arraylenth; i++ {
    for j := 0; j < arraylenth-i-1; j++ {
      if( array[j] > array[j+1] ){
        array[j] , array[j+1] = array[j+1] , array[j]
      }
    }
  }
}

func main(){
  array := []int{6, 1, 3, 5, 8, 4, 2, 0, 9, 7}
  sortor := Sortor{name:"冒泡排序--从小到大--稳定--n*n---"}
  sortor.sort(array)
  fmt.Println(sortor.name,array)
}
