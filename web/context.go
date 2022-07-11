package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	Params map[string]string
	// response info
	StatusCode int
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// info read
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// info read
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

//info write
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

//info write
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// resp info format control
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// resp info format control
// 根据评论区此处有bug，若出现err!=nil将无法正确返回错误状态码
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("COntent-Type", "application/json")
	c.Status(code)
	// 这里修改一次StatusCode
	//  c.StatusCode = code
	//	c.Writer.WriteHeader(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
		// err != nil，这里需要再修改一次code，但是因为上文的缘故
		// 无法再次修改
		//  w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		//	w.Header().Set("X-Content-Type-Options", "nosniff")
		//	w.WriteHeader(code)
		//	fmt.Fprintln(w, error)
	}
}

// resp info format control
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// resp info format control
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
