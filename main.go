package main

import (
	"fmt"
	"net/http"
)

func main(){
	myMux := http.NewServeMux()
	myMux.HandleFunc("/", rootPage)
	myMux.HandleFunc("/hello/", sayHello)
	http.ListenAndServe(":8080", myMux)
	fmt.Println("Hello, world.")
}

func rootPage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Showing your request %v \n", req)
	fmt.Println("\n\n... done ...\n")

}

func sayHello(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w,`
		<head>
			<h1>My Header</h>
		</head>

		<body>
			{{.}}
			<button>Hello there</button>
		</body>
	`)
	fmt.Println("Showed the webpage on port 8080")
}

