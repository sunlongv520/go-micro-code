package Boot

type FatalError struct {
     ErrMsg  string
  }
func (err *FatalError) Error() string {
     return err.ErrMsg
 }
func NewFatalError( msg string) *FatalError {
    return &FatalError{ ErrMsg: msg}
 }
func IsFatalError(err error) bool  {
	 if _,ok:=err.(*FatalError);ok{
	 	return true
	 }
	 return false
}