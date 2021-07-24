package middlewares

import (
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gin-gonic/gin"
)

func TestSession(t *testing.T) {
	// _, filename, _, _ := runtime.Caller(0)
	// dir := path.Join(path.Dir(filename), "..")
	// err := os.Chdir(dir)
	// if err != nil {
	// 	panic(err)
	// }
	asserts := assert.New(t)
	{
		handler := Session("2333")
		asserts.NotNil(handler)
		asserts.NotNil(Store)
		asserts.IsType(emptyFunc(), handler)
	}

}

func emptyFunc() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
