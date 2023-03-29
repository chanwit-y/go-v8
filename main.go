package main

import (
	"encoding/json"
	"fmt"

	v8 "rogchap.com/v8go"
)

func main() {
	ctx := v8.NewContext() // creates a new V8 context with a new Isolate aka VM
	ctx.RunScript(`const checkUser = (req) => {
		req.userId = 'xxxx'
		return req
	}`, "script.js") // executes a script on the global context
	ctx.RunScript(`
	const result = checkUser({'userId': 'dev-52'})`, "main.js") // any functions previously added to the context can be called
	val, _ := ctx.RunScript("result", "value.js") // return a value in JavaScript back to Go

	j, _ := json.Marshal(val)

	fmt.Printf("addition result: %+v", string(j))
}
