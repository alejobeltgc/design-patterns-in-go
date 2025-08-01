# Strategy Pattern

## 1. Definición del Patrón

El patrón Strategy define una familia de algoritmos, los encapsula y los hace intercambiables. Permite que el algoritmo varíe independientemente de los clientes que lo usan. En lugar de implementar un comportamiento directamente, el contexto delega a un objeto estrategia separado.

## 2. Diagrama

```
    Context                    Strategy Interface
┌─────────────────┐           ┌─────────────────────┐
│ - strategy      │◆─────────▶│ + Execute()         │
│ + SetStrategy() │           └─────────────────────┘
│ + DoSomething() │                      △
└─────────────────┘                      │
                                         │
                          ┌──────────────┼──────────────┐
                          │              │              │
              ┌─────────────────┐ ┌─────────────────┐ ┌─────────────────┐
              │ ConcreteStrategyA│ │ ConcreteStrategyB│ │ ConcreteStrategyC│
              │ + Execute()     │ │ + Execute()     │ │ + Execute()     │
              └─────────────────┘ └─────────────────┘ └─────────────────┘
```

## 3. Pasos para Implementar en Go

### Paso 1: Definir la Interface Strategy
```go
type FlyBehavior interface {
    Fly()
}
```

### Paso 2: Implementar Estrategias Concretas
```go
type FlyWithWings struct{}
func (f *FlyWithWings) Fly() {
    fmt.Println("I'm flying with wings!")
}

type FlyNoWay struct{}
func (f *FlyNoWay) Fly() {
    fmt.Println("I can't fly")
}
```

### Paso 3: Crear el Context
```go
type Duck struct {
    flyBehavior FlyBehavior
}

func (d *Duck) PerformFly() {
    if d.flyBehavior != nil {
        d.flyBehavior.Fly()
    }
}

func (d *Duck) SetFlyBehavior(fb FlyBehavior) {
    d.flyBehavior = fb
}
```

### Paso 4: Usar el Patrón
```go
duck := &Duck{}
duck.SetFlyBehavior(&FlyWithWings{})
duck.PerformFly() // "I'm flying with wings!"

duck.SetFlyBehavior(&FlyNoWay{})
duck.PerformFly() // "I can't fly"
```

## 4. Escenarios Recomendables

- **Múltiples algoritmos**: Cuando tienes diferentes formas de realizar una tarea
- **Evitar condicionales**: Reemplazar largos if/switch statements
- **Cambio dinámico**: Necesitas cambiar comportamiento en tiempo de ejecución
- **Extensibilidad**: Agregar nuevos algoritmos sin modificar código existente

### Casos de Uso Reales:
- Algoritmos de ordenamiento
- Métodos de pago (tarjeta, PayPal, transferencia)
- Estrategias de descuento
- Algoritmos de compresión
- Validadores de datos

## 5. Particularidades en Go

- **Interfaces implícitas**: No necesitas declarar que implementas la interfaz
- **Composición sobre herencia**: Go usa embedding en lugar de herencia clásica
- **Punteros vs valores**: Usa punteros para estrategias para evitar copias
- **Validación nil**: Siempre verifica que la estrategia no sea nil
- **Constructores**: Usa funciones `NewXxx()` para inicialización

```go
// Particularidad Go: Embedding
type MallardDuck struct {
    Duck  // Embedding, no herencia
}

// Particularidad Go: Validación nil
func (d *Duck) PerformFly() {
    if d.flyBehavior != nil {  // Importante en Go
        d.flyBehavior.Fly()
    }
}
```

## 6. Pros y Contras

### ✅ Pros
- **Flexibilidad**: Cambio de algoritmos en tiempo de ejecución
- **Extensibilidad**: Fácil agregar nuevas estrategias
- **Separación de responsabilidades**: Cada estrategia tiene una responsabilidad
- **Testabilidad**: Cada estrategia se puede testear independientemente
- **Open/Closed Principle**: Abierto para extensión, cerrado para modificación

### ❌ Contras
- **Complejidad**: Más clases/structs que una implementación simple
- **Overhead**: Puede ser excesivo para algoritmos simples
- **Conocimiento del cliente**: El cliente debe conocer las diferentes estrategias
- **Comunicación**: Puede requerir pasar contexto entre estrategia y cliente

## Ejemplo Completo

Ver implementación completa en: `behavioral/strategy/duck_simulator/`

```bash
cd behavioral/strategy/duck_simulator
go run .
```