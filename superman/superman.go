package superman

import (
	"log"

	"github.com/iEricKoh/yuncv/yuncv"
)

func CreateInstance(username, password, softKey string) *Superman {
	r := &Superman{
		Username: username,
		Password: password,
		SoftKey:  softKey,
	}

	return r
}

func (r *Superman) Resolve(bt []byte) (string, error) {
	client := yuncv.Client{U: r.Username, P: r.Password, Id: r.SoftKey}

	result, _ := client.Send(bt, 0)

	log.Println("Captcha result...")
	log.Println(result)

	return result, nil
}
