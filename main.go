package main

func main() {

	//DynamicArrays()
	//find(89, 90, 91, 95)
	//find(45, 56, 67, 45, 90, 109)

	//arrays()
	//nestedLoops()
	//fizzBuzz()
	//crapsGame()

	/*for diceOne := 1; diceOne <= 6; diceOne++ {
		for diceTwo := 1; diceTwo <= 6; diceTwo++ {
			fmt.Println("Your roll is", diceOne, diceTwo)
			rollName(diceOne, diceTwo)
		}
	}*/

	/*for row := 0; row < 5; row++ {
		for col := 0; col < 10; col++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	*/

	/*item1 := [5]int{13, 27, 39, 56, 68} //items always start at 0 when you call them
	/*for index := 0; index < len(items); index++ {
	item := items[index]*/
	/*
		for index, item := range item1 {
			fmt.Println("index", index)
			fmt.Println(item)
			fmt.Println()
		}*/

	/*items := ([]string)
		items[0] = "A"
		items{1} = "B"
		items[2] = "C"
		fmt.Println(items)

		moreItems := append(items, "D", "E")
		fmt.Println(moreItems)

	}*/
}

//morningLoop()

/*todaysPlans := "nothing planned"
Plans(false, todaysPlans) //return we can go out
Plans(true, todaysPlans)  //returns we will stay in
Plans(true, "busy")
Plans(false, "busy")
Plans(true, "we're not doing anything")
Plans(false, "Any string that is not exactly 'busy'")
*/

/* func Plans(raining bool, todaysPlans string) {
	if !raining && todaysPlans != "busy" { //!raining flips the value of raining so it actually equals true, ! means not
		fmt.Println("We can go out!")
	} else {
		fmt.Println("We will stay in")
	}
}
*/

/*
func morningLoop() {

	for i := 1; i <= 5; i++ {

		fmt.Println("Good Morning")

		if i == 3 {

			fmt.Println("Have a nice day")

		}

}

*/
/*
func advancedConditional() {
	raining := false
	todaysPlans := "nothing planned"

	if !raining && todaysPlans != "busy" { //!raining flips the value of raining so it actually equals true, ! means not
		fmt.Println("We can go out!")
	} else {
		fmt.Println("We will stay in")
	}
}
*/
