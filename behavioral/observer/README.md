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
type WeatherListener interface {
    Update(temperature float64, humidity float64, pressure float64)
}
```

### Paso 2: Definir la Interface Subject
```go
type WeatherPublisher interface {
    RegisterObserver(o WeatherListener)
    RemoveObserver(o WeatherListener)
    NotifyObservers()
}
```

### Paso 3: Implementar el Subject Concreto
```go
type WeatherData struct {
    observerList []WeatherListener
    temperature  float64
    humidity     float64
    pressure     float64
}

func (wd *WeatherData) RegisterObserver(o WeatherListener) {
    wd.observerList = append(wd.observerList, o)
}

func (wd *WeatherData) NotifyObservers() {
    for _, observer := range wd.observerList {
        observer.Update(wd.temperature, wd.humidity, wd.pressure)
    }
}

func (wd *WeatherData) SetMeasurements(temp, hum, press float64) {
    wd.temperature = temp
    wd.humidity = hum
    wd.pressure = press
    wd.NotifyObservers()
}
```

### Paso 4: Implementar Observers Concretos
```go
type CurrentConditionsDisplay struct {
    temperature float64
    humidity    float64
    pressure    float64
}

func (ccd *CurrentConditionsDisplay) Update(temp, hum, press float64) {
    ccd.temperature = temp
    ccd.humidity = hum
    ccd.pressure = press
    ccd.Display()
}

func (ccd *CurrentConditionsDisplay) Display() {
    fmt.Printf("Current conditions: %.1f°C and %.1f%% humidity\n", 
        ccd.temperature, ccd.humidity)
}
```

### Paso 5: Usar el Patrón
```go
weatherData := NewWeatherData()
currentDisplay := NewCurrentConditionsDisplay()

weatherData.RegisterObserver(currentDisplay)
weatherData.SetMeasurements(26.6, 65.0, 1013.1) // Notifica automáticamente
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
func (wd *WeatherData) NotifyObservers() {
    var wg sync.WaitGroup
    for _, observer := range wd.observerList {
        wg.Add(1)
        go func(obs WeatherListener) {
            defer wg.Done()
            obs.Update(wd.temperature, wd.humidity, wd.pressure)
        }(observer)
    }
    wg.Wait()
}

// Alternativa Go idiomática con channels
type EventChannel chan WeatherEvent
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
type WeatherStation struct {
    subscribers []chan WeatherData
}

func (ws *WeatherStation) Subscribe() <-chan WeatherData {
    ch := make(chan WeatherData, 1)
    ws.subscribers = append(ws.subscribers, ch)
    return ch
}

func (ws *WeatherStation) Publish(data WeatherData) {
    for _, ch := range ws.subscribers {
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