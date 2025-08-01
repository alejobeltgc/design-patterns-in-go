# Factory Method Pattern

## 1. Definición del Patrón

El patrón Factory Method define una interfaz para crear objetos, pero permite que las subclases decidan qué clase instanciar. Delega la creación de objetos a subclases, proporcionando un punto de extensión para la creación de productos.

## 2. Diagrama

```
    Creator Interface                     Product Interface
┌─────────────────────────┐             ┌─────────────────────────┐
│ + FactoryMethod()       │             │ + Operation()           │
│ + AnOperation()         │             └─────────────────────────┘
└─────────────────────────┘                        △
           △                                       │
           │                              ┌────────┼────────┐
┌─────────────────────────┐               │        │        │
│ ConcreteCreatorA        │               │        │        │
│ + FactoryMethod()       │──────────────▶│        │        │
└─────────────────────────┘        ┌─────────────────────────┐
┌─────────────────────────┐        │ ConcreteProductA        │
│ ConcreteCreatorB        │        │ + Operation()           │
│ + FactoryMethod()       │        └─────────────────────────┘
└─────────────────────────┘        ┌─────────────────────────┐
                                   │ ConcreteProductB        │
                                   │ + Operation()           │
                                   └─────────────────────────┘
```

## 3. Pasos para Implementar en Go

### Paso 1: Definir la Interface Product
```go
type Product interface {
    Operation() string
    Configure() error
}

type BaseProduct struct {
    Name string
    Type string
    Config map[string]interface{}
}
```

### Paso 2: Crear el Creator Base
```go
type Creator interface {
    ProcessRequest(config map[string]interface{}) (string, error)
    CreateProduct(config map[string]interface{}) Product  // Factory Method
}

// Template Method común (evita duplicación)
type BaseCreator struct{}

func (bc *BaseCreator) ProcessRequest(creator Creator, config map[string]interface{}) (string, error) {
    product := creator.CreateProduct(config)  // Factory Method
    
    if product == nil {
        return "", errors.New("failed to create product")
    }
    
    if err := product.Configure(); err != nil {  // Algoritmo común
        return "", err
    }
    
    result := product.Operation()
    return result, nil
}
```

### Paso 3: Implementar Productos Específicos
```go
type ConcreteProductAlpha struct {
    BaseProduct
}

func NewConcreteProductAlpha(config map[string]interface{}) *ConcreteProductAlpha {
    return &ConcreteProductAlpha{
        BaseProduct: BaseProduct{
            Name:   "Product Alpha",
            Type:   "Type Alpha",
            Config: config,
        },
    }
}

func (p *ConcreteProductAlpha) Operation() string {
    return fmt.Sprintf("Alpha operation with config: %v", p.Config)
}

func (p *ConcreteProductAlpha) Configure() error {
    if p.Config["required_field"] == nil {
        return errors.New("required_field missing")
    }
    return nil
}
```

### Paso 4: Crear Concrete Creators
```go
type ConcreteCreatorAlpha struct {
    BaseCreator
}

func (c *ConcreteCreatorAlpha) ProcessRequest(config map[string]interface{}) (string, error) {
    return c.BaseCreator.ProcessRequest(c, config)
}

// Factory Method - decide qué crear
func (c *ConcreteCreatorAlpha) CreateProduct(config map[string]interface{}) Product {
    return NewConcreteProductAlpha(config)
}
```

### Paso 5: Usar el Patrón
```go
func main() {
    creatorAlpha := &ConcreteCreatorAlpha{}
    creatorBeta := &ConcreteCreatorBeta{}
    
    config := map[string]interface{}{"required_field": "value"}
    
    // Mismo código, diferentes productos
    resultAlpha, _ := creatorAlpha.ProcessRequest(config)  // Product Alpha
    resultBeta, _ := creatorBeta.ProcessRequest(config)    // Product Beta
}
```

## 4. Escenarios Recomendables

- **Múltiples variantes**: Cuando tienes familias de productos relacionados
- **Extensibilidad**: Necesitas agregar nuevos tipos sin modificar código existente
- **Configuración por contexto**: Diferentes implementaciones según el contexto
- **Template Method**: Algoritmo común con pasos variables

