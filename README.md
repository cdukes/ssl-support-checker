# DuckDuckGo Smarter Encryption Checker

This package contains two functions to a) check whether a URL's domain supports encryption and b) automatically upgrade a URL string to SSL if possible.

```
import checker "github.com/cdukes/duckduckgo-smarter-encryption-checker"

log.Print( checker.SupportsSSL("http://example.com") ) // false
log.Print( checker.SupportsSSL(url) ) // true

log.Print( checker.MaybeUpgradeURL("http://example.com") ) // http://example.com
log.Print( checker.MaybeUpgradeURL(url) ) // https://google.com
```