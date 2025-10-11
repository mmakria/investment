package main

import "fmt"

func main(){


	a := 2
	b := 2
	fmt.Println( (a+b))
	result := test123(2,4)
	fmt.Println(result)
	fmt.Println(comparaison(charlie, bob, thomas))


// test des inputs

var prenom string
fmt.Print("Entrez votre nom")
fmt.Scan(&prenom)
fmt.Println("Bonjour" + prenom)
}

type Person struct {
Nom string
Age int
}

var charlie = Person{Nom: "Charlie", Age: 122}
var bob = Person{Nom: "Bob", Age: 32}
var thomas= Person{Nom: "Thomas", Age: 12}

func comparaison(a Person, b Person, c Person) string{

	var result string
	if a.Age>b.Age {
		result = fmt.Sprintf(a.Nom + " est plus vieux/vielle que " + b.Nom)
		}else if a.Age<b.Age		{
		result = fmt.Sprintf(a.Nom + " est plus jeune que " +  b.Nom) 
	}

	return result
	
}


func test123(a int, b int) int {
return a + b
}


