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
type Component interface {
    Operation() string
    GetInfo() string
}
```

### Paso 2: Implementar Componente Base (Opcional)
```go
type BaseComponent struct {
    Name string
}
```

### Paso 3: Implementar Componentes Concretos
```go
type ConcreteComponent struct {
    BaseComponent
}

func NewConcreteComponent() *ConcreteComponent {
    return &ConcreteComponent{
        BaseComponent: BaseComponent{Name: "Base Component"},
    }
}

func (c *ConcreteComponent) Operation() string {
    return "Base operation"
}

func (c *ConcreteComponent) GetInfo() string {
    return c.Name
}
```

### Paso 4: Implementar Decoradores Concretos
```go
type ConcreteDecorator struct {
    Component Component  // Wrappea otro Component
}

func NewConcreteDecorator(c Component) *ConcreteDecorator {
    if c == nil {
        panic("component cannot be nil")
    }
    return &ConcreteDecorator{Component: c}
}

func (d *ConcreteDecorator) Operation() string {
    return d.Component.Operation() + " + Extra functionality"  // Operación base + extra
}

func (d *ConcreteDecorator) GetInfo() string {
    return d.Component.GetInfo() + ", Decorated"  // Info base + extra
}
```

### Paso 5: Usar el Patrón
```go
// Crear objeto complejo combinando decoradores
component := NewConcreteComponent()                    // "Base operation"
component = NewConcreteDecorator(component)          // "Base operation + Extra functionality"  
component = NewAnotherDecorator(component)           // "Base operation + Extra functionality + More features"

fmt.Printf("%s: %s\n", component.GetInfo(), component.Operation())
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
func NewConcreteDecorator(c Component) *ConcreteDecorator {
    if c == nil {
        panic("component cannot be nil")  // Validación obligatoria
    }
    return &ConcreteDecorator{Component: c}
}

// Particularidad Go: Patrón de delegación + extensión
func (d *ConcreteDecorator) Operation() string {
    return d.Component.Operation() + " + Extra"  // Delegar + extender
}

// Particularidad Go: Método común en struct base
type BaseComponent struct {
    Name string
}

func (b *BaseComponent) DefaultBehavior() {
    fmt.Println("Default behavior")
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
type ComponentOption func(*Component)

func WithFeatureA() ComponentOption {
    return func(c *Component) {
        c.features = append(c.features, "FeatureA")
    }
}

func WithFeatureB() ComponentOption {
    return func(c *Component) {
        c.features = append(c.features, "FeatureB")
    }
}

func NewComponent(base string, options ...ComponentOption) *Component {
    c := &Component{name: base, features: []string{}}
    for _, option := range options {
        option(c)
    }
    return c
}

// Uso
component := NewComponent("BaseComponent", WithFeatureA(), WithFeatureB())
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

**Nota**: El ejemplo implementado usa el contexto de un sistema de café (Starbuzz) con bebidas y condimentos, pero los principios del patrón son aplicables a cualquier dominio donde necesites agregar funcionalidades dinámicamente.