package main

import (
	"financas/internal/router"
	"fmt"
)

func main() {
	fmt.Println("Iniciando servidor na porta 8080...")
	router.RouterInit()
}
