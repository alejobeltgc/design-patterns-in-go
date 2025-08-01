# Observer Pattern

## 1. Definición del Patrón

El patrón Observer define una dependencia uno-a-muchos entre objetos, de manera que cuando un objeto cambia de estado, todos sus dependientes son notificados y actualizados automáticamente. Permite el desacoplamiento entre el objeto que notifica (Subject) y los objetos que reciben la notificación (Observers).

## 2. Diagrama

```
     Subject Interface                    Observer Interface
┌─────────────────────────┐             ┌─────────────────────────┐
│ + RegisterObserver()    │             │ + Update()              │
│ + RemoveObserver()      │             └─────────────────────────┘
│ + NotifyObservers()     │                        △
└─────────────────────────┘                        │
           △                                       │
           │                              ┌────────┼────────┐
┌─────────────────────────┐               │        │        │
│ ConcreteSubject         │               │        │        │
│ - observerList          │◆─────────────▶│        │        │
│ - state                 │        ┌─────────────────────────┐
│ + GetState()            │        │ ConcreteObserverA       │
│ + SetState()            │        │ + Update()              │
└─────────────────────────┘        └─────────────────────────┘
                                   ┌─────────────────────────┐
                                   │ ConcreteObserverB       │
                                   │ + Update()              │
                                   └─────────────────────────┘
```

## 3. Pasos para Implementar en Go

### Paso 1: Definir la Interface Observer
```go
type Observer interface {
    Update(data interface{})
}
```

### Paso 2: Definir la Interface Subject
```go
type Subject interface {
    RegisterObserver(o Observer)
    RemoveObserver(o Observer)
    NotifyObservers()
}
```

### Paso 3: Implementar el Subject Concreto
```go
type ConcreteSubject struct {
    observerList []Observer
    state        interface{}
}

func (cs *ConcreteSubject) RegisterObserver(o Observer) {
    cs.observerList = append(cs.observerList, o)
}

func (cs *ConcreteSubject) NotifyObservers() {
    for _, observer := range cs.observerList {
        observer.Update(cs.state)
    }
}

func (cs *ConcreteSubject) SetState(state interface{}) {
    cs.state = state
    cs.NotifyObservers()
}
```

### Paso 4: Implementar Observers Concretos
```go
type ConcreteObserver struct {
    id   string
    data interface{}
}

func (co *ConcreteObserver) Update(data interface{}) {
    co.data = data
    co.Display()
}

func (co *ConcreteObserver) Display() {
    fmt.Printf("Observer %s received: %v\n", co.id, co.data)
}
```

### Paso 5: Usar el Patrón
```go
subject := NewConcreteSubject()
observer1 := NewConcreteObserver("Observer1")

subject.RegisterObserver(observer1)
subject.SetState("New State") // Notifica automáticamente
```

## 4. Escenarios Recomendables

- **Notificaciones**: Sistemas que necesitan notificar múltiples componentes
- **Event-driven architecture**: Microservicios que reaccionan a eventos
- **UI updates**: Interfaces que se actualizan cuando cambian los datos
- **Logging distribuido**: Múltiples loggers escuchando eventos
- **Cache invalidation**: Invalidar múltiples caches cuando cambian datos

### Casos de Uso Reales:
- Sistema de notificaciones (email, SMS, push)
- Webhooks en APIs
- Actualizaciones de UI en tiempo real
- Sistemas de monitoreo y alertas
- Event sourcing

## 5. Particularidades en Go

- **Slices para observers**: Usa `[]Observer` para mantener la lista
- **Interfaces implícitas**: Los observers implementan automáticamente la interfaz
- **Goroutines**: Puedes notificar observers en paralelo
- **Channels**: Alternativa idiomática usando channels de Go
- **Evitar import cycles**: Separar interfaces en paquetes apropiados

```go
// Particularidad Go: Notificación concurrente
func (cs *ConcreteSubject) NotifyObservers() {
    var wg sync.WaitGroup
    for _, observer := range cs.observerList {
        wg.Add(1)
        go func(obs Observer) {
            defer wg.Done()
            obs.Update(cs.state)
        }(observer)
    }
    wg.Wait()
}

// Alternativa Go idiomática con channels
type EventChannel chan interface{}
```

## 6. Pros y Contras

### ✅ Pros
- **Desacoplamiento**: Subject y observers son independientes
- **Extensibilidad**: Fácil agregar nuevos observers
- **Broadcast**: Un cambio notifica a múltiples objetos
- **Open/Closed**: Abierto para extensión de observers
- **Reutilización**: Observers pueden ser reutilizados

### ❌ Contras
- **Complejidad**: Puede ser difícil debuggear el flujo
- **Memory leaks**: Observers no removidos pueden causar leaks
- **Performance**: Notificar muchos observers puede ser costoso
- **Orden impredecible**: No hay garantía del orden de notificación
- **Cascading updates**: Puede causar actualizaciones en cadena

## Alternativa Go Idiomática con Channels

```go
type EventPublisher struct {
    subscribers []chan interface{}
}

func (ep *EventPublisher) Subscribe() <-chan interface{} {
    ch := make(chan interface{}, 1)
    ep.subscribers = append(ep.subscribers, ch)
    return ch
}

func (ep *EventPublisher) Publish(data interface{}) {
    for _, ch := range ep.subscribers {
        select {
        case ch <- data:
        default: // No bloquear si el channel está lleno
        }
    }
}
```

## Ejemplo Completo

Ver implementación completa en: `behavioral/observer/weather/`

```bash
cd behavioral/observer/weather
go run .
```

**Nota**: El ejemplo implementado usa el contexto de una estación meteorológica con múltiples displays, pero los principios del patrón son aplicables a cualquier dominio donde necesites notificaciones uno-a-muchos.