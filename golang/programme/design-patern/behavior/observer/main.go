package main

import observer "always-learning/golang/programme/design-patern/behavior/observer/customer"

func main() {

	shirtItem := observer.NewItem("Nike Shoe")

	observerFirst := &observer.Customer{Id: "abc@gmail.com"}
	observerSecond := &observer.Customer{Id: "xyz@gmail.com"}

	shirtItem.Register(observerFirst)
	shirtItem.Register(observerSecond)

	shirtItem.UpdateAvailability()
	// observerThird := &observer.Customer{Id: "uio@gmail.com"}
	shirtItem.Deregister(observerSecond)
	shirtItem.UpdateAvailability()
	// shirtItem = observer.NewItem("Adidas Shoe")

	// observerFirst = &observer.Customer{Id: "iop@gmail.com"}
	// observerSecond = &observer.Customer{Id: "iuy@gmail.com"}

	// shirtItem.Register(observerFirst)
	// shirtItem.Register(observerSecond)

	// shirtItem.UpdateAvailability()

}
