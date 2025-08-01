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
type Strategy interface {
    Execute() string
}
```

### Paso 2: Implementar Estrategias Concretas
```go
type ConcreteStrategyA struct{}
func (s *ConcreteStrategyA) Execute() string {
    return "Algorithm A implementation"
}

type ConcreteStrategyB struct{}
func (s *ConcreteStrategyB) Execute() string {
    return "Algorithm B implementation"
}
```

### Paso 3: Crear el Context
```go
type Context struct {
    strategy Strategy
}

func (c *Context) ExecuteStrategy() string {
    if c.strategy != nil {
        return c.strategy.Execute()
    }
    return "No strategy set"
}

func (c *Context) SetStrategy(s Strategy) {
    c.strategy = s
}
```

### Paso 4: Usar el Patrón
```go
context := &Context{}
context.SetStrategy(&ConcreteStrategyA{})
result := context.ExecuteStrategy() // "Algorithm A implementation"

context.SetStrategy(&ConcreteStrategyB{})
result = context.ExecuteStrategy() // "Algorithm B implementation"
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
type ConcreteContext struct {
    Context  // Embedding, no herencia
}

// Particularidad Go: Validación nil
func (c *Context) ExecuteStrategy() string {
    if c.strategy != nil {  // Importante en Go
        return c.strategy.Execute()
    }
    return "No strategy set"
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

**Nota**: El ejemplo implementado usa el contexto de un simulador de patos con diferentes comportamientos, pero los principios del patrón son aplicables a cualquier dominio donde necesites algoritmos intercambiables.