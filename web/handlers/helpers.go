package handlers

import (
	"bytes"
	"fmt"
	pb "github.com/sgzmd/f3/web/gen/go/flibuserver/proto/v1"
	"github.com/sgzmd/go-telegram-auth/tgauth"
	"log"
	"net/http"
	"text/template"

	"github.com/davecgh/go-spew/spew"
)

const (
	ErrorMessage = `Error: ***** {{ .Err }} ******
	
Request: {{ .R }}
	`
)

type ErrorPage struct {
	Err error
	R   string
}

func ErrorToBrowser(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Add("Content-Type", "text/plain")
	t := template.Must(template.New("ErrorMessage").Parse(ErrorMessage))
	rstring := spew.Sdump(r)
	data := ErrorPage{Err: err, R: rstring}
	var buf bytes.Buffer
	e := t.Execute(&buf, &data)
	if e != nil {
		log.Fatalf("Failed to create error page: %+v", e)
	}

	w.Write(buf.Bytes())
}

func TrackedEntryUrl(entry *pb.TrackedEntry) string {
	url := "http://flibusta.is/"
	if entry.Key.EntityType == pb.EntryType_ENTRY_TYPE_SERIES {
		url += "s"
	} else if entry.Key.EntityType == pb.EntryType_ENTRY_TYPE_AUTHOR {
		url += "a"
	} else {
		return ""
	}

	return fmt.Sprintf("%s/%d", url, entry.Key.EntityId)
}

func MakeUserKey(ui *tgauth.UserInfo) string {
	return MakeUserKeyFromUserNameAndId(ui.UserName, ui.Id)
}

func MakeUserKeyFromUserNameAndId(userName string, userId int64) string {
	return fmt.Sprintf("%s-%d", userName, userId)
}
