# DuckDuckGo Smarter Encryption Checker

This package contains two functions to a) check whether a URL's domain supports encryption and b) automatically upgrade a URL string to SSL if possible.

```
import checker "github.com/cdukes/duckduckgo-smarter-encryption-checker"

u, _ := url.Parse("http://example.com")
log.Print(checker.SupportsSSL(u)) // false
log.Print(checker.MaybeUpgradeURL("http://example.com")) // http://example.com

u, _ = url.Parse("http://google.com")
log.Print(checker.SupportsSSL(u)) // true
log.Print(checker.MaybeUpgradeURL("http://google.com"))  // https://google.com
```