package gobases4

import (
	"errors"
	"fmt"
)

type TaxableError struct {
	Err  error
	Code int
}

func (m *TaxableError) Error() string {
	return fmt.Sprintf("%s, Status code: %d", m.Err, m.Code)
}

type LowSalaryError struct {
	Err  error
	Code int
}

func (m *LowSalaryError) Error() string {
	return fmt.Sprintf("%s, Status code: %d", m.Err, m.Code)
}

func CalculateSalary(hoursWorked int, hourValue float32) (float32, error) {
	if hoursWorked < 80 {
		return 0, &TaxableError{Code: 1, Err: errors.New("Error: the worker cannot have worked less than 80 hours per month")}
	}
	salary := hourValue * float32(hoursWorked)

	if salary < 10000 {
		return salary, &LowSalaryError{Code: 3, Err: errors.New("Error: salary is less than 10000")}
	}
	if salary < 150000 {
		return salary, &TaxableError{Code: 2, Err: fmt.Errorf("Error: the minimum taxable amount is 150,000 and the salary entered is: %g", salary)}
	}

	return salary * 0.9, nil
}
