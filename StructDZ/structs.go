package main

import "fmt"

type student struct {
	learningToLearn  int
	REST_API         REST_API
	DSandAlgorithms  DSandAlgorithms
	PortfolioProject PortfolioProject
	SQL              int
	Docker           int
}

type REST_API struct {
	create bool
	read   bool
	update bool
	delete bool
}

func NewRESTAPI() REST_API {
	return REST_API{
		create: true,
		read:   true,
		update: false,
		delete: false,
	}
}

func NewStudent(api REST_API, ds DSandAlgorithms) *student {
	return &student{
		REST_API:        api,
		DSandAlgorithms: ds,
	}
}

type DSandAlgorithms struct {
	arraysAndLists int
	linkedLists    int
	trees          int
	graphs         int
}

type PortfolioProject struct {
	ReactAndGolang string
	TelegramBot    string
	Microservices  string
}

func (s *student) GetKnowledge() string {
	return fmt.Sprintf("1) create=%v, 2) read=%v 3) update=%v 4) delete=%v", s.REST_API.create, s.REST_API.read, s.REST_API.update, s.REST_API.delete)
}

func main() {
	var student1 student
	//
	api := NewRESTAPI()
	student1.REST_API = api
	//
	student1.learningToLearn = 100
	//
	student1.SQL = 20
	student1.Docker = 1
	//
	learningSkill := student1.learningToLearn
	if learningSkill == 100 {
		fmt.Println("his learning skill boosted")
	}
	//

	rest := student1.GetKnowledge()
	dsa := DSandAlgorithms{arraysAndLists: 4, linkedLists: 0, trees: 0, graphs: 0}

	fmt.Println("Is true practiced with", rest)
	fmt.Println("Values for practiced arrays and lists, linked lists, trees, and graphs:", dsa)

	fmt.Printf("SQL knowledge is %v%%\n", student1.SQL)
	fmt.Printf("Docker knowledge is %v%%\n", student1.Docker)

	student1.PortfolioProject.ReactAndGolang = "in Process"
	student1.PortfolioProject.TelegramBot = "not started yet"
	student1.PortfolioProject.Microservices = "not started yet"
	if student1.PortfolioProject.ReactAndGolang == "done" {
		fmt.Println("You are ready to apply for job")
	} else {
		fmt.Println("Keep practice")
	}
}
