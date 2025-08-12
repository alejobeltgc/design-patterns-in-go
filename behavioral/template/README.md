# Template Method Pattern

## 1. Definición del Patrón

El patrón Template Method define el esqueleto de un algoritmo en una clase base, permitiendo que las subclases redefinan ciertos pasos del algoritmo sin cambiar su estructura general. Encapsula las partes invariantes del algoritmo y delega las partes variables a las subclases.

## 2. Diagrama

```
    AbstractClass                    ConcreteClass A              ConcreteClass B
┌─────────────────────┐           ┌─────────────────────┐      ┌─────────────────────┐
│ + TemplateMethod()  │           │ + PrimitiveOp1()    │      │ + PrimitiveOp1()    │
│ + ConcreteOp1()     │           │ + PrimitiveOp2()    │      │ + PrimitiveOp2()    │
│ + ConcreteOp2()     │           └─────────────────────┘      └─────────────────────┘
│ # PrimitiveOp1()    │                      △                            △
│ # PrimitiveOp2()    │                      │                            │
└─────────────────────┘                      └────────────────────────────┘

TemplateMethod() {
    ConcreteOp1()
    PrimitiveOp1()    // Implementado por subclases
    ConcreteOp2()
    PrimitiveOp2()    // Implementado por subclases
}
```

## 3. Pasos para Implementar en Go

### Paso 1: Definir la Interface para Operaciones Variables
```go
type TemplateInterface interface {
    PrimitiveOperation1()
    PrimitiveOperation2()
}
```

### Paso 2: Crear la Clase Template Base
```go
type AbstractTemplate struct{}

// TemplateMethod define el algoritmo principal
func (a *AbstractTemplate) TemplateMethod(impl TemplateInterface) {
    a.concreteOperation1()
    impl.PrimitiveOperation1()  // Paso variable
    a.concreteOperation2()
    impl.PrimitiveOperation2()  // Paso variable
    a.concreteOperation3()
}

// Operaciones concretas (invariantes)
func (a *AbstractTemplate) concreteOperation1() {
    fmt.Println("Ejecutando operación concreta 1")
}

func (a *AbstractTemplate) concreteOperation2() {
    fmt.Println("Ejecutando operación concreta 2")
}

func (a *AbstractTemplate) concreteOperation3() {
    fmt.Println("Ejecutando operación concreta 3")
}
```

### Paso 3: Implementar Clases Concretas
```go
type ConcreteClassA struct {
    AbstractTemplate  // Embedding
}

func NewConcreteClassA() *ConcreteClassA {
    return &ConcreteClassA{}
}

func (c *ConcreteClassA) PrimitiveOperation1() {
    fmt.Println("ConcreteClassA: Implementación específica de operación 1")
}

func (c *ConcreteClassA) PrimitiveOperation2() {
    fmt.Println("ConcreteClassA: Implementación específica de operación 2")
}
```

### Paso 4: Usar el Patrón
```go
// Crear instancias concretas
classA := NewConcreteClassA()
classB := NewConcreteClassB()

// Ejecutar el algoritmo template
classA.TemplateMethod(classA)  // Mismo algoritmo, implementación A
classB.TemplateMethod(classB)  // Mismo algoritmo, implementación B
```

## 4. Escenarios Recomendables

- **Algoritmos similares**: Cuando tienes algoritmos que siguen los mismos pasos pero con implementaciones diferentes
- **Evitar duplicación**: Eliminar código duplicado en algoritmos similares
- **Control de flujo**: Cuando la clase base debe controlar el orden de ejecución
- **Extensibilidad**: Permitir variaciones del algoritmo sin modificar la estructura
- **Frameworks**: Definir puntos de extensión en frameworks y librerías

### Casos de Uso Reales:
- Algoritmos de ordenamiento con diferentes criterios de comparación
- Procesamiento de datos con diferentes formatos de entrada/salida
- Workflows de negocio con pasos variables
- Parsers con diferentes formatos pero misma estructura
- Algoritmos de validación con reglas específicas
- Pipelines de procesamiento con etapas customizables

## 5. Particularidades en Go

- **Embedding**: Usa embedding para simular herencia
- **Interfaces**: Define interfaces para las operaciones variables
- **Composición**: Prefiere composición sobre embedding cuando sea apropiado
- **Funciones como parámetros**: Alternativa idiomática usando funciones
- **Validación**: Siempre valida que las implementaciones no sean nil

