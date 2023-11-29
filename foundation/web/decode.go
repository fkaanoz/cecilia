package web

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	ApplicationJSON = "application/json"
)

func DecodeJSONBody(r *http.Request, target interface{}) error {

	if r.Header.Get("Content-Type") != ApplicationJSON {
		return errors.New("wrong content-type")
	}

	reader := io.LimitReader(r.Body, 1024*1024)
	defer r.Body.Close()

	body, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	// TODO : Think about these error types.
	err = json.Unmarshal(body, target)

	jsonSyntaxErr := &json.SyntaxError{}

	switch {
	case errors.Is(err, io.EOF):

	case errors.As(err, jsonSyntaxErr):

	}

	return nil
}
