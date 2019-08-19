package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

type LeetSpeaker struct {
	leetmap map[string]string
}
func (ls *LeetSpeaker) Init() {
	ls.leetmap = make(map[string]string)
	ls.leetmap["A"] = "4"
	ls.leetmap["B"] = "8"
	ls.leetmap["C"] = "("
	ls.leetmap["D"] = "|>"
	ls.leetmap["E"] = "3"
	ls.leetmap["F"] = "|="
	ls.leetmap["G"] = "6"
	ls.leetmap["H"] = "#"
	ls.leetmap["I"] = "!"
	ls.leetmap["J"] = "_|"
	ls.leetmap["K"] = "|<"
	ls.leetmap["L"] = "1"
	ls.leetmap["M"] = "|\\/|"
	ls.leetmap["N"] = "|\\|"
	ls.leetmap["O"] = "0"
	ls.leetmap["P"] = "|D"
	ls.leetmap["Q"] = "(,)"
	ls.leetmap["R"] = "|2"
	ls.leetmap["S"] = "$"
	ls.leetmap["T"] = "7"
	ls.leetmap["U"] = "|_|"
	ls.leetmap["V"] = "\\/"
	ls.leetmap["W"] = "\\/\\/"
	ls.leetmap["X"] = "}{"
	ls.leetmap["Y"] = "`/"
	ls.leetmap["Z"] = "2"
}
func (ls *LeetSpeaker) translateToLeet(message string) string {
	message = strings.ToUpper(message)
	result := ""
	for i := 0; i < len(message); i++{
		letter := string([]rune(message)[i])
		if val, ok := ls.leetmap[letter]; ok {
			result += val
		} else {
			result += letter
		}
	}
	return result
}
func httpleet(w http.ResponseWriter, r *http.Request){
	leetSpeaker := new(LeetSpeaker)
	leetSpeaker.Init()
	results := leetSpeaker.translateToLeet(r.FormValue("text"))
	fmt.Fprintf(w, results)
	fmt.Printf("%s was translated to %s\n", r.FormValue("text"), results)
}
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/leetspeak", httpleet).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
func main() {
	leetSpeaker := new(LeetSpeaker)
	leetSpeaker.Init()
	log.Print("Standing up webserver...")
	handleRequests()

}
