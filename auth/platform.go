package auth

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/filebrowser/filebrowser/v2/errors"
	"github.com/filebrowser/filebrowser/v2/settings"
	"github.com/filebrowser/filebrowser/v2/users"
)

// MethodPlatformAuth is used to identify platform auth.
const MethodPlatformAuth settings.AuthMethod = "platform"

// PlatformAuth is a json implementaion of an Auther.
type PlatformAuth struct {
	Endpoint string `json:"endpoint" yaml:"endpoint"`
}

var sessionIDPattern = regexp.MustCompile(`sessionId(\s+)?=(\s+)?([a-fA-F0-9]+)`)

// Auth authenticates the user via user cookies with the proactive rest services
func (a PlatformAuth) Auth(r *http.Request, sto *users.Storage, root string) (*users.User, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rest/studio/currentuser", a.Endpoint), nil)
	if err != nil {
		return nil, err
	}

	if cookies, ok := r.Header["Cookie"]; ok {
		for _, value := range cookies {
			sessionIDVals := sessionIDPattern.FindStringSubmatch(value)
			if len(sessionIDVals) >= 4 {
				req.Header["sessionid"] = []string{sessionIDVals[3]}
				break
			}
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, os.ErrPermission
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	user, err := sto.Get(root, string(bodyBytes))
	if err == errors.ErrNotExist {
		return nil, os.ErrPermission
	}

	return user, nil
}

// LoginPage tells that proactive doesn't require a login page.
func (a PlatformAuth) LoginPage() bool {
	return false
}
