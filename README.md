# monokage
In-Memory Database Using Bloom Filter

## Installation
```
go get -u github.com/seipan/monokage
```

## Usage
```go
package main

import "github.com/seipan/monokage"

func main() {
	db := monokage.NewDB(100, 10)
	db.Insert([]byte("hello"))
	db.Insert([]byte("world"))

	ok := db.Check([]byte("hello"))
	if ok {
		println("hello is in the database")
	}else{
		println("hello is not in the database")
	}

}

```
## License
Code licensed under 
[the MIT License](https://github.com/seipan/monokage/blob/main/LICENSE).

## Author
[seipan](https://github.com/seipan).
