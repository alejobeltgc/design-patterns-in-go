# Decorator Pattern

## 1. Definición del Patrón

El patrón Decorator permite agregar funcionalidades a objetos dinámicamente sin alterar su estructura. Proporciona una alternativa flexible a la herencia para extender funcionalidad, envolviendo el objeto original en una serie de objetos decoradores.

## 2. Diagrama

```
    Component Interface
┌─────────────────────────┐
│ + Operation()           │
└─────────────────────────┘
           △
           │
    ┌──────┴──────┐
    │             │
┌─────────────────────────┐    ┌─────────────────────────┐
│ ConcreteComponent       │    │ Decorator               │
│ + Operation()           │    │ - component: Component  │◆──┐
└─────────────────────────┘    │ + Operation()           │   │
                               └─────────────────────────┘   │
                                          △                  │
                                          │                  │
                               ┌─────────────────────────┐   │
                               │ ConcreteDecorator       │   │
                               │ + Operation()           │   │
                               │ + AddedBehavior()       │───┘
                               └─────────────────────────┘
```

## 3. Pasos para Implementar en Go

### Paso 1: Definir la Interface Component
```go
type Beverage interface {
    Cost() float64
    GetDescription() string
}
```

### Paso 2: Implementar Componente Base
```go
type Pizza struct {
    Name  string
    Dough string
    Sauce string
}
```

### Paso 3: Implementar Componentes Concretos
```go
type Espresso struct {
    Pizza
}

func NewEspresso() *Espresso {
    return &Espresso{
        Pizza: Pizza{Name: "Espresso"},
    }
}

func (e *Espresso) Cost() float64 {
    return 1.99
}

func (e *Espresso) GetDescription() string {
    return e.Name
}
```

### Paso 4: Implementar Decoradores Concretos
```go
type Milk struct {
    Beverage Beverage  // Wrappea otro Beverage
}

func NewMilk(b Beverage) *Milk {
    if b == nil {
        panic("beverage cannot be nil")
    }
    return &Milk{Beverage: b}
}

func (m *Milk) Cost() float64 {
    return m.Beverage.Cost() + 0.10  // Costo base + extra
}

func (m *Milk) GetDescription() string {
    return m.Beverage.GetDescription() + ", Milk"  // Descripción base + extra
}
```

### Paso 5: Usar el Patrón
```go
// Crear bebida compleja combinando decoradores
beverage := NewEspresso()                    // $1.99 - "Espresso"
beverage = NewMilk(beverage)                // $2.09 - "Espresso, Milk"  
beverage = NewMocha(beverage)               // $2.29 - "Espresso, Milk, Mocha"
beverage = NewWhip(beverage)                // $2.49 - "Espresso, Milk, Mocha, Whip"

fmt.Printf("%s: $%.2f\n", beverage.GetDescription(), beverage.Cost())
```

## 4. Escenarios Recomendables

- **Extensión dinámica**: Agregar responsabilidades a objetos en tiempo de ejecución
- **Múltiples combinaciones**: Cuando hay muchas combinaciones posibles de funcionalidades
- **Evitar herencia**: Alternativa a crear subclases para cada combinación
- **Funcionalidades opcionales**: Características que pueden o no estar presentes

### Casos de Uso Reales:
- Middleware HTTP (autenticación, logging, compresión)
- Decoradores de conexiones de base de datos (retry, logging, métricas)
- Sistemas de archivos (compresión, encriptación, cache)
- APIs con funcionalidades opcionales
- Pipelines de procesamiento de datos

## 5. Particularidades en Go

- **Composición**: Go usa composición en lugar de herencia
- **Interfaces implícitas**: Los decoradores implementan automáticamente la interfaz
- **Validación nil**: Siempre validar que el componente no sea nil
- **Constructores**: Usar funciones `NewXxx()` para crear decoradores
- **Embedding**: Puede usar embedding para compartir campos comunes

```go
// Particularidad Go: Validación en constructores
func NewMilk(b Beverage) *Milk {
    if b == nil {
        panic("beverage cannot be nil")  // Validación obligatoria
    }
    return &Milk{Beverage: b}
}

// Particularidad Go: Patrón de delegación + extensión
func (m *Milk) Cost() float64 {
    return m.Beverage.Cost() + 0.10  // Delegar + extender
}

// Particularidad Go: Método común en struct base
type Pizza struct {
    Name string
}

func (p *Pizza) DefaultBox() {
    fmt.Println("Placing in standard box")
}
```

## 6. Pros y Contras

### ✅ Pros
- **Flexibilidad**: Agregar/quitar funcionalidades dinámicamente
- **Composición**: Combinar múltiples decoradores
- **Single Responsibility**: Cada decorador tiene una responsabilidad
- **Open/Closed**: Abierto para extensión, cerrado para modificación
- **Reutilización**: Decoradores pueden ser reutilizados

### ❌ Contras
- **Complejidad**: Muchos objetos pequeños pueden ser confusos
- **Debugging**: Difícil seguir el flujo a través de múltiples decoradores
- **Performance**: Overhead de múltiples llamadas de método
- **Orden importante**: El orden de los decoradores puede afectar el resultado
- **Identidad**: El objeto decorado pierde su identidad original

## Alternativa Go con Functional Options

```go
// Patrón Go idiomático para configuración
type BeverageOption func(*Beverage)

func WithMilk() BeverageOption {
    return func(b *Beverage) {
        b.cost += 0.10
        b.description += ", Milk"
    }
}

func WithMocha() BeverageOption {
    return func(b *Beverage) {
        b.cost += 0.20
        b.description += ", Mocha"
    }
}

func NewBeverage(base string, options ...BeverageOption) *Beverage {
    b := &Beverage{description: base, cost: 1.99}
    for _, option := range options {
        option(b)
    }
    return b
}

// Uso
beverage := NewBeverage("Espresso", WithMilk(), WithMocha())
```

## Middleware HTTP Example

```go
type Middleware func(http.Handler) http.Handler

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Request: %s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Verificar autenticación
        next.ServeHTTP(w, r)
    })
}

// Uso: Decorar handler con múltiples middlewares
handler := LoggingMiddleware(AuthMiddleware(myHandler))
```

## Ejemplo Completo

Ver implementación completa en: `structural/decorator/starbuzz/`

```bash
cd structural/decorator/starbuzz
go run .
```