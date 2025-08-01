package main

import (
	"fmt"
	"sync"
)

type ChocolateBoiler struct {
	empty  bool
	boiled bool
}

var (
	instance *ChocolateBoiler
	once     sync.Once
)

func GetInstance() *ChocolateBoiler {
	once.Do(func() {
		fmt.Println("Creando una única instancia de chocolatera")

		instance = &ChocolateBoiler{
			empty:  true,
			boiled: false,
		}
	})

	return instance
}

func (cb *ChocolateBoiler) Fill() {
	if cb.empty {
		cb.empty = false
		cb.boiled = false
		fmt.Println("Llenando la chocolatera con leche y chocolate")
	} else {
		fmt.Println("Error: La chocolatera ya está llena")
	}
}

func (cb *ChocolateBoiler) Boil() {
	if !cb.empty && !cb.boiled {
		cb.boiled = true
		fmt.Println("Hirviendo el contenido de la chocolatera")
	} else {
		fmt.Println("Error: No se puede hervir - está vacía o ya hervida")
	}
}

func (cb *ChocolateBoiler) Drain() {

	if !cb.empty && cb.boiled {
		cb.empty = true
		fmt.Println("Drenando la chocolatera hervida")
	} else {
		fmt.Println("Error: No se puede drenar - está vacía o no hervida")
	}
}

func (cb *ChocolateBoiler) IsEmpty() bool {
	return cb.empty
}

func (cb *ChocolateBoiler) IsBoiled() bool {
	return cb.boiled
}

func (cb *ChocolateBoiler) String() string {
	return fmt.Sprintf("ChocolateBoiler{empty: %t, boiled: %t}", cb.empty, cb.boiled)
}
