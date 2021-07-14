package main

import "fmt"

type Data struct {
}

/*
	Not following ISP
	We have different types of Calculation in Payroll Application.
	  - Earning
	  - Deduction
	  - PF
	  - Tax
	Here for some country PF and Tax not application
*/
type Calculation interface {
	Earning(data Data) bool
	Deduction(data Data) bool
	PF(data Data) bool
	Tax(data Data) bool
}

// Here we need all the calculations
type IndiaPayroll struct {
	//...
}

func (I *IndiaPayroll) Earning(data Data) bool {
	fmt.Print("Earning Calculated for India")
	return true
}

func (I *IndiaPayroll) Deduction(data Data) bool {
	fmt.Print("Deduction Calculated for India")
	return true
}

func (I *IndiaPayroll) PF(data Data) bool {
	fmt.Print("PF Calculated for India")
	return true
}

func (I *IndiaPayroll) Tax(data Data) bool {
	fmt.Print("Tax Calculated for India")
	return true
}

// Here we don't calculate Tax and PF so which is unnecessary.
// Either we have provide message saying "Not Applicable" which is not good design as per ISP
type DubaiPayroll struct {
	//...
}

func (I *DubaiPayroll) Earning(data Data) bool {
	fmt.Print("Earning Calculated for Dubai")
	return true
}

func (I *DubaiPayroll) Deduction(data Data) bool {
	fmt.Print("Deduction Calculated for Dubai")
	return true
}

func (I *DubaiPayroll) PF(data Data) bool {
	fmt.Print("PF Calculation not applicable for Dubai")
	return false
}

func (I *DubaiPayroll) Tax(data Data) bool {
	fmt.Print("Tax Calculation not applicable for Dubai")
	return false
}

/*
	Following ISP - Better Approach to handle above Scenario
	Split into Several Interfaces
	We have different types of Calculation in Payroll Application.
	  - Earning
	  - Deduction
	  - PF
	  - Tax
	Here for some country PF and Tax not application
*/

// Split all the Calculation into several interface
type IEarningCalculation interface {
	Earning(data Data) bool
}

type IDeductionCalculation interface {
	Deduction(data Data) bool
}

type IPFCalculation interface {
	PF(data Data) bool
}

type ITaxCalculation interface {
	Tax(data Data) bool
}

type IIndiaPayroll interface {
	IEarningCalculation
	IDeductionCalculation
	IPFCalculation
	ITaxCalculation
}

type ISPIndiaPayroll struct {

}

func (i ISPIndiaPayroll) Earning(data Data) bool {
	fmt.Print("Earning Calculated for India")
	return true
}

func (i ISPIndiaPayroll) Deduction(data Data) bool {
	fmt.Print("Deduction Calculated for India")
	return true
}

func (i ISPIndiaPayroll) PF(data Data) bool {
	fmt.Print("PF Calculated for India")
	return true
}

func (i ISPIndiaPayroll) Tax(data Data) bool {
	fmt.Print("Tax Calculated for India")
	return true
}

type IDubaiPayroll interface {
	IEarningCalculation
	IDeductionCalculation
}

type ISPDubaiPayroll struct {

}

func (d ISPDubaiPayroll) Earning(data Data) bool {
	fmt.Print("Earning Calculated for Dubai")
	return true
}

func (d ISPDubaiPayroll) Deduction(data Data) bool {
	fmt.Print("Deduction Calculated for Dubai")
	return true
}


// Main
func main() {
	/*Not Followed ISP*/
	indiaPayroll := IndiaPayroll{}
	fmt.Print(indiaPayroll.Earning(Data{}))
	fmt.Print(indiaPayroll.Deduction(Data{}))
	fmt.Print(indiaPayroll.PF(Data{}))
	fmt.Print(indiaPayroll.Tax(Data{}))

	dubaiPayroll := DubaiPayroll{}
	fmt.Print(dubaiPayroll.Earning(Data{}))
	fmt.Print(dubaiPayroll.Deduction(Data{}))
	fmt.Print(dubaiPayroll.PF(Data{}))
	fmt.Print(dubaiPayroll.Tax(Data{}))

	/*Followed ISP
	  We can't access PF and Tax calculation for Dubai Payroll.
	*/
	ispIndiaPayroll := ISPIndiaPayroll{}
	fmt.Print(ispIndiaPayroll.Earning(Data{}))
	fmt.Print(ispIndiaPayroll.Deduction(Data{}))
	fmt.Print(ispIndiaPayroll.PF(Data{}))
	fmt.Print(ispIndiaPayroll.Tax(Data{}))

	ispDubaiPayroll := ISPDubaiPayroll{}
	fmt.Print(ispDubaiPayroll.Earning(Data{}))
	fmt.Print(ispDubaiPayroll.Deduction(Data{}))
}