package services

type AdminService interface {
}

type adminService struct {
}

func NewAdminService() AdminService {
	return &adminService{}
}
