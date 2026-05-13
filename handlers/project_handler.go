package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"tast-list.com/models"
	"tast-list.com/repositories"
)

type ProjectHandler struct {
	taskRepo repositories.ProjectRepository
}

func NewProjectHandler(taskRepo *repositories.TaskRepository) *TaskHandler {
	return &TaskHandler{
		taskRepo: *taskRepo,
	}
}

func (h *ProjectHandler) CreateProject(c fiber.Ctx) error {
	var task models.Task
	if err := c.Bind().Body(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"staus":   "error",
			"message": "Campos faltantes",
			"data":    err,
		})
	}

	if err := h.taskRepo.CreateProject(&task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"staus":   "error",
			"message": "No se pudo crear la tarea",
			"data":    err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"staus":   "success",
		"message": "No se pudo crear la tarea",
		"data":    task,
	})
}

func (h *ProjectHandler) GetAll(c fiber.Ctx) error {
	tasks, err := h.taskRepo.GetAll()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"staus":   "error",
			"message": "No se pudo encontrar las tareas",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"staus":   "success",
		"message": "Lista de tareas encontrada",
		"data":    tasks,
	})
}

func (h *ProjectHandler) GetById(c fiber.Ctx) error {
	id := c.Params("id")

	// 1. ParseUint devuelve (valor, error). Debes capturar ambos.
	idConvert, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "ID inválido",
		})
	}
	task, err := h.taskRepo.GetById(uint(idConvert))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"staus":   "error",
			"message": "No se pudo encontrar la tarea",
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"staus":   "success",
		"message": "Tarea encontrada",
		"data":    task,
	})
}

func (h *ProjectHandler) UpdateTaskTask(c fiber.Ctx) error {

	id := c.Params("id")

	// 1. ParseUint devuelve (valor, error). Debes capturar ambos.
	idConvert, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "ID inválido",
		})
	}

	findTask, err := h.taskRepo.GetById(uint(idConvert))

	if findTask == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"staus":   "error",
			"message": "No se pudo encontrar la tarea",
			"data":    nil,
		})
	}

	var task models.Task
	if err := c.Bind().Body(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"staus":   "error",
			"message": "Campos faltantes",
			"data":    err,
		})
	}

	if err := h.taskRepo.UpdateTask(&task); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"staus":   "error",
			"message": "No se pudo crear la tarea",
			"data":    err,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"staus":   "success",
		"message": "Tarea actualizada",
		"data":    task,
	})
}

func (h *ProjectHandler) DeleteTask(c fiber.Ctx) error {
	id := c.Params("id")

	// 1. ParseUint devuelve (valor, error). Debes capturar ambos.
	idConvert, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "ID inválido",
		})
	}
	task, err := h.taskRepo.GetById(uint(idConvert))

	if task == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Tarea no encontrada",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Tarea borrada correctamente",
		"data":    task,
	})
}
