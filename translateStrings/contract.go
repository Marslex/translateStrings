package translateStrings

type translate interface {
	Translate(inText string,fmtOrigin string, fmtDestiny string) (string,error)
}