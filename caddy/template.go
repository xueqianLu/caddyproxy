package caddy

var staticWebsiteCaddyFileTempl = `
{{ .DomainName }} {
    root * {{ .RootDir }}
    file_server

    import COMMON_CONFIG
}
`

var newWebsitePayloadTempl = `
{
  "handle": [
    {
      "handler": "subroute",
      "routes": [
        {
          "handle": [
            {
              "handler": "vars",
              "root": "{{ .RootDir }}"
            }
          ]
        },
        {
          "handle": [
            {
              "handler": "headers",
              "response": {
                "set": {
                  "Strict-Transport-Security": [
                    "max-age=63072000"
                  ]
                }
              }
            }
          ],
          "match": [
            {
              "path": [
                "/"
              ]
            }
          ]
        },
        {
          "handle": [
            {
              "encodings": {
                "gzip": {},
                "zstd": {}
              },
              "handler": "encode",
              "prefer": [
                "zstd",
                "gzip"
              ]
            },
            {
              "handler": "file_server",
              "hide": [
                "./Caddyfile",
                "{{ .WebsiteCaddyFile }}"
              ]
            }
          ]
        }
      ]
    }
  ],
  "match": [
    {
      "host": [
        "{{ .DomainName }}"
      ]
    }
  ],
  "terminal": true
}`
