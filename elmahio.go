// Package elmahio - Log errors to
// Elmah.io from a Go(lang) web application
package elmahio

import (
        "bytes"
        "encoding/json"
        "errors"
        "log"
        "net/http"
        "net/url"
        "os"
        "strconv"
        "time"
)

// Settings struct
// This will set the
// APIKey and the LogID
// for writing on Elmah.io
type Settings struct {
        APIKey  string
        LogID   string
        Source  string
        Version float64
}

// Message struct
// This will be the message
// sent to Elmah.io
type Message struct {
        Application     string              `json:"application,omitempty"`
        Detail          string              `json:"detail,omitempty"`
        Hostname        string              `json:"hostname,omitempty"`
        Title           string              `json:"title,omitempty"`
        Source          string              `json:"source,omitempty"`
        StatusCode      int                 `json:"statuscode,omitempty"`
        DateTime        time.Time           `json:"datetime,omitempty"`
        Type            string              `json:"type,omitempty"`
        User            string              `json:"user,omitempty"`
        Severity        string              `json:"severity,omitempty"`
        URL             string              `json:"url,omitempty"`
        Method          string              `json:"method,omitempty"`
        Version         string              `json:"version,omitempty"`
        Cookies         map[string]string   `json:"cookies,omitempty"`
        Form            map[string][]string `json:"form,omitempty"`
        QueryString     url.Values          `json:"querystring,omitempty"`
        ServerVariables map[string]string   `json:"servervariables,omitempty"`
        Data            map[string]string   `json:"data,omitempty"`
}

// ElmahHandlerFunc type declaration
// This handler must be used to send
// erros on Elmah.io
type ElmahHandlerFunc func(w http.ResponseWriter, r *http.Request) (*http.Response, error)

// Settings Instance
var settings *Settings

// Setup Elmah.io settings
func Setup(APIKey string, LogID string) error {
        if len(APIKey) == 0 {
                return errors.New("Please specify an API Key")
        }
        if len(LogID) == 0 {
                return errors.New("Please specify a Log ID")
        }
        settings = &Settings{APIKey: APIKey, LogID: LogID}
        return nil
}

// SetVersion - Set the application version
func SetVersion(version float64) {
        settings.Version = version
}

// SetSource - Set the application source
func SetSource(source string) {
        settings.Source = source
}

// ElmahHandler middleware
// Will log the error message and send
// the details to Elmah.io
func ElmahHandler(handler ElmahHandlerFunc) http.Handler {
        return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
                httpResponse, err := handler(response, request)
                if err != nil {
                        log.Println(err.Error())
                        requestURL, urlErr := url.Parse(request.RequestURI)
                        if urlErr != nil {
                                log.Fatal(urlErr)
                        }
                        request.ParseForm()
                        requestQueryString, _ := url.ParseQuery(requestURL.RawQuery)
                        message := Message{
                                Hostname:    requestURL.Host,
                                URL:         requestURL.Path,
                                DateTime:    time.Now(),
                                Severity:    "Error",
                                StatusCode:  httpResponse.StatusCode,
                                Type:        err.Error(),
                                Method:      request.Method,
                                User:        requestURL.User.Username(),
                                QueryString: requestQueryString,
                                Application: os.Args[0],
                                Detail:      err.Error(),
                                Title:       err.Error(),
                                Cookies:     make(map[string]string, len(request.Cookies())),
                                Form:        make(map[string][]string, len(request.Form)),
                                Source:      "",
                                Version:     "",
                        }
                        if settings.Version > 0.0 {
                                message.Version = strconv.FormatFloat(settings.Version, 'f', -1, 64)
                        }
                        if len(settings.Source) > 0 {
                                message.Source = settings.Source
                        }
                        if len(request.Cookies()) > 0 {
                                for _, cookie := range request.Cookies() {
                                        message.Cookies[cookie.Name] = cookie.Value
                                }
                        }
                        if len(request.Form) > 0 {
                                for key, value := range request.Form {
                                        message.Form[key] = value
                                }
                        }
                        messageBytes := new(bytes.Buffer)
                        json.NewEncoder(messageBytes).Encode(message)
                        _, err = http.Post("https://api.elmah.io/v3/messages/"+settings.LogID+"?api_key="+settings.APIKey, "application/json; charset=utf-8", messageBytes)
                        if err != nil {
                                log.Fatal(err.Error())
                        }
                }
        })
}