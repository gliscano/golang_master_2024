package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	name        string
	description string
	completed   bool
}

type TaskList struct {
	tasks []Task
}

// Método para agregar tarea
func (list *TaskList) addTask(t Task) {
	list.tasks = append(list.tasks, t)
}

// Método para marcar como completada una tarea
func (list *TaskList) markAsCompleted(index int) {
	list.tasks[index].completed = true
}

// Método para editar una tarea
func (list *TaskList) editTask(index int, task Task) {
	list.tasks[index] = task
}

// Método para eliminar una tarea
func (list *TaskList) deleteTask(index int) {
	list.tasks = append(list.tasks[:index], list.tasks[index+1:]...)
}

func main() {
	// Instancia de Task
	list := TaskList{}

	// Instancia de Bufio para entrada de datos
	// newReader crea un nuevo lector implementado en io.Reader
	// os.Stdin, representa la entrada estandar como el teclado
	read := bufio.NewReader(os.Stdin)
	for {
		var option int
		fmt.Println("Seleccion una opción\n",
			"1. Agregar tarea\n",
			"2. Marcar tarea como completada\n",
			"3. Editar tarea\n",
			"4. Eliminar tarea\n",
			"5. Salir")

		fmt.Print("Ingrese una opción: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			var task Task
			fmt.Print("Ingrese el nombre de la tarea: ")
			task.name, _ = read.ReadString('\n')
			fmt.Print("Ingrese la descripción de la tarea: ")
			task.description, _ = read.ReadString('\n')
			list.addTask(task)
			fmt.Println("Tarea agregada correctamente")
		case 2:
			var index int
			fmt.Print("Ingrese el índice de la tarea a marcar como completada:")
			fmt.Scan(&index)
			list.markAsCompleted(index)
			fmt.Println("Tarea marcada como completada")
		case 3:
			var index int
			var task Task
			fmt.Print("Ingrese el índice de la tarea que desea editar:")
			fmt.Scan(&index)
			fmt.Print("Ingrese el nombre de la tarea")
			task.name, _ = read.ReadString('\n')
			fmt.Print("Ingrese la descripción de la tarea")
			task.description, _ = read.ReadString('\n')
			list.editTask(index, task)
			fmt.Println("Tarea actualizada completada")
		case 4:
			var index int
			fmt.Print("Ingrese el índice de la tarea que desea eliminar:")
			fmt.Scan(&index)
			list.deleteTask(index)
			fmt.Println("Tarea eliminada correctamente.")
		case 5:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opción invalida")
		}

		// Listar todas las tareas
		fmt.Println("Lista de Tareas:")
		fmt.Println("=================================")
		for index, task := range list.tasks {
			fmt.Printf("Tarea %d. %s - %s - Completado: %t\n", index, task.name, task.description, task.completed)
		}
		fmt.Println("=================================")
	}
}
