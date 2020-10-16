// Copyright © 2020 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resources

import (
	"io/ioutil"
	"path/filepath"
)

// CACrt is the certificate authority certificate.
var CACrt = []byte(`-----BEGIN CERTIFICATE-----
MIIFEjCCAvqgAwIBAgIBATANBgkqhkiG9w0BAQsFADAoMSYwJAYDVQQDEx1UZXN0
aW5nIGNlcnRpZmljYXRlIGF1dGhvcml0eTAgFw0yMDA3MTQxOTExNDVaGA8yMTE5
MDcxNDE5MTE0M1owKDEmMCQGA1UEAxMdVGVzdGluZyBjZXJ0aWZpY2F0ZSBhdXRo
b3JpdHkwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQDKzi9DSBPK5b/K
TegB+2bQx6ExAn0goLtDAo+r0mypqWZ876ADWrG8E2dz0VTqLsT1BoSzjnlWWAn1
kWCbMvTgx3X11NVhnAMhSQPU9wuIi1sy2azaK2lTv9HCg17Q+8E/jl7BB/qlKgJ8
TtpwTlmoB7atCn4Gk2T3vvjbkDUhZyGBV0khqHBvj4E4zC8a7NgycE2kYZibLkVn
jAZwOKshU764S/ICgHAj/IP9WD50mXPXVGsa6Aje7eqydtRWpS2vvi56IPZJKPcW
CsqUCx6g0+IkUNj73NAzn0UAhgdO9PvNKuTyFzUfEpTI2HCPFzgc3pC/g6CVm47V
14KMX4uRqNyVeNYeiOP6vrJcnqu3/OTzWAYSTeII3Fqwp3m9niLMM1drWcZIUIxT
ymeUV6MQe4f49w4CwkUnpJI7FxUgBls2ar4QBzRJlj0LCP7Ri/rzjvvjmEzejZnZ
sk54z/H2zo0c2mHtxmF06abqByEpx7r5Z8uHunQH+2r8ZACNFFTX72tojVKYs45L
y6qPI/X48MuYIabi6KNDEElt2ykoYCitZ+GlIzyOR0mBxAt0TXE7QTj5xWOTWqgk
lllWuoizApUDDWbzDDANmaybB7eweVszZr8QIi7GQ8S/yO3C+Th+HBt6gw3A67br
F4rNJ6Ql/lf6fQBxNXnJgj6EhTQCZQIDAQABo0UwQzAOBgNVHQ8BAf8EBAMCAQYw
EgYDVR0TAQH/BAgwBgEB/wIBADAdBgNVHQ4EFgQU8inA9NnqIST5SBR6L3z9rbZd
1BowDQYJKoZIhvcNAQELBQADggIBAEMXJAHbRG9o4t2yuPDYqxSRo6df9wzgdFiH
FMEWuJidUn3a+RKepa6YN4BgaA9yygUm5DvMkyzUjL/hdfJ+9ORu5Sdbxe+/O1H4
nL/C1i+3p++L4ZpoNAX/jbnMmWnsGOc8qnoAw+9L3MYeTIK5r27xBTvRCnKvr5me
ynL5Tka6UYXigWxcZ+u1tvN/q/4OQ5R8vHSyPo5g4Vl7CNYQnLZewGS2B1u4f7sA
+Xw+Pelhbt6b2XyA0dO4O5JlXaHS0E+OZU8UDdbr0SL+byKBUd6ZqVLl3JnsH/Nh
bgvTZssaP0HN+deiOwTsJRJZFXtjn2MSyjWwcVTmMTJrjpcLnlTPBd/FB0L/m6W3
U5Fq9yR2h4EzyX3ZHTh+QF5G2f178WVLM99G8ZfQUiUYbEm62J1ulPfd5jN8yH3d
w5o+dHY+xMLOXjVNKv4oI5Q+2XGFAD+ACWJwvfUsN7F6MDW+5PjKaGz7fvoe4g/T
w7jsEKYyNdFQuArHZVTUALQUWLTIMJPiSQ1Qh8r0E/VH7bMw973jLeQizklz6vEw
fm+IPsuBZZDxcbeps2pW5VzxTU2q2Aww0AoucreSnPgZelk8U0WAuS3afqNjCCsd
630H9+Ssr3qrRX7iNRpLL2itWxduURi7ff4wJyDXsCIPVGAe/zyPuYKy2xdvwQcK
qN0RjB5h
-----END CERTIFICATE-----`)

// SignerTest01Crt is the certificate for singer-test01.
var SignerTest01Crt = []byte(`-----BEGIN CERTIFICATE-----
MIIEWTCCAkGgAwIBAgIQOmQoybWcXv5kykIKSdE5YjANBgkqhkiG9w0BAQsFADAo
MSYwJAYDVQQDEx1UZXN0aW5nIGNlcnRpZmljYXRlIGF1dGhvcml0eTAgFw0yMDA3
MTQxOTEyNTZaGA8yMTE5MDcxNDE5MTE0MlowGDEWMBQGA1UEAxMNc2lnbmVyLXRl
c3QwMTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKyN0GHe1PLPUvrO
jki9BJvx8QaiEyzTtmDv9N0AjsXOfgNlw8qdJ9uWDGEE4YduNIJg1pNqQuKaUedJ
sH4ddnGflVutEFGbC3awXFbOw3FYxNp6pli7E41sDJ1qihxb3E4xInEYDZpkBZkF
AxiujCN2wue7SNuaAc6J8EmUvvONArL1pAiJ7XeSYmx1HsNpkVobq7DMgnNcTr3I
D599RKX6PMpxtyAr+qgPwxXTIOno4AINIcAc3nywubi71SE8112c/fGHNU4Shx1y
hTKnbeE+Rmn5jsOLoknzqDrAkyuT9XNw6r9Nld311furo7MyTK9NX8nRXshI24dP
p9qk4gkCAwEAAaOBjDCBiTAOBgNVHQ8BAf8EBAMCA7gwHQYDVR0lBBYwFAYIKwYB
BQUHAwEGCCsGAQUFBwMCMB0GA1UdDgQWBBRq9vcOi24CL2wXKhfSo1fnuIHt6jAf
BgNVHSMEGDAWgBTyKcD02eohJPlIFHovfP2ttl3UGjAYBgNVHREEETAPgg1zaWdu
ZXItdGVzdDAxMA0GCSqGSIb3DQEBCwUAA4ICAQAV339n2B1/7zF3sRr6ZtTKX8b3
sdZD2uFzla4C/kvPQQertvlO1rTr3EfzEkvvA0X+knMqJ1glVSyVxzz/XFPif8IE
ddfPcHnnj/fObe/SysDv/R1AMlewwXwjPHggpCnNwXvxyE2ryLdSCFbsoSkZocqk
o6J0AxjlnMXXvqHVPEFYbLzMjguhRULpOy8PUX326jcxlBLjVejhwUfv7sS1pwEc
+tVH5IFHSGwGJSCfrnWh2mSkT9sxcGRXTJ2MFXH4iMs6wG9Gx6GoVtBuroXYgzcr
BLDe/5oUd+xNP47sJgYSnaTpGv2/WZnCCGangAjP+0jU6lN5CS861lIADZRmAZVB
1ZVlr0QPQ+Nmod+BM8qdw+kBSz/nCzEfTHrd4n3hBE/cFT4R9NrdekGCg6/GQIYb
sg5mdZkU+gwdpf7nUxyP6bJV6zqPZ1/ZwyRFcRqnXaLST45n2HCb1V754FUgg9jO
FQj4dtUjud7zs2J2KfDAPjXE7CWOpFlUjtDMjnaCyGh5IcRbgWSTm6KEQj38IU6R
wu/ezoj80BkYEP5Q6YOT1pDHUsNQwJW/cSV0ELGkfxovrN9VLRGJR+8uJ6j0wWMk
5cPmnLsvAuLPBO3iY+/w1ZKU40nebTwZL+2MyOAEidbLwqAtOhlsj9f8r0o4XBU0
QQ9lLQoXywGU8bfdgw==
-----END CERTIFICATE-----`)

