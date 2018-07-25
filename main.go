package main

import (
   	"app"
)

func main() {
 app := &app.App{}
 app.Init()
 app.Run(":8082")
 
 }

 