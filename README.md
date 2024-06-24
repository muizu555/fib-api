# Fibonacci API Documentation

## Overview

The Fibonacci API is designed to return the nth Fibonacci number. The Fibonacci sequence is a series where the next term is the sum of the previous two terms, starting with 1 and 1. For example, the sequence goes 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89...

<img width="297" alt="image" src="https://github.com/muizu555/fib-api/assets/109199972/cb0acfe8-0f3d-4205-96ff-c693d9e3d1e8">


## Endpoints

### Get Fibonacci Number

**Endpoint:** `/fib`

**Method:** `GET`

**Description:** Returns the nth Fibonacci number.

#### Request Parameters

- `n` (required): The position in the Fibonacci sequence. It must be a positive integer.

#### Request Example

```sh
curl -X GET http://52.68.144.36:8080/fib?n=6
```

## GET /fib?n=<number> (ex n=6)
n番目のフィボナッチ数を取得する

Status Code: 200 OK
RESPONSE
```json
{
    "result": "8"
}

Status Code: 400 Bad Request
入力値nが文字列やマイナスといった範囲外の場合
RESPONSE
```json
{
    "message": "Invalid parameter 'n'",
    "status": "400"
}

Status Code: 500 Internal Server Error
RESPONSE
```json
{
    "message": "Internal Server Error",
    "status": "500"
}


