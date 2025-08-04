# Command Pattern

## 1. Definición del Patrón

El patrón Command encapsula una petición como un objeto, permitiendo parametrizar clientes con diferentes peticiones, encolar o registrar peticiones, y soportar operaciones de deshacer. Convierte una petición en un objeto independiente que contiene toda la información sobre la petición, desacoplando el objeto que invoca la operación del objeto que la realiza.

## 2. Diagrama

```
    Client                     Invoker                    Command Interface
┌─────────────┐           ┌─────────────────┐           ┌─────────────────────┐
│             │           │ - command       │◆─────────▶│ + Execute()         │
│             │           │ + SetCommand()  │           │ + Undo()            │
│             │           │ + ButtonPressed()│           └─────────────────────┘
└─────────────┘           └─────────────────┘                      △
                                                                   │
                                                    ┌──────────────┼──────────────┐
                                                    │              │              │
                                        ┌─────────────────┐ ┌─────────────────┐ ┌─────────────────┐
                                        │ ConcreteCommandA│ │ ConcreteCommandB│ │ MacroCommand    │
                                        │ - receiver      │ │ - receiver      │ │ - commands[]    │
                                        │ + Execute()     │ │ + Execute()     │ │ + Execute()     │
                                        │ + Undo()        │ │ + Undo()        │ │ + Undo()        │
                                        └─────────────────┘ └─────────────────┘ └─────────────────┘
                                                │                      │
                                                ▼                      ▼
                                        ┌─────────────────┐    ┌─────────────────┐
                                        │ ReceiverA       │    │ ReceiverB       │
                                        │ + Action()      │    │ + Action()      │
                                        └─────────────────┘    └─────────────────┘
```

## 3. Pasos para Implementar en Go

### Paso 1: Definir la Interface Command
```go
type Command interface {
    Execute()
    Undo()
}
```

### Paso 2: Crear Receivers (Objetos que realizan el trabajo)
```go
type Receiver struct {
    state string
}

func (r *Receiver) Action() {
    fmt.Println("Ejecutando acción en el receiver")
    r.state = "active"
}

func (r *Receiver) ReverseAction() {
    fmt.Println("Revirtiendo acción en el receiver")
    r.state = "inactive"
}
```

### Paso 3: Implementar Comandos Concretos
```go
type ConcreteCommand struct {
    receiver *Receiver
}

func (c *ConcreteCommand) Execute() {
    c.receiver.Action()
}

func (c *ConcreteCommand) Undo() {
    c.receiver.ReverseAction()
}
```

### Paso 4: Crear el Invoker
```go
type Invoker struct {
    commands    []Command
    undoCommand Command
}

func (i *Invoker) SetCommand(slot int, command Command) {
    i.commands[slot] = command
}

func (i *Invoker) ExecuteCommand(slot int) {
    i.commands[slot].Execute()
    i.undoCommand = i.commands[slot]
}

func (i *Invoker) UndoLastCommand() {
    if i.undoCommand != nil {
        i.undoCommand.Undo()
    }
}
```

### Paso 5: Usar el Patrón
```go
// Crear receiver
receiver := &Receiver{}

// Crear comando
command := &ConcreteCommand{receiver: receiver}

// Crear invoker
invoker := &Invoker{}
invoker.SetCommand(0, command)

// Ejecutar
invoker.ExecuteCommand(0)  // Ejecuta la acción
invoker.UndoLastCommand()  // Deshace la acción
```

## 4. Escenarios Recomendables

- **Desacoplamiento**: Separar el objeto que invoca de quien ejecuta
- **Operaciones Undo/Redo**: Necesitas deshacer operaciones
- **Logging y Auditoría**: Registrar todas las operaciones realizadas
- **Macro comandos**: Ejecutar múltiples operaciones como una sola
- **Colas de peticiones**: Encolar, programar o transmitir peticiones
- **Transacciones**: Agrupar operaciones que deben ejecutarse juntas

### Casos de Uso Reales:
- Interfaces de usuario (botones, menús, toolbars)
- Editores de texto (Ctrl+Z, Ctrl+Y)
- Sistemas de automatización y control
- Wizards y asistentes paso a paso
- Sistemas de procesamiento por lotes
- APIs REST (cada endpoint como comando)
- Transacciones de base de datos
- Sistemas de workflow y procesos

## 5. Particularidades en Go

- **Interfaces implícitas**: Los comandos implementan automáticamente la interfaz
- **Composición**: Usa embedding para comandos complejos
- **Punteros a receivers**: Mantén referencias a los objetos que realizan el trabajo
- **Null Object Pattern**: Usa NoCommand para slots vacíos
- **Estado previo**: Guarda estado para operaciones Undo complejas

```go
// Particularidad Go: NoCommand (Null Object Pattern)
type NoCommand struct{}
func (n *NoCommand) Execute() {} // No hace nada
func (n *NoCommand) Undo() {}    // No hace nada

// Particularidad Go: Estado para Undo complejo
type StatefulCommand struct {
    receiver  *StatefulReceiver
    prevState interface{}  // Guarda estado previo
}

func (c *StatefulCommand) Execute() {
    c.prevState = c.receiver.GetState() // Guarda antes de cambiar
    c.receiver.ChangeState("new_state")
}

func (c *StatefulCommand) Undo() {
    // Restaura estado previo específico
    c.receiver.RestoreState(c.prevState)
}
```

## 6. Variantes del Patrón

### Macro Command
```go
type MacroCommand struct {
    commands []Command
}

func (m *MacroCommand) Execute() {
    for _, cmd := range m.commands {
        cmd.Execute()
    }
}

func (m *MacroCommand) Undo() {
    // Deshacer en orden inverso
    for i := len(m.commands) - 1; i >= 0; i-- {
        m.commands[i].Undo()
    }
}
```

### Command con Parámetros
```go
type ParameterizedCommand struct {
    receiver *Receiver
    params   map[string]interface{}
}

func (c *ParameterizedCommand) Execute() {
    c.receiver.ExecuteWithParams(c.params)
}
```

### Command Queue
```go
type CommandQueue struct {
    commands []Command
}

func (q *CommandQueue) AddCommand(cmd Command) {
    q.commands = append(q.commands, cmd)
}

func (q *CommandQueue) ExecuteAll() {
    for _, cmd := range q.commands {
        cmd.Execute()
    }
}
```

## 7. Pros y Contras

### ✅ Pros
- **Desacoplamiento**: Invoker y receiver están desacoplados
- **Flexibilidad**: Fácil agregar nuevos comandos sin cambiar código existente
- **Undo/Redo**: Soporte natural para operaciones de deshacer
- **Macro comandos**: Combinar comandos simples en operaciones complejas
- **Logging**: Fácil registrar y auditar operaciones
- **Colas**: Encolar, programar y transmitir peticiones

### ❌ Contras
- **Complejidad**: Más clases/structs para operaciones simples
- **Overhead de memoria**: Cada comando es un objeto
- **Indirección**: Una capa adicional entre invoker y receiver
- **Estado**: Manejar estado para Undo puede ser complejo

## Ejemplo Completo

Ver implementación completa en: `behavioral/command/remote/`

```bash
cd behavioral/command/remote
go run main.go
```

**Nota**: El ejemplo implementado demuestra los conceptos fundamentales del patrón Command con funcionalidad completa de Undo y comandos macro, pero los principios del patrón son aplicables a cualquier sistema donde necesites encapsular peticiones como objetos.