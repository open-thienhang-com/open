package service

type IService interface {
}


type Service struct {
}

func Init() IService {
	return &Service{
		
	}
}