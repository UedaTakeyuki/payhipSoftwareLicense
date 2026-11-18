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
- [Decrease]()
- [Disable]
- [Enable]
- [Usage]
- [Verify]
- [SetProcuctSecret]

### [PayhipResponse struct](go/vers.go#L13)
