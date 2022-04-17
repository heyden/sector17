#dependency-track

# API Usage

## Importing an SBOM

Take a CycloneDX or SPDX SBOM and import it into Dependency Track. SBOM processing is asynchronous. The operation returns a token that can be used to check when the SBOM processing is done.

### Request - From CycloneDX Bom

```shell
git clone https://github.com/juice-shop/juice-shop.git
npm install

npm install -g @cyclonedx/bom
cyclonedx-bom -o juiceshop-bom.json

curl -v -X "PUT" \
	"http://localhost:8080/api/v1/bom" \
	-H 'Content-Type: application/json' \
	-H "X-API-Key: $DT_API_KEY" \
	-d @juice-shop/juiceshop-bom.json
```

### Request - Alternative with Parameters

```
curl -v -X "POST" \
	"http://localhost:8080/api/v1/bom" \
	-H 'Content-Type: multipart/form-data' \
	-H "X-Api-Key: $DT_API_KEY" \
	-F "autoCreate=true" \
	-F "projectName=juiceshop" \
	-F "projectVersion=12.10.2" \
	-F "bom=@juice-shop/juiceshop-bom.json"
```

### Response

```json
{ "token" : "f4a7b70b-96ac-45df-afe9-278d68c72aef" }
```


## BOM Processing

Returns the processing status of an SBOM.

### Prequisites

- token from `POST /api/v1/bom`

### Request

```shell
curl -v -X 'GET' \
	'http://localhost:8080/api/v1/bom/token/f4a7b70b-96ac-45df-afe9-278d68c72aef' \
	-H "X-Api-Key: $DT_API_KEY" \
	-H 'accept: application/json'
```

### Response

```json
{ "processing" : false }
```


## Project Lookup

Certain project related lookups require the project UUID instead of project name or version.

### Request

```shell
curl -v -X 'GET' \
	'http://localhost:8080/api/v1/project/lookup?name=juiceshop&version=12.10.2' \
	-H "X-Api-Key: $DT_API_KEY" \
	-H 'accept: application/json'
 ```
 
### Response

```json
{
	"name":"juiceshop",
	"version":"12.10.2",
	"classifier":"APPLICATION",
	"uuid":"662f362d-66ae-48b1-abd9-ebfdd077c6df",
	"lastBomImport":1634486493956,
	"lastBomImportFormat":"CycloneDX 1.3",
	"lastInheritedRiskScore":157.0,
	"active":true
}
```


## Fetch Vulnerabilities

### Request

```shell
curl -v -X 'GET' \
	'http://localhost:8080/api/v1/vulnerability/project/662f362d-66ae-48b1-abd9-ebfdd077c6df' \
	-H "X-Api-Key: $DT_API_KEY" \
	-H 'accept: application/json'
 ```

### Response

Array of vulnerabilities. The response is too large to show an example.


## Badges

*Note: The badge capability does not require authentication.*

![[img/badge.png]]

### URL

`http://localhost:8080/api/v1/badge/vulns/project/juiceshop/12.10.2`

### Request

```shell
curl -v -X 'GET' \
 'http://localhost:8080/api/v1/badge/vulns/project/juiceshop/12.10.2' \
 -H 'accept: image/svg+xml'
 ```