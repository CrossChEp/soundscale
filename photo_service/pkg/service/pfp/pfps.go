package pfp

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"photo_service/pkg/config/path_config"
	"photo_service/pkg/logger"
	"photo_service/pkg/proto/photo_service_proto"
)

func Upload(request *photo_service_proto.PFPRequest) error {
	filePath, err := filepath.Abs(fmt.Sprintf("%s/%s.png", path_config.PFPPath, request.UserId))
	photo, err := base64.StdEncoding.DecodeString(request.Photo)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't decode photo. Details: %v", err))
		return err
	}
	err = os.WriteFile(filePath, photo, 0644)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't upload pfp!. Details: %v", err))
		return err
	}
	return nil
}

func Download(request *photo_service_proto.GetPFPReq) (string, error) {
	filePath, err := filepath.Abs(fmt.Sprintf("%s/%s.png", path_config.PFPPath, request.UserId))
	pfp, err := os.ReadFile(filePath)
	if err != nil {
		logger.ErrorLog(fmt.Sprintf("Error: couldn't read file. Details: %v", err))
		return "", err
	}
	return base64.StdEncoding.EncodeToString(pfp), nil
}
