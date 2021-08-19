package polymorphism

import "fmt"

/*
type Income interface {
    calculate() int
    source() string
}
// FixedBilling represents projects of Fixed and Billing type.
type FixedBilling struct {
    projectName string
    biddedAmount int
}
// TimeAndMaterial represents projects of Time and Material type.
type TimeAndMaterial struct {
    projectName string
    noOfHours  int
    hourlyRate int
}
func (fb FixedBilling) calculate() int {
    return fb.biddedAmount
}

func (fb FixedBilling) source() string {
    return fb.projectName
}

func (tm TimeAndMaterial) calculate() int {
    return tm.noOfHours * tm.hourlyRate
}

func (tm TimeAndMaterial) source() string {
    return tm.projectName
}

func calculateNetIncome(ic []Income) {
    var netincome int = 0
    for _, income := range ic {
        fmt.Printf("Income From %s = $%d\n", income.source(), income.calculate())
        netincome += income.calculate()
    }
    fmt.Printf("Net income of organization = $%d", netincome)
}

func Polymorphism() {
    project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
    project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
    project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}

    incomeStreams := []Income{project1, project2, project3}
    calculateNetIncome(incomeStreams)
}
*/

type PaymentMethod interface {
	calculateFederalTax() float32
	calculateProvincialTax() float32
	calculateTotal() float32
}

type CreditCard struct {
	description string `default:"Credit Card"`
	amount      int
	commission  float32
}
type Cash struct {
	description string `default:"Cash"`
	amount      int
}

func (cc CreditCard) calculateFederalTax() float32 {
	return (float32(cc.amount) * 5) / 100
}
func (cc CreditCard) calculateProvincialTax() float32 {
	return (float32(cc.amount) * 10) / 100
}
func (cc CreditCard) calculateTotal() float32 {
	return float32(cc.amount) +
		cc.calculateFederalTax() +
		cc.calculateProvincialTax() +
		cc.commission
}

func (cs Cash) calculateFederalTax() float32 {
	return (float32(cs.amount) * 5) / 100
}
func (cs Cash) calculateProvincialTax() float32 {
	return (float32(cs.amount) * 10) / 100
}
func (cs Cash) calculateTotal() float32 {
	return float32(cs.amount) +
		cs.calculateFederalTax() +
		cs.calculateProvincialTax()
}

func calculateTotalPurchases(pMethods []PaymentMethod) {
	var total float32 = 0
	for _, pMethod := range pMethods {
		// fmt.Printf("Income From %s = $%d\n", pMethod.source(), pMethod.calculate())
		fmt.Println("Purchase by:", pMethod.calculateTotal())

		total += pMethod.calculateTotal()
	}
	// fmt.Printf("Total Amount = $%d", total)
	fmt.Println("Total Purchases = $", total)
}

func PurchaseCalculator() {
	purchase1 := CreditCard{
		description: "credit card",
		amount:      100,
		commission:  5,
	}
	purchase2 := Cash{
		description: "cash",
		amount:      1000,
	}

	// One by one
	fmt.Println("\nPurchase 1 by", purchase1.description, ", Amount:", purchase1.amount, ", Total (taxes+commission):", purchase1.calculateTotal())
	fmt.Println("Purchase 2 by", purchase2.description, ", Amount:", purchase2.amount, ", Total (taxes):", purchase2.calculateTotal())
	fmt.Print("\n")

	// ALL at the same time
	purchases := []PaymentMethod{purchase1, purchase2}
	calculateTotalPurchases(purchases)

	// Only one
	var x PaymentMethod = CreditCard{
		description: "desc",
		amount:      10,
		commission:  5,
	}
	fmt.Println("x.Total:", x.calculateTotal())
}
