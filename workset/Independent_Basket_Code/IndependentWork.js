const basket = [];

	
	const items = [
	{id: "banana", name : "Banana"},
	{id: "catFood", name : "Cat Food"},
	{id: "bread", name : "Bread"},
	{id: "avocado", name : "Avocado"}
	
	]
	
	items.forEach(item => {
	const getCorrectItem = document.getElementById(item.id);
	getCorrectItem.addEventListener("click", () => {
	addItemToBasket(item.name)
	});
	});
	
		
	const addItemToBasket = (item) => {
	basket.push(item);
	console.log("Added",item, "to basket")
	return basket
	}
	

const viewBasketButton = document.getElementById("viewBasket");

viewBasketButton.addEventListener("click", () => {
console.log("Basket:", basket);
});

const clearBasketButton = document.getElementById("clearBasket");

clearBasketButton.addEventListener("click", () => {
basket.length = 0
console.log("Basket has been cleared!")
return basket
});

const removeLastItemFromBasket = document.getElementById("removeLastItem");

removeLastItemFromBasket.addEventListener("click", () => {
let removedItem = basket.pop();
console.log("Removed the last item ", removedItem ,"from Basket");
return basket
});


const userTotal = document.getElementById("user_score");
let userScore = 0

userTotal.addEventListener("click", () => {
alert(userScore + " Pence")	


});


const getPriceForItem = (item) =>{
 if (item == "Banana"){
	return 50;
  
 }
 
 if (item == "Cat Food"){
	return 90;
  
 } 
 
 if (item == "Bread"){
	return 70;
  
 }
 
 if (item == "Avocado"){
	return 150;
  
 }

}


const totalPriceForItemsInBasket = document.getElementById("total");
totalPriceForItemsInBasket.addEventListener("click", () => {
 let totalPrice = 0;
 
 for(let item of basket) {

totalPrice += getPriceForItem(item)
} 
alert("Total Price: " + totalPrice + " Pence")
});


