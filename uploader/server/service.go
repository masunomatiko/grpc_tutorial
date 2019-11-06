package main

import (
	"io"
	"log"
	"path/filepath"
	"io/ioutil"

	pb "github.com/masunomatiko/grpc_tutorial/uploader/proto"
)

type fileService struct{}

func (s *fileService) Upload(stream pb.FileService_UploadServer) error {
	var blob []byte
	var name string
	for {
		c, err := stream.Recv()
		if err == io.EOF {
			log.Printf("donw %d bytes\n", len(blob))
			break
		}
		if err != nil {
			panic(err)
		}
		name = c.GetName()
		blob = append(blob, c.GetData()...)
	}
	fp := filepath.Join("./uploader/resource", name)
	ioutil.WriteFile(fp, blob, 0644)
	stream.SendAndClose(
		&pb.FileResponse{
			Size: int64(len(blob)),
		},
	)
	return nil
}
