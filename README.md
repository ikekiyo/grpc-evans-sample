# grpc-evans-sample

Sample the gRPC development using evans library

# evans library

https://github.com/ktr0731/evans

## How to evans

### REPL mode

```bash
evans -p 50051 -r

  ______
 |  ____|
 | |__    __   __   __ _   _ __    ___
 |  __|   \ \ / /  / _. | | '_ \  / __|
 | |____   \ V /  | (_| | | | | | \__ \
 |______|   \_/    \__,_| |_| |_| |___/

 more expressive universal gRPC client


proto.employee.EmployeeService@127.0.0.1:50051> package proto.employee

proto.employee@127.0.0.1:50051> service EmployeeService

proto.employee.EmployeeService@127.0.0.1:50051> call Employees
{
  "employees": [
    {
      "age": 20,
      "id": "1",
      "name": "Yamada"
    },
    {
      "age": 25,
      "id": "2",
      "name": "Suzuki",
      "type": "PartTime"
    },
    {
      "age": 45,
      "id": "3",
      "name": "Sasaki"
    }
  ]
}
```

## protocol build

generate codes with protocol buffers

```
make proto.build
```
