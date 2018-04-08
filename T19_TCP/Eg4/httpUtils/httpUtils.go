////////////////////////////////////////////////////////////////////////////////
// WORK IN PROGRESS
// Experimental code for an http server without using the net.http package.
// This is just as a fun challenge and golang practice.
////////////////////////////////////////////////////////////////////////////////

package httpUtils

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

//////////////////////////////////////////
// CONSTANTS
//////////////////////////////////////////

const (
	//httpMethodOptions = "OPTIONS"
	CHTTPMethodGet = "GET"
	/*httpMethodHead    = "HEAD"
	httpMethodPost    = "POST"
	httpMethodPut     = "PUT"
	httpMethodDelete = "DELETE"
	httpMethodTrace = "TRACE"
	httpMethodConnect = "CONNECT"*/
)

// CHTTPMsgLineDelimeter represents the default delimeter
// used to separate parts of an HTTP message.
const CHTTPMsgLineDelimeter = "\r\n"

//////////////////////////////////////////
// STRUCT - HTTP REQUEST
//////////////////////////////////////////

// HTTPRequest is a structure for representing HTTP requests.
type HTTPRequest struct {
	// Request-Line
	Method      string
	RequestURI  string
	HTTPVersion string
	// Headers
	// TODO: expand this into more detail.
	Headers map[string]string
	// Message Body
	Body string
}

// Constructor provides default setup of HTTPRequest attributes.
func (r *HTTPRequest) Constructor() {
	r.Headers = make(map[string]string)
}

// SetRequestLine sets the request line variables of a given HTTPRequest.
func (r *HTTPRequest) SetRequestLine(httpMethod, requestURI, httpVersion string) {
	r.Method = httpMethod
	r.RequestURI = requestURI
	r.HTTPVersion = httpVersion
}

// AddHeader adds a new header to a given HTTPRequest.
func (r *HTTPRequest) AddHeader(key, value string) {
	r.Headers[key] = value
}

// SetBody sets the body content for a given HTTPRequest.
func (r *HTTPRequest) SetBody(body string) {
	r.Body = body
}

//////////////////////////////////////////
// STRUCT - HTTP RESPONSE
//////////////////////////////////////////

// HTTPResponse is a structure for representing HTTP responses.
type HTTPResponse struct {
	// Status-Line
	HTTPVersion  string
	StatusCode   int
	ReasonPhrase string
	// Headers
	// TODO: expand this into more detail.
	Headers map[string]string
	// Message Body
	Body string
}

// Constructor provides default initialisation for HTTP Response attributes.
func (r *HTTPResponse) Constructor() {
	r.SetStatusLine("HTTP/1.1", 500, "Server error")
	r.Headers = make(map[string]string)
}

// ConstructorWithStatusLine provides default initialisation for HTTP Response attributes.
// It is a helper for setting the status-line staight away.
func (r *HTTPResponse) ConstructorWithStatusLine(httpVersion string, statusCode int, reasonPhrase string) {
	r.SetStatusLine(httpVersion, statusCode, reasonPhrase)
	r.Headers = make(map[string]string)
}

// GetStatusLineAsString returns the status line as one whole string.
func (r *HTTPResponse) GetStatusLineAsString() string {
	return fmt.Sprintf("%s %d %s", r.HTTPVersion, r.StatusCode, r.ReasonPhrase)
}

// SetStatusLine setter for the status-line.
func (r *HTTPResponse) SetStatusLine(httpVersion string, statusCode int, reasonPhrase string) {
	r.HTTPVersion = httpVersion
	r.StatusCode = statusCode
	r.ReasonPhrase = reasonPhrase
}

// SetStatusCode setter for the status code.
func (r *HTTPResponse) SetStatusCode(statusCode int, reasonPhrase string) {
	r.StatusCode = statusCode
	r.ReasonPhrase = reasonPhrase
}

// GetHeadersAsStringBlock returns the headers separated by a line delimeter.
// The final header does not add on a line delimeter.
func (r *HTTPResponse) GetHeadersAsStringBlock() string {
	retStr := ""
	pairCounter := 0
	totalNumPairs := len(r.Headers)
	for k, v := range r.Headers {
		retStr += (k + ": " + v)
		if pairCounter < totalNumPairs {
			retStr += CHTTPMsgLineDelimeter
		}
		pairCounter++
	}
	return retStr
}

// AddHeader adds a header to the given HTTP Response.
func (r *HTTPResponse) AddHeader(key, value string) {
	r.Headers[key] = value
}

// SetBody sets the body content of a given HTTP Response.
func (r *HTTPResponse) SetBody(body string) {
	r.Body = body
}

//////////////////////////////////////////
// HELPER FUNCTIONS - HTTP:
//////////////////////////////////////////

// ParseHTTPRequest is a helper function to create
// an instance of HTTPRequest from a given Reader.
func ParseHTTPRequest(reader io.Reader) HTTPRequest {
	retData := HTTPRequest{}
	retData.Constructor()
	scanner := bufio.NewScanner(reader)
	lineIdx := 0
	for scanner.Scan() {
		ln := scanner.Text()
		if lineIdx == 0 {
			// Parse the request line.
			requestLineTokens := strings.Fields(ln)
			retData.SetRequestLine(
				requestLineTokens[0],
				requestLineTokens[1],
				requestLineTokens[2])
		} else if ln != "" {
			// Parse in headers.
			tokenList := strings.Split(ln, ": ")
			if len(tokenList) == 2 {
				retData.AddHeader(tokenList[0], tokenList[1])
			}
		} else if ln == "" {
			// Headers done.
			// TODO: Detect if we have a message body.
			break
		}
		lineIdx++
	}
	return retData
}

// SendHTTPResponse is a helper function to write a given HTTPResponse
// to the provided Writer. The Writer is normally an instance of net.Conn.
func SendHTTPResponse(writer io.Writer, response HTTPResponse) {
	responseAsStr := ""
	responseAsStr += response.GetStatusLineAsString()
	responseAsStr += CHTTPMsgLineDelimeter
	responseAsStr += response.GetHeadersAsStringBlock()
	responseAsStr += CHTTPMsgLineDelimeter
	responseAsStr += CHTTPMsgLineDelimeter
	if response.Body != "" {
		responseAsStr += response.Body
	}
	fmt.Fprint(writer, responseAsStr)
}

// GetResourceFileContents is a helper function for getting the entire
// contents of a file as a string.
func GetResourceFileContents(resourceRelativePath string) (string, bool) {
	fileDataAsBytes, err := ioutil.ReadFile(resourceRelativePath)
	if err != nil {
		return "", false
	}
	retStr := string(fileDataAsBytes)
	return retStr, true
}

//////////////////////////////////////////
