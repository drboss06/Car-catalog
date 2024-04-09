package handler

import (
	"carDirectory/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetCars получает список автомобилей с возможностью фильтрации по номеру и марке.
// @Summary Получение списка автомобилей
// @Description Получение списка автомобилей с возможностью фильтрации по номеру и марке.
// @Tags Cars
// @Accept json
// @Produce json
// @Param regNum query string false "Номер автомобиля для фильтрации"
// @Param mark query string false "Марка автомобиля для фильтрации"
// @Param page query string false "Номер страницы"
// @Param pageSize query string false "Размер страницы"
// @Success 200 {array} model.Car_model "Успешный ответ"
// @Failure 500  "Внутренняя ошибка сервера"
// @Router /cars [get]
func (h *Handler) GetCars(c *gin.Context) {
	l.Info().Msg("Get cars request")
	regNum := c.Query("regNum")
	mark := c.Query("mark")

	l.Debug().Msg(fmt.Sprintf("regNum: %s, mark: %s", regNum, mark))

	page := c.Query("page")
	pageSize := c.Query("pageSize")

	if page == "" {
		page = "1"
	}
	if pageSize == "" {
		pageSize = "10"
	}

	cars, err := h.service.Car.GetCars(regNum, mark, page, pageSize)

	l.Debug().Msg(fmt.Sprintf("cars: %v", cars))

	if err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	l.Info().Msg("Get cars response successfully")

	c.JSON(200, cars)
}

// DeleteCar удаляет автомобиль по его идентификатору.
// @Summary Удаление автомобиля
// @Description Удаляет автомобиль по его идентификатору.
// @Tags Cars
// @Accept json
// @Produce json
// @Param id path int true "Идентификатор автомобиля"
// @Success 200 "Car deleted successfully"
// @Failure 400 "Ошибка запроса"
// @Failure 500 "Internal server error"
// @Router /cars/{id} [delete]
func (h *Handler) DeleteCar(c *gin.Context) {
	l.Info().Msg("Delete request")

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	l.Debug().Msg(fmt.Sprintf("id: %d", idInt))

	if err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	if err := h.service.Car.DeleteCar(idInt); err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	l.Info().Msg("Car deleted successfully")

	c.JSON(200, gin.H{
		"message": "Car deleted successfully",
	})

}

// UpdateCar изменяет информацию об автомобиле по его идентификатору.
// @Summary Изменение информации об автомобиле
// @Description Изменяет информацию об автомобиле по его идентификатору.
// @Tags Cars
// @Accept json
// @Produce json
// @Param id path int true "Идентификатор автомобиля"
// @Param model.CarUpdate body model.CarUpdate true "Данные для обновления автомобиля"
// @Success 200 "Car updated successfully"
// @Failure 400 "Ошибка запроса"
// @Failure 500 "Internal server error"
// @Router /cars/{id} [put]
func (h *Handler) UpdateCar(c *gin.Context) {
	l.Info().Msg("Update car request")

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	carUpdate := model.CarUpdate{}

	l.Debug().Msg(fmt.Sprintf("model carUpdate: %v, id: %d", carUpdate, idInt))

	err = c.BindJSON(&carUpdate)

	if err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	if err := h.service.Car.UpdateCar(idInt, carUpdate); err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	l.Info().Msg("Car updated successfully")

	c.JSON(200, gin.H{
		"message": "Car updated successfully",
	})
}

// AddCar добавляет новый автомобиль.
// @Summary Добавление нового автомобиля
// @Description Добавляет новый автомобиль.
// @Tags Cars
// @Accept json
// @Produce json
// @Param model.CarAdd body model.CarAdd true "Данные нового автомобиля"
// @Success 200 "Car added successfully"
// @Failure 400 "Ошибка запроса"
// @Failure 500 "Internal server error"
// @Router /cars [post]
func (h *Handler) AddCar(c *gin.Context) {
	l.Info().Msg("Add car request")

	carAdd := model.CarAdd{}

	l.Debug().Msg(fmt.Sprintf("model carAdd: %v", carAdd))

	err := c.BindJSON(&carAdd)

	if err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	if err := h.service.Car.AddCar(carAdd); err != nil {
		newErrorResponse(c, 500, err.Error())
		return
	}

	l.Info().Msg("Car added successfully")

	c.JSON(200, gin.H{
		"message": "Car added successfully",
	})

}
