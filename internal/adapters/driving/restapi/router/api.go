// Package router ...
package router

import (
	"github.com/andriykusevol/aktemplategorm/internal/adapters/driving/restapi/middleware"
	"github.com/gin-gonic/gin"
)

func (r *router) RegisterAPI(app *gin.Engine) {

	// // Handle OPTIONS request
	// app.OPTIONS("/*prefligtrequest", func(c *gin.Context) {
	// 	// Set status to 200 OK with no response body
	// 	c.Status(http.StatusOK)
	// })

	// login := app.Group("login"){
	// 	login.POST("", )
	// }

	auth := app.Group("/auth")
	{
		auth.POST("", r.authHandler.Login)
	}

	adminlevel := app.Group("/adminlevel")
	{
		management := adminlevel.Group("management/country")
		{
			management.POST(":country_code_a2/region", r.administrativeregionHandler.AddRegion)
			management.PATCH(":country_code_a2/region/:region_code", r.administrativeregionHandler.AddRegion)
			management.DELETE(":country_code_a2/region/:region_code", r.administrativeregionHandler.AddRegion)

			management.POST(":country_code_a2/region/:region_code/subregion", r.administrativeregionHandler.AddRegion)
			management.PATCH(":country_code_a2/region/:region_code/subregion/:sub_region_code", r.administrativeregionHandler.AddRegion)
			management.DELETE(":country_code_a2/region/:region_code/subregion/:sub_region_code", r.administrativeregionHandler.AddRegion)

			management.POST(":country_code_a2/region/:region_code/subregion/:sub_region_code/adminlevel4", r.administrativeregionHandler.AddRegion)
			management.PATCH(":country_code_a2/region/:region_code/subregion/:sub_region_code/adminlevel4/:admin_level_4_code", r.administrativeregionHandler.AddRegion)
			management.DELETE(":country_code_a2/region/:region_code/subregion/:sub_region_code/adminlevel4/:admin_level_4_code", r.administrativeregionHandler.AddRegion)

			management.POST(":country_code_a2/region/:region_code/subregion/:sub_region_code/adminlevel4/city", r.administrativeregionHandler.AddRegion)
			management.PATCH(":country_code_a2/region/:region_code/subregion/:sub_region_code/adminlevel4/city/:id", r.administrativeregionHandler.AddRegion)
			management.DELETE(":country_code_a2/region/:region_code/subregion/:sub_region_code/adminlevel4/city/:id", r.administrativeregionHandler.AddRegion)

		}

		functional := adminlevel.Group("functional/country")
		{
			functional.GET(":country_code_a2/region", r.administrativeregionHandler.AddRegion)
			functional.GET(":country_code_a2/region/:region_code/subregion", r.administrativeregionHandler.AddRegion)
			functional.GET(":country_code_a2/region/:region_code/subregion/:sub_region_code", r.administrativeregionHandler.AddRegion)
			functional.GET(":country_code_a2/region/:region_code/city", r.administrativeregionHandler.AddRegion)
			functional.GET(":country_code_a2/region/:region_code/city/:id", r.administrativeregionHandler.AddRegion)

			functional.GET(":country_code_a2/subregion/byregion", r.administrativeregionHandler.AddRegion)
			functional.GET(":country_code_a2/region/:region_code/adminlevel4/bysubregion", r.administrativeregionHandler.AddRegion)
			functional.GET(":country_code_a2/city/byregion", r.administrativeregionHandler.AddRegion)
		}
	}

	pspfee := app.Group("/pspfee")
	{

		//--------------------------------------------------
		// New Version
		//--------------------------------------------------

		// g.Use(middleware.UserAuthMiddleware(a.auth,
		// 	middleware.AllowPathPrefixSkipper("/api/v1/pub"),
		// ))

		// Protected route.
		pspfee.POST("", middleware.AuthMiddleware(r.authRepo), r.mobileMoneyWithdrawalFeeHandler.CreatePSP)
		pspfee.GET(":id", r.mobileMoneyWithdrawalFeeHandler.GetPSP)

		// =========================================================
		// This is more hard to handle in backend because we do not use c.ShouldBind
		// But it is more explicitly and convenient in Postman and SwaggerUi.
		// It pops Pagination and Sort parameters from query, and treats all other
		// parameters as Filter fields (It also recognize arrays and do query accordigly)
		// PspQueryByMap:
		// 		//type Values map[string][]string
		// 		queryParams := c.Request.URL.Query()
		// 			domainQueryFilter, errObj := request.ToDomain_QueryMapFilter(queryParams,
		//											psp.ValidateFilterFields, psp.ValidateOrderFields)
		// 		h.feeSetApp.QueryFilterPSP(ctx, *domainQueryFilter)

		// This is more easy to tackle in backend (the c.ShouldbindQuery do a lot of job)
		// and we can define the "var queryfilter request.QueryFilter"
		// instead of just "queryParams := c.Request.URL.Query()" like in PspQueryByMap
		// But using json as a parameter requires to base64 encode string, and it is not obvious
		// what parameters available.
		// PspQueryByJson:
		//		var queryfilter request.QueryFilter
		//		c.ShouldBindQuery(&queryfilter)
		//			domainQueryFilter, apierr := queryfilter.ToDomain_QueryFilter(
		//											psp.ValidateFilterFields, psp.ValidateOrderFields)
		// 		h.feeSetApp.QueryFilterPSP(ctx, *domainQueryFilter)

		// From the Postman and SwaggerUi it looks exacly the same, but on backend it differs
		// from the PspQueryByMap in that, it requires to define request.SomeQueryParam for each request.
		// That is why I designed PspQueryByMap - to have a domainQueryFilter and build it right from request.
		// without keeping request.SomeQueryParam for each request.
		// PspQueryByRequest:
		// 		var pspqueryparams request.PSPQueryParam		// Separate QueryParam for each object
		//		c.ShouldBindQuery(&pspqueryparams)
		// 		domainQueryParam := pspqueryparams.ToDomain_PSPQueryParam()
		//  	h.feeSetApp.QueryPSP(ctx, *domainQueryParam)

		pspfee.GET("querypspmap", r.mobileMoneyWithdrawalFeeHandler.PspQueryByMap)
		pspfee.GET("querypspjson", r.mobileMoneyWithdrawalFeeHandler.PspQueryByJson)
		pspfee.GET("querypsprequest", r.mobileMoneyWithdrawalFeeHandler.PspQueryByRequest)

		pspfee.PATCH("patchbyid/:ID", middleware.AuthMiddleware(r.authRepo), r.mobileMoneyWithdrawalFeeHandler.PspPatchByID)
		pspfee.PATCH("patchbyarray", r.mobileMoneyWithdrawalFeeHandler.PspPatchByrray)
		pspfee.PATCH("patchbyquery", r.mobileMoneyWithdrawalFeeHandler.PspPatchByQuery)

		// =========================================================

		//pspfee.GET("", r.mobileMoneyWithdrawalFeeHandler.ListPSP)

		pspfee.DELETE("", r.mobileMoneyWithdrawalFeeHandler.DeletePSP)

		//Protected route.
		pspfee.POST("feeset", middleware.AuthMiddleware(r.authRepo), r.mobileMoneyWithdrawalFeeHandler.AddMobileProviderFeeSet)

		//Protected route.
		pspfee.POST("feerange", middleware.AuthMiddleware(r.authRepo), r.mobileMoneyWithdrawalFeeHandler.AddMobileProviderFeeRange)

		pspfee.GET("feerange", r.mobileMoneyWithdrawalFeeHandler.GetFeeSetRange)

		//pspfee.GET("calculate-single-fee/:psp_code/:amount", r.mobileMoneyWithdrawalFeeHandler.CalculateFeeForAmount)
		pspfee.GET("calculate-single-fee", r.mobileMoneyWithdrawalFeeHandler.CalculateFeeForAmount)
		pspfee.POST("calculate-bulk-fee", r.mobileMoneyWithdrawalFeeHandler.CalculateListFeeForAmount)

		//--------------------------------------------------
		//--------------------------------------------------

		set := pspfee.Group("set")
		{
			management := set.Group("management")
			{
				management.POST(":psp_code", r.mobileMoneyWithdrawalFeeHandler.AddMobileProviderFeeSet)
				management.GET(":psp_code", r.mobileMoneyWithdrawalFeeHandler.ListMobileProviderFeeSet)
				management.GET(":psp_code/:pspfeeset_id", r.mobileMoneyWithdrawalFeeHandler.GetMobileProviderFeeSet)
				management.DELETE(":psp_code/:pspfeeset_id", r.mobileMoneyWithdrawalFeeHandler.DeleteMobileProviderFeeSet)
			}
			functional := set.Group("functional")
			{
				functional.GET(":psp_code/:amount", r.mobileMoneyWithdrawalFeeHandler.CalculateFeeForAmount)
				functional.POST(":psp_code", r.mobileMoneyWithdrawalFeeHandler.CalculateListFeeForAmount)
			}
		}
		feerange := pspfee.Group("range")
		{
			management := feerange.Group("management")
			{
				management.POST(":pspfeeset_id", r.mobileMoneyWithdrawalFeeHandler.AddMobileProviderFeeRange)
				management.GET(":pspfeeset_id", r.mobileMoneyWithdrawalFeeHandler.ListFeeSetRange)
				management.GET(":pspfeeset_id/:feerange_id", r.mobileMoneyWithdrawalFeeHandler.GetFeeSetRange)
				management.PATCH(":pspfeeset_id/:feerange_id", r.mobileMoneyWithdrawalFeeHandler.PatchFeeRange)
				management.DELETE("pspfeeset_id/:feerange_id", r.mobileMoneyWithdrawalFeeHandler.DeleteMobileProviderFeeRange)
			}
		}
	}

}