// SignerTest01Key is the certificate for singer-test01.
var SignerTest01Key = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEArI3QYd7U8s9S+s6OSL0Em/HxBqITLNO2YO/03QCOxc5+A2XD
yp0n25YMYQThh240gmDWk2pC4ppR50mwfh12cZ+VW60QUZsLdrBcVs7DcVjE2nqm
WLsTjWwMnWqKHFvcTjEicRgNmmQFmQUDGK6MI3bC57tI25oBzonwSZS+840CsvWk
CIntd5JibHUew2mRWhursMyCc1xOvcgPn31Epfo8ynG3ICv6qA/DFdMg6ejgAg0h
wBzefLC5uLvVITzXXZz98Yc1ThKHHXKFMqdt4T5GafmOw4uiSfOoOsCTK5P1c3Dq
v02V3fXV+6ujszJMr01fydFeyEjbh0+n2qTiCQIDAQABAoIBACcRPJrMDr6ivhDW
71P8p6x/DKkJzPmbPXGZIoFe/PRAGju+sKORDVMkF81ng4BcLTtPnYVmy5nugpix
EGqdVRHIpdJJzqYLSn8m0uE8kvd4t6kXl84DHRwp2HlTg82D5s81RK0CWyIXf0Tz
442VB1mIK/y3ZHmD1uDiTir6qHkPltYh1ZZyY5rd06321Nyi32CMAC3Sah6Jk9xI
czhmMADiHjFoysJDX4Q5FBZqgThYhDNaFNATOiAb8vOD31NjpWBJJXab6Mm2GugF
YLqVudND0vUXLHpwZMlt9sX31j7AOo1tricHsFZjGDd4S2a/8ugMRwAdV2gogjUv
Kdj1TzkCgYEA1+NSGMqQpjTFXaHuRv6NRkzyd6MDl4KDfdAeIZCS+6X6DFzkZUgA
Ir70+m36Xvh3NoPX+HPUkQQP/5ZmdkRLz9DGSJmec+y6EdW6mzXSPCYXyBNqC4CR
Sj6zv08O31FfI4jMWhVXTQfU3dDRJtwii7zZYwo7hAPgEUVC7qo54PcCgYEAzJ1R
SVSt2OxxMFdfE+M6+jbAXJEjkOXywf1b0SCpVoIEajL6dZMNsF+oLxzYqQy1UaSb
pCvPxQvCl/uAAWYntVLTOkPEvdUqsUit/3N04YWjOpZKeUL2fhtrhZ0UNqeLGNUS
k9OCj/O10V93XxPw+vOj7rtwUf58AOu5/qXJlP8CgYEAkd1NguKaiTHuiCz4yY8D
9RPYX20M6DmOjlsngJYmOVETeXbp+mSDcvaCnxHfsHtAUN0T9xKL9M9B3/bGk/Hh
JzBwSG1C2iCAN7yosGXU/j7eopg7djoP8JIc8I2CBvD7zw6Gw/bXXXyFHroQFql1
zeRzHK+1NqVHp2OcrZTmNlECgYBfn9XxU2W7zyRG105T2QojDZtwp1Pbz9tX1bQn
VABPsYumphRvBj6Lgujyu6R1vL6wXSFv5BnBmPFkXeFAxiEgmIim3i0AGrNNDw5i
J+8jxnS2WK6NZUIVRXNGilZ8BGj5PfrkoyCNbfQ22UrMYGFqppqiY27mouwI8iGG
ZKyEuQKBgQCvmHrSdXM+YvtYeCATDT7r5G0CsUPuxsAXrkdRocoE5Nxkkub3hNmK
K3sG1MRrPsTtBgvdnBnvP4m2KiSCXfrsDNfWXGcxPfsLbJhjtmdmQ3roKdiEab3O
/usXIC7wM2aweTKmN9A1nBcWugIVop7yS3JzQm1uNWfzPWTWuL4RMA==
-----END RSA PRIVATE KEY-----`)

// SignerTest02Crt is the certificate for singer-test02.
var SignerTest02Crt = []byte(`-----BEGIN CERTIFICATE-----
MIIEWjCCAkKgAwIBAgIRAJ+RF/lFEwUcn+ENsCNJPQAwDQYJKoZIhvcNAQELBQAw
KDEmMCQGA1UEAxMdVGVzdGluZyBjZXJ0aWZpY2F0ZSBhdXRob3JpdHkwIBcNMjAw
NzE0MTkxMzA3WhgPMjExOTA3MTQxOTExNDJaMBgxFjAUBgNVBAMTDXNpZ25lci10
ZXN0MDIwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCnqL2UQZ0JSMyW
BbckBIGUN+jXLCi0zalhK8X8fZPRgLubgfkhqJInawNy5yCD6qVELaopJ0YViwJF
CtXKRzMuyMkH6w4Z94PoPTMMnAykfXuLzmUco3ijnhowzkRmD7hgj1JRCscKurkT
FxZeYyo6vQsfFJjuGx6DS70Z6UsEHcf7XVAn9u+5f26sGyeTvHqlPGAAdfCMKVRC
KVXbTwyrqjn/wFBy7K7BjxjbOBPjlnmDMbS8JBRlKw3xOKRQ9BoHtHBHnmo2rZHs
mrFlRYHZQbgxiRpcPGgVufWBKS8LFdNtTuafBpmlXCjcX5ijg2avjfFx0wZ2AB/d
OX/C3sSBAgMBAAGjgYwwgYkwDgYDVR0PAQH/BAQDAgO4MB0GA1UdJQQWMBQGCCsG
AQUFBwMBBggrBgEFBQcDAjAdBgNVHQ4EFgQUMXrb/F4hPpq7mEm7nGFaDxVfioww
HwYDVR0jBBgwFoAU8inA9NnqIST5SBR6L3z9rbZd1BowGAYDVR0RBBEwD4INc2ln
bmVyLXRlc3QwMjANBgkqhkiG9w0BAQsFAAOCAgEAC6DkcHVHegfZ42+FU7fLw81z
Jm/ngs4SbH6db249Q9lt00R70m3OU+vpAp3P/tIc/H6Jn7up1LItJvWlwUiqCyxn
YQj13r/lOWR9aqjS5jM4XvcjaiC1vfQdfghT/yHQT3HRsWSCsv+nDch4d80e20e+
ukPqLuDTTn7+nIWip2cHPIgaYOE6ZP9G3f3PORx8UwnTQvHiqV0tmBg61bD4xJmb
FIUFFmvG7LVmomu9GPdciXd9aOiBYmk+vHNXy5ekgDUR7xVM2Lt0uAoqegTZOWVh
UlAE1AAqL2YJcmej+g30by/ZvS/6GISYTGr+OhYSv+JAOY2ITBSrLD4c3CNzcq1d
SK6+hcVHRd9fv7Da3vpVVrNf4RcJGlzLFVfxCsrHeRrVfojjWW7voV8KpIqrBdX+
n1eWn1yiIonhZXvGo/UenPQ9QLjh8/l9ALn4DoefxdHnFtgFeyq5btAGKn5hEzOF
0Dbgob0B6s3JPQxk4DQiPwGFGGHLQGM9z3cbf7iOzf+3R2roTx+3s8WEljteYLOJ
F35AG+it80P21FzlMuLpnI3zbFomlx1U7Acf2lSpYshdRrcVpIfet3M8ZaraSDNO
mdY+QOjjjLpHwI64KqeJuCPbzdUrZW2Vth5DkDUwyX+Cld4G2HI3362NeI4gzWzG
tnOTVaeHiYpxrArnLno=
-----END CERTIFICATE-----`)

// SignerTest02Key is the key for singer-test02.
var SignerTest02Key = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAp6i9lEGdCUjMlgW3JASBlDfo1ywotM2pYSvF/H2T0YC7m4H5
IaiSJ2sDcucgg+qlRC2qKSdGFYsCRQrVykczLsjJB+sOGfeD6D0zDJwMpH17i85l
HKN4o54aMM5EZg+4YI9SUQrHCrq5ExcWXmMqOr0LHxSY7hseg0u9GelLBB3H+11Q
J/bvuX9urBsnk7x6pTxgAHXwjClUQilV208Mq6o5/8BQcuyuwY8Y2zgT45Z5gzG0
vCQUZSsN8TikUPQaB7RwR55qNq2R7JqxZUWB2UG4MYkaXDxoFbn1gSkvCxXTbU7m
nwaZpVwo3F+Yo4Nmr43xcdMGdgAf3Tl/wt7EgQIDAQABAoIBAE0yhRgenhAIVSs+
tnT/HisBE1UAID4f9D2pyh/YMpfkjn2r+upkk+dSfuQJSA91m2MpI5CPZNXGi+T+
eDILVqXUDbx3nqaWMUZd54OG1stme9yDzErDemjcA3M1hoj39A2B8IUgUUW/dDg2
CTassmUZZmWJNmFsW0BZP1kO8luSHAnGqXoi62JdLGdcmrRIoeCdIhkCTlV/ciQr
WotnsaEF5Kz9ta0nnxB/Rr5DSR4W8x22ALo0bxKVWHDpKF/0t9ElBtzQat13p5NV
OIp/V32DUI2AGBem82lDxYCS0aAotxEljyB5TqmJQSJfXQUxRc0/dXK4Kdecl/gH
6dRcfYECgYEAyfRo6eo4WxcLpiKj6vMbGHcEneCaiwjxJD9KAKntXxeJ5NSDxTfK
7Apio0wRyCSo3mt7vNyVNhxOz2uO9mo5CqgSY5fNHlZu971fNEXp0vg6GMNS768x
/K94cbv5u1nOnIXa5SEixWr0Lh2ZsG1vavTLxHcMamv+7N437Ak7V7sCgYEA1IbL
yVKK9CYLaXuLTvHF2DF74BH5d8VXnRz0hMd82cFDTfNK6MwfNrOomN8mO7eZCAzG
RRsaU8X47UxeRDh4E2VcWToSYo/d6i0Wd7IF4WZeGol9PtXFGYxWSv7NkR9F7nj9
6Y0FXQblCdlfkQeI5TWxDMqcBGw5QxwODJAMmvMCgYEAqulMDIIq1XQIaL4yKPk6
ehbnclENmRKlOT7SewNUHsDF80Gijror4lzbo0USW2Yi/7DI1El9gYCtdb6aC5JT
2d/pSB3E+qK9YK4ELzHns6JdUG2k97E1xZoefWpO664SO8bQPE8xpQ9hvNFSKsxC
maq38/moKTxiTsW1X+1kKlcCgYEAqs6remKcyxzIjz57+DbYi6k5phzMIza3884R
t6Wc8mYhccTXr8JeU9iQ6Exwrg4hMBcUQvZFco7qQc3e9XVtDCmqzwudOxnlgRA7
vneVwlJDz5Aw2Q92GdJwiFXBYaGA0ujrKYnthZbE/eV2qVkk7RL5+Q3d1rAkVYt2
vyIG/4ECgYEAvMFs1kpWmIGuW1LI+gbXil+LLpR6l/4Lr9dbcW2oSTF88TW7Ijs6
kIoeB/eLMWt8jDnyp9G/UoB9jfFgtvZ+gGyPizZQWy02sY0laQnmq9RXA8dkIPtL
8zLX+YwCSclX5slV0OWwqTlzp757AYHx1rLFgaFt3rVwY6lveVITjH0=
-----END RSA PRIVATE KEY-----`)

