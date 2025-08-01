# Simple Factory Pattern

## 1. Definición del Patrón

El patrón Simple Factory encapsula la creación de objetos en una clase dedicada. Proporciona una interfaz para crear objetos sin especificar sus clases concretas, centralizando la lógica de creación en un solo lugar.

## 2. Diagrama

```
     Client                    Factory                     Product Interface
┌─────────────────┐         ┌─────────────────┐           ┌─────────────────────┐
│                 │────────▶│ + CreateProduct()│          │ + Operation()       │
│                 │         │   (type: string) │          └─────────────────────┘
└─────────────────┘         └─────────────────┘                     △
                                      │                             │
                                      │                             │
                                      ▼                    ┌────────┼────────┐
                                 creates                   │        │        │
                                                  ┌─────────────────────┐ ┌─────────────────────┐
                                                  │ ConcreteProductA    │ │ ConcreteProductB    │
                                                  │ + Operation()       │ │ + Operation()       │
                                                  └─────────────────────┘ └─────────────────────┘
```

## 3. Pasos para Implementar en Go

### Paso 1: Definir la Interface Product
```go
type Product interface {
    Operation() string
    GetInfo() string
}

type BaseProduct struct {
    Name string
    Type string
}
```

### Paso 2: Implementar Productos Concretos
```go
type ConcreteProductA struct {
    BaseProduct
}

func NewConcreteProductA() *ConcreteProductA {
    return &ConcreteProductA{
        BaseProduct: BaseProduct{
            Name: "Product A",
            Type: "Type A",
        },
    }
}

func (p *ConcreteProductA) Operation() string {
    return "Operation from Product A"
}

func (p *ConcreteProductA) GetInfo() string {
    return fmt.Sprintf("%s (%s)", p.Name, p.Type)
}
```

### Paso 3: Crear el Simple Factory
```go
type SimpleFactory struct{}

func NewSimpleFactory() *SimpleFactory {
    return &SimpleFactory{}
}

func (f *SimpleFactory) CreateProduct(productType string) Product {
    switch productType {
    case "A":
        return NewConcreteProductA()
    case "B":
        return NewConcreteProductB()
    case "C":
        return NewConcreteProductC()
    default:
        fmt.Printf("Unknown product type: %s\n", productType)
        return nil
    }
}
```

### Paso 4: Usar el Factory
```go
func main() {
    factory := NewSimpleFactory()
    
    product := factory.CreateProduct("A")
    if product != nil {
        result := product.Operation()
        info := product.GetInfo()
        fmt.Printf("%s: %s\n", info, result)
    }
}
```

## 4. Escenarios Recomendables

- **Creación centralizada**: Cuando la lógica de creación es compleja
- **Múltiples productos relacionados**: Familia de objetos similares
- **Desacoplamiento**: Cliente no debe conocer clases concretas
- **Configuración dinámica**: Crear objetos basado en parámetros

### Casos de Uso Reales:
- Conexiones de base de datos (MySQL, PostgreSQL, MongoDB)
- Parsers de archivos (JSON, XML, CSV)
- Loggers (file, console, remote)
- Procesadores de pago (Stripe, PayPal, Square)
- Adaptadores de APIs externas

## 5. Particularidades en Go

- **Switch statements**: Forma idiomática de decidir qué crear
- **Constructores**: Usar funciones `NewXxx()` para inicialización
- **Interfaces implícitas**: Los productos implementan automáticamente la interfaz
- **Validación**: Manejar casos de tipos desconocidos
- **Embedding**: Usar embedding para compartir campos comunes

```go
// Particularidad Go: Constructor del factory
func NewPizzaFactory() *PizzaFactory {
    return &PizzaFactory{}
}

// Particularidad Go: Switch para decisión
func (f *PizzaFactory) CreatePizza(pizzaType string) IPizza {
    switch pizzaType {
    case "cheese":
        return NewCheesePizza()
    default:
        return nil  // Manejo explícito de casos no válidos
    }
}

// Particularidad Go: Embedding para reutilización
type CheesePizza struct {
    Pizza  // Embedding del struct base
}

// Particularidad Go: Validación en cliente
pizza := factory.CreatePizza("cheese")
if pizza != nil {  // Verificación necesaria en Go
    pizza.Prepare()
}
```

## 6. Pros y Contras

### ✅ Pros
- **Centralización**: Lógica de creación en un solo lugar
- **Desacoplamiento**: Cliente no conoce clases concretas
- **Mantenibilidad**: Fácil modificar lógica de creación
- **Reutilización**: Factory puede ser usado por múltiples clientes
- **Consistencia**: Garantiza inicialización correcta

### ❌ Contras
- **Violación OCP**: Agregar productos requiere modificar el factory
- **Acoplamiento**: Factory conoce todas las clases concretas
- **Complejidad**: Puede ser excesivo para casos simples
- **Single Point of Failure**: Problemas en factory afectan todo
- **Escalabilidad**: Difícil manejar muchos tipos de productos

## Comparación con Factory Method

| Simple Factory | Factory Method |
|----------------|----------------|
| Una clase crea objetos | Subclases deciden qué crear |
| Centralizado | Distribuido |
| Menos flexible | Más extensible |
| Fácil de implementar | Más complejo |

## Ejemplo con Configuración

```go
type DatabaseFactory struct {
    config map[string]string
}

func NewDatabaseFactory(config map[string]string) *DatabaseFactory {
    return &DatabaseFactory{config: config}
}

func (f *DatabaseFactory) CreateConnection(dbType string) Database {
    switch dbType {
    case "mysql":
        return NewMySQLConnection(f.config["mysql_url"])
    case "postgres":
        return NewPostgresConnection(f.config["postgres_url"])
    case "mongodb":
        return NewMongoConnection(f.config["mongo_url"])
    default:
        return nil
    }
}
```

## Ejemplo con Validación Avanzada

```go
func (f *PizzaFactory) CreatePizza(pizzaType string) (IPizza, error) {
    pizzaType = strings.ToLower(strings.TrimSpace(pizzaType))
    
    switch pizzaType {
    case "cheese":
        return NewCheesePizza(), nil
    case "pepperoni":
        return NewPepperoniPizza(), nil
    case "veggie":
        return NewVeggiePizza(), nil
    default:
        return nil, fmt.Errorf("unknown pizza type: %s", pizzaType)
    }
}
```

## Ejemplo Completo

Ver implementación completa en: `creational/factory/simple_factory/`

```bash
cd creational/factory/simple_factory
go run .
```

**Nota**: El ejemplo implementado usa el contexto de una pizzería con diferentes tipos de pizza, pero los principios del patrón son aplicables a cualquier dominio donde necesites centralizar la creación de objetos.