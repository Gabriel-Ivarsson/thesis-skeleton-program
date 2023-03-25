package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Gabriel-Ivarsson/thesis-skeleton-program/example/mathCalc"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func routeAdd(w http.ResponseWriter, req *http.Request) {
	a := req.URL.Query().Get("a")
	b := req.URL.Query().Get("b")
	aInt, err := strconv.Atoi(a)
	if err != nil {
		return
	}
	bInt, err := strconv.Atoi(b)
	if err != nil {
		return
	}
	result := mathcalc.Addition(aInt, bInt)
	fmt.Println(result)
}

func routeSub(w http.ResponseWriter, req *http.Request) {
	a := req.URL.Query().Get("a")
	b := req.URL.Query().Get("b")
	aInt, err := strconv.Atoi(a)
	if err != nil {
		return
	}
	bInt, err := strconv.Atoi(b)
	if err != nil {
		return
	}
	result := mathcalc.Substract(aInt, bInt)
	fmt.Println(result)
}

func routeMul(w http.ResponseWriter, req *http.Request) {
	a := req.URL.Query().Get("a")
	b := req.URL.Query().Get("b")
	aInt, err := strconv.Atoi(a)
	if err != nil {
		return
	}
	bInt, err := strconv.Atoi(b)
	if err != nil {
		return
	}
	result := mathcalc.Multiplication(aInt, bInt)
	fmt.Println(result)
}

func routeDiv(w http.ResponseWriter, req *http.Request) {
	a := req.URL.Query().Get("a")
	b := req.URL.Query().Get("b")
	aInt, err := strconv.Atoi(a)
	if err != nil {
		return
	}
	bInt, err := strconv.Atoi(b)
	if err != nil {
		return
	}
	result := mathcalc.Division(aInt, bInt)
	fmt.Println(result)
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func Serve(port int) {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.HandleFunc("/add", routeAdd)
	http.HandleFunc("/multiplication", routeMul)
	http.HandleFunc("/division", routeDiv)
	http.HandleFunc("/substraction", routeSub)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
