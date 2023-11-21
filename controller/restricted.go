package controller

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/RehanAfridikkk/API-Authentication/structure"
	"github.com/RehanAfridikkk/word-count-Echo-API-fileupload/cmd"
	
	"strconv"
)

func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*structure.JwtCustomClaims)
	name := claims.Name
 
	
	file, err := c.FormFile("file")
	if err != nil {
	   return c.JSON(http.StatusBadRequest, echo.Map{
		  "error": "File not provided",
	   })
	}
 
	routines, err := strconv.Atoi(c.FormValue("routines"))
	if err != nil {
	   return c.JSON(http.StatusBadRequest, echo.Map{
		  "error": "Invalid value for routines",
	   })
	}
 

	fileContent, err := file.Open()
	if err != nil {
	   return c.JSON(http.StatusInternalServerError, echo.Map{
		  "error": "Failed to open file",
	   })
	}
	defer fileContent.Close()
 
	totalCounts, _, runTime, err := cmd.ProcessFile(fileContent, routines)
	if err != nil {
	   return c.JSON(http.StatusInternalServerError, echo.Map{
		  "error": "Failed to process file",
	   })
	}
 
	
	return c.JSON(http.StatusOK, echo.Map{
	   "name":              name,
	   "lineCount":         totalCounts.LineCount,
	   "wordsCount":        totalCounts.WordsCount,
	   "vowelsCount":       totalCounts.VowelsCount,
	   "punctuationCount":  totalCounts.PunctuationCount,
	   "runTime":           runTime.String(),
	})
 }
 