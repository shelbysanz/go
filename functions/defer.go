package main

import (
	"fmt"
)

func bootup() {
	/*
		Defer is like saying take this line and ensure to execute before returning
		This line will for sure run when return is called and before return actually returns
		In this example instead of writing textio boot up done before each return statement,
		we only have to write it once deferred and it will always run before returning!
	*/
	defer fmt.Println("TEXTIO BOOTUP DONE")
	ok := connectToDB()
	if !ok {
		return
	}
	ok = connectToPaymentProvider()
	if !ok {
		return
	}
	fmt.Println("All systems ready!")
}

// don't touch below this line

var shouldConnectToDB = true

func connectToDB() bool {
	fmt.Println("Connecting to database...")
	if shouldConnectToDB {
		fmt.Println("Connected!")
		return true
	}
	fmt.Println("Connection failed")
	return false
}

var shouldConnectToPaymentProvider = true

func connectToPaymentProvider() bool {
	fmt.Println("Connecting to payment provider...")
	if shouldConnectToPaymentProvider {
		fmt.Println("Connected!")
		return true
	}
	fmt.Println("Connection failed")
	return false
}

func test(dbSuccess, paymentSuccess bool) {
	shouldConnectToDB = dbSuccess
	shouldConnectToPaymentProvider = paymentSuccess
	bootup()
	fmt.Println("====================================")
}

func main() {
	test(true, true)
	test(false, true)
	test(true, false)
	test(false, false)
}
