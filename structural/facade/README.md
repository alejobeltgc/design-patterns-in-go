# Facade Pattern

## 1. Definición del Patrón

El patrón Facade proporciona una interfaz unificada y simplificada para un conjunto de interfaces en un subsistema. Define una interfaz de nivel superior que hace que el subsistema sea más fácil de usar, ocultando la complejidad del sistema subyacente al cliente.

## 2. Diagrama

```
    Client                     Facade                    Subsystem Classes
┌─────────────┐           ┌─────────────────┐      ┌─────────────────┐
│             │──────────▶│ + Operation1()  │─────▶│ SubsystemA      │
│             │           │ + Operation2()  │──┐   │ + MethodA()     │
└─────────────┘           │ + Operation3()  │  │   └─────────────────┘
                          └─────────────────┘  │   ┌─────────────────┐
                                               ├──▶│ SubsystemB      │
                                               │   │ + MethodB()     │
                                               │   └─────────────────┘
                                               │   ┌─────────────────┐
                                               └──▶│ SubsystemC      │
                                                   │ + MethodC()     │
                                                   └─────────────────┘
```

## 3. Pasos para Implementar en Go

### Paso 1: Crear las Clases del Subsistema
```go
type SubsystemA struct {
    name string
}

func NewSubsystemA(name string) *SubsystemA {
    return &SubsystemA{name: name}
}

func (s *SubsystemA) OperationA() {
    fmt.Printf("SubsystemA: %s realizando operación A\n", s.name)
}

type SubsystemB struct {
    name string
}

func NewSubsystemB(name string) *SubsystemB {
    return &SubsystemB{name: name}
}

func (s *SubsystemB) OperationB() {
    fmt.Printf("SubsystemB: %s realizando operación B\n", s.name)
}
```

### Paso 2: Crear la Facade
```go
type Facade struct {
    subsystemA *SubsystemA
    subsystemB *SubsystemB
    subsystemC *SubsystemC
}

func NewFacade() *Facade {
    return &Facade{
        subsystemA: NewSubsystemA("ComponenteA"),
        subsystemB: NewSubsystemB("ComponenteB"),
        subsystemC: NewSubsystemC("ComponenteC"),
    }
}
```

### Paso 3: Implementar Operaciones Simplificadas
```go
func (f *Facade) SimpleOperation() {
    fmt.Println("Facade: Iniciando operación compleja...")
    
    // Coordina múltiples subsistemas
    f.subsystemA.OperationA()
    f.subsystemB.OperationB()
    f.subsystemC.OperationC()
    
    fmt.Println("Facade: Operación compleja completada")
}

func (f *Facade) AnotherOperation() {
    fmt.Println("Facade: Iniciando otra operación...")
    
    // Diferente combinación de subsistemas
    f.subsystemB.OperationB()
    f.subsystemA.OperationA()
    
    fmt.Println("Facade: Otra operación completada")
}
```

### Paso 4: Usar el Patrón
```go
// Cliente usa la interfaz simplificada
facade := NewFacade()

// Una sola llamada ejecuta múltiples operaciones complejas
facade.SimpleOperation()
facade.AnotherOperation()
```

## 4. Escenarios Recomendables

- **Sistemas complejos**: Simplificar el acceso a subsistemas complejos
- **Múltiples dependencias**: Cuando el cliente necesita interactuar con muchas clases
- **Puntos de entrada**: Crear puntos de entrada únicos para funcionalidades relacionadas
- **Migración gradual**: Introducir nuevas interfaces sin cambiar el código existente
- **Desacoplamiento**: Reducir dependencias entre cliente y subsistemas

### Casos de Uso Reales:
- APIs de alto nivel sobre librerías complejas
- Interfaces de usuario que coordinan múltiples servicios
- Sistemas de configuración que manejan múltiples componentes
- Wrappers para librerías de terceros
- Servicios de orquestación en microservicios
- Interfaces simplificadas para frameworks complejos

## 5. Particularidades en Go

- **Composición**: Usa composición para agrupar subsistemas
- **Constructores**: Inicializa todos los subsistemas en el constructor
- **Interfaces opcionales**: La facade puede implementar interfaces para mayor flexibilidad
- **Error handling**: Centraliza el manejo de errores de múltiples subsistemas
- **Configuración**: Puede encapsular configuración compleja