// SignerTest03Crt is the certificate for singer-test03.
var SignerTest03Crt = []byte(`-----BEGIN CERTIFICATE-----
MIIEWjCCAkKgAwIBAgIRAJu28fjs8COygbgpPnST4icwDQYJKoZIhvcNAQELBQAw
KDEmMCQGA1UEAxMdVGVzdGluZyBjZXJ0aWZpY2F0ZSBhdXRob3JpdHkwIBcNMjAw
NzE0MTkxMzM0WhgPMjExOTA3MTQxOTExNDJaMBgxFjAUBgNVBAMTDXNpZ25lci10
ZXN0MDMwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDGg2WKF/2VcMBs
YMpgTAiouP9eoXDkvVQd97GUyMfTH8Oz1k+yNZ/WUU/k5z33HTkAyJkLxreDgOri
k+wuZeLToEbB414N3m7JL6ixbo5jt0A134v8zNetGasBz8a3l+JqY9reQPyyTrMi
va5zcGkDJTTVaerVrdclfPq7AF4NnsFP1qNFZoU9pqFoZN76Oz0lUQdv+qy/VxCz
r4QuF/1TNvTU/k/3yPM9uIOfoCMaCdexRCoa3BcLF0QiC3KOTOAUkSdWVz0pYZW7
aoaAo7a+PWYbpm/FAUGPXWbEWWHdtWkoeUyygjMrjqrLQsdbqOl4QALOKJxrS8sR
ueHYlCv1AgMBAAGjgYwwgYkwDgYDVR0PAQH/BAQDAgO4MB0GA1UdJQQWMBQGCCsG
AQUFBwMBBggrBgEFBQcDAjAdBgNVHQ4EFgQUJCvuUovWLCLq6IMuJcgeoFFT734w
HwYDVR0jBBgwFoAU8inA9NnqIST5SBR6L3z9rbZd1BowGAYDVR0RBBEwD4INc2ln
bmVyLXRlc3QwMzANBgkqhkiG9w0BAQsFAAOCAgEAGkQb15A9nctJ+aq+tc9wtkgZ
Quku/o/boQ04hWrVaMj4kKw0cVXRApa5LdHr0YHIcOddvCXq73N38scIgmMlXS5O
kaLnwM2J42zlze1BtPxJyqynfz5ifqrRcuH89CKznENbUdWmCwsh2jU9Jxme7yIp
2J0uU472PyLdQcDPD0ah89tIcQ51fPaChY8RGlndAzjkOPUHm9UWjE13Y49ca16h
xFHGpxHLWmjvOKn5k3Y3q2hgXR67DP9sXYEA1iRPZzhWztJXzWePRzPjYfaC9orm
y8YzQa1MgnuGbrAsJ5xyEWsn5Px7ZdiEPNmNvxBeWtmJGL0w4EeZ9UR5P55dVdnz
PAkju3MUqqb2uY/tKzD5Ps5UUuZT8Q/aZUt6kQjeu8Li3bvIQJsdPZ/AghTwQZek
IL5whpCMLJIpK5PKHzXUK3fCHWKLvEZRLhjS8FOtLdqUM5lGkdkhmw+xIxgfRSu/
oPGrJzTEWohoRKUcmTRNmwq53OJxSqS82tnXJFNG7HMhzp9qwEQ+IfJMZBNSEA9y
EdgVypD4J+EsKHM4vS150qV0Ipm0TbXPQNMcttm1MaBUWoiOl2ZWhhnLH3eIMTVP
AtjX7nZ2WsXL62pOOPRzmq4mDRklrZOGJLALlAbh2GV6Vlgl0INLhGKMCg4zR9zC
1PBzRN/jGLFUtPWDcfo=
-----END CERTIFICATE-----`)

