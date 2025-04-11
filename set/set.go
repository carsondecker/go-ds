package set

// Set represents a collection of unordered, unique elements of a specified type
type Set[T comparable] struct {
	elements map[T]struct{}
}

// NewSet returns a pointer to an empty set of a specified type
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		elements: map[T]struct{}{},
	}
}

// Adds a given element to the set
func (s *Set[T]) Add(element T) {
	s.elements[element] = struct{}{}
}

// Removes a given element from the set
func (s *Set[T]) Remove(element T) {
	delete(s.elements, element)
}

// Checks if the set contains the given element
func (s *Set[T]) Contains(element T) bool {
	_, ok := s.elements[element]
	return ok
}

// Returns the length (size) of the set
func (s *Set[T]) Len() int {
	return len(s.elements)
}

// Clears the set
func (s *Set[T]) Clear() {
	s.elements = map[T]struct{}{}
}

// Checks if the set is empty
func (s *Set[T]) IsEmpty() bool {
	return len(s.elements) == 0
}
