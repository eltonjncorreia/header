# Header

Setting common headers

```go

package main

import (
	"net/http"
)

func main() {
    
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {})

	http.ListenAndServe(":8000", Headers(mux))  // Call Header Middlewares
}
```
