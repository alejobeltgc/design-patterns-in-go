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
type IPizza interface {
    Prepare()
    Bake()
    Cut()
    Box()
}

type Pizza struct {
    Name     string
    Dough    string
    Sauce    string
    Toppings []string
}
```

### Paso 2: Crear el Creator Base
```go
type PizzaStore interface {
    OrderPizza(pizzaType string) IPizza
    CreatePizza(pizzaType string) IPizza  // Factory Method
}

// Template Method común (evita duplicación)
type BasePizzaStore struct{}

func (ps *BasePizzaStore) OrderPizza(creator PizzaStore, pizzaType string) IPizza {
    pizza := creator.CreatePizza(pizzaType)  // Factory Method
    
    if pizza != nil {
        pizza.Prepare()  // Algoritmo común
        pizza.Bake()
        pizza.Cut()
        pizza.Box()
    }
    
    return pizza
}
```

### Paso 3: Implementar Productos Específicos
```go
type NyStyleCheesePizza struct {
    Pizza
}

func NewNyStyleCheesePizza() *NyStyleCheesePizza {
    return &NyStyleCheesePizza{
        Pizza: Pizza{
            Name:     "NY Style Cheese Pizza",
            Dough:    "Thin Crust",
            Sauce:    "Marinara",
            Toppings: []string{"Fresh Mozzarella"},
        },
    }
}

func (ny *NyStyleCheesePizza) Prepare() {
    fmt.Printf("Preparing %s\n", ny.Name)
    fmt.Printf("Adding %v\n", ny.Toppings)
}
```

### Paso 4: Crear Concrete Creators
```go
type NyStylePizzaStore struct {
    BasePizzaStore
}

func (ps *NyStylePizzaStore) OrderPizza(pizzaType string) IPizza {
    return ps.BasePizzaStore.OrderPizza(ps, pizzaType)
}

// Factory Method - decide qué crear
func (ps *NyStylePizzaStore) CreatePizza(pizzaType string) IPizza {
    switch pizzaType {
    case "cheese":
        return NewNyStyleCheesePizza()
    case "pepperoni":
        return NewNyStylePepperoniPizza()
    default:
        return nil
    }
}
```

### Paso 5: Usar el Patrón
```go
func main() {
    nyStore := &NyStylePizzaStore{}
    chicagoStore := &ChicagoStylePizzaStore{}
    
    // Mismo código, diferentes productos
    nyPizza := nyStore.OrderPizza("cheese")      // NY Style
    chicagoPizza := chicagoStore.OrderPizza("cheese")  // Chicago Style
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