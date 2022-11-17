package gin

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestServer(t *testing.T) {
	ctx := context.Background()

	srv := NewServer(
		WithAddress(":8800"),
	)

	srv.Use(gin.Recovery())
	srv.Use(gin.Logger())

	srv.GET("/login/*param", func(c *gin.Context) {
		if len(c.Params.ByName("param")) > 1 {
			c.AbortWithStatus(404)
			return
		}
		c.String(200, "Hello World!")
	})

	if err := srv.Start(ctx); err != nil {
		panic(err)
	}

	defer func() {
		if err := srv.Stop(ctx); err != nil {
			t.Errorf("expected nil got %v", err)
		}
	}()
}
