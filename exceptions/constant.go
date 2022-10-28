package exceptions

type Code string

const (
	ResourceNotEncodedCode   Code = "fwork_rne"
	ResourcesNotPairedCode        = "fwork_rnp"
	ResourceNotGeneratedCode      = "fwork_rng"
	ResourceDuplicatedCode        = "fwork_rd"
	ResourceNotFoundCode          = "fwork_rnf"
	ResourceInvalidCode           = "fwork_ri"
	ResourceClosedCode            = "fwork_rc"
)

type Message string

const (
	ResourceNotEncodedMessage   Message = "resource not encoded"
	ResourcesNotPairedMessage           = "resources not paired"
	ResourceDuplicatedMessage           = "resource duplicated"
	ResourceNotGeneratedMessage         = "resource not generated"
	ResourceNotFoundMessage             = "resource not found"
	ResourceInvalidMessage              = "resource invalid"
	ResourceClosedMessage               = "resource closed"
)