```go
// Particularidad Go: Constructor con inicialización completa
func NewFacade(config Config) *Facade {
    return &Facade{
        subsystemA: NewSubsystemA(config.ConfigA),
        subsystemB: NewSubsystemB(config.ConfigB),
        subsystemC: NewSubsystemC(config.ConfigC),
    }
}

// Particularidad Go: Manejo centralizado de errores
func (f *Facade) ComplexOperation() error {
    if err := f.subsystemA.OperationA(); err != nil {
        return fmt.Errorf("error en subsistema A: %w", err)
    }
    
    if err := f.subsystemB.OperationB(); err != nil {
        return fmt.Errorf("error en subsistema B: %w", err)
    }
    
    return nil
}

// Particularidad Go: Facade que implementa interface
type SystemManager interface {
    Start() error
    Stop() error
}

func (f *Facade) Start() error {
    // Coordina el inicio de múltiples subsistemas
    return f.ComplexStartupSequence()
}
```

## 6. Variantes del Patrón

### Facade con Interface
```go
type SystemFacade interface {
    PerformTask() error
    GetStatus() string
}

type ConcreteFacade struct {
    subsystems []Subsystem
}

func (c *ConcreteFacade) PerformTask() error {
    for _, subsystem := range c.subsystems {
        if err := subsystem.Execute(); err != nil {
            return err
        }
    }
    return nil
}
```

### Facade con Factory
```go
type FacadeFactory struct{}

func (ff *FacadeFactory) CreateFacade(facadeType string) SystemFacade {
    switch facadeType {
    case "basic":
        return NewBasicFacade()
    case "advanced":
        return NewAdvancedFacade()
    default:
        return NewDefaultFacade()
    }
}
```

### Facade Jerárquica
```go
type HighLevelFacade struct {
    lowLevelFacade *LowLevelFacade
}

func (h *HighLevelFacade) HighLevelOperation() {
    h.lowLevelFacade.LowLevelOperation1()
    h.lowLevelFacade.LowLevelOperation2()
}
```

## 7. Facade vs Otros Patrones

### Facade vs Adapter
- **Facade**: Simplifica una interfaz compleja
- **Adapter**: Hace compatibles interfaces incompatibles

### Facade vs Mediator
- **Facade**: Simplifica acceso a subsistemas (unidireccional)
- **Mediator**: Coordina comunicación entre objetos (bidireccional)

### Facade vs Proxy
- **Facade**: Simplifica acceso a múltiples objetos
- **Proxy**: Controla acceso a un objeto específico

## 8. Pros y Contras

### ✅ Pros
- **Simplicidad**: Interfaz simple para sistemas complejos
- **Desacoplamiento**: Reduce dependencias entre cliente y subsistemas
- **Flexibilidad**: Permite cambios en subsistemas sin afectar clientes
- **Reutilización**: Operaciones comunes encapsuladas y reutilizables
- **Mantenibilidad**: Centraliza lógica de coordinación
- **Testing**: Más fácil testear la facade que múltiples subsistemas

### ❌ Contras
- **Limitaciones**: Puede no exponer toda la funcionalidad de los subsistemas
- **God Object**: Riesgo de convertirse en un objeto que hace demasiado
- **Overhead**: Capa adicional de indirección
- **Rigidez**: Puede limitar la flexibilidad si es muy específica

## Ejemplo con Manejo de Errores

```go
type DatabaseFacade struct {
    userRepo    *UserRepository
    orderRepo   *OrderRepository
    emailSvc    *EmailService
    cacheSvc    *CacheService
}

func (d *DatabaseFacade) ProcessOrder(userID int, orderData OrderData) error {
    // Operación compleja que coordina múltiples servicios
    user, err := d.userRepo.GetUser(userID)
    if err != nil {
        return fmt.Errorf("error obteniendo usuario: %w", err)
    }
    
    order, err := d.orderRepo.CreateOrder(user, orderData)
    if err != nil {
        return fmt.Errorf("error creando orden: %w", err)
    }
    
    if err := d.emailSvc.SendConfirmation(user.Email, order); err != nil {
        // Log error but don't fail the operation
        log.Printf("Error enviando email: %v", err)
    }
    
    d.cacheSvc.InvalidateUserCache(userID)
    
    return nil
}
```

## Ejemplo Completo

Ver implementación completa en: `structural/facade/home_theater/`

```bash
cd structural/facade/home_theater
go run .
```

**Nota**: El ejemplo implementado demuestra los conceptos fundamentales del patrón Facade con coordinación de múltiples subsistemas y operaciones simplificadas, pero los principios del patrón son aplicables a cualquier sistema donde necesites simplificar el acceso a funcionalidades complejas.