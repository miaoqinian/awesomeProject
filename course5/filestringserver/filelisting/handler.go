package filelisting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//const prefix = "/list/"
const prefix = "/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError(
			fmt.Sprintf("path %s mast start with %s", request.URL.Path, prefix))
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
