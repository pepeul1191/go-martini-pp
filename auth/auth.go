package auth

import "io/ioutil"
import "net/http"
import "encoding/json"

import "github.com/martini-contrib/sessions"

type Account struct {
    Name     string
    Password string
}

type SuccessfullResponse struct {
    Result string
}

//
// Login handler
//
func Login(w http.ResponseWriter, r *http.Request, session sessions.Session) (int, string) {
    // read request body
    body, _ := ioutil.ReadAll(r.Body)
    // credentials of current user
    var credentials map[string]interface{}
    // decode json
    json.Unmarshal(body, &credentials)

    //
    // @TODO check user in the database
    //

    // save in session
    session.Set(credentials["username"], credentials["password"])

    // build response
    response, _ := json.Marshal(&SuccessfullResponse{Result : "ok"})

    // return response
    return 200, string(response)
}

//
// Logout handler
//
func Logout(w http.ResponseWriter, r *http.Request, session sessions.Session) (int, string) {
    // read request body
    body, _ := ioutil.ReadAll(r.Body)
    // credentials of current user
    var credentials map[string]interface{}
    // decode json
    json.Unmarshal(body, &credentials)
    // delete from sessions
    session.Delete(credentials["username"])
    // build response
    response, _ := json.Marshal(&SuccessfullResponse{Result : "ok"})
    // return response
    return 200, string(response) 
}
