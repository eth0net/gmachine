// Package gmachine implements a simple virtual CPU, known as the G-machine.
package gmachine

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new G-machine by default.
const DefaultMemSize = 1024

// Opcodes for gmachines.
const (
	OpHALT uint64 = iota
	OpNOOP
	OpINCA
	OpDECA
	OpSETA
)

// Program is a set of gmachine instructions.
type Program []uint64

// Machine stores the details of a virtual CPU, gmachine.
type Machine struct {
	A, P   uint64
	Memory []uint64
}

// New creates a new gmachine and returns it.
func New() *Machine {
	return &Machine{
		Memory: make([]uint64, DefaultMemSize),
	}
}

// Run executes the operations in memory starting from the point indicated by
// the Program Counter, returning on execution termination or on reaching the
// end of memory.
func (m *Machine) Run() {
	for {
		operation := m.Memory[m.P]
		m.P++
		switch operation {
		case OpHALT:
			return
		case OpINCA:
			m.A++
		case OpDECA:
			m.A--
		case OpSETA:
			operand := m.Memory[m.P]
			m.A = operand
			m.P++
		}
	}
}

// RunProgram copies the given program to the machine memory and calls Run.
func (m *Machine) RunProgram(p Program) {
	copy(m.Memory, p)
	m.P = 0
	m.Run()
}
