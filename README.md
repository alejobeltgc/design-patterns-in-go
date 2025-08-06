# Design Patterns in Go

Este repositorio contiene implementaciones de patrones de diseño clásicos en Go, organizados por categorías y con ejemplos prácticos.

## 📁 Estructura del Proyecto

```
design-patterns-in-go/
├── behavioral/          # Patrones de comportamiento
│   ├── command/         # Command Pattern
│   ├── observer/        # Observer Pattern
│   └── strategy/        # Strategy Pattern
├── structural/          # Patrones estructurales
│   ├── adapter/         # Adapter Pattern
│   └── decorator/       # Decorator Pattern
├── creational/          # Patrones creacionales
│   ├── factory/         # Factory Patterns
│   │   ├── simple_factory/
│   │   └── factory_method/
│   └── singleton/       # Singleton Pattern
└── README.md
```

## 🎯 Patrones Implementados

### Behavioral Patterns (Patrones de Comportamiento)

#### [Command Pattern](behavioral/command/README.md)

Encapsula una petición como un objeto, permitiendo parametrizar clientes y soportar operaciones de deshacer.

- **Ejemplo**: Control remoto universal con funcionalidad Undo y comandos macro
- **Casos de uso**: Interfaces de usuario, operaciones Undo/Redo, colas de peticiones, logging

#### [Observer Pattern](behavioral/observer/README.md)

Define una dependencia uno-a-muchos entre objetos para notificaciones automáticas.

- **Ejemplo**: Sistema de estación meteorológica con múltiples displays
- **Casos de uso**: Notificaciones, event-driven architecture, UI updates

#### [Strategy Pattern](behavioral/strategy/README.md)

Define una familia de algoritmos intercambiables encapsulados.

- **Ejemplo**: Simulador de patos con diferentes comportamientos de vuelo y sonido
- **Casos de uso**: Algoritmos de ordenamiento, métodos de pago, validadores

### Structural Patterns (Patrones Estructurales)

#### [Adapter Pattern](structural/adapter/README.md)

Permite que interfaces incompatibles trabajen juntas actuando como un puente entre ellas.

- **Ejemplo**: Sistema de adaptación bidireccional con transformaciones inteligentes
- **Casos de uso**: Integración de APIs, código legacy, gateways de pago, adaptadores de DB

#### [Decorator Pattern](structural/decorator/README.md)

Agrega funcionalidades a objetos dinámicamente sin alterar su estructura.

- **Ejemplo**: Sistema de café Starbuzz con condimentos opcionales
- **Casos de uso**: Middleware HTTP, decoradores de conexiones DB, pipelines

### Creational Patterns (Patrones Creacionales)

#### [Simple Factory](creational/factory/simple_factory/README.md)

Encapsula la creación de objetos en una clase dedicada.

- **Ejemplo**: Fábrica de pizzas con diferentes tipos
- **Casos de uso**: Conexiones DB, parsers, loggers

#### [Factory Method](creational/factory/factory_method/README.md)

Permite que las subclases decidan qué clase instanciar.

- **Ejemplo**: Tiendas de pizza de diferentes estilos (NY vs Chicago)
- **Casos de uso**: Procesadores de pago por región, adaptadores de APIs

#### [Singleton Pattern](creational/singleton/README.md)

Garantiza una única instancia con acceso global.

- **Ejemplo**: Chocolatera industrial con control de estado
- **Casos de uso**: Configuración, conexiones DB, caches globales

## 🚀 Cómo Ejecutar los Ejemplos

### Prerrequisitos

- Go 1.19 o superior

### Ejecutar un patrón específico

```bash
# Command Pattern
cd behavioral/command/remote
go run main.go

# Observer Pattern
cd behavioral/observer/weather
go run .

# Strategy Pattern
cd behavioral/strategy/duck_simulator
go run .

# Adapter Pattern
cd structural/adapter/ducks
go run main.go

# Decorator Pattern
cd structural/decorator/starbuzz
go run .

# Simple Factory
cd creational/factory/simple_factory
go run .

# Factory Method
cd creational/factory/factory_method
go run .

# Singleton Pattern
cd creational/singleton/chocolate
go run .
```

