package types

// Field - обертка для доступа к полю с признаком, установлено ли поле.
// Инициировать значение поля можно только через конструктор WrapField
type Field[T any] struct {
	value T
	isSet bool
}

// IsSet - установлено ли поле
func (f Field[T]) IsSet() bool {
	return f.isSet
}

// Value - значение поля
func (f Field[T]) Value() T {
	return f.value
}

// WrapField - используется для полей, для которых важен признак, установленно ли значение поля
func WrapField[T any](v T) Field[T] {
	return Field[T]{
		value: v,
		isSet: true,
	}
}
