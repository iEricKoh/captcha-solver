package ruokuai

import (
	"os"

	"github.com/buger/jsonparser"
	"github.com/iEricKoh/httpClient"
)

const url = "http://api.ruokuai.com/create.json"

func CreateInstance(username, password, softKey string) *Ruokuai {
	r := &Ruokuai{
		Username: username,
		Password: password,
		SoftKey:  softKey,
	}

	return r
}

func (r *Ruokuai) Resolve(typeId, softId string, pic *os.File) (string, error) {
	form := &httpClient.Form{
		"username": r.Username,
		"password": r.Password,
		"typeid":   typeId,
		"softid":   softId,
		"softkey":  r.SoftKey,
		"image":    pic,
	}

	resp, err := httpClient.Post(url, &httpClient.Options{
		Header: &httpClient.Header{
			"Content-Type": "multipart/form-data",
		},
		Form: form,
	})

	result, err := jsonparser.GetString(resp.Body, "Result")

	if err != nil {
		return "", err
	}

	return result, nil
}
