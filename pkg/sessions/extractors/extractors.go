package extractors

import (
	"pasteBin/pkg/sessions"

	"github.com/gin-gonic/gin"
)
func ExtractUserSessionPayload(c *gin.Context) ( *sessions.SessionPayload,bool ){
 currentUser,exists:= c.Get("currentUser")
	if !exists{
		return nil,false;
	}
	parsedUser:=currentUser.(*sessions.SessionPayload)
	return parsedUser,true;

}
