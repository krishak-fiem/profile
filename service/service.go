package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	profilemodels "github.com/krishak-fiem/models/go/profile"
	"github.com/krishak-fiem/profile/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedProfileServiceServer
}

var validate = validator.New()

func (s Server) UpdateProfile(ctx context.Context, message *pb.UpdateProfileMessage) (*pb.UpdateProfileResponse, error) {
	err := validate.Struct(message)
	if err != nil {
		return &pb.UpdateProfileResponse{}, status.Error(codes.Aborted, err.Error())
	}

	profile := new(profilemodels.Profile)
	profile.Email = message.Email
	profile.City = message.City
	profile.Name = message.Name
	profile.State = message.State
	profile.PhoneNumber = message.PhoneNumber
	profile.StreetAddress = message.StreetAddress
	profile.Pincode = message.Pincode

	err = profile.UpdateProfile()
	if err != nil {
		return &pb.UpdateProfileResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.UpdateProfileResponse{Payload: "Profile Updated.", Status: true}, nil
}

func (s Server) GetProfile(ctx context.Context, message *pb.GetProfileMessage) (*pb.ProfileResponse, error) {
	profile := new(profilemodels.Profile)
	profile.Email = message.Email
	err := profile.GetProfile()
	if err != nil {
		return &pb.ProfileResponse{}, status.Error(codes.NotFound, err.Error())
	}

	return &pb.ProfileResponse{
		Status:        true,
		Email:         profile.Email,
		StreetAddress: profile.StreetAddress,
		State:         profile.State,
		Name:          profile.Name,
		City:          profile.City,
		PhoneNumber:   profile.PhoneNumber,
		Pincode:       profile.Pincode}, nil
}
