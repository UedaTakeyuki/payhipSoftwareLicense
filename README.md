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
### [PayhipResponse struct](https://github.com/UedaTakeyuki/payhipSoftwareLicense/blob/a52658ab354c84992b7ab88ad5f9d3b324588d39/go/vers.go#L13)
