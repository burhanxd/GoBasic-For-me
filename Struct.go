package main // learning go through this cool gitbook xD := https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/02.4.html

import "fmt"

type person struct{ //struct intialized
	name string
	age int
}

func older(p1,p2 person) (person, int){ // func to check who is older by passing struct values
	if p1.age > p2.age {
	return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main(){
	var burhan person
	burhan.name, burhan.age = "Burhan", 18 // method 1 to intialize struct
	myLappy := person{name: "myLappy", age: 9} // method 2 to intialize struct
	myPhone := person{"myPhone", 3} // method 3 to intialize struct
	
	  tb_Older, tb_diff := older(burhan, myLappy)
    tp_Older, tp_diff := older(burhan, myPhone)
    bp_Older, bp_diff := older(myLappy, myPhone)

    fmt.Printf("Of %s and %s, %s is older by %d years\n", burhan.name, myLappy.name, tb_Older.name, tb_diff)
    fmt.Printf("Of %s and %s, %s is older by %d years\n", burhan.name, myPhone.name, tp_Older.name, tp_diff)
    fmt.Printf("Of %s and %s, %s is older by %d years\n", myLappy.name, myPhone.name, bp_Older.name, bp_diff)

}