// SignerTest03Key is the certificate for singer-test03.
var SignerTest03Key = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAxoNlihf9lXDAbGDKYEwIqLj/XqFw5L1UHfexlMjH0x/Ds9ZP
sjWf1lFP5Oc99x05AMiZC8a3g4Dq4pPsLmXi06BGweNeDd5uyS+osW6OY7dANd+L
/MzXrRmrAc/Gt5fiamPa3kD8sk6zIr2uc3BpAyU01Wnq1a3XJXz6uwBeDZ7BT9aj
RWaFPaahaGTe+js9JVEHb/qsv1cQs6+ELhf9Uzb01P5P98jzPbiDn6AjGgnXsUQq
GtwXCxdEIgtyjkzgFJEnVlc9KWGVu2qGgKO2vj1mG6ZvxQFBj11mxFlh3bVpKHlM
soIzK46qy0LHW6jpeEACziica0vLEbnh2JQr9QIDAQABAoIBAHD7j+sMxHMtfRcx
73znSXZd9FozoEFP0HRN6XA45iIvTf8o1qsjAhnEpaguRIJEy5C6mwCs1P2vff//
GMk/i93OtvbbJUjXdE5lJan2tmvifFUtKktzer02grHBQ+RyOZc7xfIBItGJIUOA
ma1yKRJ1NqDQa6u6CV14yGtuRhmRYyiO+k43n9ig5GNtnhMN1Oi3jii6BoKxBZGE
ooFm3iQx3yqRLWo6nvvumjIGerdvhtACKI0SqlXacbzeJxa4WlLnlj3ccN8mK+kO
bQYR/HeGCn5m+0of39cBFooZjO1iC+jiy1bI10QpJ3Re0dEn2lKf4Smet2ZWgLoS
/nAHlPECgYEAyGNTjCDYOJk8ZzqRGfwdzwmDOxTeujD1kKBufusDmPjZc5YSWY85
iFUr7NihZZyfo4S6AsQlbwswG+vY74eHe33sK9FCepkymHTCgiE1SY1GlhZBV3lB
9hERuSz2VA72x6AuoIN0NaQJ4MyzDqY+jm4NQ+XUKO75B393RXjc+psCgYEA/Zrh
IAB/KJg4gLQrOYa/5dBjGIlIs4TbnB1lgWsYyaJDySR0YITIIJpTD2bgkG0eqjm1
4fFOo/znvISapEQh7hslusaJzRum2yExvCPrfECtcbBx52T0dUHWalY9ejMew0eQ
83Dw2TrBytzHTIsuX8y3NVRw4rDa/hjrY41+VK8CgYAaPl/nvzlyGCeAAXyVYZ5p
yf4k07PjwJu4iDpQZuj+tMCN0b6vegF36GerSifBDgUePji6OgQJCfcQARBVNnO4
6aHvjyVctwmYS9pZfo3jBxySdXGzSg8occ3XaZsNITSSqljQ7sZebBBbH5PnvD//
GUylcskZX0q//6KYN58BNwKBgAiNSn4rxh32VCFy8eo0sw/q4QyYxIzZNBalnyCo
HePexu5nfk0q7Ry8V4SzcWstYtVWsN13p7E7/AyNZDGZ4pMG81TDES6Leir1iZnQ
lEZSYAbvbkfhTaUOnU5krhoK00S+ixLKgjSxGIvgug/Iub2dR6hSuqPKVvgumvF2
egYdAoGAFH2UHTH4pa0U0PXG0GMZG40gMpBLN4jJTruynf/viwdvoKkpUWdjbjl8
GNwlttIq9YzYfrUTLn1sYhROrEUIUFmxc2ttFGPjMUifzosVDUbzMRbwhOpSgzof
K6dgyN0czk0YccR6L6TCQ+lGL1zNdE567TUMT8DcujPRzPp6NhA=
-----END RSA PRIVATE KEY-----`)

// SignerTest04Crt is the certificate for singer-test04.
var SignerTest04Crt = []byte(`-----BEGIN CERTIFICATE-----
MIIEWTCCAkGgAwIBAgIQSVCO/y+/drLLF1/Ehl5HAjANBgkqhkiG9w0BAQsFADAo
MSYwJAYDVQQDEx1UZXN0aW5nIGNlcnRpZmljYXRlIGF1dGhvcml0eTAgFw0yMDA3
MTQxOTEzNDVaGA8yMTE5MDcxNDE5MTE0MlowGDEWMBQGA1UEAxMNc2lnbmVyLXRl
c3QwNDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAONujJphNIMRsSMC
NGFGip151lMGNUGCOimjCvy/DVmBGIvMacqCp267rUbogz1fyVt8JMwjgbaA1Uou
iJoMPC5PFnz9+/jETbjopGGxMXI103R/FWIGlpTozWq6G7u9ou7mBS/zWPen5aBt
IwVZKQFNBDiGUn50w3F5Kt/otBLwFHL9qlHfsYeIGWzsI0y3j3oDSGSh6JJKGZDv
MBSh0bxHnZPoVkHU15otF4uttSheYp6+otZZrVTb/C3Lr87CretjkCCqTFdqossn
bmRfkFy7X+c1LaTiT8IfHEiOWehsJA6Cy3F3dxBj1knccBjShBHTyKiOzJSxY9Ug
CG8AgE8CAwEAAaOBjDCBiTAOBgNVHQ8BAf8EBAMCA7gwHQYDVR0lBBYwFAYIKwYB
BQUHAwEGCCsGAQUFBwMCMB0GA1UdDgQWBBS9d8lNgz2ImVGa49a7JIeFY1R8IDAf
BgNVHSMEGDAWgBTyKcD02eohJPlIFHovfP2ttl3UGjAYBgNVHREEETAPgg1zaWdu
ZXItdGVzdDA0MA0GCSqGSIb3DQEBCwUAA4ICAQAUFKQqRVcsOJmKcrNhOeCPfjdR
1Xt4S9XYGEtdxOSHcnwiOvj/p3lCJAtqphVEDxe4THieeL05xMcOS9Q01TNTiNV5
dIm/N/hxcKKVYEeM7HtFSHMRIRuxb9wD1lgDV+ZhXXYNbs+dnSt3mENP5+FanVRP
JhfLPxbtaG/QR+N7ZQyvcdWv4jHfvPGB6l+ySP8IZOo5SktDTeEJYsjUgKZbSyvb
T6CTpftItgJo8yaRdbE5TsWx017KWh0vlk/dQAvvxQakN9aO3/JgtxcLPEkmNJxN
mIJA1xAcxV5/ixW+mX13lDOTsQc/PwrvzKW6OFvlvgPSfCMW0CVaFJf/yNCkngoc
DQIKPqRX46VLhgkGhsl8f1L7iuMn3tkERzrQx8XQr8f0zYTVyp5DbzdhZ51e9MmS
LNeC4NWDGhBhlepSpnWISiCgR7zkQmnbAL2CRHhNXu7f0aElRrvQ3LyT45Zvncy7
763lVGpdsdSEY1kKISMyFzGWMSDQI68J+5rEtuLews7KQimaKrEcp1b7bH5O6IjJ
hedkCs/Yc/n4FDuudYvdsDHFT2x0HTZn+qkIX57gpzSek3+1b3I5hJQtWgcSx/NR
neIRkZVTCW0R6hfKW7lea9rDTHMiyJ8jpscZjqg5KITLxJJ3ipg3miXxjeUMOZRW
6u0QVYXCMU+um5odFQ==
-----END CERTIFICATE-----`)

// SignerTest04Key is the key for singer-test04.
var SignerTest04Key = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA426MmmE0gxGxIwI0YUaKnXnWUwY1QYI6KaMK/L8NWYEYi8xp
yoKnbrutRuiDPV/JW3wkzCOBtoDVSi6Imgw8Lk8WfP37+MRNuOikYbExcjXTdH8V
YgaWlOjNarobu72i7uYFL/NY96floG0jBVkpAU0EOIZSfnTDcXkq3+i0EvAUcv2q
Ud+xh4gZbOwjTLePegNIZKHokkoZkO8wFKHRvEedk+hWQdTXmi0Xi621KF5inr6i
1lmtVNv8LcuvzsKt62OQIKpMV2qiyyduZF+QXLtf5zUtpOJPwh8cSI5Z6GwkDoLL
cXd3EGPWSdxwGNKEEdPIqI7MlLFj1SAIbwCATwIDAQABAoIBAEOkiaUQ7DFffbPP
4AxekrOrnRcsbYKCXRItMANkP3AzeT17GqvkmC/TGgJQ/VevuY/AKKGeneHOB6H8
nRxUL0IGu8WJNPwURpmMd/emX/J8F9w6P/3bv8WiSBZKLB12lZNATqKoWfmushk4
3IWmsw9z/KFjMf3ydH63bw0RhR1sQr5TT0BLje/0z1neaoVNfocTgySpM4DH+Mpv
ZE38VR4+MlujKpbdGwR0qcBwvA5tG4jkGLBf3TPTYYMpjnDZZ9ZRYQ0BDJhWEA7a
3Dpy2A1OGsXIQ5wt3MjPqK4el1mIcTaaUvaIHNEzB9Og6y07uMMzJJtTJsz4pgcq
VauWWgECgYEA++RN5Ii7Bf4J/PyesUusgkMpu22Umc4UImipJzGhwLCauZdKn7Cy
ZwvZH4aVGfoooCWlYTwDlv2exWinz1RvEGsP9SMykOk66A+RU/Ej7PP7zqQcZN5k
P/Qn3BfyO5gksbDhnm+DJl1i7Gre+T09ZsHCunGWCS5jatlfUhfMZMECgYEA5yQe
tFl4icOPPxOA7ISWy5dhleTD97DikqCgntqxA52SgqtF2LAHXBl+6bdQBf5w4zpu
FBoXtVBZdbmzK2ms3LQpeyzuYuRNJLwtrVGs+YLUXrm7aZPHBPKGH+kRTaU62t6A
ou+hytMNMXDDYICbf1snEemKC6H9qSI0A6Vf2Q8CgYEAk1ss4ivG1SuXJkOWhnuR
kKa/zCC/1PZEuxhlFEOpr7Lg4P+LrT3OMBAzVYkCwq6gg4diZy6XnwYBktS1jsmD
K6SmMi8EDtgSGN26k9O2w5C711gUMEIVfYBUrSHpGEnZ9YVXh7sOiywIieu/Qyk7
OVjSlQWL0xUrHb/KTpkTfkECgYEA0xZWgCfxgAa7bERYiiewOy/9q9Fm8m51DKlq
5ogb/oxJv9HAkNp5bi/OioyhpUewOqQi4XLO0gQWHwA/U1dHyasy8s2ey2tp1DGS
mWUszhUf8341XH7b03XU7ZGA/uL1s+pdme+0VzGVK+CqXRg2agGJ1b7tFiCTMoIM
9INz3BsCgYEA+29TYqx2UTun9PxELEbJEZ/IC8jMmJROpx6Ido7mky5huJlgfhtU
C7YbhCkChxM6hbgehwGWH8xj1jbVkliPe0XpABVkIACXKtQMzhg4shHlCmRTgZ+h
xNFJxVdoGRAVUthJyUTcAxWZMH4K5giMdzcQk4YxANekoBmUw3itinA=
-----END RSA PRIVATE KEY-----`)

