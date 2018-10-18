package bank

var deposits = make(chan int) // send amount to deposit
var withdraw = make(chan int) // send amount to withdraw
var balances = make(chan int) // receive balance

func Deposit(amount int)  { deposits <- amount }
func Withdraw(amount int) { withdraw <- amount }
func Balance() int        { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case amount := <-withdraw:
			balance -= amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
