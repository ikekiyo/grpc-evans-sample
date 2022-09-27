package handler

import (
	"context"
	"math/rand"
	"time"

	api "evans.sample/gen"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type EmployeeHandler struct {
	api.EmployeeServiceServer
}

func NewEmployeeHandler() *EmployeeHandler {
	return &EmployeeHandler{}
}

func employeeList() []*api.Employee {
	return []*api.Employee{
		{
			Id:   1,
			Name: "Yamada",
			Age:  20,
			Type: api.EmployeeType_FullTime,
		},
		{
			Id:   2,
			Name: "Suzuki",
			Age:  25,
			Type: api.EmployeeType_PartTime,
		},
		{
			Id:   3,
			Name: "Sasaki",
			Age:  45,
			Type: api.EmployeeType_FullTime,
		},
	}
}

func (h *EmployeeHandler) Employee(
	ctx context.Context,
	req *api.EmployeeRequest,
) (*api.EmployeeResponse, error) {
	employees := employeeList()
	employee := api.Employee{}
	for _, e := range employees {
		if e.Id == req.Id {
			employee = api.Employee{
				Id:   e.Id,
				Name: e.Name,
				Age:  e.Age,
				Type: e.Type,
			}
		}
	}
	return &api.EmployeeResponse{
		Employee: &employee,
	}, nil
}

func (h *EmployeeHandler) Employees(
	ctx context.Context,
	req *api.EmployeesRequest,
) (*api.EmployeesResponse, error) {
	employees := employeeList()
	return &api.EmployeesResponse{
		Employees: employees,
	}, nil
}