// SignerTest05Crt is the certificate for singer-test05.
var SignerTest05Crt = []byte(`-----BEGIN CERTIFICATE-----
MIIEWTCCAkGgAwIBAgIQFx8Gc/4fzxgTzS4MtpoovTANBgkqhkiG9w0BAQsFADAo
MSYwJAYDVQQDEx1UZXN0aW5nIGNlcnRpZmljYXRlIGF1dGhvcml0eTAgFw0yMDA3
MTQxOTEzNTNaGA8yMTE5MDcxNDE5MTE0MlowGDEWMBQGA1UEAxMNc2lnbmVyLXRl
c3QwNTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKuCo3uGGnlWK2Mr
qaX8Ep8USj+298tjTUogBUe1dvlkgNBiY9MiUuaBg3VaAqJy70G/nMuWos/Z/+m7
c+SyJaZaMGtOf91gSpSWWN2Ugh8PivgJMi+FWiSKwK3sYEN7V3g3iABGkNCfduA1
4eTAdKCBowQ7Xxz5wpWL5DjjgW6B0KbjL4iREzTwjl+bPHdOFEHzIdxLh18CRU1D
dvdwEwn8ST0VFdjl2hQVm+EYlMv8gyyl9mFaQ16YGW4XJSIjXNDIcxpe7ux6157o
IO9FS21A4aiNKWLrLtJEErX2F/PHZPb2i5L8kbumhom186VJssIlMGGFIfO8UiYC
ilwOWaECAwEAAaOBjDCBiTAOBgNVHQ8BAf8EBAMCA7gwHQYDVR0lBBYwFAYIKwYB
BQUHAwEGCCsGAQUFBwMCMB0GA1UdDgQWBBTYCgsnT8K6EV9hMxe5laZJHW4mNTAf
BgNVHSMEGDAWgBTyKcD02eohJPlIFHovfP2ttl3UGjAYBgNVHREEETAPgg1zaWdu
ZXItdGVzdDA1MA0GCSqGSIb3DQEBCwUAA4ICAQCOhEuD8fMm5iEU+416MoTQpLwX
irHPRym/5fGQxV5AG7ij9X/Pb9WPa7EO24K8/zxcmr7zdkMM/g89PmT/+Aa2Gy7+
XGNhGOQd3xgNcCk0cQ9GL/4HeyPWvLh681difsUEx3Gazh2NXlRAw2rScHmPuu06
+qZwu+zJkL8BUEXMgP4zh10B26zQQ/uFyE4M6Y0h4s2bOC06+vK8OBasklbC36ry
4x4BKlliADx9jtzHmsZJ30jMvjh6Uc2aPo91CGNXSfw4Zs8Ncz1fbyYlWd9goB9C
gxQsUK7o8JARioQQ0JPkMwh/4C2d41G7UFSDesOzbOlx+llNnE7JiWJiV6MFX7VV
Sy92p9/n0MV4pHYtiiftPhvzdLXRhqdKb+hAeYjEAquvHU7MSEv2EAB4Cen0wgh8
CAFqYrQNjZgiTZuPwZljoMlYmuGue91GX17+UvNA4hpUKXWxzampkKnO/Pp/TbCn
DQenKzBwoeY+faTZq30wAyvxfVWh1uebTxIf/7nUdUkVPVZDCwhbQCMxomXHwOhF
4xMTHOmCFxwOuUWqUwQeXIYGgQx0sm/qjx48+jetmJMOOvE4NEXVuRsP5Ggylqa0
vVTqf0eZISAIaBg+dVBe3NNCAfT89DtCE5j/pnqwbcK5sQYF1oxLLs0Wchz31L7Z
8lyeVv/66rxC8CrFSQ==
-----END CERTIFICATE-----`)