### Casos de Uso Reales:
- Procesadores de pago por región (US, EU, Asia)
- Adaptadores de base de datos por proveedor
- Parsers por formato de archivo
- Conectores de APIs por versión
- Estrategias de autenticación por tipo

## 5. Particularidades en Go

- **Composición**: Usar embedding en lugar de herencia
- **Interfaces implícitas**: Los creators implementan automáticamente la interfaz
- **Template Method**: Evitar duplicación con struct base
- **Constructores**: Funciones `NewXxx()` para cada producto
- **Validación**: Manejar casos donde no se puede crear el producto

```go
// Particularidad Go: Composición en lugar de herencia
type NyStylePizzaStore struct {
    BasePizzaStore  // Composición
}

// Particularidad Go: Delegación al template method
func (ps *NyStylePizzaStore) OrderPizza(pizzaType string) IPizza {
    return ps.BasePizzaStore.OrderPizza(ps, pizzaType)  // Pasa self
}

// Particularidad Go: Factory method con switch
func (ps *NyStylePizzaStore) CreatePizza(pizzaType string) IPizza {
    switch pizzaType {
    case "cheese":
        return NewNyStyleCheesePizza()
    default:
        fmt.Printf("No disponible: %s estilo NY\n", pizzaType)
        return nil
    }
}

// Particularidad Go: Evitar duplicación con BasePizzaStore
type BasePizzaStore struct{}

func (ps *BasePizzaStore) OrderPizza(creator PizzaStore, pizzaType string) IPizza {
    // Template method común
    pizza := creator.CreatePizza(pizzaType)
    // ... algoritmo común
    return pizza
}
```

## 6. Pros y Contras

### ✅ Pros
- **Extensibilidad**: Fácil agregar nuevos creators y productos
- **Polimorfismo**: Mismo código, diferentes comportamientos
- **Single Responsibility**: Cada creator maneja un tipo específico
- **Open/Closed**: Abierto para extensión, cerrado para modificación
- **Template Method**: Reutilización del algoritmo común

### ❌ Contras
- **Complejidad**: Más clases que Simple Factory
- **Overhead**: Puede ser excesivo para casos simples
- **Acoplamiento**: Creator conoce productos específicos
- **Jerarquía**: Requiere jerarquía de creators y products
- **Debugging**: Más difícil seguir el flujo de creación

## Comparación con Simple Factory

| Simple Factory | Factory Method |
|----------------|----------------|
| Una clase crea objetos | Subclases deciden qué crear |
| Centralizado | Distribuido |
| Menos flexible | Más extensible |
| Fácil de implementar | Más complejo |
| Viola OCP | Cumple OCP |

## Ejemplo con Configuración Avanzada

```go
type DatabaseCreator interface {
    CreateConnection(config Config) Database
    ValidateConfig(config Config) error
}

type PostgresCreator struct {
    BaseDatabaseCreator
}

func (pc *PostgresCreator) CreateConnection(config Config) Database {
    if err := pc.ValidateConfig(config); err != nil {
        return nil
    }
    
    return &PostgresConnection{
        host:     config.Host,
        port:     config.Port,
        database: config.Database,
    }
}

func (pc *PostgresCreator) ValidateConfig(config Config) error {
    if config.Host == "" {
        return errors.New("host is required for Postgres")
    }
    return nil
}
```

## Ejemplo con Registro Dinámico

```go
type CreatorRegistry struct {
    creators map[string]PizzaStore
}

func (r *CreatorRegistry) Register(name string, creator PizzaStore) {
    r.creators[name] = creator
}

func (r *CreatorRegistry) CreatePizza(storeName, pizzaType string) IPizza {
    if creator, exists := r.creators[storeName]; exists {
        return creator.CreatePizza(pizzaType)
    }
    return nil
}

// Uso
registry := &CreatorRegistry{creators: make(map[string]PizzaStore)}
registry.Register("ny", &NyStylePizzaStore{})
registry.Register("chicago", &ChicagoStylePizzaStore{})

pizza := registry.CreatePizza("ny", "cheese")
```

## Ejemplo Completo

Ver implementación completa en: `creational/factory/factory_method/`

```bash
cd creational/factory/factory_method
go run .
```

**Nota**: El ejemplo implementado usa el contexto de tiendas de pizza con diferentes estilos regionales, pero los principios del patrón son aplicables a cualquier dominio donde las subclases necesiten decidir qué objetos crear.