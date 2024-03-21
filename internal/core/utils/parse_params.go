package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPageParams(c *gin.Context, defaultPage, defaultSize int) (int, int, error) {
	page, err := ParseIntegerQueryParam(c, "page", defaultPage)
	if err != nil {
		return 0, 0, errors.New("invalid page input")
	}

	size, err := ParseIntegerQueryParam(c, "size", defaultSize)
	if err != nil {
		return 0, 0, errors.New("invalid size input")
	}

	return page, size, nil
}
func ParseIntegerQueryParam(c *gin.Context, paramName string, defaultValue int) (int, error) {
	param := c.Query(paramName)
	if param == "" {
		return defaultValue, nil
	}

	number, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}

	return number, nil
}