// SignerTest05Key is the key for singer-test05.
var SignerTest05Key = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAq4Kje4YaeVYrYyuppfwSnxRKP7b3y2NNSiAFR7V2+WSA0GJj
0yJS5oGDdVoConLvQb+cy5aiz9n/6btz5LIlplowa05/3WBKlJZY3ZSCHw+K+Aky
L4VaJIrArexgQ3tXeDeIAEaQ0J924DXh5MB0oIGjBDtfHPnClYvkOOOBboHQpuMv
iJETNPCOX5s8d04UQfMh3EuHXwJFTUN293ATCfxJPRUV2OXaFBWb4RiUy/yDLKX2
YVpDXpgZbhclIiNc0MhzGl7u7HrXnugg70VLbUDhqI0pYusu0kQStfYX88dk9vaL
kvyRu6aGibXzpUmywiUwYYUh87xSJgKKXA5ZoQIDAQABAoIBAHHPie6xKaY04Su4
0KXdpPm9PBwVrU9hAOvcXcSdDWsnHzeo4rc7gqmFbheUeEUWYeBOZ5zNxBKLhf+z
LY+oC8Xc+V5rw/vCJxt8vVGkd8hmxwAbEhtosU2oaX3Aaoy9L2kN6IjOStC10Out
tc3XTE0qWWO4hkAzGMWmA/cgOUZ7dfSe+qrBiMcBcQiDWozeSztl1FR1Wbq34rYI
ql0J3at0Ca9hiDR82T7DYAEZHkaASCvJt5wtRYmT+JbkqgSWfQO3F6NEXk7hOBJI
RYjx/enPko8JDTY/OQS7JtB78WX4n9ndeEdM7W9rJUAiz+08oODIlPtYV5Iqrov8
Q7cF+QECgYEAx/Y3HYza2DbychO8f0vjH/tdsjySA+BGxF/M0dQ1GpMceIsNdhPM
B3vSSsQaROIhi29++g2QjFRqn7o6K0yw5DfecerzCVo1CVSIoTZQDLSiP3pV+vfN
C01uWS3wnoRlnAYIHbyIfz7nGLu56JQ2HNnBIt1I6Rb8tLBjRhFpIlECgYEA25M9
Y8Q/Y8Pcl6GtDNxAJAj7DRrhBGRwgyu8Ka/MuqLWM6DN8bvsZdHP45FXG8dSU0Sz
0ybKuSjohnopacncBOnpjh/sEqqv3S+4/j8jghmIa+EjsIPlxIKjX7dwXmwXXMRf
zBUpXJpOfFe9ulacauv6mkOobdmZA0ZuSV01HlECgYEAsiiODLcN2SyDsN4iySw8
7abRYVeUJP3zL03HuIAg3E+MQ15pHPYgh95rjA/S8+KiTpCFipcBXfZslWmgICoR
RRT+DXNvHLsRnAGERlaU1e0uze3ao2ObfeF8WtqSkzmVKoE5Q/1RhEXwMto3lqBO
4j+lU8HjD0Ja30Z0/N3QVTECgYEAzCPOFW0lChnUgtz7SFOVBmubDIoK3cyLlkgY
/iDFlFdEEmfUUIRIujkgsBA2Dkt3zJa5IPyVySOxWyVET2guuBrI2yvujURHkLqH
oUkOOCyI2tNMHRXjjpluTtT7Ea4o4kRoVBLwi3mispPYft97OST/rmBsvQRq79KR
Qn/3nKECgYASJ8lZrUMTl74amFSGVrd4xRPhKJTer4C1HeLlwUaAKlEksNvpNecw
l3inYSityX8BXZzOje0pINg9Hd4+ylPzh1PxqCrRjyNNb5R4z2nJrVb8hDCcoG/p
Mia3unc6Qwos7FDwj9okULGK9P+lq+3vDvtiFflTZZsq46qmzrEmtQ==
-----END RSA PRIVATE KEY-----`)

// ClientTest01Crt is the certificate for client-test01.
var ClientTest01Crt = []byte(`-----BEGIN CERTIFICATE-----
MIIEWTCCAkGgAwIBAgIQEmQiGczWl8GehMItbxxtpTANBgkqhkiG9w0BAQsFADAo
MSYwJAYDVQQDEx1UZXN0aW5nIGNlcnRpZmljYXRlIGF1dGhvcml0eTAgFw0yMDA3
MTQxOTE0MTBaGA8yMTE5MDcxNDE5MTE0MlowGDEWMBQGA1UEAxMNY2xpZW50LXRl
c3QwMTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMgqnlKXWgKe2/tl
OYQg+szYJKNCOnzq9h2LgDyGAg7weKToPFhqDDJeONV5Hs8Oe5J4hLw8nTy3QVc5
Ir4sqjdBE522X0ZKSVXkP6WXT2DknLfFqiYHPHfK1t8fjhGcIExfo3sU7x+tyZT+
tQwe3+wMd1/deEl/Y9imKvwJihTCtAORKlemyVzEdGLxlQkZ+OKIbmZ5WD1eru5R
O/sLkohoLqsF6kKwO5nsbs3ovD02kTbXCavofN8lXmk3Kp7e6Bbj0ZWtM1sx/1KD
iDKsX1eobdpl6cMhPiBsAbxBG4kQqvAuC8MtIWQrvHO+jPglo3Hfl0FvYhit5iFJ
uodoSLUCAwEAAaOBjDCBiTAOBgNVHQ8BAf8EBAMCA7gwHQYDVR0lBBYwFAYIKwYB
BQUHAwEGCCsGAQUFBwMCMB0GA1UdDgQWBBQjib0wXi+YdCOBflTX/Frp1tD/xjAf
BgNVHSMEGDAWgBTyKcD02eohJPlIFHovfP2ttl3UGjAYBgNVHREEETAPgg1jbGll
bnQtdGVzdDAxMA0GCSqGSIb3DQEBCwUAA4ICAQCKW9Da1WZ+TimbjhjKclm60OjT
yHX5pH3eZrG3iqkpSgoBMFKCtIEulHk38DTqEqKESFubU80A4H8lkQSWYxx/OV+3
OLYh3QgXv3R92pVqrnRD1Dspy+YtePrECART+RSCVoN2+4Oc+rqEMJtshQDGMlVJ
wN9mp1OiGiAU0EyNNUDmxoMaWs3+BGMSngbgTY2I3U76RFwWqXi2FK9gGiAfH/Vf
XQ+AKXtPq+pQdLPjGWZ0KfzmHUtaxoR3A3u6G4hi+eh7PY4N64g/3lO5FwzEK2ke
VSn6xq50PWk1OkULyWVobo9VsXiJaAltBV6Le6Eyk3dofFb8zp+mKedDEC/anhXw
tzO7OFltqOiyaJNOdAXgsskom71PTDM0aJEG/8Jl5eOwlEoKi4A7KRFt6rq09SHj
Ww8yjX6EbTUmT9plKYQmKCshJ65e1Y5rGElxTzfzeQjDAV+584xzF7heUcKbBn+o
Sndol4HdgMo+yLwbyPL/GdB2m+Itzfid/cFUd57X/9MuPrRvMHmHoas1e7/DtuqK
riQ2B5PmKjvDz+nrR7aIR3nMFSlA2UAM4sGU0hmAR9XQRf7unKW95tcKRLx/+Blt
b71wse/U4f4QKz+JUntuY/x+Q13/np/DUDz0koYRtwqe+J4nlZdSH9on1i8lTXcj
yWJYDKQTImOKFINmag==
-----END CERTIFICATE-----`)

// ClientTest01Key is the key for client-test01.
var ClientTest01Key = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAyCqeUpdaAp7b+2U5hCD6zNgko0I6fOr2HYuAPIYCDvB4pOg8
WGoMMl441Xkezw57kniEvDydPLdBVzkiviyqN0ETnbZfRkpJVeQ/pZdPYOSct8Wq
Jgc8d8rW3x+OEZwgTF+jexTvH63JlP61DB7f7Ax3X914SX9j2KYq/AmKFMK0A5Eq
V6bJXMR0YvGVCRn44ohuZnlYPV6u7lE7+wuSiGguqwXqQrA7mexuzei8PTaRNtcJ
q+h83yVeaTcqnt7oFuPRla0zWzH/UoOIMqxfV6ht2mXpwyE+IGwBvEEbiRCq8C4L
wy0hZCu8c76M+CWjcd+XQW9iGK3mIUm6h2hItQIDAQABAoIBAQCgORUQEVyYWmbF
3anjbK60x5LUJ9A/a6hjz+VvDOma4bwAbHDZaNGog6lEnzYdEX+yp8vADVjjX+Dq
m66GwaPipLG5/WBYGGCJYzHbL7n3WidkJtHirMonLXM/mLeUuv9Tgv6OKByco/SG
0jdDo3ckMHphfxqo4lKe+avQoSYpJKKRINefHUDTuVbybutAssAJODpfPdrZuoGy
W0ZYNKHUx8vA+yYPL+8LoVaw1C8sekeoVNAIUKAPOHXkqcy4E2qyMlgR0+VDHLHq
7bGnBt1CnnIGyE+5fFTDxzCNwJsHzBB9B6F9W6/eX7QWktVOkt8P+T0uCQukGFgh
s+jX2pCFAoGBAN3l+cTPduSoFxGKVjpE0kyjLIVnohMcDVC74ftCXu0NVVwk3PMV
CqqqBVqMVKUHAzOdnYrvL/YG3XFVV1jpZAZUcp2c3OFpFTlicNRt3UgddxVy7WuY
FsKhXQXh6RWjvptLL6/6ADx+ZrfzrvoiYMNTgJ1Gm852Zge7b1ygSahXAoGBAObt
qHQUWKHFk2+Vgyp54SPy/tc636JIGjPM24RhWM2MZjvOae7MNnFIXY3ibwbshmmu
7bT1UNvp7ex1D58ulhqHzcy8EpIQZdxffEI6JqQ1TK9Q3qc4Pbdg138ezzpW1XQp
jamBOohaVxzhHpwEaAaD5yC2BbDmCiU3WewpMB/TAoGAE/8vnQ7dNgn32jrPPn0J
PqLN1k/aiUJT6NylptD6YP58nMstpjJVPcAIr3pJ/n09Sp9/nQ+lENTZi+cW8gpG
W1Os6ItEVIP1x7AZXutvr5oIK0SqJLIWCwAjs+4B5VNWUARcjc1HzCvP8e/h6uTC
N5gE1SeRzu9YjoXTqVNAkA0CgYBn2+1OP7RZFYYowkKawPQL+ga5gCYCU5FNSM5V
rH6G+6UjMsOb/cZijpxc6sDqiUgukdkg8M/sCDrUhRWAjzA9QGTDtrZXcP6O7Xby
RjsI0Vvq4WEyLe474lcpOg8TeuhHdUTcPl2344GYYHsmyiiK/ZnesV6/38YiVNGh
kCivtwKBgQC0EKA/1BQ5dN30iz6GyBneCtJauJFnBsiNkn8ERpnfgMh1b1XlNRtj
mCJtFtCCoOFlEW09up2daZew9yI7dJVcFqlJgSNAoQGze+1WZXEA2QhHfJprSmGi
MWxGRw6f60FdvZss49OMsMDyRetDDL/dWzwXVQ4QfEAYpDBUUSSMFg==
-----END RSA PRIVATE KEY-----`)

