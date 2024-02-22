package handler

import (
	"context"
	"fmt"
	"log"

	model "github.com/grpc/gobus/service/pkg/Model"
	pb "github.com/grpc/gobus/service/pkg/PB"
	"github.com/grpc/gobus/service/pkg/service/interfaces"
)

type ProviderHandler struct {
	service interfaces.ProviderServiceInter
	pb.ProviderServicesServer
}

func (ph *ProviderHandler) ProviderLogin(ctx context.Context, pbl *pb.LoginRequest) (*pb.Result, error) {
	providerLogin := &model.ProviderLogin{
		Username: pbl.Username,
		Password: pbl.Password,
	}
	result, err := ph.service.ProviderLogin(providerLogin)
	if err != nil {
		log.Print("Error while provider logging in - Handler")
		response := &pb.Result{
			Status:  "Failed",
			Error:   "Provider Login Failed",
			Message: "",
		}
		return response, err
	}
	msg := fmt.Sprintf("%s - Provider successfully logged in.", result.Username)
	response := &pb.Result{
		Status:  "Success",
		Error:   "No Error",
		Message: msg,
	}
	return response, nil
}
func (ph *ProviderHandler) ProviderSignUp(ctx context.Context, pbd *pb.ProviderDataRequest) (*pb.Result, error) {
	providerData := &model.Provider{
		Username: pbd.UserName,
		Password: pbd.Password,
		Email:    pbd.Email,
	}
	result, err := ph.service.CreateProvider(providerData)
	if err != nil {
		log.Print("Error while provider signing up - Handler")
		response := &pb.Result{
			Status:  "Failed",
			Error:   "Provider Signup Failed",
			Message: "",
		}
		return response, err
	}
	msg := fmt.Sprintf("%s - Provider successfully Signed in.", result.Username)
	response := &pb.Result{
		Status:  "Success",
		Error:   "No Error",
		Message: msg,
	}
	return response, nil
}
func (ph *ProviderHandler) EditProvider(ctx context.Context, pbd *pb.ProviderDataRequest) (*pb.Result, error) {
	providerData := &model.Provider{
		Username: pbd.UserName,
		Password: pbd.Password,
		Email:    pbd.Email,
	}
	result, err := ph.service.EditProvider(providerData)
	if err != nil {
		log.Print("Error while editing provider info - Handler")
		response := &pb.Result{
			Status:  "Failed",
			Error:   "Edit Provider Failed",
			Message: "",
		}
		return response, err
	}
	msg := fmt.Sprintf("%s - ProviderInfo successfully edited.", result.Username)
	response := &pb.Result{
		Status:  "Success",
		Error:   "No Error",
		Message: msg,
	}
	return response, nil
}
func (ph *ProviderHandler) DeleteProviderById(ctx context.Context, pbi *pb.IdRequest) (*pb.Result, error) {
	result, err := ph.service.DeleteProviderById(uint(pbi.Id))
	if err != nil {
		log.Print("Error while Deleting provider - Handler")
		response := &pb.Result{
			Status:  "Failed",
			Error:   "Provider deletion Failed",
			Message: "",
		}
		return response, err
	}
	msg := fmt.Sprintf("%s - ProviderInfo successfully Deleted.", result.Username)
	response := &pb.Result{
		Status:  "Success",
		Error:   "No Error",
		Message: msg,
	}
	return response, nil
}
func (ph *ProviderHandler) FindProviderById(ctx context.Context, pbi *pb.IdRequest) (*pb.Result, error) {
	result, err := ph.service.FindProviderById(uint(pbi.Id))
	if err != nil {
		log.Print("Error while Finding provider - Handler")
		response := &pb.Result{
			Status:  "Failed",
			Error:   "Provider find Failed",
			Message: "",
		}
		return response, err
	}
	msg := fmt.Sprintf("Fetched Provider data successfully -- ProviderName=%s , ProviderEmail=%s , ProviderPassword= %s", result.Username,result.Email,result.Password)
	response := &pb.Result{
		Status:  "Success",
		Error:   "No Error",
		Message: msg,
	}
	return response, nil
}
func (ph *ProviderHandler) CreateProvider(ctx context.Context, pbd *pb.ProviderDataRequest) (*pb.Result, error) {
	providerData := &model.Provider{
		Username: pbd.UserName,
		Password: pbd.Password,
		Email:    pbd.Email,
	}
	result, err := ph.service.CreateProvider(providerData)
	if err != nil {
		log.Print("Error while adding provider - Handler")
		response := &pb.Result{
			Status:  "Failed",
			Error:   "Provider Addition Failed",
			Message: "",
		}
		return response, err
	}
	msg := fmt.Sprintf("%s - Provider successfully added.", result.Username)
	response := &pb.Result{
		Status:  "Success",
		Error:   "No Error",
		Message: msg,
	}
	return response, nil
}

func NewProviderHandler(svc interfaces.ProviderServiceInter) *ProviderHandler {
	return &ProviderHandler{
		service: svc,
	}
}
