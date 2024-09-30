package ardan.rego

import rego.v1

default auth := false

auth if {
	[valid, _, _] := verify_jwt
	valid = true
}

verify_jwt := io.jwt.decode_verify(input.Token, {
	"cert": input.Key,
	"iss": input.ISS,
})

-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtOmdbAhYKIOmMtzIMJ+Q
btC8+jL9fr/cW/DjvDmhOUdRulikGa7G8lE1eD8ggkwn+4UCAx5JnK/1OGN/tVNE
TiRjOrw01cnOarwhXJhd9TS7eMsS8k9njx03x0WC9YxwocZQKXAxZZfmgfRLeZBt
FDa/fmrxtfsk+kHkfYiEgTuLbAJrQsL6PWKYkQhqb7juB/Q52mTMYIfh3s7v6WvO
JouRfpWPpefY3QI0In+UG3ucqpPaT1SWTqazxes+ehbFG8USpmBcKuh+8/xQjIID
diUPnP+1jsrDXsK/JtfmGtRJ26VfzZfDTOXSQEg0OO99vihaHaAwD/7GtpxmKjD1
3wIDAQAB
-----END PUBLIC KEY-----

{
    "Token": "eyJhbGciOiJSUzI1NiIsImtpZCI6IjU0YmIyMTY1LTcxZTEtNDFhNi1hZjNlLTdkYTRhMGUxZTJjMSIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzZXJ2aWNlIHByb2plY3QiLCJzdWIiOiIxMjM0NTY3ODkwIiwiZXhwIjoxNzU5MTgyNDkyLCJpYXQiOjE3Mjc2NDY0OTIsInJvbGVzIjpbIkFETUlOIl19.TQaFAb5pnJspUbHBuYa7rkq2-W3XzHoDFwQebAUyTKxFOCCzGKhGmjRtpQiy-zNPBLSEBCze4H12eqvUhFoluj7CE4mpSxKZ1MLol9RYXcA2BqSJbXIcmaA3Eu7jwubOBw7pmzcKnaXcsnZkDHLsE-_9KLz5iW9MFIB9sX1tI7MAaf9izmyZ_FyUyXe8DpAybUkoClpLhwrhqvRtuE0yhGE1qRvL_yo9bnC1pY-OYbR7OOJh7pEAFpjauGpIVP7RC0ALTrCCPMGbysQ5hnTwNM5tVc3uef0eM0CRn1fi10eCnAeIXhMS8fmhwDsNgXj_rX0Ge5RX4vTj6kdeO93WHA",
    "Key:": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtOmdbAhYKIOmMtzIMJ+Q\nbtC8+jL9fr/cW/DjvDmhOUdRulikGa7G8lE1eD8ggkwn+4UCAx5JnK/1OGN/tVNE\nTiRjOrw01cnOarwhXJhd9TS7eMsS8k9njx03x0WC9YxwocZQKXAxZZfmgfRLeZBt\nFDa/fmrxtfsk+kHkfYiEgTuLbAJrQsL6PWKYkQhqb7juB/Q52mTMYIfh3s7v6WvO\nJouRfpWPpefY3QI0In+UG3ucqpPaT1SWTqazxes+ehbFG8USpmBcKuh+8/xQjIID\ndiUPnP+1jsrDXsK/JtfmGtRJ26VfzZfDTOXSQEg0OO99vihaHaAwD/7GtpxmKjD1\n3wIDAQAB\n-----END PUBLIC KEY-----",
    "ISS": "service project"
}