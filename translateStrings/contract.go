package translateStrings

type translate interface {
	Translate(inText String,fmtOrigin String, fmtDestiny String) (String,error)
}