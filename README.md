# Problem
Hosted landing page need Host to be set to obscure name BACKEND_HOST; we want body content to reflect our OFFICIAL_DOMAIN.

# Idea
Run proxy on small instance

```bash
	OFFICIAL_DOMAIN=https://domain.tld BACKEND_HOST=http://lp.technical.domain go run main.go
```

