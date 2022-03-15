package file_utils

import (
	"bytes"
	"container/list"
	"context"
	"errors"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"os"
	"strings"
)

type FileOperator interface {
	ReadFile(string, string) (*[]byte, error)
	WriteFile(*[]byte, string, string) error
	DelFile(string, string) error
	DelFiles(string, *list.List) (*list.List, *list.List)
}

type BaseFileOperator struct {
	retryTimes int
}

func (self *BaseFileOperator) ReadFile(fileDir string, fileName string) (*[]byte, error) {

	return nil, nil
}
func (self *BaseFileOperator) WriteFile(fileBytes *[]byte, fileDir string, fileName string) error {

	return nil
}
func (self *BaseFileOperator) DelFile(fileDir string, fileName string) error {

	return nil
}
func (self *BaseFileOperator) DelFiles(fileDir string, fileNames *list.List) (*list.List, *list.List) {

	return nil, nil
}

type LocalFileOperator struct {
	BaseFileOperator
}

func (self *LocalFileOperator) ReadFile(fileDir string, fileName string) (*[]byte, error) {

	filePath := fileDir + "/" + fileName

	fileInfo, err := os.Stat(filePath)

	if err != nil {
		return nil, err
	}

	if fileInfo.IsDir() {
		return nil, errors.New("目标是一个目录，无法读取。")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	filesize := fileInfo.Size()
	byteBuffer := make([]byte, filesize)

	_, err = file.Read(byteBuffer)

	if err != nil {
		return nil, err
	}

	return &byteBuffer, err
}

func (self *LocalFileOperator) ScanFiles(dirPath string, recursion bool, suffix *[]string) (*list.List, error) {

	fileList := new(list.List)

	fileInfo, err := os.Stat(dirPath)

	if err != nil {
		return nil, err
	}

	if !fileInfo.IsDir() {
		return nil, errors.New("输入的路径不是一个目录，无法扫描。")
	}

	dir, err := os.OpenFile(dirPath, os.O_RDONLY, os.ModeDir)

	if err != nil {
		return nil, err
	}
	defer dir.Close()

	//读取目录
	rds, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	for _, fod := range rds {

		fodName := fod.Name()
		fodPath := dirPath + "/" + fod.Name()

		if fod.IsDir() {

			if fodName == "." || fodName == ".." {
				continue
			}

			if recursion {
				subDirFileList, subDirErr := self.ScanFiles(fodPath, recursion, suffix)

				if subDirErr == nil && subDirFileList != nil {
					fileList.PushBackList(subDirFileList)
				}
			}

		} else if suffix != nil && len(*suffix) > 0 {

			currentFileNameArray := strings.Split(fod.Name(), ".")

			currentSuffix := strings.ToLower(currentFileNameArray[len(currentFileNameArray)-1])

			for _, needSuffix := range *suffix {

				if currentSuffix == needSuffix {
					fileList.PushBack(fodPath)
				}

			}

		} else {

			fileList.PushBack(fodPath)
		}

	}

	return fileList, err

}

type RemoteFileOperator struct {
	BaseFileOperator
}

type DiskFileOperator struct {
	LocalFileOperator
}

type SambaFileOperator struct {
	RemoteFileOperator
}

type S3FileOperator struct {
	RemoteFileOperator
}

type MinioFileOperator struct {
	RemoteFileOperator
	bucket string
	client *minio.Client
}

func (self *MinioFileOperator) init(endpoint string, accessKeyID string, secretAccessKey string, useSSL bool) error {

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err == nil {
		self.client = client
	}

	return err
}

func (self *MinioFileOperator) ReadFile(bucket string, filePath string) (*[]byte, error) {

	if self.client == nil {
		return nil, errors.New("客户端未初始化。")
	}

	var object *minio.Object
	var objInfo minio.ObjectInfo
	var err error

	for i := 0; i < self.retryTimes; i++ {

		err = nil

		object, err = self.client.GetObject(context.Background(), bucket, filePath, minio.GetObjectOptions{})
		if err != nil {

			log.Println("read minio object error", bucket, filePath, "retry times", i, err)

			continue

		}

		objInfo, err = object.Stat()
		if err != nil {

			log.Println("read minio object info error", bucket, filePath, "retry times", i, err)

			continue
		}

		if err == nil {
			break
		}

	}

	if err != nil {
		return nil, err
	}

	fileBytes := make([]byte, objInfo.Size)

	object.Read(fileBytes)

	return &fileBytes, err
}

func (self *MinioFileOperator) WriteFile(fileBytes *[]byte, bucket string, filePath string) error {

	var err error

	for i := 0; i < self.retryTimes; i++ {

		err = nil

		r := bytes.NewReader(*fileBytes)

		_, err = self.client.PutObject(context.Background(), bucket, filePath, r, r.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})

		if err != nil {

			log.Println("write minio error", bucket, filePath, "retry times", i, err)

			continue

		} else {
			return err
		}
	}

	return err
}

func (self *MinioFileOperator) DelFile(bucket string, filePath string) error {

	opts := minio.RemoveObjectOptions{
		ForceDelete:      true,
		GovernanceBypass: true,
	}

	var err error

	for i := 0; i < self.retryTimes; i++ {

		err = nil

		err = self.client.RemoveObject(context.Background(), bucket, filePath, opts)

		if err != nil {

			log.Println("delete minio error", bucket, filePath, "retry times", i, err)

			continue

		} else {
			return err
		}

	}

	return err
}

func (self *MinioFileOperator) DelFiles(bucket string, fileNames *list.List) (*list.List, *list.List) {

	objectsCh := make(chan minio.ObjectInfo)

	for fileName := fileNames.Front(); fileName != nil; fileName = fileName.Next() {
		oi := minio.ObjectInfo{
			Key: fileName.Value.(string),
		}
		objectsCh <- oi
	}

	opts := minio.RemoveObjectsOptions{
		GovernanceBypass: true,
	}

	nameList := list.New()
	errList := list.New()

	for rErr := range self.client.RemoveObjects(context.Background(), bucket, objectsCh, opts) {

		nameList.PushBack(rErr.ObjectName)
		errList.PushBack(rErr)
	}

	return nameList, errList
}

func NewMinioFileOperator(endpoint string, accessKeyID string, secretAccessKey string, useSSL bool) (*MinioFileOperator, error) {

	mfo := new(MinioFileOperator)
	mfo.retryTimes = 3
	err := mfo.init(endpoint, accessKeyID, secretAccessKey, useSSL)

	if err != nil {
		mfo = nil
	}

	return mfo, err
}

type MultiObjectStorageFileOperator struct {
	BaseFileOperator
}
