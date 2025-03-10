---
title: "Epik"
date: 2019-03-03T16:39:46+01:00
draft: false
slug: epik
dnsprovider:
  since:    "v4.5.0"
  code:     "epik"
  url:      "https://www.epik.com/"
---

<!-- THIS DOCUMENTATION IS AUTO-GENERATED. PLEASE DO NOT EDIT. -->
<!-- providers/dns/epik/epik.toml -->
<!-- THIS DOCUMENTATION IS AUTO-GENERATED. PLEASE DO NOT EDIT. -->


Configuration for [Epik](https://www.epik.com/).


<!--more-->

- Code: `epik`
- Since: v4.5.0


Here is an example bash command using the Epik provider:

```bash
EPIK_SIGNATURE=xxxxxxxxxxxxxxxxxxxxxxxxxx \
lego --email you@example.com --dns epik -d '*.example.com' -d example.com run
```




## Credentials

| Environment Variable Name | Description |
|-----------------------|-------------|
| `EPIK_SIGNATURE` | Epik API signature (https://registrar.epik.com/account/api-settings/) |

The environment variable names can be suffixed by `_FILE` to reference a file instead of a value.
More information [here]({{% ref "dns#configuration-and-credentials" %}}).


## Additional Configuration

| Environment Variable Name | Description |
|--------------------------------|-------------|
| `EPIK_HTTP_TIMEOUT` | API request timeout in seconds (Default: 30) |
| `EPIK_POLLING_INTERVAL` | Time between DNS propagation check in seconds (Default: 2) |
| `EPIK_PROPAGATION_TIMEOUT` | Maximum waiting time for DNS propagation in seconds (Default: 60) |
| `EPIK_TTL` | The TTL of the TXT record used for the DNS challenge in seconds (Default: 3600) |

The environment variable names can be suffixed by `_FILE` to reference a file instead of a value.
More information [here]({{% ref "dns#configuration-and-credentials" %}}).




## More information

- [API documentation](https://docs-userapi.epik.com/v2/)

<!-- THIS DOCUMENTATION IS AUTO-GENERATED. PLEASE DO NOT EDIT. -->
<!-- providers/dns/epik/epik.toml -->
<!-- THIS DOCUMENTATION IS AUTO-GENERATED. PLEASE DO NOT EDIT. -->
