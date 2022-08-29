# Silverplate

Simple HTTP/S proxy server made with [goproxy](https://github.com/elazarl/goproxy/). Support for JSON-configurable rejecting rules.

## Configuration
Silverplate comes with 2 JSON config file : **config.json** and **rules.json**. **config.json** is for general proxy configuration and **rules.json** is where you will be defining your rejecting rules. An example rule is already provided.

```json
{
  "rules": [
    {
      "type": "DstHostIs",
      "value": "www.reddit.com",
      "conditions": {
        "hourRange": [9, 18]
      },
      "rejectMessage": "Not permitted during working hours",
      "httpStatusCode": 403
    }
  ]
}
```

### config.json
| Key                        | Description                                                                                                                                                             | Value type | Default value | Accepted value        |
|----------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------|---------------|-----------------------|
| **verboseMode**            | Sets verbose mode. Verbose mode will print all requests output to the console.                                                                                          | boolean    | true          | true / false          |
| **keepDestinationHeaders** | Whether or not the headers in the response should be kept before proxying.                                                                                              | boolean    | true          | true / false          |
| **keepHeader**             | Sets if proxy headers should be kept.                                                                                                                                   | boolean    | true          | true / false          |
| **port**                   | Port on which the proxy server will be started on.                                                                                                                      | number     | 3128          | Any valid port number |
| **handleNonProxyRequests** | When set to true, all requests that cannot be proxied will be forwarded to a separate HTTP handler function and served as HTTP. Otherwise these requests are discarded. | boolean    | true          | true / false          |

### rules.json
| Key                | Description                                                                                                                           | Value type  |
|--------------------|---------------------------------------------------------------------------------------------------------------------------------------|-------------|
| **type**           | The type of rule on which to evaluate the value                                                                                       | RuleType    |
| **value**          | Value of the rule. This is dependent on the chosen type of rule. For example, for a DstHostIs rule, the value should be a valid FQDN. | string      |
| **conditions**     | Additionnal conditions to check in order to evaluate the rule to true                                                                 | Condition[] |
| **rejectMessage**  | The message displayed to the client if the request is rejected                                                                        | string      |
| **httpStatusCode** | The returned HTTP status code to set in the headers if the request is rejected                                                        | number      |

#### Types

##### RuleType
| Rule               | Description                                                                      |
|--------------------|----------------------------------------------------------------------------------|
| **UrlHasPrefix**   | Whether the URL has a prefix that matches [value]                                |
| **UrlIs**          | The URL is a perfect match with [value]                                          |
| **ReqHostMatches** | The request hostname matches [value]. Should be one or more regular expressions. |
| **ReqHostIs**      | The request hostname is a perfect match with value [value]. Should be a string.  |
| **UrlMatches**     | The URL matches one or more regular expressions                                  |
| **DstHostIs**      | The destination FQDN is a perfect match with [value]                             |
| **SrcIpIs**        | The source IP is a perfect match with [value]                                    |

##### Condition
| Condition | Description                                       | Value type |
|-----------|---------------------------------------------------|------------|
| hourRange | Evaluates to true if the current hour is in range | number[]   |

## License
Distributed under the MIT License. See `LICENSE.txt` for more information.
