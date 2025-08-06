# Adapter Pattern

## 1. Definición del Patrón

El patrón Adapter permite que interfaces incompatibles trabajen juntas. Actúa como un puente entre dos interfaces incompatibles, convirtiendo la interfaz de una clase en otra interfaz que el cliente espera. Permite reutilizar clases existentes sin modificar su código fuente.

## 2. Diagrama

```
    Client                     Target Interface              Adapter                    Adaptee
┌─────────────┐           ┌─────────────────────┐      ┌─────────────────┐      ┌─────────────────┐
│             │──────────▶│ + Request()         │◀─────│ - adaptee       │◆────▶│ + SpecificReq() │
│             │           └─────────────────────┘      │ + Request()     │      │                 │
└─────────────┘                      △                 └─────────────────┘      └─────────────────┘
                                     │
                          ┌─────────────────────┐
                          │ ConcreteTarget      │
                          │ + Request()         │
                          └─────────────────────┘
```

## 3. Pasos para Implementar en Go

### Paso 1: Definir la Interface Target (Objetivo)
```go
type Target interface {
    Request() string
}
```

### Paso 2: Implementar el Adaptee (Clase a Adaptar)
```go
type Adaptee struct {
    data string
}

func NewAdaptee(data string) *Adaptee {
    return &Adaptee{data: data}
}

func (a *Adaptee) SpecificRequest() string {
    return "Specific: " + a.data
}
```

### Paso 3: Crear el Adapter
```go
type Adapter struct {
    adaptee *Adaptee
}

func NewAdapter(adaptee *Adaptee) *Adapter {
    if adaptee == nil {
        panic("adaptee cannot be nil")
    }
    return &Adapter{adaptee: adaptee}
}

func (a *Adapter) Request() string {
    // Adapta la llamada específica a la interfaz esperada
    return a.adaptee.SpecificRequest()
}
```

### Paso 4: Implementar Target Concreto (Opcional)
```go
type ConcreteTarget struct {
    name string
}

func NewConcreteTarget(name string) *ConcreteTarget {
    return &ConcreteTarget{name: name}
}

func (c *ConcreteTarget) Request() string {
    return "Standard: " + c.name
}
```

### Paso 5: Usar el Patrón
```go
// Cliente que espera trabajar con Target
func ClientCode(target Target) {
    fmt.Println(target.Request())
}

// Uso
adaptee := NewAdaptee("legacy data")
adapter := NewAdapter(adaptee)

ClientCode(adapter)  // Funciona transparentemente
```

## 4. Escenarios Recomendables

- **Integración de librerías**: Usar librerías de terceros con interfaces incompatibles
- **Código legacy**: Integrar sistemas antiguos con nuevas arquitecturas
- **APIs incompatibles**: Conectar servicios con diferentes contratos
- **Migración gradual**: Transición de una API antigua a una nueva
- **Reutilización**: Aprovechar funcionalidad existente sin modificarla

### Casos de Uso Reales:
- Adaptadores de bases de datos (MySQL, PostgreSQL, MongoDB)
- Gateways de pago (Stripe, PayPal, Square)
- Servicios de almacenamiento (AWS S3, Google Cloud, Azure)
- APIs de terceros con diferentes formatos
- Sistemas de logging con diferentes interfaces
- Convertidores de formatos de datos

## 5. Particularidades en Go

- **Interfaces implícitas**: El adapter implementa automáticamente la interfaz target
- **Composición**: Usa composición para envolver el adaptee
- **Validación nil**: Siempre validar que el adaptee no sea nil
- **Constructores**: Usar funciones `NewXxx()` para crear adapters
- **Múltiples adaptaciones**: Un adapter puede adaptar múltiples métodos

```go
// Particularidad Go: Validación en constructor
func NewAdapter(adaptee Adaptee) *Adapter {
    if adaptee == nil {
        panic("adaptee cannot be nil")  // Validación obligatoria
    }
    return &Adapter{adaptee: adaptee}
}

// Particularidad Go: Adaptación con transformación
func (a *Adapter) Request() string {
    // Puede hacer transformaciones complejas
    result := a.adaptee.SpecificRequest()
    return strings.ToUpper(result)  // Transformar datos
}

// Particularidad Go: Adapter bidireccional
type ReverseAdapter struct {
    target Target
}

func (r *ReverseAdapter) SpecificRequest() string {
    return r.target.Request()  // Adapta en dirección opuesta
}
```

## 6. Variantes del Patrón

### Object Adapter (Composición)
```go
type ObjectAdapter struct {
    adaptee *Adaptee  // Composición
}

func (o *ObjectAdapter) Request() string {
    return o.adaptee.SpecificRequest()
}
```

### Class Adapter (Embedding en Go)
```go
type ClassAdapter struct {
    Adaptee  // Embedding (similar a herencia)
}

func (c *ClassAdapter) Request() string {
    return c.SpecificRequest()  // Acceso directo
}
```

### Two-Way Adapter (Bidireccional)
```go
type TwoWayAdapter struct {
    target  Target
    adaptee *Adaptee
}

func (t *TwoWayAdapter) Request() string {
    return t.adaptee.SpecificRequest()
}

func (t *TwoWayAdapter) SpecificRequest() string {
    return t.target.Request()
}
```

## 7. Adapter vs Otros Patrones

### Adapter vs Decorator
- **Adapter**: Cambia la interfaz de un objeto
- **Decorator**: Agrega funcionalidad manteniendo la interfaz

### Adapter vs Facade
- **Adapter**: Adapta una interfaz específica
- **Facade**: Simplifica múltiples interfaces complejas

### Adapter vs Bridge
- **Adapter**: Resuelve incompatibilidad después del diseño
- **Bridge**: Separa abstracción e implementación desde el diseño

## 8. Pros y Contras

### ✅ Pros
- **Reutilización**: Permite usar código existente sin modificarlo
- **Desacoplamiento**: Separa el cliente del código adaptado
- **Flexibilidad**: Múltiples adapters para diferentes adaptees
- **Transparencia**: El cliente no sabe que está usando un adapter
- **Single Responsibility**: Cada adapter tiene una responsabilidad específica

### ❌ Contras
- **Complejidad**: Agrega una capa adicional de indirección
- **Performance**: Overhead de llamadas adicionales
- **Mantenimiento**: Más clases para mantener
- **Debugging**: Puede ser difícil seguir el flujo de llamadas

## Ejemplo de Adapter para APIs

```go
// API antigua
type LegacyAPI struct{}
func (l *LegacyAPI) OldMethod(data string) string {
    return "Legacy: " + data
}

// Nueva interfaz esperada
type ModernAPI interface {
    NewMethod(request Request) Response
}

// Adapter
type APIAdapter struct {
    legacy *LegacyAPI
}

func (a *APIAdapter) NewMethod(request Request) Response {
    // Transforma request nuevo a formato antiguo
    result := a.legacy.OldMethod(request.Data)
    // Transforma respuesta antigua a formato nuevo
    return Response{Result: result}
}
```

## Ejemplo Completo

Ver implementación completa en: `structural/adapter/ducks/`

```bash
cd structural/adapter/ducks
go run main.go
```

**Nota**: El ejemplo implementado demuestra los conceptos fundamentales del patrón Adapter con adaptación bidireccional y transformaciones inteligentes, pero los principios del patrón son aplicables a cualquier sistema donde necesites hacer compatibles interfaces incompatibles.