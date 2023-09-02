package route

import (
	"auth-services/controller"
	"auth-services/middleware"
	"auth-services/repository"
	"fmt"

	"github.com/casbin/casbin/v2"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes : all the routes are defined here
func SetupRoutes(db *gorm.DB) {
	httpRouter := gin.Default()

	httpRouter.Use(middleware.CORSMiddleware())

	// Initialize  casbin adapter
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize casbin adapter: %v", err))
	}

	// Load model configuration file and policy store adapter
	enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	enforcer.EnableAutoSave(true)
	if err != nil {
		panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	}

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	policyRepository := repository.NewPolicySkkRepository()
	userRepository := repository.NewUserSkkRepository()
	userController := controller.NewUserSkkController(userRepository, policyRepository)
	apiRoutes := httpRouter.Group("/api")
	{
		// apiRoutes.POST("/admin/register", userController.RegisterAdmin)
		//apiRoutes.POST("/signin", userController.Loginuserskk)
		apiRoutes.POST("/signin", userController.NewLoginuserskk)
		apiRoutes.POST("/otp", userController.SendOtp)
	}

	//kkpProtectedRoutes := apiRoutes.Group("/kkp")
	// historyRepository := repository.NewHistoryRepository()
	// historyController := controller.NewHistoryController(historyRepository)
	//{
	// pengajuanRekomendasiRepository := repository.NewPengajuanRekomendasiRepository()
	// pengajuanRekomendasiController := controller.NewPengajuanRekomendasiController(pengajuanRekomendasiRepository)
	// pengajuanRekomendasiProtectedRoutes := kkpProtectedRoutes.Group("/api/pengajuan/rekomendasi", middleware.AuthorizeJWT())
	// {
	// 	pengajuanRekomendasiProtectedRoutes.POST("/add", middleware.Authorize("r_satker", "kkp", "km_pengajuan_rekomendasi_kkp", "write", enforcer), pengajuanRekomendasiController.AddRekomendasi)
	// 	pengajuanRekomendasiProtectedRoutes.POST("/update", middleware.Authorize("r_satker", "kkp", "km_pengajuan_rekomendasi_kkp", "write", enforcer), pengajuanRekomendasiController.UpdateRecomendasi)
	// 	pengajuanRekomendasiProtectedRoutes.GET("/detail", middleware.Authorize("r_satker", "kkp", "km_pengajuan_rekomendasi_kkp", "read", enforcer), pengajuanRekomendasiController.DetailPengajuanRekomendasi)
	// 	pengajuanRekomendasiProtectedRoutes.GET("/list", middleware.Authorize("r_satker", "kkp", "km_pengajuan_rekomendasi_kkp", "read", enforcer), pengajuanRekomendasiController.ListRecomendPerUser)
	// 	pengajuanRekomendasiProtectedRoutes.GET("/history/list", middleware.Authorize("r_satker", "kkp", "km_pengajuan_rekomendasi_kkp", "read", enforcer), historyController.ListHistory)
	// }

	//}

	skkProtectedRoutes := apiRoutes.Group("/skk")
	{
		holdingRepository := repository.NewHoldingRepository()
		holdingController := controller.NewHoldingController(holdingRepository)
		holdingProtectedRoutes := skkProtectedRoutes.Group("/master/holding", middleware.AuthorizeJWT())
		{
			holdingProtectedRoutes.POST("/add", middleware.Authorize("r_skk", "skk", "m_holding", "write", enforcer), holdingController.AddHolding)
			holdingProtectedRoutes.GET("/list", middleware.Authorize("r_skk", "skk", "m_holding", "read", enforcer), holdingController.ListHolding)
			holdingProtectedRoutes.GET("/detail", middleware.Authorize("r_skk", "skk", "m_holding", "read", enforcer), holdingController.DetailHolding)
			holdingProtectedRoutes.PUT("/update", middleware.Authorize("r_skk", "skk", "m_holding", "write", enforcer), holdingController.UpdateHolding)
			holdingProtectedRoutes.GET("/total", middleware.Authorize("r_skk", "skk", "m_holding", "read", enforcer), holdingController.TotalHolding)
		}

		accountSkkRepository := repository.NewAccountSkkRepository()
		accountSkkController := controller.NewAccountSkkController(accountSkkRepository)
		accountProtectedRoutes := skkProtectedRoutes.Group("/master/account")
		{
			accountProtectedRoutes.GET("/list", accountSkkController.List)
			accountProtectedRoutes.GET("/detail", accountSkkController.DetailAccount)
			accountProtectedRoutes.POST("/add", accountSkkController.AddAccount)
			accountProtectedRoutes.PUT("/update", accountSkkController.UpdateAccount)
		}

		kkksRepository := repository.NewKkksRepository()
		kkksController := controller.NewKkksSkkController(kkksRepository)
		kkksProtectedRoutes := skkProtectedRoutes.Group("/master/kkks", middleware.AuthorizeJWT())
		{
			kkksProtectedRoutes.POST("/add", middleware.Authorize("r_skk", "skk", "m_kkks", "write", enforcer), kkksController.AddKkks)
			kkksProtectedRoutes.GET("/list", middleware.Authorize("r_skk", "skk", "m_kkks", "read", enforcer), kkksController.List)
			kkksProtectedRoutes.GET("/detail", middleware.Authorize("r_skk", "skk", "m_kkks", "read", enforcer), kkksController.DetailAccount)
			kkksProtectedRoutes.PUT("/update", middleware.Authorize("r_skk", "skk", "m_kkks", "write", enforcer), kkksController.UpdateKkks)
			kkksProtectedRoutes.GET("/total", middleware.Authorize("r_skk", "skk", "m_kkks", "read", enforcer), kkksController.TotalKkks)
			//kkksProtectedRoutes.GET("/edit", kkksController.EditKkks)
			kkksProtectedRoutes.POST("/edit", middleware.Authorize("r_skk", "skk", "m_kkks", "write", enforcer), kkksController.UpdateAccountKkks)
		}

		wkRepository := repository.NewWkRepository()
		wkController := controller.NewWkSkkController(wkRepository)
		wkProtectedRoutes := skkProtectedRoutes.Group("/master/wk", middleware.AuthorizeJWT())
		{
			wkProtectedRoutes.POST("/add", middleware.Authorize("r_skk", "skk", "m_wk", "write", enforcer), wkController.AddWk)
			wkProtectedRoutes.GET("/list", middleware.Authorize("r_skk", "skk", "m_wk", "read", enforcer), wkController.ListWk)
			wkProtectedRoutes.GET("/detail", middleware.Authorize("r_skk", "skk", "m_wk", "read", enforcer), wkController.DetailWk)
			wkProtectedRoutes.PUT("/update", middleware.Authorize("r_skk", "skk", "m_wk", "write", enforcer), wkController.EditWk)
			wkProtectedRoutes.GET("/total", middleware.Authorize("r_skk", "skk", "m_wk", "read", enforcer), wkController.TotalWk)
			wkProtectedRoutes.GET("/edit", middleware.Authorize("r_skk", "skk", "m_wk", "write", enforcer), wkController.GetByIdWk)
			wkProtectedRoutes.GET("/drop-holding", middleware.Authorize("r_skk", "skk", "m_wk", "read", enforcer), wkController.DropdownHodling)
			wkProtectedRoutes.GET("/drop-kkks", middleware.Authorize("r_skk", "skk", "m_wk", "read", enforcer), wkController.DropdownKkkks)
			wkProtectedRoutes.DELETE("/delete", middleware.Authorize("r_skk", "skk", "m_wk", "write", enforcer), wkController.DeleteWk)
		}

		policyskkRepository := repository.NewPolicySkkRepository()
		policyskkController := controller.NewPolicySkkController(policyskkRepository)
		policyskkProtectedRoutes := skkProtectedRoutes.Group("/role", middleware.AuthorizeJWT())
		{
			policyskkProtectedRoutes.POST("/add", middleware.Authorize("r_skk", "skk", "s_role", "write", enforcer), policyskkController.AddRoleSkk)
			policyskkProtectedRoutes.GET("/list", middleware.Authorize("r_skk", "skk", "s_role", "read", enforcer), policyskkController.ListRoleSkk)
			policyskkProtectedRoutes.GET("/detail", middleware.Authorize("r_skk", "skk", "s_role", "read", enforcer), policyskkController.DetailRoleSkk)
			policyskkProtectedRoutes.POST("/update", middleware.Authorize("r_skk", "skk", "s_role", "write", enforcer), policyskkController.UpdateRoleSkk)
			policyskkProtectedRoutes.DELETE("/delete", middleware.Authorize("r_skk", "skk", "s_role", "write", enforcer), policyskkController.DeleteRoleSkk)
		}

		userskkRepository := repository.NewUserSkkRepository()
		userskkController := controller.NewUserSkkController(userskkRepository, policyskkRepository)
		userskkProtectedRoutes := skkProtectedRoutes.Group("/user", middleware.AuthorizeJWT())
		{
			userskkProtectedRoutes.POST("/add", middleware.Authorize("r_skk", "skk", "s_user", "write", enforcer), userskkController.Adduserskk)
			userskkProtectedRoutes.GET("/list", middleware.Authorize("r_skk", "skk", "s_user", "read", enforcer), userskkController.Listuserskk)
			userskkProtectedRoutes.GET("/detail", middleware.Authorize("r_skk", "skk", "s_user", "read", enforcer), userskkController.Detailuserskk)
			userskkProtectedRoutes.PUT("/update", middleware.Authorize("r_skk", "skk", "s_user", "write", enforcer), userskkController.Updateuserskk)
			userskkProtectedRoutes.GET("/group", middleware.Authorize("r_skk", "skk", "s_user", "read", enforcer), userskkController.TypeUserRole)
			apiRoutes.POST("/send-otp", userskkController.SendOtpSoa)
			apiRoutes.POST("/validasi-otp", userskkController.ValidasiOtpSoa)
			apiRoutes.POST("/reset-pass", userskkController.ChangePasswordSKk)
		}

		profilkkRepository := repository.NewUserSkkRepository()
		profilkkController := controller.NewUserSkkController(profilkkRepository, policyskkRepository)
		profilkkProtectedRoutes := skkProtectedRoutes.Group("/profil", middleware.AuthorizeJWT())
		{
			profilkkProtectedRoutes.POST("/update-pass", profilkkController.UpdatePasswordSKk)
			profilkkProtectedRoutes.POST("/update-profil", profilkkController.UpdateProfil)
			profilkkProtectedRoutes.GET("/detail-profil", profilkkController.DetailuserskkProfil)
		}

		ratekkRepository := repository.NewRateRepository()
		ratekkController := controller.NewRateController(ratekkRepository)
		rateskkProtectedRoutes := skkProtectedRoutes.Group("/rate", middleware.AuthorizeJWT())
		{
			rateskkProtectedRoutes.GET("/list-lps", middleware.Authorize("r_skk", "skk", "d_sukubunga", "read", enforcer), ratekkController.ListLps)
			rateskkProtectedRoutes.GET("/detail-lps", middleware.Authorize("r_skk", "skk", "d_sukubunga", "read", enforcer), ratekkController.DetailLps)
			rateskkProtectedRoutes.POST("/update-lps", middleware.Authorize("r_skk", "skk", "d_sukubunga", "write", enforcer), ratekkController.UpdateLps)
			rateskkProtectedRoutes.GET("/list-libor", middleware.Authorize("r_skk", "skk", "d_sukubunga", "read", enforcer), ratekkController.ListLibor)
			rateskkProtectedRoutes.GET("/detail-libor", middleware.Authorize("r_skk", "skk", "d_sukubunga", "read", enforcer), ratekkController.DetailLibor)
			rateskkProtectedRoutes.POST("/update-libor", middleware.Authorize("r_skk", "skk", "d_sukubunga", "write", enforcer), ratekkController.UpdateLibor)
			rateskkProtectedRoutes.GET("/list-sofr", middleware.Authorize("r_skk", "skk", "d_sukubunga", "read", enforcer), ratekkController.ListSofr)
			rateskkProtectedRoutes.GET("/detail-sofr", middleware.Authorize("r_skk", "skk", "d_sukubunga", "read", enforcer), ratekkController.DetailSofr)
			rateskkProtectedRoutes.POST("/update-sofr", middleware.Authorize("r_skk", "skk", "d_sukubunga", "write", enforcer), ratekkController.UpdateSofr)
		}

		asrskkRepository := repository.NewAsrRepository()
		asrskkController := controller.NewAsrController(asrskkRepository)
		asrskkProtectedRoutes := skkProtectedRoutes.Group("/asr", middleware.AuthorizeJWT())
		{
			asrskkProtectedRoutes.GET("/card", middleware.Authorize("r_skk", "skk", "d_totaldanaasr", "read", enforcer), asrskkController.CardDashboard)
			asrskkProtectedRoutes.GET("/grafik-total", middleware.Authorize("r_skk", "skk", "d_totaldanaasr", "read", enforcer), asrskkController.GrafikTotalDanaAsr)
			asrskkProtectedRoutes.GET("/giro-depo", middleware.Authorize("r_skk", "skk", "d_totaldanaasr", "read", enforcer), asrskkController.GrafikGiroDepo)
			asrskkProtectedRoutes.GET("/tren-pp", middleware.Authorize("r_skk", "skk", "d_totaldanaasr", "read", enforcer), asrskkController.TrenPenerimaanPengeluaran)
			asrskkProtectedRoutes.GET("/tren-bunga", middleware.Authorize("r_skk", "skk", "d_totaldanaasr", "read", enforcer), asrskkController.TrenBunga)
		}
		rekeningsaldoRepository := repository.NewRekeningSaldoRepository()
		rekeningsaldoController := controller.NewRekeningSaldoController(rekeningsaldoRepository)
		rekeningsaldoProtectedRoutes := skkProtectedRoutes.Group("/rekening-saldo", middleware.AuthorizeJWT())
		{
			rekeningsaldoProtectedRoutes.GET("/list-wk", middleware.Authorize("r_skk", "skk", "d_daftarrekening", "read", enforcer), rekeningsaldoController.GetDropdownWk)
			rekeningsaldoProtectedRoutes.GET("/summary-rek", middleware.Authorize("r_skk", "skk", "d_daftarrekening", "read", enforcer), rekeningsaldoController.SummaryRekening)
			rekeningsaldoProtectedRoutes.GET("/list-table", middleware.Authorize("r_skk", "skk", "d_daftarrekening", "read", enforcer), rekeningsaldoController.ListTable)
			rekeningsaldoProtectedRoutes.GET("/excel", middleware.Authorize("r_skk", "skk", "d_daftarrekening", "read", enforcer), rekeningsaldoController.Excel)
			rekeningsaldoProtectedRoutes.GET("/csv", middleware.Authorize("r_skk", "skk", "d_daftarrekening", "read", enforcer), rekeningsaldoController.Csv)
		}

		daftartrxRepository := repository.NewDaftarTrxRepository()
		daftartrxController := controller.NewDaftarTrxController(daftartrxRepository)
		daftartrxProtectedRoutes := skkProtectedRoutes.Group("/daftar-trx", middleware.AuthorizeJWT())
		{
			daftartrxProtectedRoutes.GET("/list-account", middleware.Authorize("r_skk", "skk", "d_daftartransaksi", "read", enforcer), daftartrxController.GetDropdownAccount)
			daftartrxProtectedRoutes.GET("/list-wk", middleware.Authorize("r_skk", "skk", "d_daftartransaksi", "read", enforcer), daftartrxController.GetDropdownWk)
			daftartrxProtectedRoutes.GET("/table-trx", middleware.Authorize("r_skk", "skk", "d_daftartransaksi", "read", enforcer), daftartrxController.ListTableTrx)
			daftartrxProtectedRoutes.GET("/card-bunga", middleware.Authorize("r_skk", "skk", "d_daftartransaksi", "read", enforcer), daftartrxController.CardBunga)
			daftartrxProtectedRoutes.GET("/card-lain", middleware.Authorize("r_skk", "skk", "d_daftartransaksi", "read", enforcer), daftartrxController.CardLainLain)
			daftartrxProtectedRoutes.GET("/card-saldo", middleware.Authorize("r_skk", "skk", "d_daftartransaksi", "read", enforcer), daftartrxController.CardSaldo)
			daftartrxProtectedRoutes.GET("/card-penerimaan", middleware.Authorize("r_skk", "skk", "d_daftartransaksi", "read", enforcer), daftartrxController.CardPenerimaan)
			daftartrxProtectedRoutes.GET("/card-pengeluaran", middleware.Authorize("r_skk", "skk", "d_daftartransaksi", "read", enforcer), daftartrxController.CardPengeluaran)
			daftartrxProtectedRoutes.GET("/excel", middleware.Authorize("r_skk", "skk", "d_daftartransaksi", "read", enforcer), daftartrxController.ExportExcel)
			daftartrxProtectedRoutes.GET("/csv", middleware.Authorize("r_skk", "skk", "d_daftartransaksi", "read", enforcer), daftartrxController.ExportCsv)
		}
	}

	//httpRouter.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
	httpRouter.Run(":8081")
}
