package src

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/marlon/test-api/gen"
	"github.com/marlon/test-api/utils"
)

func (r *MutationResolver) Upload(ctx context.Context, files []*gen.FileField) (resData interface{}, err error) {
	if len(files) <= 0 {
		return nil, fmt.Errorf("上传文件不能为空")
	}
	tx := gen.GetTransaction(ctx)
	unix := time.Unix(time.Now().Unix(), 0)
	fileDate := unix.Format("2006-01")

	err = os.MkdirAll("./uploads/"+fileDate, os.ModePerm)
	if err != nil {
		return resData, err
	}

	var forError error = nil
	uploadImage := []string{}

	for _, v := range files {
		hash := v.Hash
		if hash == "" {
			hash = utils.EncryptMd5(uuid.Must(uuid.NewV4()).String())
		}
		uploadFile := gen.UploadFile{}
		if err := tx.Select("id, name").Where("hash = ?", hash).First(&uploadFile).Error; err == nil {
			uploadImage = append(uploadImage, uploadFile.Name)
		} else {
			file := v.File
			filesuffix := strings.ToLower(path.Ext(file.Filename)) // 文件后缀名

			filePath := "./uploads/" + fileDate + "/" + uuid.Must(uuid.NewV4()).String() + filesuffix
			// filePath := "./uploads/" + fileDate + "/" + v.Filename
			saveFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
			defer saveFile.Close()

			if err == nil {
				uploadFile.ID = uuid.Must(uuid.NewV4()).String()
				uploadFile.Name = filePath
				uploadFile.Hash = hash
				tx.Create(&uploadFile)

				uploadImage = append(uploadImage, filePath)
				io.Copy(saveFile, file.File)
			} else {
				forError = err
				break
			}
		}
	}

	if forError != nil {
		return resData, forError
	}

	resData = uploadImage

	return resData, err
}
