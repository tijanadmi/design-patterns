package service

//import "log"

// singleton service
/*type IdServiceSingleton struct{
	idService *IdService
}

func (s *IdServiceSingleton) GetService() *IdService{
	 if s.idService == nil{
		log.Print("instantiation")
		s.idService = newIsService()
	 }
	 return s.idService
}
type IdService struct {
	counter int
}

func newIsService() *IdService {
	return &IdService{counter: 0}
}

func (s *IdService) Next() int {
	s.counter++
	return s.counter
}*/

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}

	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}