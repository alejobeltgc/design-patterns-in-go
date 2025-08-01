# Singleton Pattern

## 1. Definición del Patrón

El patrón Singleton garantiza que una clase tenga exactamente una instancia y proporciona un punto de acceso global a ella. Controla la instanciación de una clase y asegura que solo exista una instancia en toda la aplicación.

## 2. Diagrama

```
                    Singleton
        ┌─────────────────────────────┐
        │ - instance: *Singleton      │◆──┐
        │ - once: sync.Once           │   │
        │                             │   │
        │ + GetInstance(): *Singleton │───┘
        │ + DoSomething()             │
        │ - new() // Constructor privado
        └─────────────────────────────┘
                        △
                        │
                   ┌─────────┐
                   │ Client  │
                   └─────────┘
```

## 3. Pasos para Implementar en Go

### Paso 1: Definir el Tipo y Variables Globales
```go
type ChocolateBoiler struct {
    empty  bool
    boiled bool
}

var (
    instance *ChocolateBoiler  // Puntero a la única instancia
    once     sync.Once         // Garantiza ejecución única
)
```

### Paso 2: Implementar GetInstance() con sync.Once
```go
func GetInstance() *ChocolateBoiler {
    once.Do(func() {
        // Esta función se ejecuta SOLO una vez
        fmt.Println("Creando una única instancia de chocolatera")
        instance = &ChocolateBoiler{
            empty:  true,
            boiled: false,
        }
    })
    return instance
}
```

### Paso 3: Agregar Métodos de Negocio
```go
func (cb *ChocolateBoiler) Fill() {
    if cb.empty {
        cb.empty = false
        cb.boiled = false
        fmt.Println("Llenando la chocolatera con leche y chocolate")
    } else {
        fmt.Println("Error: La chocolatera ya está llena")
    }
}

func (cb *ChocolateBoiler) Boil() {
    if !cb.empty && !cb.boiled {
        cb.boiled = true
        fmt.Println("Hirviendo el contenido de la chocolatera")
    } else {
        fmt.Println("Error: No se puede hervir - está vacía o ya hervida")
    }
}

func (cb *ChocolateBoiler) Drain() {
    if !cb.empty && cb.boiled {
        cb.empty = true
        fmt.Println("Drenando la chocolatera hervida")
    } else {
        fmt.Println("Error: No se puede drenar - está vacía o no hervida")
    }
}
```

### Paso 4: Usar el Singleton
```go
func main() {
    // Obtener la instancia desde cualquier lugar
    boiler1 := GetInstance()
    boiler2 := GetInstance()
    
    // Ambas variables apuntan al mismo objeto
    fmt.Printf("Same instance: %t\n", boiler1 == boiler2)  // true
    
    // Usar métodos
    boiler1.Fill()
    boiler1.Boil()
    boiler1.Drain()
}
```

## 4. Escenarios Recomendables

- **Recursos únicos**: Cuando solo debe existir una instancia de algo
- **Acceso global**: Necesitas acceso desde cualquier parte de la aplicación
- **Recursos costosos**: Objetos caros de crear (conexiones, archivos)
- **Estado compartido**: Información que debe ser consistente globalmente

### Casos de Uso Reales:
- Conexiones de base de datos (pools)
- Configuración de aplicación
- Loggers (archivos de log)
- Caches globales
- Administradores de recursos
- Clientes HTTP (connection pooling)

## 5. Particularidades en Go

- **sync.Once**: Forma idiomática y más eficiente que double-checked locking
- **Variables globales**: Usar variables de paquete para la instancia
- **Constructores**: No hay constructores privados, se confía en la convención
- **Thread safety**: sync.Once garantiza seguridad en concurrencia
- **Inicialización lazy**: La instancia se crea solo cuando se necesita