// ClientTest02Crt is the certificate for client-test02.
var ClientTest02Crt = []byte(`-----BEGIN CERTIFICATE-----
MIIEWjCCAkKgAwIBAgIRANEPhPoPTEoCoNdq5yZgyDcwDQYJKoZIhvcNAQELBQAw
KDEmMCQGA1UEAxMdVGVzdGluZyBjZXJ0aWZpY2F0ZSBhdXRob3JpdHkwIBcNMjAw
NzE0MTkxNDE4WhgPMjExOTA3MTQxOTExNDJaMBgxFjAUBgNVBAMTDWNsaWVudC10
ZXN0MDIwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCo4MTVF60qU2Mv
ccV2mP/3ILM40hzThDgRpJd+IjD1KBvNuWbRl9vJlRT3NuwmjjYgG4TEzYMxEwnp
mOHSCmtmGy4sQKPot7gjrgqXdVKt7/9/uqOl2hv9B92AYiK6Ob2ev3dsfwvGCq0I
+CpoaQqdIYwN5ZQVMXOMzeFlu74/jK/HFC1n2THaA1EWWI/FZywsbUA7g41ha2ls
7fOfKnNhh1ovQryoMJeZjiuWbZqzPgvDP4Dszyp7tuMYLQtQ6vvK3MkLzvzYgsN6
1fiGwwYiBu5FYgiQ44MPGaJ5UgR5Pn2ocGmbGdCVm5qfWFtgrVoNtjmquAR5+6Ai
ttB06Kw/AgMBAAGjgYwwgYkwDgYDVR0PAQH/BAQDAgO4MB0GA1UdJQQWMBQGCCsG
AQUFBwMBBggrBgEFBQcDAjAdBgNVHQ4EFgQUVvxTtaJNcnQbbkR3vVwHLsFQiqow
HwYDVR0jBBgwFoAU8inA9NnqIST5SBR6L3z9rbZd1BowGAYDVR0RBBEwD4INY2xp
ZW50LXRlc3QwMjANBgkqhkiG9w0BAQsFAAOCAgEAWXb+03oCqQP7lR4Ijl3k8Bj7
E4C6qbTjAgSUo3JKVCrSIvCpK/f295feyeUhb6Gd6fGhgEsz/HxOUD+p+QS5Z3tq
7Ydx3Wb/EJ3uu3AUZf35ClRsf/bMsanzvKvc94pkCfqbMSDa8s+F73wgZQzCqQ3j
5645BKM0f7W1Lq0O1htl8fpVb8g+vhVhwvBHuK6IIb1Mqh9gSuY+OGR5l0v7B2s7
Lzt2Yj0yDkXvfnrLsTxf3u/O02Xxk5pN34BgRLPP6Rv6Nht3KfIO/mRX1RetLRo5
wkCeMt86on8APWJxlfS7WmB12jXvzNjznlGzyA5Bg8coqq6imqXGa5F2JEclc4Xy
J1afWsJCrZ+e1TF65oEe2c61z6QFWMLPgftKUhEd1nve8xxTroFckx/DSJw155XL
ond2+dpD63znGQYATFRUk2dJfhNRZVcqLABsqHuNJlTHrZInU9FrfoMIQl7HCT7b
Y+cmqQd+Z6c/RhASv6Ej6PQYNfcj7QO4OwloazGBnWllBgmAmERdE0d+65ZYTIf+
jjWljvDx7qK/t4gwk23/9SQoPp66Yr60Jxm2vK1o70wxxKAe0G/skMO7oH8+1td/
L/7LyA4miSuym3OotnmL54xOAIZvwdr7ULlHaADt6N3QUpmjIfF5za7dzCek8iWv
+gv/m7jLA2aihJxgcOM=
-----END CERTIFICATE-----`)

