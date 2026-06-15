# api-students
i'm create this api for improve my skills with Golang

## Routes
GET "/students" - List of all students
POST "/students" - Create Student
GET "/students/:id" - Get student by id
PUT "/students/:id" - Update student
DELETE "/students/:id" - Delete student

## Structs
Struct Student
Id
Name
Mail
CPF
Age
IsActive


## Payload example

{
    "name": "Heitor Oliveira",
    "cpf": "123456",
    "mail": "exemplo@exemplo.com",
    "age": 19,
    "isActive": false
}