```go
// Particularidad Go: sync.Once es la forma idiomática
var (
    instance *MySingleton
    once     sync.Once  // Más eficiente que mutex manual
)

func GetInstance() *MySingleton {
    once.Do(func() {
        // Se ejecuta solo una vez, garantizado por Go runtime
        instance = &MySingleton{}
    })
    return instance
}

// Particularidad Go: No hay constructores privados
// Se confía en la convención de usar GetInstance()

// Particularidad Go: Thread safety en métodos si es necesario
func (s *MySingleton) SafeMethod() {
    s.mu.Lock()
    defer s.mu.Unlock()
    // operación thread-safe
}

// Particularidad Go: Test de concurrencia
func TestSingleton() {
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            instance := GetInstance()
            // Todas las goroutines obtienen la misma instancia
        }()
    }
    wg.Wait()
}
```

## 6. Pros y Contras

### ✅ Pros
- **Instancia única**: Garantiza que solo existe una instancia
- **Acceso global**: Disponible desde cualquier parte del código
- **Inicialización lazy**: Se crea solo cuando se necesita
- **Memoria**: Ahorra memoria al no crear múltiples instancias
- **Consistencia**: Estado consistente en toda la aplicación

### ❌ Contras
- **Testing**: Dificulta los unit tests (estado global)
- **Acoplamiento**: Crea dependencias implícitas
- **Concurrencia**: Puede crear cuellos de botella
- **Violación SRP**: Controla su creación y su funcionalidad
- **Herencia**: Difícil de extender o subclasificar

## Alternativas en Go

### Dependency Injection (Recomendada)
```go
// En lugar de Singleton global
type UserService struct {
    db     Database
    logger Logger
    cache  Cache
}

func NewUserService(db Database, logger Logger, cache Cache) *UserService {
    return &UserService{db: db, logger: logger, cache: cache}
}

// Inyectar dependencias en main
func main() {
    db := NewDatabase()
    logger := NewLogger()
    cache := NewCache()
    
    userService := NewUserService(db, logger, cache)
    // ...
}
```

### Package-level Variables
```go
// Para configuración simple
var Config = struct {
    DatabaseURL string
    APIKey      string
}{
    DatabaseURL: os.Getenv("DB_URL"),
    APIKey:      os.Getenv("API_KEY"),
}

// Uso directo
func connectDB() {
    db, err := sql.Open("postgres", Config.DatabaseURL)
    // ...
}
```

### Channels para Recursos Únicos
```go
// Pool de recursos usando channels
type ResourcePool struct {
    resources chan Resource
}

func NewResourcePool(size int) *ResourcePool {
    pool := &ResourcePool{
        resources: make(chan Resource, size),
    }
    
    // Llenar el pool
    for i := 0; i < size; i++ {
        pool.resources <- NewResource()
    }
    
    return pool
}

func (p *ResourcePool) Get() Resource {
    return <-p.resources
}

func (p *ResourcePool) Put(r Resource) {
    p.resources <- r
}
```

## Ejemplo con Cleanup

```go
type DatabaseSingleton struct {
    conn *sql.DB
}

var (
    database *DatabaseSingleton
    dbOnce   sync.Once
)

func GetDatabase() *DatabaseSingleton {
    dbOnce.Do(func() {
        conn, err := sql.Open("postgres", "connection_string")
        if err != nil {
            panic("Failed to connect to database: " + err.Error())
        }
        
        database = &DatabaseSingleton{conn: conn}
        
        // Registrar cleanup
        runtime.SetFinalizer(database, (*DatabaseSingleton).cleanup)
    })
    return database
}

func (db *DatabaseSingleton) cleanup() {
    if db.conn != nil {
        db.conn.Close()
    }
}
```

## Test de Concurrencia

```go
func TestSingletonConcurrency(t *testing.T) {
    var wg sync.WaitGroup
    instances := make([]*ChocolateBoiler, 1000)
    
    // Lanzar 1000 goroutines simultáneas
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func(index int) {
            defer wg.Done()
            instances[index] = GetInstance()
        }(i)
    }
    
    wg.Wait()
    
    // Verificar que todas son la misma instancia
    firstInstance := instances[0]
    for i := 1; i < 1000; i++ {
        if instances[i] != firstInstance {
            t.Errorf("Singleton failed: multiple instances created")
        }
    }
}
```

## Ejemplo Completo

Ver implementación completa en: `creational/singleton/chocolate/`

```bash
cd creational/singleton/chocolate
go run .
```