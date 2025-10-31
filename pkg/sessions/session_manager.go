package sessions

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)
type SessionManager struct {
	sessionStore *sessions.CookieStore
	storeName  string
}
type SessionPayload struct{

	UserID uint 
	Email string
	IsAdmin bool 
}
func NewSessionManager( storeName string) *SessionManager {
	envSecret:=os.Getenv("SESSION_SECRET")
	if envSecret==""{
		envSecret="default-secret"
	}

	store := sessions.NewCookieStore([]byte(envSecret))
	return &SessionManager{
		sessionStore: store,
		storeName: storeName,
	}
}


func (sm *SessionManager) Set(r *http.Request,w http.ResponseWriter ,payload *SessionPayload)(error){
	session,_:=sm.sessionStore.Get(r,sm.storeName)
	session.Values["user_id"]=payload.UserID
	session.Values["email"]=payload.Email
	session.Values["is_admin"]=payload.IsAdmin
	err:=session.Save(r,w)
	return err;
}
func (sm *SessionManager) Get(r *http.Request) (*SessionPayload,error) {
	session,err:=sm.sessionStore.Get(r,sm.storeName)
	if err!=nil{
		return nil,err
	}
	if session.IsNew{
		return nil,nil
	}
	if session.Values["user_id"]==nil{
		return nil,nil
	}
	if session.Values["email"]==nil{
		return nil,nil
	}
	if session.Values["is_admin"]==nil{
		return nil,nil
	}
	userId:=session.Values["user_id"].(uint)
	email:=session.Values["email"].(string)
	isAdmin:=session.Values["is_admin"].(bool)
	payload:=&SessionPayload{UserID: userId,Email:email,IsAdmin: isAdmin}
	return payload,nil
}
func (sm *SessionManager) Destroy(r *http.Request, w http.ResponseWriter)(error){
	session,_:=sm.sessionStore.Get(r,sm.storeName)
	session.Options.MaxAge=-1
	err:=	session.Save(r,w)
	if err!=nil{
		return err
	}
	return nil;
}
