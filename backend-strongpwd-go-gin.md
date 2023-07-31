# Backend Strong Password Recommendation steps

## Problem

Build backend server with below API Spec and also store log on postgres database  
**Input**: password string  
1 <= password.length <= 40  
Password contains letters, digits, ‘.’ Dot, or ‘!’ Exclamation mark  
**Output**: minimum number of actions to make the password strong  
BASE URL /api/strong_password_steps  

## Example 1

Request

```json
{
 "init_password": "aA1"
}
```

Response

```json
{ 
 "num_of_steps": 3
}
```

## Example 2

Request

```json
{
 "init_password": "1445D1cd"
}
```

Response

```json
{ 
 "num_of_steps": 0
}
```

## Logic

### Criteria of Strong password

1. Password length >=6, < 20 characters.

2. Contains at least 1 lowercase letter, at least 1 uppercase letter, and at least 1 digit

3. Does not contain 3 repeating characters in a row e.g. 11123

### Type of actions

1. Add one character to password

2. Remove one character from password,

3. Replace one character of password with another character

## Tech stack

Go with Gin framework

## Bonus

1. Unit test
2. HTTPStatus(success, bad request)
3. Easy to maintain code structure

## Deliverables

Reply me via email with your git repo naming = {your_name}_agnos_backend_internship_2023 and README how to deploy on local and how to run unit test. However, the code does not need to be completed. (We also consider attempts)
