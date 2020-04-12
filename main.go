package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/timpark0807/restapi/handler"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig *oauth2.Config
)

func init() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://127.0.0.1:8000/callback",
		ClientID:     "156573644182-d7k95dbl1cj0j188d4mcui32oa5hsmcn.apps.googleusercontent.com",
		ClientSecret: "Nw42Ltddcm7zxCdC7tjBZbR2",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	var htmlIndex = `<html>
<body>
	<a href="/login">Google Log In</a>
</body>
</html>`
	fmt.Fprintf(w, htmlIndex)
}

var (
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	fmt.Fprintf(w, "Content: %s\n", content)
}

func getUserInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}

func main() {
	router := mux.NewRouter()

	// people
	router.HandleFunc("/api/tenant", handler.ListTenants).Methods("GET")
	router.HandleFunc("/api/tenant/{ss}", handler.GetTenant).Methods("GET")
	router.HandleFunc("/api/tenant", handler.CreateTenant).Methods("POST")
	router.HandleFunc("/api/tenant/{id}", handler.UpdateTenant).Methods("PUT")
	router.HandleFunc("/api/tenant/{ss}", handler.DeleteTenant).Methods("DELETE")

	// property
	router.HandleFunc("/api/property", handler.ListProperties).Methods("GET")
	router.HandleFunc("/api/property", handler.CreateProperty).Methods("POST")
	router.HandleFunc("/api/property/{_id}", handler.GetProperty).Methods("GET")
	router.HandleFunc("/api/property/{_id}", handler.UpdateProperty).Methods("PUT")
	router.HandleFunc("/api/property/{_id}", handler.DeleteProperty).Methods("DELETE")

	// login oAuth
	router.HandleFunc("/", handleMain)
	router.HandleFunc("/login", handleGoogleLogin)
	router.HandleFunc("/callback", handleGoogleCallback)
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
