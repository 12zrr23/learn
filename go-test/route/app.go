package route
import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine{

	r := gin.Default()

	r.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	return r

}