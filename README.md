# Payhip Software License handler

## In Go lang
### How to use 
- Import rep ***github.com/UedaTakeyuki/payhipSoftwareLicense/go***
- Set ```payhip product key``` with function ***SetProcuctSecret***.
- Use handler functions like ***Verify***

Fx:
```
import(
  // import rep
	license "github.com/UedaTakeyuki/payhipSoftwareLicense/go"
)

// set payhip product key
license.SetProcuctSecret("prod_sk_xXXXX_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")

func verifyLicesne(key string) bool {
  // use handler function
  result, err := license.Verify(key)
  if !result.Data.Enabled {
    return false
  }
  return true
}
```

### functions
- [Decrease](go/decrease.go#L17)
- [Disable](go/disable.go#L17)
- [Enable](go/enable.go#L17)
- [Usage](go/usage.go#L17)
- [Verify](go/verify.go#L16)
- [SetProcuctSecret](go/vers.go#L26)

### structs
- [PayhipResponse struct](go/vers.go#L13)
