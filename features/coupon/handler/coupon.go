package handler

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"user_barang/features/coupon"

	gojwt "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type CouponHandler struct {
	s coupon.Service
}

func New(s coupon.Service) coupon.Handler {
	return &CouponHandler{
		s: s,
	}
}

func (cc *CouponHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(CouponRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}

		file, err := c.FormFile("gambar")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		dst, err := os.Create("image/" + file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		var inputProcess = new(coupon.Coupon)
		inputProcess.NamaProgram = input.NamaProgram
		inputProcess.Link = input.Link
		inputProcess.Gambar = file.Filename

		result, err := cc.s.TambahCoupon(c.Get("user").(*gojwt.Token), *inputProcess)

		if err != nil {
			c.Logger().Error("ERROR Register, explain:", err.Error())
			var statusCode = http.StatusInternalServerError
			var message = "terjadi permasalahan ketika memproses data"

			if strings.Contains(err.Error(), "terdaftar") {
				statusCode = http.StatusBadRequest
				message = "data yang diinputkan sudah terdaftar ada sistem"
			}

			return c.JSON(statusCode, map[string]any{
				"message": message,
			})
		}

		var response = new(CouponResponse)
		response.NamaProgram = result.NamaProgram
		response.Link = result.Link
		response.Gambar = result.Gambar
		response.ID = result.ID
		response.UserID = result.UserID

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success create data",
			"data":    response,
		})
	}
}

func (cc *CouponHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		page, err := strconv.Atoi(c.QueryParam("page"))
		if page <= 0 {
			page = 1
		}
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		if limit <= 0 {
			limit = 5
		}
		results, err := cc.s.DapatCoupon(c.Get("user").(*gojwt.Token), page, limit)
		if err != nil {
			c.Logger().Error("Error fetching coupons: ", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Failed to retrieve coupon data",
			})
		}
		var response []CouponResponse
		for _, result := range results {
			response = append(response, CouponResponse{
				ID:          result.ID,
				NamaProgram: result.NamaProgram,
				Link:        result.Link,
				Gambar:      result.Gambar,
				UserID:      result.UserID,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":    "Success fetching all coupon data",
			"data":       response,
			"pagination": map[string]interface{}{"page": page, "limit": limit},
		})
	}
}

func (cc *CouponHandler) GetAllTanpa() echo.HandlerFunc {
	return func(c echo.Context) error {
		page, err := strconv.Atoi(c.QueryParam("page"))
		if page <= 0 {
			page = 1
		}
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		if limit <= 0 {
			limit = 5
		}
		results, err := cc.s.SemuaCoupon(page, limit)
		if err != nil {
			c.Logger().Error("Error fetching coupons: ", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Failed to retrieve coupon data",
			})
		}
		var response []CouponResponse
		for _, result := range results {
			response = append(response, CouponResponse{
				ID:          result.ID,
				NamaProgram: result.NamaProgram,
				Link:        result.Link,
				Gambar:      result.Gambar,
				UserID:      result.UserID,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":    "Success fetching all coupon data",
			"data":       response,
			"pagination": map[string]interface{}{"page": page, "limit": limit},
		})
	}
}