// ClientTest02Key is the key for client-test02.
var ClientTest02Key = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAqODE1RetKlNjL3HFdpj/9yCzONIc04Q4EaSXfiIw9Sgbzblm
0ZfbyZUU9zbsJo42IBuExM2DMRMJ6Zjh0gprZhsuLECj6Le4I64Kl3VSre//f7qj
pdob/QfdgGIiujm9nr93bH8LxgqtCPgqaGkKnSGMDeWUFTFzjM3hZbu+P4yvxxQt
Z9kx2gNRFliPxWcsLG1AO4ONYWtpbO3znypzYYdaL0K8qDCXmY4rlm2asz4Lwz+A
7M8qe7bjGC0LUOr7ytzJC8782ILDetX4hsMGIgbuRWIIkOODDxmieVIEeT59qHBp
mxnQlZuan1hbYK1aDbY5qrgEefugIrbQdOisPwIDAQABAoIBAGKzU3bzqoqFR//v
r6f8DgXCf9ziuGRK73YoSz41/4UZFv7qsBQpfBRZ5HIEEIsMLMghLTzUnCtvZCi5
6KmY22JG2nqUoPefgKihzKDuug7cqOhfDcweKKN/GE8zi/ZpTtDcOJEZw6jQjoci
e66gTcq+U5u/pQep6k3N+kE6NBvxS40+6nvWpqDNShGXf4Wf4H+ZFlueZn9ExQJ+
SzGpbuh7/kQSN9KiE0TZPwuRHxUG9S/jyGqpoUx3m9LPcymLE3sW1z9JQokhmyUM
nncSj/3A3bCdJ+Gk3vN78UmDp8hs+naKXqYpohFLrLzf1euojeTUUCTk3/ZUNLd3
wMtbhrECgYEA10evXE6d0xzT90WF0bgOBxxgVo7ArWngjOBNg2snF4OMvMiCusng
3ACYb3r3FP7XTnmtX0pzDhRiR2Dtdj+iKUQj3Edlqze4Jl1lapcP00oGHaL0aWHS
sSJ1dLiKOX+J3NipOb6SMVyRC2TCtfDOrVKjTh2IL/r12tjEQZ4bwncCgYEAyNIz
KqBYfLfnz50zzovbxGpXXioWCZnbegTIEuav24KVA2Dnxm6+X3jSa2y6Jha+KKcn
ck7xLlyXNTiFDteCoNOHDoFD29deThIqhVENaaj99v6uqa2pFKaLguKSpmVTYGfN
nzgc0TrH24EotKwUKNQOPeeim40M2SXaakyiznkCgYAA50EaHw5Ue3N8PyNvnNka
OIlC87hlQeN6U9qfaYyxcZDenezGgeBaq/n3xclRojwfDS2oD7Tp4zYCXKrCa0Pv
7mREIVyQ2lwAdTXeu6GKXm4mI6/o3Us4CQ/7HGcgFKPsdFBJeL4+TMgjBxo2Dzue
yJD82+zdXq03bN9t8w7kwwKBgQCff/QLoxck15xSFXWUUxjVw/BoZdzi+0SKgUm3
WWnLWseny0vLGyIxCfF2SQEAR4f5GeoGrBR60id2qdFknF6wBdF6/8g5z7CjKWKE
SH/yr7omdtmbAscME8syxWncpxW1uDxkfVjGBX2/JbKC7TmFxAcbu+I20aymu53i
V8PnYQKBgHJ6g/YmCVsTETxiL7jQlLhmiitI4DBtu7FJmh8L75eyZs7JyXrrsmu7
+IMc+2eF2yB0Po9TInzCyN/eLVANZEL0uqQjT1F8MoftJspe5GxRtTtT/iv02B/O
5DKMSYKqDJVSH4D+Rw9RJNYQvElXBXo2/jwnTxaEXUarNP8WDRb1
-----END RSA PRIVATE KEY-----`)

// ClientTest03Crt is the certificate for client-test03.
var ClientTest03Crt = []byte(`-----BEGIN CERTIFICATE-----
MIIEWjCCAkKgAwIBAgIRAOEjdJ1nM7Zr50Ht9TA4s4EwDQYJKoZIhvcNAQELBQAw
KDEmMCQGA1UEAxMdVGVzdGluZyBjZXJ0aWZpY2F0ZSBhdXRob3JpdHkwIBcNMjAw
NzE0MTkxNDI2WhgPMjExOTA3MTQxOTExNDJaMBgxFjAUBgNVBAMTDWNsaWVudC10
ZXN0MDMwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCgiezBVyx8iy0i
RCzmht//pvKRnzRMBCau0Bb5VLhYkzAMnQCKxu7k9+bi70rgV3HOOzFfYEJRFHqN
DKqodRQu8rmCMBSwyS/uALJk2ms/QcUVXD73Q6uD2m2rARPXa9AroNJ0ONnaOTRT
GHr6Eh+obiCj7m5yuHqaX7QBcM4AXqdN8H57xsH6zBs8xPKozI0QRadtMF0OFbwe
bHvf9CJ+dn0KsxnJlhP+Lj5b95LYCazOAiDZdWTVn7CAZr2L/5kW25VxPlrrVfEj
i1MxEvOkdq4tH36+VjUu/NqC71RkLiNP+cDnokkn4ygzD+hKcUKkpm68Aox0yrBA
c13KDmUpAgMBAAGjgYwwgYkwDgYDVR0PAQH/BAQDAgO4MB0GA1UdJQQWMBQGCCsG
AQUFBwMBBggrBgEFBQcDAjAdBgNVHQ4EFgQUWEUXCwHFixg2RMkafbFr4MtqqXgw
HwYDVR0jBBgwFoAU8inA9NnqIST5SBR6L3z9rbZd1BowGAYDVR0RBBEwD4INY2xp
ZW50LXRlc3QwMzANBgkqhkiG9w0BAQsFAAOCAgEABcbrz/GuHwXE5aGxwbRPYK4D
UoMgAxfKGdMeq/q/KFF4sFUdPL+9+OEw0mo7IJDKzSDyxAyMa+ZZ0/gRuMYVPDtJ
VUt2PVpf3njhwbXqrvZaZ7zqCa0Lcsie0MReS7PPW+rBYLE4kXEROl/VpxMVO3XD
TDlmjfcpsOmB3APHUSRoKvf6wBTngZqyUD4+MI/ICLJTZ48BMAZQnsdfWsCrv3pY
LLICO+ie+T+Q8SA5kj/IhQNwFPtmDtT5b0JMA9OCXZxgi+AjxJd/fC3j0kAYX5/W
IlR+I5OMns8y5Cn69BvkQsaR0QusS11FiSlYLpvu9ZNCaeCBl49GDlKJHkurFJbf
J9sj0Wri23i1rO4BX7iUmxHORj+YkxVIu5tWSRG5pvhrzUGsi9kZUnwuZBH95ctL
4RybnKldaEtEFVhpgcRITqxl0KLSDIQKYXM/E6BkBP9sxUijw8kdn+elDYvYKRbc
tegroRGgQlvv+tmfywnbl3wgqhJNofA6F1w3B5O/osu9kS1d1lTcZRUTBPfz445T
yaDE2hKJ3EAMU19pENsoQ03RUdXZdNcIgV9bfQXjVNQgR0M3BJsvtmD1Axc2iz8Y
muV8tNq6HtzTVQENwY9XjOQuMA0eIhjfScNtopgVqzKuH+OtmkMULK/MREce6G3D
ZDPJlY1l7uBTkvtcg/4=
-----END CERTIFICATE-----`)

// ClientTest03Key is the key for client-test03.
var ClientTest03Key = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAoInswVcsfIstIkQs5obf/6bykZ80TAQmrtAW+VS4WJMwDJ0A
isbu5Pfm4u9K4FdxzjsxX2BCURR6jQyqqHUULvK5gjAUsMkv7gCyZNprP0HFFVw+
90Org9ptqwET12vQK6DSdDjZ2jk0Uxh6+hIfqG4go+5ucrh6ml+0AXDOAF6nTfB+
e8bB+swbPMTyqMyNEEWnbTBdDhW8Hmx73/QifnZ9CrMZyZYT/i4+W/eS2AmszgIg
2XVk1Z+wgGa9i/+ZFtuVcT5a61XxI4tTMRLzpHauLR9+vlY1Lvzagu9UZC4jT/nA
56JJJ+MoMw/oSnFCpKZuvAKMdMqwQHNdyg5lKQIDAQABAoIBAFFhWZkw3aZMzxZd
xP0fRhHMyY40TEh0mj9n0R2XatPL2UGrnQi8i4GagXpsn2JWMaS/sOmOXE/Jt84k
q6Y1o4OhG5T5VqO/eQKHa+dgKZmpd4S6fFjP/vQaGBnls+8yAhb7sXOlFhpfgTzW
XhDCezXgjb9MMMUK62uyNAe+7U6V3xsuHojwd16CV3xDC2q7gbSwsMZm2HChBYA2
NLxgE7C7WJWTvch7nvZYYprFs982P8MZWgbRkwhiw0EdO2baEQNTMpJwDhT7Gcl2
dKVM4zNLtTA4nw7V38TZDpTlGBYtFcHknogisPVxkUGiUK+kNgWRFyduFjDyQu2n
m58QxWECgYEAwEd1TK2WfRe1YuCN4A7nUcQByJorZqYvaXSCyZwhIPIN+MqRhtef
vDfMVDCgUqtUqlK8JXnz0arV+rYdgVHcgB7mDOTdrz8wxjlOl5GoYDMhCP52vDYa
1VRNQ1lELAOrdo8wPSKoxXtGA8p0LDFm38zB6UbLhQzsIDcooDmnzx0CgYEA1b2u
8s+Bl9L4d6xxdCxEOX5FtV/aNGEjYdWjMmwag8jcB+dT5bvfYlEPWR98riDisnxQ
EVk4qhGjUqoeW0fTmy9QxSYp9LMOSoUDRzDfl2bswELB2vap0uU0oHDbT57HeN7a
IAFRCwg/FKIoeEmkNeusI/3f1W8MMm16Uj+wFH0CgYACxQS7hgSU6LKEKhfhElXi
p6Ae85mMcPhd3H/Fx6nyf4oT+1b9Sj3SyDr5O3oTtsQRb/+lyovoiT1rzxO9uSAj
+E34AZPv4kkhkdG//SkfuZzQNFohe+YHDJ/QSIji2Wqu2oEnYEhuD8iCZXgm6s9A
igKBCbQExprgG+tJ44q3QQKBgQCn46TGCQSMUyTkK8m6LQMx/eOXgkENn7eBI6Ra
+Nsi5OUgOC6IZ7ghq/ZYVQlEZYRsGoVx+xktTUlypznNDXBDlzjkgwO1t+fj1PMs
OlGGxUv2APnwmovuoidiVwONWPTqFnJTbXVRKxsRhAYx5fZcfE/svX/SULXN7nyE
mo4eqQKBgQCIPXpp9KhH40dHfDnOdTi+xzUBTnJwxA2yETFhmFsrc245c19saZNK
XDCQeaEUOuLjokRojHrN27ghz3wJoUJzTdjtHZ+yE9s2cW9VXRUxkun0aLcQd/iY
xvXEII1qCvhI4mmWvbcepvtswQmTEeOroM/27/IIEAOvTyqd6rDi2A==
-----END RSA PRIVATE KEY-----`)

// SignerCerts gives access to signer certificates by ID.
var SignerCerts = map[uint64][]byte{
	1: SignerTest01Crt,
	2: SignerTest02Crt,
	3: SignerTest03Crt,
	4: SignerTest04Crt,
	5: SignerTest05Crt,
}

// SignerKeys gives access to signer keys by ID.
var SignerKeys = map[uint64][]byte{
	1: SignerTest01Key,
	2: SignerTest02Key,
	3: SignerTest03Key,
	4: SignerTest04Key,
	5: SignerTest05Key,
}

// SetupCerts sets up a number of certificates on-disk in the provided location.
func SetupCerts(base string) error {
	if err := ioutil.WriteFile(filepath.Join(base, "ca.crt"), CACrt, 0600); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(base, "signer-test01.crt"), SignerTest01Crt, 0600); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(base, "signer-test01.key"), SignerTest01Key, 0600); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(base, "signer-test02.crt"), SignerTest02Crt, 0600); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(base, "signer-test02.key"), SignerTest02Key, 0600); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(base, "signer-test03.crt"), SignerTest03Crt, 0600); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(base, "signer-test03.key"), SignerTest03Key, 0600); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(base, "signer-test04.crt"), SignerTest04Crt, 0600); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(base, "signer-test04.key"), SignerTest04Key, 0600); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(base, "signer-test05.crt"), SignerTest05Crt, 0600); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(base, "signer-test05.key"), SignerTest05Key, 0600); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(base, "client-test01.crt"), ClientTest01Crt, 0600); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(base, "client-test01.key"), ClientTest01Key, 0600); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(base, "client-test02.crt"), ClientTest02Crt, 0600); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(base, "client-test02.key"), ClientTest02Key, 0600); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(base, "client-test03.crt"), ClientTest03Crt, 0600); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(base, "client-test03.key"), ClientTest03Key, 0600); err != nil {
		return err
	}

	return nil
}
