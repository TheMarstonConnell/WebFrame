package frame

import "honnef.co/go/js/dom/v2"

const (
	VARIANT_PRIMARY = "primary"
	VARIANT_SUCCESS = "success"
	VARIANT_NEUTRAL = "neutral"
	VARIANT_DEFAULT = "default"
	VARIANT_WARNING = "warning"
	VARIANT_DANGER  = "danger"
)

type Builder struct {
	Doc  dom.Document
	Body dom.Element
}
