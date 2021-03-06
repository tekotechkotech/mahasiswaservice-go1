package http

import (
	"net/http"
	"os"

	"github.com/apex/log"
	"github.com/tekotechkotech/mahasiswaservice-go1/internal/services"
	mhsConst "github.com/tekotechkotech/mahasiswaservice-go1/pkg/common/const"
	"github.com/tekotechkotech/mahasiswaservice-go1/pkg/dto"
	mhsErrors "github.com/tekotechkotech/mahasiswaservice-go1/pkg/errors"

	"github.com/labstack/echo"
)

type HttpHandler struct {
	service services.Services
}

func NewHttpHandler(e *echo.Echo, srv services.Services) {
	handler := &HttpHandler{
		srv,
	}
	e.GET("api/v1/latihan/ping", handler.Ping)
	e.POST("api/v1/latihan/mahasiswa-alamat", handler.SaveMahasiswaAlamat)
	e.PATCH("api/v1/latihan/mahasiswa", handler.UpdateMahasiswa)
	e.POST("api/v1/latihan/alamat", handler.SaveAlamatId)
	e.GET("api/v1/latihan/mahasiswa-alamat", handler.ShowAllMahasiswaAlamat)
}

// cek koneksi
func (h *HttpHandler) Ping(c echo.Context) error {

	version := os.Getenv("VERSION")
	if version == "" {
		version = "pong"
	}

	data := version

	return c.JSON(http.StatusOK, data)
}

// menyimpan
func (h *HttpHandler) SaveMahasiswaAlamat(c echo.Context) error {
	postDTO := dto.MahasiswaReqDTO{}
	if err := c.Bind(&postDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err := postDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = h.service.SaveMahasiswaAlamat(&postDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.SaveSuccess,
		Data:    nil,
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HttpHandler) SaveAlamatId(c echo.Context) error {
	postDTO := dto.AlamatIdReqDTO{}
	if err := c.Bind(&postDTO); err != nil { //bind = req ke variabel
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err := postDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = h.service.SaveAlamatId(&postDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.SaveSuccess,
		Data:    nil,
	}

	return c.JSON(http.StatusOK, resp)
}

// UPDATE
func (h *HttpHandler) UpdateMahasiswa(c echo.Context) error {
	postDTO := dto.UpadeMahasiswaNamaReqDTO{}
	if err := c.Bind(&postDTO); err != nil {
		log.Error(err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	err := postDTO.Validate()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	err = h.service.UpdateMahasiswaNama(&postDTO)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(getStatusCode(err), dto.ResponseDTO{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.SaveSuccess,
		Data:    nil,
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HttpHandler) ShowAllMahasiswaAlamat(c echo.Context) error {
	Data, _ := h.service.ShowAllMahasiswaAlamat()

	var resp = dto.ResponseDTO{
		Success: true,
		Message: mhsConst.GetDataSuccess,
		Data:    Data,
	}

	return c.JSON(http.StatusOK, resp)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case mhsErrors.ErrInternalServerError:
		return http.StatusInternalServerError
	case mhsErrors.ErrNotFound:
		return http.StatusNotFound
	case mhsErrors.ErrConflict:
		return http.StatusConflict
	case mhsErrors.ErrInvalidRequest:
		return http.StatusBadRequest
	case mhsErrors.ErrFailAuth:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
