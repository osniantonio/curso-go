package main

import "net/http"

func main() {

	// servidor 1
	mux := http.NewServeMux()

	// para adicionar rotas tenho HandleFunc e o Handle

	// criando um endpoint com função anônima
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello, World!"))
	// })

	// ou criando um endpoint chamando uma função criada
	mux.HandleFunc("/", HomeHandler)

	mux.Handle("/blog", blog{title: "My Blog!"})

	http.ListenAndServe(":8080", mux)

	// servidor 2
	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Osni!"))
	})
	http.ListenAndServe(":8081", mux2)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
