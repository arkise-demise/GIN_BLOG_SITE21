package routing

import (
	"context"
	"fmt"
	"path"
	"ride-plus-v2/internal/constant/errors"
	"ride-plus-v2/internal/constant/state"
	"ride-plus-v2/internal/handler/middleware"
	"ride-plus-v2/platform/logger"
	"ride-plus-v2/platform/opa"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm/utils"
)

type Router struct {
	Method      string
	Path        string
	Handler     gin.HandlerFunc
	Middlewares []gin.HandlerFunc
	UnAuthorize bool
}

func RegisterRoute(
	grp *gin.RouterGroup,
	routes []Router,
	zapLogger logger.Logger,
) {
	for _, route := range routes {
		for _, domain := range route.Domain {
			exist := false
			var handler []gin.HandlerFunc

			var endpoint string

			switch domain.Name {
			case authDomains.System.Name:
				endpoint = path.Join("system", route.Path)
			case authDomains.Corporate.Name:
				endpoint = path.Join("corporate/:company_id", route.Path)
			case authDomains.Team.Name:
				endpoint = path.Join("team/:team_id", route.Path)
			case authDomains.User.Name:
				endpoint = route.Path
			default:
				zapLogger.Fatal(context.Background(), fmt.Sprintf("Invalid Domain %s Registered on Route", domain))
			}

			if !route.UnAuthorize {
				for _, permission := range route.Permission {
					if err := permission.Validate(); err != nil {
						err := errors.ErrInvalidUserInput.Wrap(err, "invalid input")
						zapLogger.Fatal(context.Background(), "invalid input", zap.Error(err), zap.Any("permission",
							permission))
						return
					}
					if utils.Contains(permission.Domains, string(domain.ID)) {
						exist = true

						//set permission context
						handler = append(handler, middleware.AddPermissionData(permission.Statement))
					}
				}
			}

			if !exist && !route.UnAuthorize {
				zapLogger.Fatal(
					context.Background(),
					"permission is not set for path: "+
						endpoint+
						" method "+
						route.Method)
			}
			handler = append(handler, middleware.AddDomainAndTenantInfo(domain))

			handler = append(handler, route.Middlewares...)
			handler = append(handler, route.Handler)
			grp.Handle(route.Method, endpoint, handler...)
		}
	}
}
