package main

import (
	//"errors"
	"fmt"
)

//type error interface {
//	Error() string
//}
//
//
//func Sqrt(f float64)(float64, error){
//	if f < 0{
//		return 0, errors.New("math: square root of negative number")
//	}
//
//	return 1, nil
//}

func main() {
	//_, err := Sqrt(-1)
	//if err != nil{
	//	fmt.Println(err)
	//}

	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	if _, errMsg := Divide(100, 0); errMsg != "" {
		fmt.Println("errorMsg is: ", errMsg)
	}
}

type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string {
	strFormat := `
	Cannot proceed, the divider is zero.
	divedee: %d
	diveder: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}

func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}
