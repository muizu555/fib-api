# Fibonacci API Documentation

## Overview

The Fibonacci API is designed to return the nth Fibonacci number. The Fibonacci sequence is a series where the next term is the sum of the previous two terms, starting with 1 and 1. For example, the sequence goes 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89...

<img width="297" alt="image" src="https://github.com/muizu555/fib-api/assets/109199972/cb0acfe8-0f3d-4205-96ff-c693d9e3d1e8">


## Architecture
### cmd層
main.go
アプリケーションのエントリーポイント。
### Domain層
ここにはフィボナッチ数を計算するロジック (fibonacci.go) とそのユニットテスト (fibonacci_test.go) が含まれます。
この層は他の層に依存しません。
### Usecase層
fibonacci_usecase.go
アプリケーションのユースケースを実装。Domain層に依存しますが、他の層には依存しません。(現状は...)
仮にデータベースとのやり取りが必要になるとrepository層などを追加し、その層に依存することになります。
### Handler層
fibonacci_handler.go
HTTPリクエストを処理し、Usecase層を呼び出します。
### Router層
router.go
ルーティングの設定を行います。

## Endpoints

### Get Fibonacci Number

**Endpoint:** `/fib`

**Method:** `GET`

**Description:** Returns the nth Fibonacci number.

#### Request Parameters

- `n` (required): The position in the Fibonacci sequence. It must be a positive integer.

#### Request Example

```sh
curl -X GET http://54.168.163.143/fib?n=6

DNSの不具合の影響が出ています curl -X GET http://cipher-vault.com/fib?n=6
```

## GET /fib?n=<number> (ex n=6)
n番目のフィボナッチ数を取得する
```
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


