#dependency-track, #application-security, #source-composition-analysis 

# Getting Started

## Deployment

### Local

#### Docker

```shell
curl -LO https://dependencytrack.org/docker-compose.yml

docker-compose up -d
```


#### JAR

*Note: this method is being deprecated. Preferable to use the Docker method.*

```shell
java -jar dependency-track-bundle.jar
```


## Usage

### Get Projects

#### NodeJs

```shell
git clone https://github.com/juice-shop/juice-shop.git

cd juice-shop
npm install

npm install -g @cyclonedx/bom
cyclonedx-bom -o juiceshop-bom.json
```

#### Golang

```shell
# Download cyclonedx-gomod from https://github.com/CycloneDX/cyclonedx-gomod/releases
git clone git@github.com:Contrast-Security-OSS/go-test-bench.git

cd go-test-bench
cyclonedx-gomod.exe mod -output gotestbench-bom.json -json=true -type application
```

### Importing an SBOM

#### From CycloneDX Bom

```shell
git clone https://github.com/juice-shop/juice-shop.git
npm install

npm install -g @cyclonedx/bom
cyclonedx-bom -o juiceshop-bom.json

curl -X "PUT" \
	"http://localhost:8080/api/v1/bom" \
	-H 'Content-Type: application/json' \
	-H "X-API-Key: $DT_API_KEY" \
	-d @juice-shop/juiceshop-bom.json
```


#### Alternative with Parameters

```
curl -X "POST" \
	"http://localhost:8080/api/v1/bom" \
     -H 'Content-Type: multipart/form-data' \
     -H "X-Api-Key: $DT_API_KEY" \
     -F "autoCreate=true" \
	 -F "projectName=juiceshop" \
	 -F "projectVersion=12.10.2" \
     -F "bom=@juice-shop/juiceshop-bom.json"
```