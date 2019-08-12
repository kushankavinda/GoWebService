package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/webAPi/models"
	"github.com/webAPi/servises"
)

type userControllers struct {
	userIDPattern *regexp.Regexp
}

func (uc userControllers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else if r.URL.Path == "/biometricVerification" {
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte(servises.UserService()))
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		switch r.Method {
		case http.MethodGet:
			uc.get(id, w)
		}
	}
}
func (uc *userControllers) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetUsers(), w)
}
func (uc *userControllers) get(id int, w http.ResponseWriter) {
	u, err := models.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not find user name"))
		return
	}
	encodeResponseAsJSON(u, w)
}
func (uc *userControllers) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User Object"))
		return
	}
	u, err = servises.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}
func (uc *userControllers) parseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}
func newUserController() *userControllers {
	return &userControllers{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}

func (c *userControllers) ListAccounts(ctx *gin.Context) {
	/*	q := ctx.Request.URL.Query().Get("q")
		accounts, err := model.AccountsAll(q)
		if err != nil {
			httputil.NewError(ctx, http.StatusNotFound, err)
			return
		}*/
	ctx.JSON(http.StatusOK, servises.UserService())
}
