package main

import "fmt"

func main() {
	var name string
	name = "Hello this is shreyas"
	fmt.Println(name)
	name = ` 
	This
	is 
	shreyas 
	from the  
	terminal`
	//The value within a raw string literal is not processed by GO so it will even print the new lines
	var dirIn string
	dirIn = `
		 <html>
	        <h3>This is just a sample heading</h>
	        <p>Hello This is shreyas from the terminal</p>
	        <body>
	          What's up
	        </body>
	      </html>`
	//Raw string will be useful when we are typing HTML code, as it makes it more readable
	fmt.Println(dirIn)
	fmt.Println(name)
}
