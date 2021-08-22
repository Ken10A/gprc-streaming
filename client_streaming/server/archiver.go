package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"

	pb "client_streaming/proto"
)

type ArchiverService struct{}

func (as *ArchiverService) Zip(stream pb.Archiver_ZipServer) error {
	buf := new(bytes.Buffer)
	zf := zip.NewWriter(buf)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			err = zf.Close()
			if err != nil {
				log.Print("Error creating the zip file: %v", err)
				return err
			}
			return stream.SendAndClose(&pb.ZipResponse{
				ZippedContents: buf.Bytes(),
			})
		}

		if err != nil {
			log.Print("Error reading the request: %v", err)
			return err
		}

		f, err := zf.Create(req.FileName)
		if err != nil {
			log.Print("Error creating the zip file entry: %v", err)
			return err
		}

		_, err = f.Write(req.Contents)
		if err != nil {
			fmt.Println("Error writing zip file: %v", err)
			return err
		}
	}
}
