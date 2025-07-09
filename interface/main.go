package main

import _case "interface/case"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	/*
		cat := _case.NewCat()
		animalLife(cat)

	*/
	_case.Init()

}
func animalLife(a _case.AnimalI) {
	a.Each()
	a.Drink()
	a.Sleep()
	a.Run()
}
