package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"

	pb "bidirectional_streaming/proto"
)

type TokenizerService struct{}

func (ts *TokenizerService) Tokenize(s pb.Tokenizer_TokenizeServer) error {
	for {
		req, err := s.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			fmt.Printf("Got an error receiving fron the client: %+v", err)
			return err
		}

		rdr := bufio.NewReader(bytes.NewReader(req.FileContents))
		scanner := bufio.NewScanner(rdr)
		scanner.Split(bufio.ScanWords)

		results := &pb.TokenizeResponse{
			Words: make(map[string]int64),
		}

		for scanner.Scan() {
			word := strings.TrimSpace(scanner.Text())

			_, ok := results.Words[word]
			if ok {
				results.Words[word]++
			} else {
				results.Words[word] = 1
			}
		}

		err = s.Send(results)
		if err != nil {
			fmt.Printf("Got an error sending to the client: %+v", err)
			return err
		}
	}
}
