- It does not implement a primary port (primary port is implemented by application layer itself)
  But it is an adapter. The primary port exposes methods, that the driving adapter (this) will call.
  

- The driving adapter(e.g, HTTP/gRPC/GraphQl controller)
  interacts with the application service through the primary port

    // internal/adapters/driving/restapi/handler/auth_handler.go

        type AuthHandler interface {
            Login(c *gin.Context)
        }

        type authImpl struct {
            impl pport.Auth
        }

        // We could do it like this, and just substitute it
        // in the di_app.go, but in this case wi'll lose
        // the AuthHandler interface restriction.
        func NewAuthSimple(authApp pport.Auth) *authImpl {
            return &authImpl{impl: authApp}
        }

        // But we also want to define the interface
        // JUST TO MAKE SURE THAT WE'RE IMPLEMENTED ALL
        // HANDLER METHODS.
        // IT DOES NOT IMPACT ON OTHER LAYERS.
        func NewAuth(authApp pport.Auth) AuthHandler {
            return &authImpl{impl: authApp}
        }

    // internal/application/di_app.go
            authHandler := handler.NewAuthSimple(authApp)

        // Now just inject handlers to router.
        routerRouter := router.NewRouter(
            authRepo,
            authHandler,
            administrativeregionHandler,
            feeSetHandler,
        )

    // But indeed, we expect an interface in router (not just an implementation(structure))
    // it is because we can return a structure instead of interface
    // Any type that implements the specified methods is said to satisfy the interface.
        // There are two reasons to use functions with receivers:
	        // You will invoke the method through an interface type.
            // You really like the method-call syntax.

    // internal/adapters/driving/restapi/router/router.go
        // NewRouter ...
        func NewRouter(
            authRepo auth_domain.Repository,
            authHandler handler.AuthHandler,                    // you see, indeed we expect an interface here,
                                                                // but if we're just return a structure that implements that interface 
                                                                // it'll work also
            administrativeregionHandler handler.AdminRegion,
            mobileMoneyWithdrawalFeeHandler handler.PSPFee,
        )