## 📚 Particularidades de Go

### Interfaces Implícitas

```go
// No necesitas declarar que implementas una interfaz
type Duck struct{}
func (d *Duck) Fly() {} // Automáticamente implementa FlyBehavior

type FlyBehavior interface {
    Fly()
}
```

### Composición sobre Herencia

```go
// Go usa embedding en lugar de herencia
type MallardDuck struct {
    Duck  // Embedding
}
```

### Concurrencia con Goroutines

```go
// Singleton thread-safe con sync.Once
var (
    instance *Singleton
    once     sync.Once
)

func GetInstance() *Singleton {
    once.Do(func() {
        instance = &Singleton{}
    })
    return instance
}
```

### Constructores Idiomáticos

```go
// Patrón de constructor en Go
func NewPizza(name string) *Pizza {
    return &Pizza{
        Name: name,
        // inicialización...
    }
}
```

## 🎯 Cuándo Usar Cada Patrón

| Patrón             | Cuándo Usar                                | Evitar Cuando                         |
| ------------------ | ------------------------------------------ | ------------------------------------- |
| **Command**        | Undo/Redo, colas, logging, desacoplamiento | Operaciones simples, sin historial    |
| **Observer**       | Notificaciones múltiples, eventos          | Pocos observers, relaciones simples   |
| **Strategy**       | Múltiples algoritmos, cambio dinámico      | Algoritmo único, lógica simple        |
| **Adapter**        | Interfaces incompatibles, integración      | Interfaces ya compatibles             |
| **Decorator**      | Funcionalidades opcionales, combinaciones  | Pocas variaciones, estructura fija    |
| **Simple Factory** | Creación centralizada, pocos tipos         | Muchos tipos, lógica compleja         |
| **Factory Method** | Múltiples familias, extensibilidad         | Tipos simples, creación directa       |
| **Singleton**      | Recurso único, acceso global               | Testing, múltiples instancias válidas |

## 🔧 Mejores Prácticas en Go

### 1. Usa Interfaces Pequeñas

```go
// ✅ Bueno: Interface específica
type Writer interface {
    Write([]byte) (int, error)
}

// ❌ Evitar: Interface muy grande
type MegaInterface interface {
    Write([]byte) (int, error)
    Read([]byte) (int, error)
    Close() error
    Seek(int64, int) (int64, error)
    // ... muchos más métodos
}
```

### 2. Valida Parámetros Nil

```go
func NewDecorator(component Component) *Decorator {
    if component == nil {
        panic("component cannot be nil")
    }
    return &Decorator{component: component}
}
```

### 3. Usa Constructores

```go
// ✅ Constructor explícito
func NewUserService(db Database) *UserService {
    return &UserService{db: db}
}

// ❌ Evitar: Inicialización manual
service := &UserService{}
service.db = db
```

### 4. Prefiere Composición

```go
// ✅ Composición
type EnhancedService struct {
    BaseService
    logger Logger
}

// ❌ No hay herencia en Go
```

## 📖 Recursos Adicionales

- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Head First Design Patterns](https://www.oreilly.com/library/view/head-first-design/0596007124/)

## 🤝 Contribuciones

Las contribuciones son bienvenidas. Por favor:

1. Fork el repositorio
2. Crea una rama para tu feature (`git checkout -b feature/nuevo-patron`)
3. Commit tus cambios (`git commit -am 'Agregar nuevo patrón'`)
4. Push a la rama (`git push origin feature/nuevo-patron`)
5. Crea un Pull Request

## 📄 Licencia

Este proyecto está bajo la Licencia MIT. Ver el archivo [LICENSE](LICENSE) para más detalles.

---

**Nota**: Estos ejemplos están diseñados con fines educativos para entender los patrones de diseño en Go. En aplicaciones reales, considera las particularidades específicas de tu caso de uso.
