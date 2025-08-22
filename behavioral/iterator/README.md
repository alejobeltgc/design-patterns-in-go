# Iterator Pattern

## 1. Definición del Patrón

El patrón Iterator proporciona una forma de acceder secuencialmente a los elementos de una colección sin exponer su representación subyacente. Permite recorrer colecciones de manera uniforme, independientemente de su estructura interna.

## 2. Diagrama

```
    Iterator Interface                  Aggregate Interface
┌──────────────────────┐             ┌─────────────────────────┐
│ + HasNext() bool     │             │ + CreateIterator()      │
│ + Next() T           │             └─────────────────────────┘
└──────────────────────┘                        △
           △                                     │
           │                                     │
┌──────────────────────┐             ┌─────────────────────────┐
│ ConcreteIterator     │             │ ConcreteAggregate       │
│ - collection         │             │ - items                 │
│ - position           │             │ + CreateIterator()      │
│ + HasNext() bool     │             │                         │
│ + Next() T           │             │                         │
└──────────────────────┘             └─────────────────────────┘
```

## 3. Pasos para Implementar en Go

### Paso 1: Definir la Interfaz Iterator
```go
type Iterator[T any] interface {
    HasNext() bool
    Next() (T, error)
}
```

### Paso 2: Implementar el Iterador Concreto
```go
type ConcreteIterator struct {
    items    []T
    position int
}

func NewConcreteIterator(items []T) *ConcreteIterator {
    return &ConcreteIterator{items: items, position: 0}
}

func (it *ConcreteIterator) HasNext() bool {
    return it.position < len(it.items)
}

func (it *ConcreteIterator) Next() (T, error) {
    if !it.HasNext() {
        return nil, errors.New("no more items")
    }
    item := it.items[it.position]
    it.position++
    return item, nil
}
```

### Paso 3: Definir la Interfaz Aggregate
```go
type Aggregate[T any] interface {
    CreateIterator() Iterator[T]
}
```

### Paso 4: Implementar el Aggregate Concreto
```go
type ConcreteAggregate struct {
    items []T
}

func NewConcreteAggregate(items []T) *ConcreteAggregate {
    return &ConcreteAggregate{items: items}
}

func (ca *ConcreteAggregate) CreateIterator() Iterator[T] {
    return NewConcreteIterator(ca.items)
}
```

### Paso 5: Usar el Patrón
```go
aggregate := NewConcreteAggregate([]string{"Item1", "Item2", "Item3"})
iterator := aggregate.CreateIterator()

for iterator.HasNext() {
    item, err := iterator.Next()
    if err != nil {
        log.Println(err)
        break
    }
    fmt.Println(item)
}
```

## 4. Escenarios Recomendables

- **Colecciones heterogéneas**: Cuando necesitas recorrer diferentes tipos de colecciones de manera uniforme.
- **Encapsulación**: Cuando quieres ocultar la estructura interna de una colección.
- **Acceso secuencial**: Cuando necesitas recorrer elementos en un orden específico.
- **Extensibilidad**: Cuando quieres agregar nuevas formas de recorrer colecciones sin modificar su implementación.

### Casos de Uso Reales:
- Recorrer listas, mapas o árboles.
- Implementar paginación en sistemas de bases de datos.
- Procesar flujos de datos (streams).
- Recorrer estructuras de datos personalizadas.

## 5. Particularidades en Go

- **Generics**: Go 1.18+ permite implementar iteradores genéricos para mayor flexibilidad.
- **Interfaces implícitas**: Los iteradores concretos implementan automáticamente la interfaz `Iterator`.
- **Errores**: Es común devolver errores en el método `Next()` para manejar el final de la colección.
- **Slices**: Los slices son una estructura común para implementar colecciones en Go.
- **Concurrencia**: Puedes usar goroutines para recorrer colecciones en paralelo.

```go
// Iteración concurrente
func (it *ConcreteIterator) IterateConcurrently(process func(T)) {
    for it.HasNext() {
        item, err := it.Next()
        if err != nil {
            log.Println(err)
            break
        }
        go process(item)
    }
}
```

## 6. Pros y Contras

### ✅ Pros
- **Desacoplamiento**: Separa la lógica de iteración de la colección.
- **Uniformidad**: Proporciona una forma consistente de recorrer colecciones.
- **Extensibilidad**: Fácil agregar nuevos tipos de iteradores.
- **Encapsulación**: Oculta la estructura interna de la colección.

### ❌ Contras
- **Complejidad**: Puede ser excesivo para colecciones simples.
- **Overhead**: Los iteradores pueden agregar sobrecarga en términos de memoria y rendimiento.
- **Concurrencia**: No es seguro para concurrencia sin sincronización adicional.

## 7. Alternativas Idiomáticas en Go

### Iteración Directa con Rango
```go
items := []string{"Item1", "Item2", "Item3"}
for _, item := range items {
    fmt.Println(item)
}
```

### Uso de Channels
```go
func GenerateItems() <-chan string {
    ch := make(chan string)
    go func() {
        defer close(ch)
        for _, item := range []string{"Item1", "Item2", "Item3"} {
            ch <- item
        }
    }()
    return ch
}

for item := range GenerateItems() {
    fmt.Println(item)
}
```

## 8. Ejemplo Completo

Ver implementación completa en: `behavioral/iterator/dinermerge/`

```bash
cd behavioral/iterator/dinermerge
go run .
```

**Nota**: El ejemplo implementado utiliza menús de restaurantes como contexto, pero los principios del patrón son aplicables a cualquier dominio donde necesites recorrer colecciones de manera uniforme.