# grpc-evans-sample

Sample the gRPC development using evans library

# evans library

https://github.com/ktr0731/evans

## protocol build

generate codes with protocol buffers

```
make proto.build
```

## How to evans

### command and option

| command | description |
| ------- | ----------- |
| repl    | REPL mode   |
| cli     | CLI mode    |

| option           | description                           |
| ---------------- | ------------------------------------- |
| --proto          | comma-separated proto file names      |
| --host           | gRPC server host                      |
| --port, -p       | gRPC server port (default "50051")    |
| --reflection, -r | use gRPC reflection (default "false") |

### REPL mode

```bash
$ evans -r repl

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

proto.employee.EmployeeService@127.0.0.1:50051>
```

```bash
proto.employee.EmployeeService@127.0.0.1:50051> show message
+-------------------------+
|         MESSAGE         |
+-------------------------+
| EmployeeRequest         |
| EmployeeResponse        |
| EmployeesByTypeRequest  |
| EmployeesByTypeResponse |
| EmployeesRequest        |
| EmployeesResponse       |
+-------------------------+

proto.employee.EmployeeService@127.0.0.1:50051> desc EmployeeRequest
+-------+------------+----------+
| FIELD |    TYPE    | REPEATED |
+-------+------------+----------+
| id    | TYPE_INT64 | false    |
+-------+------------+----------+
```

### CLI mode

```bash
echo '{"id": 1}' | evans -r cli proto.employee.EmployeeService.Employee
{
  "employee": {
    "age": 20,
    "id": "1",
    "name": "Yamada"
  }
```