```go
// Particularidad Go: Validación en template method
func (a *AbstractTemplate) TemplateMethod(impl TemplateInterface) {
    if impl == nil {
        panic("implementation cannot be nil")
    }
    
    a.concreteOperation1()
    impl.PrimitiveOperation1()
    a.concreteOperation2()
    impl.PrimitiveOperation2()
}

// Particularidad Go: Alternativa con funciones
func TemplateMethodFunc(
    primitiveOp1 func(),
    primitiveOp2 func(),
) {
    concreteOperation1()
    primitiveOp1()  // Función pasada como parámetro
    concreteOperation2()
    primitiveOp2()  // Función pasada como parámetro
}

// Particularidad Go: Hook methods opcionales
type OptionalHooks interface {
    PreHook() bool   // Retorna false para saltar
    PostHook()
}

func (a *AbstractTemplate) TemplateMethodWithHooks(impl TemplateInterface, hooks OptionalHooks) {
    if hooks != nil && !hooks.PreHook() {
        return  // Saltar ejecución
    }
    
    a.TemplateMethod(impl)
    
    if hooks != nil {
        hooks.PostHook()
    }
}
```

## 6. Variantes del Patrón

### Template Method con Hook Methods
```go
type HookableTemplate interface {
    TemplateInterface
    PreProcessHook() bool   // Hook opcional antes del procesamiento
    PostProcessHook()       // Hook opcional después del procesamiento
}

func (a *AbstractTemplate) TemplateMethodWithHooks(impl HookableTemplate) {
    if !impl.PreProcessHook() {
        return  // Saltar si el hook retorna false
    }
    
    a.TemplateMethod(impl)
    impl.PostProcessHook()
}
```

### Template Method Funcional (Idiomático Go)
```go
type ProcessingStep func() error

type ProcessingTemplate struct {
    steps []ProcessingStep
}

func (p *ProcessingTemplate) Execute() error {
    for i, step := range p.steps {
        if err := step(); err != nil {
            return fmt.Errorf("error en paso %d: %w", i+1, err)
        }
    }
    return nil
}

// Uso
template := &ProcessingTemplate{
    steps: []ProcessingStep{
        func() error { return preProcess() },
        func() error { return mainProcess() },
        func() error { return postProcess() },
    },
}
```

### Template Method con Strategy
```go
type ProcessingStrategy interface {
    Process(data interface{}) (interface{}, error)
}

type TemplateWithStrategy struct {
    AbstractTemplate
    strategy ProcessingStrategy
}

func (t *TemplateWithStrategy) TemplateMethod(data interface{}) (interface{}, error) {
    // Pasos fijos del template
    t.preProcess()
    
    // Paso variable usando Strategy
    result, err := t.strategy.Process(data)
    if err != nil {
        return nil, err
    }
    
    t.postProcess()
    return result, nil
}
```

## 7. Template Method vs Otros Patrones

### Template Method vs Strategy
- **Template Method**: Define estructura del algoritmo, subclases implementan pasos
- **Strategy**: Encapsula algoritmos completos intercambiables

### Template Method vs Factory Method
- **Template Method**: Define algoritmo con pasos variables
- **Factory Method**: Delega creación de objetos a subclases

### Template Method vs Command
- **Template Method**: Estructura de algoritmo fija con pasos variables
- **Command**: Encapsula operaciones completas como objetos

## 8. Pros y Contras

### ✅ Pros
- **Reutilización**: Elimina duplicación de código en algoritmos similares
- **Control**: La clase base controla la estructura del algoritmo
- **Extensibilidad**: Fácil agregar nuevas variaciones del algoritmo
- **Mantenibilidad**: Cambios en la estructura se hacen en un solo lugar
- **Inversión de Control**: "No nos llames, nosotros te llamamos"
- **Consistencia**: Garantiza que todos sigan la misma estructura

### ❌ Contras
- **Rigidez**: Difícil cambiar la estructura del algoritmo base
- **Complejidad**: Puede ser excesivo para algoritmos simples
- **Acoplamiento**: Subclases están acopladas a la estructura de la clase base
- **Debugging**: Puede ser difícil seguir el flujo de ejecución
- **Limitaciones de Go**: Go no tiene herencia real, requiere workarounds

## Ejemplo con Manejo de Errores

```go
type ErrorHandlingTemplate interface {
    ValidateInput() error
    ProcessData() error
    SaveResults() error
}

type RobustTemplate struct{}

func (r *RobustTemplate) Execute(impl ErrorHandlingTemplate) error {
    // Template method con manejo robusto de errores
    if err := impl.ValidateInput(); err != nil {
        return fmt.Errorf("validación falló: %w", err)
    }
    
    if err := impl.ProcessData(); err != nil {
        return fmt.Errorf("procesamiento falló: %w", err)
    }
    
    if err := impl.SaveResults(); err != nil {
        return fmt.Errorf("guardado falló: %w", err)
    }
    
    return nil
}
```

## Ejemplo Completo

Ver implementación completa en: `behavioral/template/barista/`

```bash
cd behavioral/template/barista
go run .
```

**Nota**: El ejemplo implementado demuestra los conceptos fundamentales del patrón Template Method con algoritmos estructurados y pasos variables, pero los principios del patrón son aplicables a cualquier sistema donde necesites definir la estructura de un algoritmo permitiendo variaciones en pasos específicos.