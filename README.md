# SSL Support Checker

This package contains two functions to a) check whether a URL's domain supports SSL and b) automatically upgrade a URL string to SSL if supported.

```
import checker "github.com/cdukes/ssl-support-checker"

u, _ := url.Parse("http://example.com")
log.Print(checker.SupportsSSL(u)) // false
log.Print(checker.MaybeUpgradeURL("http://example.com")) // http://example.com

u, _ = url.Parse("http://google.com")
log.Print(checker.SupportsSSL(u)) // true
log.Print(checker.MaybeUpgradeURL("http://google.com"))  // https://google.com
```

Uses [DuckDuckGo Smarter Encryption](https://github.com/duckduckgo/smarter-encryption) ([Docs](https://help.duckduckgo.com/duckduckgo-help-pages/privacy/smarter-encryption/))