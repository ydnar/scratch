package cm

var (
	_ Variant2[struct{}, struct{}] = &UntypedVariant2{}
	_ Variant2[struct{}, struct{}] = &UnsizedVariant2[struct{}, struct{}]{}
	_ Variant2[string, bool]       = &SizedVariant2[Shape[string], string, bool]{}
	_ Variant2[bool, string]       = &SizedVariant2[Shape[string], bool, string]{}
)
