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
)

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
		op := m.Memory[m.P]
		m.P++
		switch op {
		case OpHALT:
			return
		case OpINCA:
			m.A++
		case OpDECA:
			m.A--
		}
	}
}
